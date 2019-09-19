<template>
  <v-container grid-list-md>
    <!-- 설정 Input Forms -->
    <v-layout row wrap>
      <v-flex xs12>
        <v-card class="mx-auto">
          <gm-config-card
            ref="configCard"
            v-if="config"
            :config="config"
          ></gm-config-card>
          <v-divider></v-divider>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="primary" @click="update">Update</v-btn>
          </v-card-actions>
        </v-card>
      </v-flex>
    </v-layout>

    <!-- Etc -->
    <v-snackbar
      v-model="snackBar"
      :multi-line="true"
      :timeout="2000"
      :top="true"
      >{{ snackMessage }}</v-snackbar
    >

    <!-- System 상태가 IDLE이 아니면 보여지는 다이얼로그 -->
    <v-dialog
      v-model="disable"
      hide-overlay
      persistent
      width="1200"
      height="600"
    >
      <v-card color="primary" dark>
        <v-layout row align-center>
          <v-flex mx-2>
            <v-card-text>
              Please stand by
              <v-progress-linear
                indeterminate
                color="white"
                class="mb-0"
              ></v-progress-linear>
            </v-card-text>
          </v-flex>
          <v-flex px-2 shrink>
            <v-btn tile large text dark @click="closeLogDialog">Close</v-btn>
          </v-flex>
        </v-layout>
        <pre class="language-js logs">{{ logs }}</pre>
      </v-card>
    </v-dialog>
  </v-container>
</template>
<script>

import ConfigCard from './ConfigCard.vue';

export default {
  components: {
    "gm-config-card": ConfigCard,
  },
  data: function () {
    return {
      config: null,
      state: "UNSET",
      logs: "",
      dialog: false,
      disable: false,
      manualHideDialog: false,
      snackBar: false,
      snackMessage: "",
    };
  },
  created: async function () {
    await this.fetchSystem();
  },
  beforeDestroy() {
    clearTimeout(this.timer);
  },
  methods: {
    /**
     * 로그 다이얼로그를 닫습니다.
     */
    closeLogDialog: function() {
      this.manualHideDialog = true;
    },
    /**
     * fetchSystem는 1초마다 System 상태를 조회합니다.
     *
     * IDLE이 아닌 경우 편집할 수 없음 (Modal Dialog 출력하여 로그 표시)
     */
    fetchSystem: async function () {
      try {
        let res = await this.gomlapi.getSystem();

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
      tempImages = this._.map(tempImages, n => n.trim());
      tempImages = this._.filter(tempImages, n => this._.isEmpty(n) == false);

      var request = {
        config: {
          node_id: this.config.nodeId,
          nodes: this.config.nodes,
          images: tempImages
        }
      };
      await this.gomlapi.setConfig(request);

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

          if (this._.isNil(log["id"]) === false) {
            let values = this._.values(logs);
            let value = this._.find(values, function (value) {
              return value.id === log["id"];
            });
            if (this._.isNil(value) === true) {
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

      let valueLogs = this._.values(logs);
      let lineList = this._.map(valueLogs, l => {
        if (this._.isNil(l.id) !== true) {
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
};
</script>

<style>
.logs {
  font-size: 0.5em;
  box-sizing: border-box;
  display: block;
  /* width: 700px; */
  height: 700px;
}
</style>
