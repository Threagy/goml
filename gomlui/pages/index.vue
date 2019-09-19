<template>
  <div>
    <v-app-bar
      dense
      app
      tabs
      class="elevation-2"
    >
      <v-toolbar-title>Dashboard</v-toolbar-title>

      <v-spacer></v-spacer>

    </v-app-bar>
    <v-card
      v-for="node in nodes"
      :key="node.id"
      class="mb-2"
    >
      <v-card-text>
        <v-subheader>{{node.alias}}</v-subheader>
        <v-divider></v-divider>
        <v-row>
          <v-col
            v-for="(gpu, i) in node.gpu_statuses"
            :key="node.id + '_' + i"
          >
            <v-row no-gutters>
              <v-progress-circular
                class="ma-2"
                :rotate="360"
                :size="90"
                :width="15"
                :value="gpu.mem_used_percent"
                color="primary"
              >
                {{ gpu.mem_used_percent }}%
              </v-progress-circular>
            </v-row>
            <v-row no-gutters>
              <v-col>
                <div class="overline font-weight-black">index: {{i}}</div>
                <div class="overline font-weight-black">total mem: {{formatBytes(gpu.memory_total_size).join(' ')}}</div>
                <div class="overline font-weight-black">used mem: {{formatBytes(gpu.memory_used_size).join(' ')}}</div>
                <div class="overline font-weight-black">temperature: {{gpu.temperature}}°C</div>
                <!-- <p
                  style="width:110px; text-align:center"
                  :color="primary"
                >
                  {{i}}
                </p> -->
              </v-col>
            </v-row>
          </v-col>
        </v-row>
      </v-card-text>
    </v-card>
  </div>
</template>

<script>
const jwt = process.client ? require('jsonwebtoken') : undefined

export default {
  layout: 'dashboard',
  middleware: 'authenticated',
  data: () => ({
    nodes: [],
  }),
  created: async function () {
    try {
      if (_.isNil(this.$gomlapi) === true || _.isNil(jwt) === true) return;
      var decoded = jwt.decode(this.$store.state.auth);
      this.username = decoded.Name;

      await this.fetchNodes();
    } catch (err) {
      console.error(err);
    }
  },
  methods: {
    formatBytes (kiloBytes, decimals){
        if(kiloBytes == 0) return ['0', 'KB'];

        var k = 1024,
            dm = decimals + 1 || 1,
            sizes = ['KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'],
            i = Math.floor(Math.log(kiloBytes) / Math.log(k));

        return [parseFloat((kiloBytes / Math.pow(k, i)).toFixed(dm)).toLocaleString() , sizes[i]]; // parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + ' ' +
    },
    getNodeColor(node) {
      var index = _.findIndex(this.nodes, n => n.alias === node);
      if (index === -1)
        return 'black';

      return this.nodeColors[index];
    },
    async fetchNodes() {
      if (_.isNil(this.$gomlapi) === true) return;

      var res = await this.$gomlapi.getNodes();
      var nodes = res.data.nodes;

      this.addAdditionalProp(nodes);

      this.nodes = _.orderBy(nodes, (n => n.id));;
    },
    addAdditionalProp(nodes) {
      for (var node of nodes) {
        for (var gpu of node.gpu_statuses) {
          var percent = (gpu.memory_used_size / gpu.memory_total_size) * 100;

          gpu.mem_used_percent = Math.round(percent);
          gpu.tooltip = `total: ${gpu.memory_total_size}, used: ${gpu.memory_used_size}`;
        }
      }
    },
  },
}
</script>

<style>
</style>
