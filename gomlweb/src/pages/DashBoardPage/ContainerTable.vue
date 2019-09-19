
<template>
  <v-data-table
    v-model="selectedList"
    :headers="headers"
    :items="containers"
    single-expand
    show-select
    show-expand
    class="elevation-1"
  >
    <template v-slot:item.name="{ item }">
      <div class="text-truncate">{{ item.names[0].substring(1) }}</div>
    </template>
    <template v-slot:item.image="{ item }">
      <div class="text-truncate container-image-cell">
        {{ getImageName(item.image) }}
      </div>
    </template>
    <template v-slot:expanded-item="props">
      <td :colspan="headers.length + 1" class="ma-0 pa-0">
        <v-card flat>
          <v-item-group multiple>
            <v-container grid-list-md>
              <v-layout wrap>
                <v-flex xs12 md8>
                  <v-item>
                    <v-card flat dense>
                      <v-list subheader two-line dense>
                        <v-subheader>Info</v-subheader>

                        <v-list-item>
                          <v-list-item-content>
                            <v-list-item-title>Name</v-list-item-title>
                            <v-list-item-subtitle>{{
                              props.item.names[0].substring(1)
                            }}</v-list-item-subtitle>
                          </v-list-item-content>
                        </v-list-item>
                        <v-list-item>
                          <v-list-item-content>
                            <v-list-item-title>Created</v-list-item-title>
                            <v-list-item-subtitle>{{
                              new Date(
                                Number(props.item.created) * 1000
                              ).toLocaleString()
                            }}</v-list-item-subtitle>
                          </v-list-item-content>
                        </v-list-item>
                      </v-list>
                      <v-divider></v-divider>
                      <v-list subheader two-line dense>
                        <v-subheader>Volumes</v-subheader>
                        <v-list-item
                          v-ripple="{ class: 'primary--text' }"
                          v-for="mount in props.item.mounts"
                          :key="mount.destination"
                        >
                          <v-list-item-content>
                            <v-list-item-title v-if="mount.source === ''"
                              >&#128449;&#8194;-- unset --</v-list-item-title
                            >
                            <v-list-item-title v-if="mount.source !== ''"
                              >&#128449;&#8194;{{
                                mount.source
                              }}</v-list-item-title
                            >

                            <v-list-item-subtitle
                              >&#8194;&#8194;&#10132;&#8194;{{
                                mount.destination
                              }}</v-list-item-subtitle
                            >
                          </v-list-item-content>
                          <v-list-item-content
                            v-if="props.item.mounts.length === 0"
                            >empty</v-list-item-content
                          >
                        </v-list-item>
                        <v-list-item
                          v-ripple="{ class: 'primary--text' }"
                          v-if="props.item.mounts.length === 0"
                        >
                          <v-list-item-content>
                            <v-list-item-title>-- empty --</v-list-item-title>
                          </v-list-item-content>
                        </v-list-item>
                      </v-list>
                    </v-card>
                  </v-item>
                </v-flex>
                <v-flex xs12 md4>
                  <v-item>
                    <v-card flat>
                      <v-list subheader dense>
                        <v-subheader>Ports</v-subheader>
                        <v-list-item
                          v-for="port in props.item.ports"
                          v-ripple="{ class: 'primary--text' }"
                          :key="port.private_port"
                          @click="
                            actionPort(
                              props.item.address,
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
                    </v-card>
                  </v-item>
                </v-flex>
              </v-layout>
            </v-container>
          </v-item-group>
        </v-card>
      </td>

      <input type="hidden" ref="copyTextBox" />
    </template>
  </v-data-table>
</template>

<script>
import { getShortImageName } from "@/utils/utils-function";

export default {
  name: "ContainerTable",
  inheritAttrs: false,
  props: {
    selected: {
      type: [Array],
      default: () => []
    },
    containers: {
      type: [Array],
      default: () => []
    },
    snackBar: {
      type: Boolean,
      default: () => false
    },
    snackMessage: {
      type: String,
      default: () => ""
    }
  },
  data: function () {
    return {
      pagination: {
        sortBy: "hostname",
        rowsPerPage: 10
      },
      headers: [
        { text: "Node", align: "left", value: "nodeAlias", width: "10%" },
        { text: "Name", align: "left", value: "name", width: "30%" },
        { text: "Image", align: "left", value: "image" },
        { text: "Status", value: "status", width: "20%" },
        { text: "", value: "data-table-expand", sortable: false }
      ],
    };
  },
  computed: {
    selectedList: {
      get: function () {
        return this.selected;
      },
      set: function (selected) {
        this.$emit('update:selected', selected)
      }
    }
  },
  methods: {
    getPortText: function (port) {
      if (port.public_port === 0)
        return `:${port.private_port}`;
      else {
        return `${port.public_port}:${port.private_port}`;
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
        this.$emit('update:snackBar', true)
        this.$emit('update:snackMessage', "Copy the SSH Address")
      } catch (err) {
        console.error(err);
      }

      /* unselect the range */
      copyTextBox.setAttribute("type", "hidden");
      window.getSelection().removeAllRanges();
    },
  }
};
</script>

<style>
.container-image-cell {
  box-sizing: border-box;
  display: block;
  width: 250px;
}
</style>
