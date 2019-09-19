<template>
  <div>
    <v-app-bar dense app tabs flat>
      <v-toolbar-title>Create Containers</v-toolbar-title>

      <v-spacer></v-spacer>
    </v-app-bar>
    <v-card-text>
      <v-form ref="form" v-model="valid" lazy-validation>
        <v-card flat class="pa-4">
          <v-row>
            <v-col cols="7" class="py-0">
              <v-text-field
                v-model="containerNewName"
                label="Container Name*"
                :disabled="autoname"
              ></v-text-field>
            </v-col>
            <v-col cols="2" class="py-0">
              <v-switch v-model="autoname" class="px-4" label="Auto"></v-switch>
            </v-col>
          </v-row>

          <v-row>
            <v-col cols="4" class="py-0">
              <v-select
                attach
                v-model="image"
                :rules="[v => !!v || 'Image is required']"
                :items="images"
                label="Images*"
                @change="selectedImage"
                required
              >
                <template slot="selection" slot-scope="data">
                  {{ shortImageName(data.item) }}
                </template>
                <template slot="item" slot-scope="data">
                  {{ shortImageName(data.item) }}
                </template>
              </v-select>
            </v-col>
            <v-col cols="4" class="py-0">
              <v-text-field v-model="tag" label="Tag"></v-text-field>
            </v-col>
            <v-col cols="4" class="py-0">
              <v-select
                attach
                v-model="nodeID"
                :items="nodes"
                item-value="id"
                item-text="hostname"
                :rules="[v => !!v || 'Node is required']"
                label="Node*"
                required
              ></v-select>
            </v-col>
          </v-row>

          <v-subheader>Ports configuration</v-subheader>
          <v-divider></v-divider>
          <v-row>
            <v-col cols="12" class="py-0">
              <v-switch
                v-model="autoport"
                class="px-4"
                label="Publish all exposed ports [random]"
                :disabled="image === ''"
              ></v-switch>
            </v-col>
          </v-row>

          <v-row v-if="!autoport">
            <v-col cols="12" class="py-0">
              <v-row
                v-for="port in networkPorts"
                :key="port.container"
                class="align-center"
              >
                <v-col cols="4" class="py-0"
                  ><v-text-field
                    v-model="port.host"
                    :disabled="autoport"
                    label="Host"
                  ></v-text-field
                ></v-col>
                <v-col cols="1" class="py-0"><span> ➡ </span></v-col>
                <v-col cols="4" class="py-0"
                  ><v-text-field
                    v-model="port.container"
                    label="Container"
                    disabled
                  ></v-text-field
                ></v-col>
              </v-row>
            </v-col>
          </v-row>

          <v-subheader>Volumes</v-subheader>
          <v-divider></v-divider>

          <v-row>
            <v-col cols="12" class="py-0">
              <v-btn
                class="ma-2"
                x-small
                outlined
                @click="addVolume"
                :disabled="image === ''"
                >+ map additional volume</v-btn
              >
            </v-col>
          </v-row>

          <v-row v-for="(volume, i) in volumes" :key="'volume_' + i">
            <v-col cols="12" class="py-0">
              <v-row class="align-center" v-if="!volume.removed">
                <v-col cols="4" class="py-0"
                  ><v-text-field
                    v-model="volume.host"
                    label="Host"
                  ></v-text-field
                ></v-col>
                <v-col cols="1" class="py-0"><span> ➡ </span></v-col>
                <v-col cols="4" class="py-0">
                  <v-combobox
                    v-model="volume.container"
                    label="Container"
                    attach
                    :items="volumeItems"
                  ></v-combobox
                ></v-col>
                <v-col cols="1" class="py-0"
                  ><v-btn icon text @click="deleteVolume(volume)"
                    ><v-icon>delete</v-icon></v-btn
                  ></v-col
                >
              </v-row>
            </v-col>
          </v-row>

          <v-subheader>Env</v-subheader>
          <v-divider></v-divider>

          <v-row>
            <v-col cols="12" class="py-0">
              <v-btn
                class="ma-2"
                x-small
                outlined
                @click="addEnvironment"
                :disabled="image === ''"
                >+ add environment variable</v-btn
              >
            </v-col>
          </v-row>

          <v-row v-for="(env, i) in envs" :key="'env_' + i">
            <v-col cols="12" class="py-0">
              <v-row class="align-center" v-if="!env.removed">
                <v-col cols="4" class="py-0"
                  ><v-combobox
                    v-model="env.name"
                    label="Name"
                    attach
                    :items="envItems"
                  ></v-combobox
                ></v-col>
                <v-col cols="1" class="py-0"><span> = </span></v-col>
                <v-col cols="4" class="py-0"
                  ><v-text-field
                    v-model="env.value"
                    label="Value"
                  ></v-text-field
                ></v-col>
                <v-col cols="1" class="py-0"
                  ><v-btn icon text @click="deleteEnvironment(env)"
                    ><v-icon>delete</v-icon></v-btn
                  ></v-col
                >
              </v-row>
            </v-col>
          </v-row>

          <v-row class="my-4">
            <v-col cols="2" class="py-0">
              <v-btn
                class="ma-2"
                block
                outlined
                color="primary"
                @click="cancel"
                >Cancel</v-btn
              >
            </v-col>
            <v-spacer></v-spacer>
            <v-col cols="2" class="py-0">
              <v-btn
                class="ma-2"
                block
                color="primary"
                @click="createContainer"
                :disabled="image === '' || nodeID === ''"
                >Create</v-btn
              >
            </v-col>
          </v-row>

          <v-row>
            <v-col>
              <v-alert
                :value="alert"
                type="error"
                transition="scale-transition"
                >{{ errorMessage }}</v-alert
              >
            </v-col>
          </v-row>
        </v-card>
      </v-form>
    </v-card-text>

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
  </div>
</template>

<script>
const jwt = process.client ? require('jsonwebtoken') : undefined

export default {
  layout: 'dashboard',
  middleware: 'authenticated',
  data: () => ({
    loadingDialog: false,
    valid: true,
    username: '',
    tag: '',
    alert: false,
    errorMessage: "",
    containerName: '',
    autoname: true,
    autoport: true,
    snackBar: false,
    nodeID: '',
    nodes: [],
    image: "",
    images: [],
    imageInspects: [],
    snackMessage: "",
    networkPorts: [],
    volumes: [],
    volumeItems: [],
    envs: [],
    envItems: [],
  }),
  computed: {
    containerNewName: {
      get: function () {
        if (this.autoname === false)
          return this.containerName;

        if (_.isNil(this.image) === true || this.image === "") {
          this.containerName = "";
          return this.containerName;
        }

        var split = this.image.lastIndexOf("/");
        var tailPart = this.image.substring(split + 1);

        var colon = tailPart.indexOf(":");
        var imageName = tailPart.substring(0, colon);

        if (this.tag !== '')
          this.containerName = `${imageName}_${this.username}_${this.tag}`;
        else
          this.containerName = `${imageName}_${this.username}`;

        return this.containerName;
      },
      set: function (value) {
        this.containerName = value;
      }
    },
    /**
     * 입력받은 Env List에서 무효한 값은 제외합니다.
     */
    validEvns: {
      get: function () {
        var list = [];
        for (var env of this.envs) {
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
        for (var volume of this.volumes) {
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

        if (this.autoport === false) {
          for (var port of this.networkPorts) {
            if (port.host.trim() === "") continue;

            var hostPort = parseInt(port.host.trim());
            var containerPort = port.container;
            map[hostPort] = containerPort;
          }
        } else {
          var index = 0;
          for (var port of this.networkPorts) {
            var containerPort = port.container;
            map[index++] = containerPort;
          }
        }
        return map;
      }
    }
  },
  created: async function () {
    try {
      if (_.isNil(this.$gomlapi) === true || _.isNil(jwt) === true) return;

      var decoded = jwt.decode(this.$store.state.auth);
      this.username = decoded.Name;

      // Docker Images 조회
      var res = await this.$gomlapi.getSystem();
      this.images = res.data.system.config.images;
      this.imageInspects = res.data.system.config.image_inspects;

      // Node 조회
      res = await this.$gomlapi.getNodes();
      this.nodes = res.data.nodes;
    } catch (err) {
      console.error(err);
    }

    // await this.fetchNodes();
  },
  methods: {
    cancel() {
      var container = this.$route.query.container;
      if (_.isNil(container) === true) return;

      this.$router.push('/containers?container=' + container);
    },
    /**
     * 이미지 이름을 짧게 변환합니다.
     */
    shortImageName: function (imageFullName) {
      var index = imageFullName.lastIndexOf("/");
      if (index === -1)
        return imageFullName;

      return imageFullName.substring(index + 1);
    },
    async createContainer() {
      try {
        if (this.$refs.form.validate() === false) {
          this.valid = false;
          return;
        }

        var request = {
          node_id: this.nodeID,
          image_name: this.image,
          container_name: this.containerName,
          envs: this.validEvns,
          volumes: this.validVolumes,
          auto_port: this.autoport,
          ports: this.validPorts,
          labels: {
            "goml.system": "true",
            "goml.owner": this.username
          }
        };
        console.log(request)
        this.loadingDialog = true;

        var res = await this.$gomlapi.createContainer(request);

        this.$router.push('/containers?container=' + res.data.id);
      } catch (err) {
        console.error(err);
        if (_.isNil(err.response.data.message) === true)
          this.errorMessage = err.response.data;
        else
          this.errorMessage = err.response.data.message;

        this.alert = true;
      } finally {
        this.loadingDialog = false;
      }
    },
    addVolume() {
      this.volumes.push({ host: "", container: "", removed: false })
    },
    deleteVolume(volume) {
      volume.removed = true;
    },
    addEnvironment() {
      this.envs.push({ name: "", value: "", removed: false })
    },
    deleteEnvironment(env) {
      env.removed = true;
    },
    selectedImage(selected) {
      this.networkPorts = [];
      this.envItems = [];

      var imageInspect = _.find(this.imageInspects, i =>
        _.some(i.repo_tags, t => selected.indexOf(t) !== -1)
      );

      if (_.isNil(imageInspect) === true) return;

      // initialize ports
      for (var port of imageInspect.container_config.exposed_ports) {
        this.networkPorts.push({
          host: "",
          container: port
        });
      }

      // initialize envs
      for (var env of imageInspect.container_config.env) {
        var pos = env.indexOf("=");
        this.envItems.push(env.substring(0, pos));
      }

      // initialize volumes
      for (var containerPath of imageInspect.container_config.volumes) {
        this.volumeItems.push(containerPath);
      }

    },
  },
  asyncData(context) {
    return { project: 'nuxt' }
  }
}
</script>

<style lang="scss" scoped>
</style>

