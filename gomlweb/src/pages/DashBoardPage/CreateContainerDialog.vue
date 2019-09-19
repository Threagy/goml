<template>
  <!-- Dialogs -->
  <v-dialog
    v-model="visible"
    fullscreen
    hide-overlay
    persistent
    transition="dialog-bottom-transition"
  >
    <v-card>
      <v-toolbar dark color="primary">
        <v-btn icon dark @click="closeDialog">
          <v-icon>close</v-icon>
        </v-btn>
        <v-toolbar-title>Create Container</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-toolbar-items>
          <v-btn v-if="step > 1" text dark @click="step -= 1">Previous</v-btn>
          <v-btn v-if="step != completeStep" text dark @click="step += 1"
            >Next</v-btn
          >
          <v-btn text dark @click="createContainer">Finish</v-btn>
        </v-toolbar-items>
      </v-toolbar>

      <v-card-text>
        <v-form ref="form" v-model="valid" lazy-validation>
          <v-card flat>
            <v-container>
              <v-alert
                :value="alert"
                type="error"
                transition="scale-transition"
                >{{ errorMessage }}</v-alert
              >
              <v-stepper non-linear v-model="step">
                <v-stepper-header>
                  <v-stepper-step editable :complete="step > 1" step="1"
                    >General Information</v-stepper-step
                  >
                  <v-divider></v-divider>
                  <v-stepper-step editable :complete="step > 2" step="2"
                    >Volumes</v-stepper-step
                  >
                  <v-divider></v-divider>
                  <v-stepper-step editable :complete="step > 3" step="3"
                    >Environments</v-stepper-step
                  >
                  <v-divider></v-divider>
                  <v-stepper-step editable step="4">Networking</v-stepper-step>
                </v-stepper-header>

                <v-stepper-items>
                  <v-stepper-content step="1">
                    <v-card flat class="mb-5">
                      <v-layout wrap>
                        <v-flex xs4 px-2>
                          <v-text-field
                            xs4
                            v-model="containerNewName"
                            label="Container Name*"
                            disabled
                          ></v-text-field>
                        </v-flex>
                        <v-flex xs4 px-2>
                          <v-select
                            v-model="image"
                            :rules="rules.imageRule"
                            :items="images"
                            label="Images"
                            @change="selectedImage"
                            required
                          >
                            <template slot="selection" slot-scope="data">
                              {{ shortImageName(data.item)}}
                            </template>
                            <template slot="item" slot-scope="data">
                              {{ shortImageName(data.item)}}
                            </template>
                          </v-select>
                        </v-flex>
                        <v-flex xs4 px-2>
                          <v-text-field
                            v-model="userName"
                            :rules="rules.userNameRules"
                            label="User Name"
                            required
                          ></v-text-field>
                        </v-flex>
                        <v-flex xs4 px-2>
                          <v-select
                            v-model="node"
                            :items="nodes"
                            item-text="hostname"
                            :rules="[v => !!v || 'Node is required']"
                            label="Node"
                            required
                            @change="selectedNode"
                          ></v-select>
                        </v-flex>
                        <!-- <v-flex xs4 px-2>
                          <v-select
                            :items="this.gpuList"
                            v-model="selectedGPUs"
                            item-text="index"
                            :rules="[v => !!v || 'Node is required']"
                            label="Select GPU"
                            multiple
                            open-on-clear
                            required
                          ></v-select>
                        </v-flex> -->
                      </v-layout>
                    </v-card>
                  </v-stepper-content>

                  <v-stepper-content step="2">
                    <v-layout ma-2>
                      <v-btn tile class="mx-2" @click="addVolume"
                        >Add Volume</v-btn
                      >
                    </v-layout>
                    <v-card flat class="mb-5">
                      <v-layout xs12>
                        <v-data-table
                          dense
                          :page="volume.page"
                          :items-per-page="volume.itemsPerPage"
                          :headers="volume.headers"
                          :items="volume.volumes"
                        >
                          <template v-slot:item.host="props">
                            <v-edit-dialog
                              :return-value.sync="props.item.host"
                              large
                              persistent
                            >
                              <v-layout row wrap align-center class="ma-0">
                                <v-btn class="mx-2" x-small icon ripple>
                                  <v-icon>create</v-icon>
                                </v-btn>
                                <div class="text-truncate key-cell">
                                  {{ props.item.host }}
                                </div>
                              </v-layout>
                              <template v-slot:input>
                                <v-text-field
                                  v-model="props.item.host"
                                  label="Update Host Path"
                                  single-line
                                  @focus="$event.target.select()"
                                ></v-text-field>
                              </template>
                            </v-edit-dialog>
                          </template>

                          <template v-slot:item.container="props">
                            <v-edit-dialog
                              :return-value.sync="props.item.container"
                              large
                              persistent
                            >
                              <v-layout row wrap align-center class="ma-0">
                                <v-btn class="mx-2" x-small icon ripple>
                                  <v-icon>create</v-icon>
                                </v-btn>
                                <div class="text-truncate value-cell">
                                  {{ props.item.container }}
                                </div>
                              </v-layout>
                              <template v-slot:input>
                                <v-text-field
                                  v-model="props.item.container"
                                  label="Update Container Path"
                                  single-line
                                  autofocus
                                  @focus="$event.target.select()"
                                ></v-text-field>
                              </template>
                            </v-edit-dialog>
                          </template>
                        </v-data-table>
                      </v-layout>
                    </v-card>
                  </v-stepper-content>

                  <v-stepper-content step="3">
                    <v-layout ma-2>
                      <v-btn tile class="mx-2" @click="addEnv"
                        >Add Environment</v-btn
                      >
                    </v-layout>
                    <v-card flat class="mb-5">
                      <v-layout xs12>
                        <v-data-table
                          dense
                          :page="env.page"
                          :items-per-page="env.itemsPerPage"
                          :headers="env.headers"
                          :items="env.envs"
                        >
                          <template v-slot:item.name="props">
                            <v-edit-dialog
                              :return-value.sync="props.item.name"
                              large
                              persistent
                            >
                              <v-layout row wrap align-center class="ma-0">
                                <v-btn class="mx-2" x-small icon ripple>
                                  <v-icon>create</v-icon>
                                </v-btn>
                                <div>
                                  <div class="text-truncate key-cell">
                                    {{ props.item.name }}
                                  </div>
                                </div>
                              </v-layout>
                              <template v-slot:input>
                                <v-text-field
                                  v-model="props.item.name"
                                  label="Update Name"
                                  single-line
                                  autofocus
                                  @focus="$event.target.select()"
                                ></v-text-field>
                              </template>
                            </v-edit-dialog>
                          </template>

                          <template v-slot:item.value="props">
                            <v-edit-dialog
                              :return-value.sync="props.item.value"
                              large
                              persistent
                            >
                              <v-layout row wrap align-center class="ma-0">
                                <v-btn class="mx-2" x-small icon ripple>
                                  <v-icon>create</v-icon>
                                </v-btn>
                                <div class="text-truncate value-cell">
                                  {{ props.item.value }}
                                </div>
                              </v-layout>
                              <template v-slot:input>
                                <v-text-field
                                  v-model="props.item.value"
                                  label="Update Value"
                                  single-line
                                  autofocus
                                  @focus="$event.target.select()"
                                ></v-text-field>
                              </template>
                            </v-edit-dialog>
                          </template>
                        </v-data-table>
                      </v-layout>
                    </v-card>
                  </v-stepper-content>

                  <v-stepper-content step="4">
                    <v-card flat class="mb-5">
                      <v-layout xs12>
                        <v-data-table
                          dense
                          :headers="network.headers"
                          :items="network.ports"
                        >
                          <template v-slot:item.host="props">
                            <v-edit-dialog
                              :return-value.sync="props.item.host"
                              large
                              persistent
                            >
                              <v-layout row wrap align-center class="ma-0">
                                <v-btn class="mx-2" x-small icon ripple>
                                  <v-icon>create</v-icon>
                                </v-btn>
                                <div>
                                  <div class="text-truncate key-cell">
                                    {{ props.item.host }}
                                  </div>
                                </div>
                              </v-layout>
                              <template v-slot:input>
                                <v-text-field
                                  v-model="props.item.host"
                                  label="Update Host Port"
                                  single-line
                                  autofocus
                                  @focus="$event.target.select()"
                                ></v-text-field>
                              </template>
                            </v-edit-dialog>
                          </template>
                        </v-data-table>
                      </v-layout>
                    </v-card>
                  </v-stepper-content>
                </v-stepper-items>
              </v-stepper>
            </v-container>
          </v-card>
        </v-form>
      </v-card-text>
    </v-card>

    <!-- Loading -->
    <v-dialog v-model="loadingDialog" hide-overlay persistent width="300">
      <v-card color="primary" dark>
        <v-card-text>
          Please stand by
          <v-progress-linear
            indeterminate
            color="white"
            class="mb-0"
          ></v-progress-linear>
        </v-card-text>
      </v-card>
    </v-dialog>
  </v-dialog>
</template>

<script>
import { getShortImageName } from "@/utils/utils-function";

export default {
  name: "CreateContainerDialog",
  inheritAttrs: false,
  props: {
    visible: {
      type: Boolean,
      default: false
    },
    nodes: {
      type: [Array],
      default: () => []
    }
  },
  data: function () {
    return {
      loadingDialog: false,
      dialog: false,
      images: [],
      image: "",
      imageInspects: [],
      userName: "",
      node: null,
      nodeID: "",
      gpuList: [],
      selectedGPUs: [],
      errorMessage: "",
      alert: false,
      valid: true,
      rules: {
        imageRule: [v => !!v || "Image is required"],
        userNameRules: [v => !!v || "User Name is required"]
      },
      step: 1,
      completeStep: 4,
      volume: {
        page: 1,
        itemsPerPage: -1,
        headers: [
          {
            text: "Host Path",
            value: "host",
            align: "left",
            width: "350px",
            sortable: false
          },
          {
            text: "Container Path",
            value: "container",
            align: "left",
            width: "450px",
            sortable: false
          }
        ],
        volumes: []
      },
      env: {
        page: 1,
        itemsPerPage: -1,
        headers: [
          {
            text: "Name",
            value: "name",
            align: "left",
            width: "350px",
            sortable: false
          },
          {
            text: "Value",
            value: "value",
            align: "left",
            width: "450px",
            sortable: false
          }
        ],
        envs: []
      },
      network: {
        headers: [
          {
            text: "Host Port",
            value: "host",
            align: "left",
            width: "350px",
            sortable: false
          },
          {
            text: "Container Port",
            value: "container",
            align: "left",
            width: "450px",
            sortable: false
          }
        ],
        ports: []
      }
    };
  },
  created: async function () {
    try {
      // Docker Images 조회
      var res = await this.gomlapi.getSystem();
      this.images = res.data.system.config.images;
      this.imageInspects = res.data.system.config.image_inspects;
    } catch (err) {
      console.error(err);
    }
  },
  computed: {
    /**
     * Image와 UserName을 조합하여 컨테이너 이름을 생성합니다.
     */
    containerNewName: {
      get: function () {
        if (this._.isNil(this.image) === true || this.image === "")
          return "";

        var split = this.image.lastIndexOf("/");
        var tailPart = this.image.substring(split + 1);

        var colon = tailPart.indexOf(":");
        var imageName = tailPart.substring(0, colon);

        return `${imageName}_${this.userName}`;
      },
      set: function () { }
    },
    /**
     * 입력받은 Env List에서 무효한 값은 제외합니다.
     */
    validEvns: {
      get: function () {
        var list = [];
        for (var env of this.env.envs) {
          if (env.name.trim() === "") continue;

          list.push(`${env.name}=${env.value}`);
        }
        return list;
      }
    },
    /**
     * 입력받은 Volume List에서 무효한 값은 제외합니다.
     */
    validVolumes: {
      get: function () {
        var list = {};
        for (var volume of this.volume.volumes) {
          if (volume.host.trim() === "" || volume.container === "") continue;

          list[volume.host] = volume.container;
        }
        return list;
      }
    },
    /**
     * 입력받은 Port List에서 무효한 값은 제외합니다.
     */
    validPorts: {
      get: function () {
        var map = {};
        for (var port of this.network.ports) {
          if (port.host.trim() === "") continue;

          var hostPort = parseInt(port.host.trim());
          var containerPort = port.container;
          map[hostPort] = containerPort;
        }
        return map;
      }
    }
  },
  methods: {
    /**
     * 이미지 이름을 짧게 변환합니다.
     */
    shortImageName: function (imageFullName) {
      return getShortImageName(imageFullName)
    },
    /**
     * CreateContainer Dialog의 변수를 초기화 합니다.
     */
    initVariable: function () {
      this.$refs.form.resetValidation();

      this.dialog = false;
      this.image = null;
      this.node = null;
      this.nodeID = "";
      this.userName = "";
      this.selectedGPUs = [];
      this.errorMessage = "";
      this.alert = false;
      this.valid = true;
      this.step = 1;
      this.volume.volumes = [];
      this.volume.page = 1;
      this.env.envs = [];
      this.env.page = 1;
    },
    /**
     * Close Dialog
     */
    closeDialog: function () {
      this.$emit("update:visible", false);
      this.initVariable();
    },
    /**
     * 컨테이너를 생성합니다.
     */
    createContainer: async function () {
      try {
        if (this.$refs.form.validate() === false) {
          this.valid = false;
          return;
        }

        var request = {
          node_id: this.nodeID,
          image_name: this.image,
          container_name: this.containerNewName,
          envs: this.validEvns,
          volumes: this.validVolumes,
          ports: this.validPorts,
          labels: {
            "goml.system": "true",
            "goml.owner": this.userName
          }
        };

        this.loadingDialog = true;

        await this.gomlapi.createContainer(request);

        this.initVariable();
        this.$emit("created");
        this.$emit("update:visible", false);
      } catch (err) {
        console.error(err.response.data);
        if (this._.isNil(err.response.data.message) === true)
          this.errorMessage = err.response.data;
        else
          this.errorMessage = err.response.data.message;

        this.alert = true;
      } finally {
        this.loadingDialog = false;
      }
    },
    /**
     * Env를 추가합니다.
     */
    addEnv() {
      this.env.envs.push({
        name: "",
        value: ""
      });
    },
    /**
     * Volume을 추가합니다.
     */
    addVolume() {
      this.volume.volumes.push({
        host: "",
        container: ""
      });
    },
    /**
     * Image를 선택합니다.
     *
     * Env, Volume, Port의 기본 값을 초기화합니다.
     */
    selectedImage(selected) {

      this.volume.volumes = [];
      this.env.envs = [];
      this.network.ports = [];

      var imageInspect = this._.find(this.imageInspects, i =>
        this._.some(i.repo_tags, t => selected.indexOf(t) !== -1)
      );

      if (this._.isNil(imageInspect) === true) return;

      // initialize volumes
      for (var containerPath of imageInspect.container_config.volumes) {
        this.volume.volumes.push({
          host: "",
          container: containerPath
        });
      }

      // initialize envs
      var tempEnvs = [];
      for (var env of imageInspect.container_config.env) {
        var pos = env.indexOf("=");
        tempEnvs.push({
          name: env.substring(0, pos),
          value: env.substring(pos + 1)
        });
      }
      this.env.envs = this._.sortBy(tempEnvs, e => e.name);

      // initialize ports
      for (var port of imageInspect.container_config.exposed_ports) {
        this.network.ports.push({
          host: "",
          container: port
        });
      }
    },
    /**
     * Node를 선택합니다.
     */
    selectedNode(selected) {
      this.gpuList = [];
      this.selectedGPUs = [];

      var hostname = selected;
      var node = this._.find(this.nodes, m => m.hostname === hostname);

      if (this._.isNil(node) === true) return;

      this.gpuList = this._.map(node.gpu_statuses, g => g.index);
      this.nodeID = node.id;
    }
  }
};
</script>

<style>
.key-cell {
  width: 270px;
}

.value-cell {
  width: 350px;
}
</style>

