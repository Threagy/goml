<template>
  <v-container grid-list-md>
    <v-layout row wrap>
      <!-- Buttons -->
      <v-flex xs12>
        <gm-container-toolbar
          :nodes="nodes"
          :selected.sync="selected"
          :filters.sync="filters"
          v-on="{
            create: containerCreate,
            start: containerStart,
            stop: containerStop,
            delete: containerDelete,
            filterChanged: filterChanged,
            refreshContainers: refreshContainers,
          }"
        ></gm-container-toolbar>
      </v-flex>

      <!-- Container List -->
      <v-flex xs12>
        <gm-container-table
          :containers="containers"
          :selected.sync="selected"
          :snackBar.sync="snackBar"
          :snackMessage.sync="snackMessage"
        ></gm-container-table>
      </v-flex>

      <!-- Nodes -->
      <v-flex xs12>
        <gm-node-table :nodes="nodes"></gm-node-table>
      </v-flex>
    </v-layout>

    <!-- Create Container Dialog -->
    <gm-create-container-dialog
      :visible.sync="visibleCreateContainerDialog"
      :nodes="nodes"
      @created="containerCreated"
    ></gm-create-container-dialog>

    <!-- Etc -->
    <v-snackbar
      v-model="snackBar"
      :multi-line="true"
      :timeout="2000"
      :top="true"
      >{{ snackMessage }}</v-snackbar
    >
  </v-container>
</template>

<script>
import CreateContainerDialog from "./CreateContainerDialog.vue";
import ContainerToolbar from "./ContainerToolbar.vue";
import ContainerTable from "./ContainerTable.vue";
import NodeTable from "./NodeTable.vue";
import { addAdditionalProp, filterContainers } from "@/utils/utils-function";
import { mapMutations } from "vuex";

export default {
  components: {
    "gm-container-toolbar": ContainerToolbar,
    "gm-create-container-dialog": CreateContainerDialog,
    "gm-container-table": ContainerTable,
    "gm-node-table": NodeTable,
  },
  data: function () {
    return {
      // global settings
      selected: [],
      filters: [],
      snackBar: false,
      snackMessage: "",
      visibleCreateContainerDialog: false,
    };
  },
  computed: {
    containers: {
      get() {
        return this.$store.state.containers;
      },
    },
    nodes: {
      get() {
        return this.$store.state.nodes;
      },
    }
  },
  created: async function() {
    this.filters.push({text:"goml"})
    await this.fetchNodeList();
  },
  methods: {
    ...mapMutations({ setNodes: "SET_NODES" }),
    ...mapMutations({ setContainers: "SET_CONTAINERS" }),
    async filterChanged() {
      await this.fetchNodeList();
    },
    async refreshContainers() {
      await this.fetchNodeList();
    },
    containerCreate() {
      this.visibleCreateContainerDialog = true;
    },
    async containerCreated() {
      this.selected = [];
      await this.fetchNodeList();
    },
    async containerStart() {
      this.selected = [];
      await this.fetchNodeList();
    },
    async containerStop() {
      this.selected = [];
      await this.fetchNodeList();
    },
    async containerDelete() {
      this.selected = [];
      await this.fetchNodeList();
    },
    async fetchNodeList() {
      try {
        var res = await this.gomlapi.getNodes();

        var nodes = res.data.nodes;

        addAdditionalProp(nodes);

        var pickContainers = this._.map(nodes, "containers");
        var containers = this._.flatten(pickContainers);

        var filterdContainers = containers;

        if (this._.isNil(this.filters) === false && this.filters.length > 0)
          filterdContainers = filterContainers(this.filters, containers);

        this.setNodes(nodes);
        this.setContainers(filterdContainers);
      } catch (err) {
        console.error(err);
      } finally {
        //this.timer = setTimeout(this.fetchNodeList, 10000);
      }
    },
  },
};
</script>
