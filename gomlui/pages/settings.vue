<template>
  <div>
    <v-app-bar
      dense
      app
      tabs
      class="elevation-2"
    >
      <v-toolbar-title>Settings</v-toolbar-title>

      <v-spacer></v-spacer>
    </v-app-bar>
    <v-card-text>
      <v-form
        ref="form"
        v-model="valid"
        lazy-validation
      >
        <v-card>
          <v-container>
            <v-row
              no-gutters
              class="d-inline-flex align-center"
              fill-height
            >
              <!-- <v-col cols=3> -->
                <v-text-field
                  class="mx-1"
                  v-model="nodeAlias"
                  label="Node Alias"
                  required
                ></v-text-field>
              <!-- </v-col> -->
              <!-- <v-col cols=3> -->
                <v-text-field
                  class="mx-1"
                  v-model="nodeAddress"
                  label="Node Address(IP)"
                  required
                ></v-text-field>
              <!-- </v-col>
              <v-col cols=2> -->
                <v-btn
                  class="mx-1"
                  color="primary"
                  tile
                  :disabled="disabledAddButton"
                  @click="addNode"
                >Add Node</v-btn>
              <!-- </v-col>
              <v-col cols=2> -->
                <v-btn
                  class="mx-1"
                  color="warning"
                  tile
                  :disabled="disabledDeleteButton"
                  @click="deleteNode"
                >Delete Node</v-btn>
              <!-- </v-col>
              <v-spacer></v-spacer> -->
            </v-row>
            <no-ssr>
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
            </no-ssr>
            <v-row>
              <v-col>
                <v-textarea
                  v-model="config.images"
                  auto-grow
                  label="Docker Images"
                  rows="5"
                ></v-textarea>
                <ul class="caption">
                  <li>
                    <small class="caption">Docker Hub Official Image:
                      docker.io/library/{name}:{tag}</small>
                  </li>
                  <li>
                    <small class="caption">Docker Hub User Image:
                      docker.io/{account}/{name}:{tag}</small>
                  </li>
                  <li>
                    <small class="caption">Private Registry Image:
                      {address}:{port}/{name}:{tag}</small>
                  </li>
                </ul>
              </v-col>
            </v-row>
            <v-row>
              <v-col>
                <v-alert
                  :value="alert"
                  type="error"
                  transition="scale-transition"
                >{{ errorMessage }}</v-alert>
              </v-col>
            </v-row>
            <v-row>
              <v-spacer></v-spacer>
              <v-col
                cols="2"
                class="py-0"
              >
                <v-btn
                  block
                  color="primary"
                  @click="update"
                >Update</v-btn>
              </v-col>
            </v-row>
          </v-container>
        </v-card>
      </v-form>
    </v-card-text>

    <!-- Etc -->
    <v-snackbar
      v-model="snackBar"
      :multi-line="true"
      :timeout="2000"
      :top="true"
    >{{ snackMessage }}</v-snackbar>

    <!-- System 상태가 IDLE이 아니면 보여지는 다이얼로그 -->
    <v-dialog
      v-model="disable"
      persistent
      width="1150"
      height="600"
    >
      <!-- <v-container> -->
      <v-card
        color="primary"
        dark
      >
        <v-card-text>
          <v-row align-center>
            <v-col cols="11">
              Please stand by
              <v-progress-linear
                indeterminate
                color="white"
                class="mb-0"
              ></v-progress-linear>
            </v-col>
            <v-col cols="1">
              <v-btn
                tile
                large
                text
                dark
                @click="closeLogDialog"
              >Close</v-btn>
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>
      <v-card>
        <v-card-text class="pa-0 ma-0">
          <v-row class="pa-0 ma-0">
            <v-col cols="12">
              <pre class="language-js logs">{{ logs }}</pre>
            </v-col>
          </v-row>
        </v-card-text>
      </v-card>
      <!-- </v-container> -->
    </v-dialog>
  </div>
</template>

<script>
export default {
  layout: 'dashboard',
  middleware: 'authenticated',
  data: () => ({
    valid: true,
    alert: false,
    errorMessage: "",
    config: { nodes: [] },
    state: "UNSET",
    logs: "",
    dialog: false,
    disable: false,
    manualHideDialog: false,
    snackBar: false,
    snackMessage: "",
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
          width: "100px",
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
  }),
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
  created: async function () {
    if (_.isNil(this.$gomlapi) === true) return;

    await this.fetchSystem();
  },
  beforeDestroy() {
    clearTimeout(this.timer);
  },
  methods: {
    /**
     * 로그 다이얼로그를 닫습니다.
     */
    closeLogDialog: function () {
      this.manualHideDialog = true;
    },
    addNode: async function () {
      try {
        if (_.findIndex(this.config.nodes, n => n.alias === this.nodeAlias) !== -1)
          return this.showAlert(`Alias[${this.nodeAlias}] is already exist`)
        if (_.findIndex(this.config.nodes, n => n.address === this.nodeAddress) !== -1)
          return this.showAlert(`Address[${this.nodeAddress}] is already exist`)

        var res = await this.$gomlapi.getNode(this.nodeAddress);
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
        this.config.nodes = _.remove(this.config.nodes, n => {
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
    },
    /**
     * 로그 다이얼로그를 닫습니다.
     */
    closeLogDialog: function () {
      this.manualHideDialog = true;
    },
    /**
     * fetchSystem는 1초마다 System 상태를 조회합니다.
     *
     * IDLE이 아닌 경우 편집할 수 없음 (Modal Dialog 출력하여 로그 표시)
     */
    fetchSystem: async function () {
      try {
        if (_.isNil(this.$gomlapi) === true) return;

        let res = await this.$gomlapi.getSystem();

        let system = res.data.system;

        // 상태 확인
        if (system.state !== "IDLE" && system.logs.length > 0 && this.manualHideDialog === false) {
          // IDLE이 아니고 Logs가 있다면 입력을 Disable하고 Modal Dialog를 보여준다.
          this.disable = true;
          this.logs = this.convertToDockerLog(system.logs);
        } else {
          this.disable = false;
        }

        // 첫 상태 확인 또는 IDLE이 아닌 상태에서 IDLE이 된 경우 config 값을 갱신한다.
        if (
          this.state === "UNSET" ||
          (this.state !== "IDLE" && system.state === "IDLE")
        ) {
          let nodeId = system.config.node_id;
          let nodes = [];
          for (var node of system.config.nodes) {
            nodes.push({
              alias: node.alias,
              id: node.id,
              address: node.address,
              hostname: node.hostname,
            });
          }

          let images = system.config.images.join("\n");

          this.config = {
            nodeId: nodeId,
            nodes: nodes,
            images: images
          };
        }
        this.state = system.state;
      } catch (err) {
        console.error(err)
      } finally {
        this.timer = setTimeout(this.fetchSystem, 1000);
      }
    },
    /**
     * update는 변경된 설정값을 업데이트 합니다.
     */
    update: async function () {
      // 줄바꿈으로 설정한 images를 배열로 변경합니다.
      let tempImages = this.config.images.split("\n");
      tempImages = _.map(tempImages, n => n.trim());
      tempImages = _.filter(tempImages, n => _.isEmpty(n) == false);

      var request = {
        config: {
          node_id: this.config.nodeId,
          nodes: this.config.nodes,
          images: tempImages
        }
      };
      await this.$gomlapi.setConfig(request);

      this.manualHideDialog = false;
      this.snackBar = true;
      this.snackMessage = "Successfully update.";
    },
    /**
     * convertToDockerLog는 Docker Log를 보기 좋게 수정합니다.
     *
     * System의 Logs의 경우 Docker Log가 대부분이며, Docker Log는 ID 기반으로 갱신합니다.
     */
    convertToDockerLog: function (rawLogs) {
      let lines = rawLogs;
      let logIndex = 0;
      let logs = {};

      for (var line of lines) {
        if (line.trim() === "") {
          let log = { status: "" };
          log.index = logIndex++;
          logs[log.index] = log;
          continue;
        }

        try {
          let log = JSON.parse(line);

          if (_.isNil(log["id"]) === false) {
            let values = _.values(logs);
            let value = _.find(values, function (value) {
              return value.id === log["id"];
            });
            if (_.isNil(value) === true) {
              log.index = logIndex++;
            } else {
              log.index = value.index;
            }

            logs[log.index] = log;
          } else {
            log.index = logIndex++;
            logs[log.index] = log;
          }
        } catch {
          let log = { status: line };
          log.index = logIndex++;
          logs[log.index] = log;
        }
      }

      let valueLogs = _.values(logs);
      let lineList = _.map(valueLogs, l => {
        if (_.isNil(l.id) !== true) {
          if (l.status === "Extracting" || l.status === "Downloading") {
            return `${l.id}: ${l.status}: ${l.progress} [${l.progressDetail.current}/${l.progressDetail.total}]`;
          } else {
            return `${l.id}: ${l.status}`;
          }
        } else {
          return l.status;
        }
      });

      return lineList.join("\n");
    },
  }
}
</script>
<style>
.logs {
  box-sizing: border-box;
  display: block;
  /* width: 700px; */
  height: 700px;
}
</style>
