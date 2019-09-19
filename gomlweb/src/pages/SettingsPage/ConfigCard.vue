
<template>
  <v-form class="pa-3 pt-4">
    <v-layout wrap>
      <v-flex xs12>
        <v-alert :value="alert" type="error" transition="scale-transition">{{
          errorMessage
        }}</v-alert>
      </v-flex>
      <v-flex xs12>
        <v-subheader>Nodes</v-subheader>
        <v-layout row align-center>
          <v-flex xs3 px-2>
            <v-text-field
              v-model="nodeAlias"
              label="Node Alias"
              required
            ></v-text-field>
          </v-flex>

          <v-flex xs3 px-2>
            <v-text-field
              v-model="nodeAddress"
              label="Node Address(IP)"
              required
            ></v-text-field>
          </v-flex>

          <v-btn
            class="mx-1"
            color="primary"
            tile
            :disabled="disabledAddButton"
            @click="addNode"
            >Add Node</v-btn
          >
          <v-btn
            class="mx-1"
            color="warning"
            tile
            :disabled="disabledDeleteButton"
            @click="deleteNode"
            >Delete Node</v-btn
          >
        </v-layout>
      </v-flex>
      <v-flex xs12>
        <v-layout xs5>
          <v-data-table
            v-model="selectedList"
            show-select
            dense
            :page="nodeTable.page"
            :items-per-page="nodeTable.itemsPerPage"
            :headers="nodeTable.headers"
            :items="config.nodes"
          >
          </v-data-table>
        </v-layout>
      </v-flex>
      <v-flex xs12>
        <v-subheader>Images</v-subheader>
        <v-textarea
          v-model="config.images"
          auto-grow
          label="Docker Images"
          rows="5"
        ></v-textarea>
        <v-layout wrap row>
          <ul class="caption">
            <li>
              <small class="caption"
                >Docker Hub Official Image:
                docker.io/library/{name}:{tag}</small
              >
            </li>
            <li>
              <small class="caption"
                >Docker Hub User Image: docker.io/{account}/{name}:{tag}</small
              >
            </li>
            <li>
              <small class="caption"
                >Private Registry Image: {address}:{port}/{name}:{tag}</small
              >
            </li>
          </ul>
        </v-layout>
      </v-flex>
    </v-layout>
  </v-form>
</template>

<script>
export default {
  name: "ConfigCard",
  inheritAttrs: false,
  props: {
    config: {
      type: Object,
      default: null
    }
  },
  data: function () {
    return {
      alert: false,
      errorMessage: "",
      selectedList: [],
      nodeAlias: "",
      nodeAddress: "",
      nodeTable: {
        page: 1,
        itemsPerPage: -1,
        headers: [
          {
            text: "Alias",
            value: "alias",
            align: "left",
            width: "200px",
            sortable: false
          },
          {
            text: "IP",
            value: "address",
            align: "left",
            width: "200px",
            sortable: false
          },
          {
            text: "Hostname",
            value: "hostname",
            align: "left",
            width: "200px",
            sortable: false
          },
          {
            text: "ID",
            value: "id",
            align: "left",
            width: "200px",
            sortable: false
          },
          { text: "", value: "delete", sortable: false }
        ],
      },
    };
  },
  computed: {
    disabledAddButton: {
      get: function () {
        return this.nodeAlias === "" || this.nodeAddress === "";
      }
    },
    disabledDeleteButton: {
      get: function () {
        return this.selectedList.length === 0;
      }
    },
  },
  methods: {
    addNode: async function () {
      try {
        if (this._.findIndex(this.config.nodes, n => n.alias === this.nodeAlias) !== -1)
          return this.showAlert(`Alias[${this.nodeAlias}] is already exist`)
        if (this._.findIndex(this.config.nodes, n => n.address === this.nodeAddress) !== -1)
          return this.showAlert(`Address[${this.nodeAddress}] is already exist`)

        var res = await this.gomlapi.getNode(this.nodeAddress);
        this.config.nodes.push({
          alias: this.nodeAlias,
          id: res.data.node_id,
          address: res.data.address,
          hostname: res.data.hostname,
        });

        this.nodeAlias = "";
        this.nodeAddress = "";
        this.initAlert();
      } catch (err) {
        this.alert = true;
        this.errorMessage = err;
        console.error(err)
      }
    },
    deleteNode: function () {
      for (var node of this.selectedList) {
        // remove는 삭제하지 않은 멤버를 반환
        this.config.nodes = this._.remove(this.config.nodes, n => {
          return n.alias !== node.alias
        });
      }
      this.selectedList = [];
    },
    initAlert: function () {
      this.alert = false;
      this.errorMessage = "";
    },
    showAlert: function (message) {
      this.alert = true;
      this.errorMessage = message;
    }
  },
};
</script>
