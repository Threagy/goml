<template>
  <div class="box" ref="box">
    <div ref="terminal"></div>
  </div>
</template>

<script>
const Terminal = process.client ? require('xterm/dist/xterm.js') : undefined
const Cookie = process.client ? require('js-cookie') : undefined
import axios from 'axios';

export default {
  data: () => ({
    term: null,
    socket: null,
    execID: null,
  }),
  middleware: 'authenticated',
  methods: {
    handleWindowResize(event) {
      if (_.isNil(this.execID) === true) return;

      var cols = Math.floor(this.$refs.box.offsetWidth / this.term._core._renderCoordinator.dimensions.actualCellWidth);
      var rows = Math.floor(this.$refs.box.offsetHeight / this.term._core._renderCoordinator.dimensions.actualCellHeight);

      var url = `http://${this.$route.query.host}:${this.$route.query.port}/v1/exec/resizeTty?execID=${this.execID}&width=${cols}&height=${rows}`;
      axios.get(url)
        .then((res) => {
          this.term.resize(cols, rows);
        })
        .catch((err) => {
          console.error(err);
        })
    }
  },
  beforeDestroy: function () {
    window.removeEventListener('resize', this.handleWindowResize)
  },
  mounted() {
    window.addEventListener('resize', this.handleWindowResize);
    this.handleWindowResize();
  },
  created: function () {
    if (_.isNil(Terminal) === true) return;

    if (_.isNil(this.term) === false) return;

    this.term = new Terminal({
      rows: 50,
      screenKeys: true,
      useStyle: true,
      cursorBlink: true,
      fontSize: 14,
      fontFamily: 'Ubuntu Mono, courier-new, courier, monospace',
    });
    this.term.open(this.$refs.terminal);
    this.term.focus()

    var xterm = this.term;
    var url = `ws://${this.$route.query.host}:${this.$route.query.port}/v1/exec/${this.$route.query.container},${window.btoa(this.$route.query.cmd)}`;
    var that = this;

    var websocket = new WebSocket(url);
    websocket.onopen = function (evt) {
      console.log('connected')

      xterm.on('data', function (data) {
        websocket.send(data);
      });

      xterm.on('title', function (title) {
        document.title = title;
      });

      websocket.onmessage = function (evt) {
        if (_.isNil(that.execID) === true) {
          that.execID = evt.data;
          that.handleWindowResize();
        }

        xterm.write(evt.data);
      }

      websocket.onclose = function (evt) {
        xterm.write("Session terminated");
        xterm.destroy();
      }

      websocket.onerror = function (evt) {
        if (typeof console.log == "function") {
          console.log(evt)
        }
      }
    }

    this.socket = websocket;
  },
}
</script>

<style lang="scss" scoped>
.box {
  display: block;
  width: calc(100% - 1 px);
  margin: 0 auto;
  padding: 2px;
  color: #fafafa;
  height: 100%;
  width: 100%;
  position: fixed;
  overflow: hidden;
}
.box .box:focus .terminal-cursor {
  background-color: #fafafa;
}
</style>

