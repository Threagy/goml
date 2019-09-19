<template>
  <div>
    <v-app-bar dense app tabs flat>
      <v-toolbar-title>Containers</v-toolbar-title>

      <v-spacer></v-spacer>

      <!-- <v-btn icon color="pink" fab>
        <v-icon>add</v-icon>
      </v-btn> -->

      <template v-slot:extension>
        <v-tabs id="contTab" v-model="tab" class="elevation-1" show-arrows>
          <v-tab
            v-for="container in containers"
            :key="container.id"
            :href="`#tab-${container.id}`"
          >
            {{ container.names[0].substring(1) }}
          </v-tab>
          <!-- <v-tabs-slider color="pink"></v-tabs-slider> -->
        </v-tabs>
      </template>
    </v-app-bar>
    <v-card-text>
      <div id="cont">
        <v-speed-dial v-model="fab" bottom right fixed absolute open-on-hover>
          <template v-slot:activator>
            <v-btn v-model="fab" color="blue darken-2" dark fab>
              <v-icon v-if="fab">close</v-icon>
              <v-icon v-else>build</v-icon>
            </v-btn>
          </template>
          <v-btn
            fab
            dark
            small
            color="blue-grey"
            v-if="showFabTerminal"
            @click="launchTerminalOptions"
          >
            <v-icon>code</v-icon>
          </v-btn>
          <v-btn fab dark small color="pink" @click="createContainer">
            <v-icon>add</v-icon>
          </v-btn>
          <v-btn
            fab
            dark
            small
            color="green"
            v-if="showFabStart"
            @click="confirmStartContainer"
          >
            <v-icon>play_arrow</v-icon>
          </v-btn>
          <v-btn
            fab
            dark
            small
            color="red"
            v-if="showFabStop"
            @click="confirmStopContainer"
          >
            <v-icon>stop</v-icon>
          </v-btn>
          <v-btn
            fab
            dark
            small
            color="indigo"
            v-if="showFabDelete"
            @click="confirmDeleteContainer"
          >
            <v-icon>delete</v-icon>
          </v-btn>
        </v-speed-dial>

        <v-tabs-items v-model="tab">
          <v-tab-item
            v-for="container in containers"
            :key="container.id"
            :value="`tab-${container.id}`"
          >
            <v-layout>
              <v-expansion-panels v-model="panels" multiple>
                <v-expansion-panel>
                  <v-expansion-panel-header>
                    <v-row no-gutters align-center fill-height>
                      <v-col cols="4"><v-subheader>Status</v-subheader></v-col>
                      <v-spacer></v-spacer>
                      <v-col cols="4">
                        <v-layout justify-end align-center fill-height>
                          <v-chip
                            small
                            class="mx-1"
                            :color="getContainerStateColor(container.state)"
                            text-color="white"
                          >
                            {{ container.state }}
                          </v-chip>
                          <v-chip
                            small
                            class="mx-1"
                            :color="getNodeColor(container.nodeAlias)"
                            text-color="white"
                          >
                            {{ container.nodeAlias }}
                          </v-chip>
                        </v-layout>
                      </v-col>
                    </v-row>
                  </v-expansion-panel-header>
                  <!-- <v-expansion-panel-header>
                    <v-subheader>Status</v-subheader>
                  </v-expansion-panel-header> -->
                  <v-expansion-panel-content>
                    <v-data-table
                      :headers="statusTableHeaders"
                      :items="container.statusTable.items"
                      hide-default-header
                      hide-default-footer
                      class="elevation-0"
                    >
                    </v-data-table>
                  </v-expansion-panel-content>
                </v-expansion-panel>

                <v-expansion-panel>
                  <v-expansion-panel-header
                    ><v-subheader
                      >Details</v-subheader
                    ></v-expansion-panel-header
                  >
                  <v-expansion-panel-content>
                    <v-tabs vertical class="elevation-1">
                      <v-tab>Network</v-tab>
                      <v-tab>ENV</v-tab>
                      <v-tab>Volumes</v-tab>
                      <!-- <v-tabs-slider color="pink"></v-tabs-slider> -->

                      <v-tab-item>
                        <v-list subheader dense>
                          <!-- <v-subheader>Network Mode</v-subheader> -->
                          <v-list-item v-ripple="{ class: 'primary--text' }">
                            <v-list-item-content>
                              <v-data-table
                                :headers="networkTableHeaders"
                                :items="container.networkTable.items"
                                hide-default-footer
                              >
                              </v-data-table>
                            </v-list-item-content>
                          </v-list-item>
                          <v-subheader>Ports</v-subheader>
                          <v-list-item
                            v-for="port in container.ports"
                            v-ripple="{ class: 'primary--text' }"
                            :key="port.private_port"
                            @click="
                              actionPort(
                                container.address,
                                port.public_port,
                                port.private_port
                              )
                            "
                          >
                            <v-list-item-icon>
                              <v-icon v-if="port.private_port === 22"
                                >desktop_mac</v-icon
                              >
                              <v-icon v-if="port.private_port !== 22"
                                >http</v-icon
                              >
                            </v-list-item-icon>
                            <v-list-item-content>
                              <v-list-item-title>{{
                                getPortText(port)
                              }}</v-list-item-title>
                            </v-list-item-content>
                          </v-list-item>
                        </v-list>
                      </v-tab-item>
                      <v-tab-item>
                        <v-data-table
                          style="width: 900px"
                          :headers="envTableHeaders"
                          :items="container.envTable.items"
                          hide-default-footer
                        >
                        </v-data-table>
                      </v-tab-item>
                      <v-tab-item>
                        <v-data-table
                          :headers="volumeTableHeaders"
                          :items="container.volumeTable.items"
                          hide-default-footer
                        >
                        </v-data-table>
                      </v-tab-item>
                    </v-tabs>
                  </v-expansion-panel-content>
                </v-expansion-panel>
              </v-expansion-panels>
            </v-layout>
          </v-tab-item>
        </v-tabs-items>
      </div>
    </v-card-text>
    <div>
      <input type="hidden" ref="copyTextBox" />
      <v-snackbar
        v-model="snackBar"
        :multi-line="true"
        :timeout="2000"
        :top="true"
        >{{ snackMessage }}</v-snackbar
      >
      <!-- Confirm Dialog -->
      <v-dialog v-model="startConfirmDialog" persistent max-width="500">
        <v-card>
          <v-card-title class="headline"
            >Do you want to start the container below?</v-card-title
          >
          <v-card-text>
            <ul>
              <li>{{ getContainerName() }}</li>
            </ul>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
              color="green darken-1"
              text
              @click="startConfirmDialog = false"
              >Cancel</v-btn
            >
            <v-btn color="green darken-1" text @click="startContainer"
              >Start</v-btn
            >
          </v-card-actions>
        </v-card>
      </v-dialog>

      <v-dialog v-model="stopConfirmDialog" persistent max-width="500">
        <v-card>
          <v-card-title class="headline"
            >Do you want to stop the container below?</v-card-title
          >
          <v-card-text>
            <ul>
              <li>{{ getContainerName() }}</li>
            </ul>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
              color="green darken-1"
              text
              @click="stopConfirmDialog = false"
              >Cancel</v-btn
            >
            <v-btn color="green darken-1" text @click="stopContainer"
              >Stop</v-btn
            >
          </v-card-actions>
        </v-card>
      </v-dialog>

      <v-dialog v-model="deleteConfirmDialog" persistent max-width="500">
        <v-card>
          <v-card-title class="headline"
            >Do you want to delete the container below?</v-card-title
          >
          <v-card-text>
            <ul>
              <li>{{ getContainerName() }}</li>
            </ul>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
              color="green darken-1"
              text
              @click="deleteConfirmDialog = false"
              >Cancel</v-btn
            >
            <v-btn color="green darken-1" text @click="deleteContainer"
              >Delete</v-btn
            >
          </v-card-actions>
        </v-card>
      </v-dialog>

      <v-dialog
        v-model="launchTerminalOptionsDialog"
        persistent
        max-width="500"
      >
        <v-card>
          <v-card-title class="headline">Terminal</v-card-title>
          <v-card-text>
            <v-text-field
              v-model="terimalCommand"
              label="Command"
              required
            ></v-text-field>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
              color="green darken-1"
              text
              @click="launchTerminalOptionsDialog = false"
              >Cancel</v-btn
            >
            <v-btn color="green darken-1" text @click="launchTerminal"
              >Launch Terminal</v-btn
            >
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>
  </div>
</template>

<script>
const jwt = process.client ? require('jsonwebtoken') : undefined

export default {
  layout: 'dashboard',
  middleware: 'authenticated',
  data: () => ({
    username: '',
    tab: null,
    fab: false,
    panels: [0, 1],
    nodes: [],
    containers: [],
    snackBar: false,
    snackMessage: "",
    startConfirmDialog: false,
    deleteConfirmDialog: false,
    stopConfirmDialog: false,
    launchTerminalOptionsDialog: false,
    terimalCommand: "/bin/bash",
    nodeColors: [ 'pink', 'teal', 'deep-orange', 'orange', 'light-green', 'cyan', 'indigo', 'brown', 'purple' ],
    statusTableHeaders: [
      { text: "Name", align: "left", value: "name" },
      { text: "Value", align: "left", value: "value" },
    ],
    networkTableHeaders: [
      { text: "Network", align: "left", value: "network", sortable: false },
      { text: "Host IP", align: "left", value: "public", sortable: false },
      { text: "Private IP", align: "left", value: "private", sortable: false },
      { text: "Gateway", align: "left", value: "gateway", sortable: false },
    ],
    envTableHeaders: [
      { text: "Name", align: "left", value: "name", sortable: false },
      { text: "Value", align: "left", value: "value", sortable: false },
    ],
    volumeTableHeaders: [
      { text: "Host", align: "left", value: "host", sortable: false },
      { text: "Container", align: "left", value: "container", sortable: false },
    ],
  }),
  computed: {
    getStatusItems: {
      get: function (item) {
        console.log("item");
        console.log(item);
      }
    },
    showFabTerminal: {
      get: function () {
        var container = this.getCurrentContainer();
        if (_.isNil(container) === true) return false;

        return container.state === "running";
      }
    },
    showFabStart: {
      get: function () {
        var container = this.getCurrentContainer();
        if (_.isNil(container) === true) return false;

        return container.state === "exited";
      }
    },
    showFabStop: {
      get: function () {
        var container = this.getCurrentContainer();
        if (_.isNil(container) === true) return false;

        return container.state === "running";
      }
    },
    showFabDelete: {
      get: function () {
        var container = this.getCurrentContainer();
        if (_.isNil(container) === true) return false;

        return container.state === "exited";
      }
    },
  },
  created: async function () {
    try {
      if (_.isNil(this.$gomlapi) === true || _.isNil(jwt) === true) return;
      var decoded = jwt.decode(this.$store.state.auth);
      this.username = decoded.Name;

      await this.fetchNodes();
      var container = this.$route.query.container;
      if (_.isNil(container) === true) return;

      this.tab = 'tab-' + container;
    } catch (err) {
      console.error(err);
    }

  },
  methods: {
    getContainerStateColor(state) {
      switch(state) {
        case 'exited':
          return 'grey';
        case 'created':
          return 'blue';
        default:
          return 'green'
      }
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

      var pickContainers = _.map(nodes, "containers");
      var containers = _.flatten(pickContainers);
      containers = _.filter(containers, c => {
        if (c.labels["goml.owner"] === this.username)
          return c;
      })
      containers = _.orderBy(containers, (c => c.names[0]));

      this.containers = containers;
      this.nodes = _.orderBy(nodes, (n => n.id));;
    },
    createContainer() {
      var container = this.getCurrentContainer();
      if (_.isNil(container) === true) return false;

      this.$router.push('/create-container?container=' + container.id)
    },
    confirmStartContainer() {
      this.startConfirmDialog = true;
    },
    launchTerminalOptions() {
      this.terimalCommand = this.getCurrentContainer().command;
      this.launchTerminalOptionsDialog = true;
    },
    confirmStopContainer() {
      this.stopConfirmDialog = true;
    },
    confirmDeleteContainer() {
      this.deleteConfirmDialog = true;
    },
    launchTerminal() {
      var container = this.getCurrentContainer();
      window.open(`/term?host=${container.address}&port=7777&container=${container.id}&cmd=${this.terimalCommand}`, "_blank");
      this.launchTerminalOptionsDialog = false;
    },
    // 선택한 Container를 Start합니다.
    async startContainer() {
      var container = this.getCurrentContainer();
      await this.$gomlapi.startContainer(container.nodeID, container.id);
      await this.fetchNodes();

      this.startConfirmDialog = false;
    },
    // 선택한 Container를 Stop합니다.
    async stopContainer() {
      var container = this.getCurrentContainer();
      await this.$gomlapi.stopContainer(container.nodeID, container.id);
      await this.fetchNodes();

      this.stopConfirmDialog = false;
    },
    // 선택한 Container를 삭제합니다.
    async deleteContainer() {
      var container = this.getCurrentContainer();
      await this.$gomlapi.deleteContainer(container.nodeID, container.id);
      await this.fetchNodes();

      this.deleteConfirmDialog = false;
    },
    getCurrentContainer: function () {
      if (_.isNil(this.tab) === true) return null;

      var containerID = this.tab.substring("tab-".length);
      return _.find(this.containers, c => c.id === containerID);
    },
    getContainerName: function () {
      var container = this.getCurrentContainer();
      if (_.isNil(container) === true) return "";

      return container.names[0].substring(1);
    },
    getPortText: function (port) {
      if (port.public_port === 0)
        return `(null) ➡ ${port.private_port}/${port.type}`;
      else {
        return `${port.ip}:${port.public_port} ➡ ${port.private_port}/${port.type}`;
      }
    },
    getImageName: function (imageFullName) {
      return getShortImageName(imageFullName);
    },
    actionPort: function (address, publicPort, privatePort) {
      if (privatePort === 22) {
        this.copySSHAddress(address, publicPort);
      } else if (publicPort !== 0) {
        window.open(`http://${address}:${publicPort}`, "_blank");
      }
    },
    copySSHAddress: function (address, port) {
      let copyTextBox = this.$refs.copyTextBox;
      copyTextBox.setAttribute("type", "text");
      copyTextBox.value = `ssh coder@${address} -p ${port}`;
      copyTextBox.select();

      try {
        document.execCommand("copy");
        this.snackBar = true;
        this.snackMessage = "Copy the SSH Address";
      } catch (err) {
        console.error(err);
      }

      /* unselect the range */
      copyTextBox.setAttribute("type", "hidden");
      window.getSelection().removeAllRanges();
    },
    addAdditionalProp(nodes) {
      for (var node of nodes) {
        for (var container of node.containers) {
          container.nodeID = node.id;
          container.nodeAlias = node.alias;
          container.address = node.address;
          container.command = "/bin/bash";
          container.statusTable = {
            items: [
              { name: "ID", value: container.id },
              { name: "Name", value: container.names[0].substring(1) },
              { name: "Status", value: container.status },
              // { name: "State", value: container.state },
              { name: "Created", value: new Date(Number(container.created) * 1000).toLocaleString() },
              { name: "Image", value: `${container.image}@${container.image_id}` },
            ]
          };
          container.networkTable = {
            items: [
              { network: container.host_config.network_mode, public: node.address }
            ]
          };
          container.envTable = { items: [] };
          for (var e of container.env) {
            var index = e.indexOf('=');
            if (index !== -1) {
              container.envTable.items.push({
                name: e.substring(0, index),
                value: e.substring(index + 1),
              })
            } else {
              container.envTable.items.push({
                name: e,
              })
            }
          }

          container.volumeTable = { items: [] };
          for (var mount of container.mounts) {
            container.volumeTable.items.push({ host: mount.source, container: mount.destination })
          }

          //         container.statusTable = {
          //     items: [
          //       { name: "ID", value: container.id },
          //       { name: "Name", value: container.names[0].substring(1) },
          //       { name: "Status", value: container.status },
          //     ]
          //   }
          // };
        }
      }
    },
  },
  asyncData(context) {
    return { project: 'nuxt' }
  }
}
</script>

<style lang="scss" scoped>
// @import "@/assets/material/_color.scss";
// @import "@/assets/material/_variables.scss";

#cont .v-speed-dial {
  position: fixed;
}

// #create .v-btn--floating {
//   position: relative;
// }
</style>

