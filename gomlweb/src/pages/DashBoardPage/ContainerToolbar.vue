
<template>
  <v-layout row align-center>
    <v-flex shrink>
      <v-btn class="mx-1" color="primary" tile dark @click="createContainer"
        >Create</v-btn
      >
      <v-btn
        class="mx-1"
        color="success"
        tile
        :disabled="disabledStartButton"
        @click="confirmStartContainer"
        >Start</v-btn
      >
      <v-btn
        class="mx-1"
        color="error"
        tile
        :disabled="disabledStopButton"
        @click="confirmStopContainer"
        >Stop</v-btn
      >
      <v-btn
        class="mx-1"
        color="warning"
        tile
        :disabled="disabledDeleteButton"
        @click="confirmDeleteContainer"
        >Delete</v-btn
      >
    </v-flex>
    <v-flex>
      <div id="tag-div">
        <vue-tags-input
          class="tags-input"
          v-model="tag"
          :tags="tags"
          @tags-changed="tagsChanged"
        />
      </div>
    </v-flex>

    <v-flex shrink class="text-center">
      <v-btn icon tile small @click="refreshContainers">
        <v-icon>refresh</v-icon>
      </v-btn>
    </v-flex>

    <!-- Confirm Dialog -->
    <v-dialog v-model="startConfirmDialog" persistent max-width="500">
      <v-card>
        <v-card-title class="headline"
          >Do you want to start the container below?</v-card-title
        >
        <v-card-text>
          <ul>
            <li :key="item.id" v-for="item in selected">
              {{ item.names[0].substring(1) }}
            </li>
          </ul>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="green darken-1" text @click="startConfirmDialog = false"
            >Cancel</v-btn
          >
          <v-btn color="green darken-1" text @click="startContainers"
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
            <li :key="item.id" v-for="item in selected">
              {{ item.names[0].substring(1) }}
            </li>
          </ul>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="green darken-1" text @click="stopConfirmDialog = false"
            >Cancel</v-btn
          >
          <v-btn color="green darken-1" text @click="stopContainers"
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
            <li :key="item.id" v-for="item in selected">
              {{ item.names[0].substring(1) }}
            </li>
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
          <v-btn color="green darken-1" text @click="deleteContainers"
            >Delete</v-btn
          >
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-layout>
</template>

<script>
import VueTagsInput from '@johmun/vue-tags-input';

export default {
  name: "ContainerButtonGroup",
  inheritAttrs: false,
  props: {
    selected: {
      type: Array,
      default: () => [],
    },
    filters: {
      type: Array,
      default: () => [],
    },
  },
  components: {
    VueTagsInput,
  },
  data: function () {
    return {
      tag: '',
      tags: [],
      startConfirmDialog: false,
      deleteConfirmDialog: false,
      stopConfirmDialog: false,
    };
  },
  created: function () {
    this.tags = this.filters;
    this.$emit('filterChanged')
  },
  computed: {
    disabledStartButton() {
      var exitedContainers = this.selected.filter(
        container => container.state === "exited"
      );
      return (
        this.selected.length === 0 ||
        exitedContainers.length !== this.selected.length
      );
    },
    disabledStopButton() {
      var runningContainers = this.selected.filter(
        container => container.state === "running"
      );
      return (
        this.selected.length === 0 ||
        runningContainers.length !== this.selected.length
      );
    },
    disabledDeleteButton() {
      var exitedContainers = this.selected.filter(
        container => container.state === "exited" || container.state === "created"
      );
      return (
        this.selected.length === 0 ||
        exitedContainers.length !== this.selected.length
      );
    }
  },
  methods: {
    tagsChanged: function (newTags) {
      this.tags = newTags;

      this.$emit('update:selected', [])
      this.$emit('update:filters', newTags)
      this.$emit('filterChanged')
    },
    async refreshContainers() {
      this.$emit("refreshContainers");
    },
    // 컨테이너 생성 이벤트를 발생합니다.
    async createContainer() {
      this.$emit("create");
    },
    confirmStartContainer() {
      this.startConfirmDialog = true;
    },
    confirmStopContainer() {
      this.stopConfirmDialog = true;
    },
    confirmDeleteContainer() {
      this.deleteConfirmDialog = true;
    },
    // 선택한 Container를 Start합니다.
    async startContainers() {
      this.startConfirmDialog = false;

      for (var container of this.selected) {
        await this.gomlapi.startContainer(container.nodeID, container.id);
      }
      this.$emit("start", this.selected);
      this.$emit('update:selected', [])
    },
    // 선택한 Container를 Stop합니다.
    async stopContainers() {
      this.stopConfirmDialog = false;

      for (var container of this.selected) {
        await this.gomlapi.stopContainer(container.nodeID, container.id);
      }
      this.$emit("stop", this.selected);
      this.$emit('update:selected', [])
    },
    // 선택한 Container를 삭제합니다.
    async deleteContainers() {
      this.deleteConfirmDialog = false;

      for (var container of this.selected) {
        await this.gomlapi.deleteContainer(container.nodeID, container.id);
      }
      this.$emit("delete", this.selected);
      this.$emit('update:selected', [])
    },
  }
};
</script>

<style>
#tag-div .vue-tags-input {
  max-width: none;
  width: 100%;
}

/* style the background and the text color of the input ... */
.vue-tags-input {
  max-width: none;
  background: #324652;
}

.vue-tags-input .ti-new-tag-input {
  background: transparent;
  color: #283944;
}

.vue-tags-input .ti-input {
  padding: 4px 10px;
  transition: border-bottom 200ms ease;
}

/* we cange the border color if the user focuses the input */
.vue-tags-input.ti-focus .ti-input {
  border: 1px solid #ebde6e;
}

/* some stylings for the autocomplete layer */
.vue-tags-input .ti-autocomplete {
  background: #283944;
  border: 1px solid #8b9396;
  border-top: none;
}

/* the selected item in the autocomplete layer, should be highlighted */
.vue-tags-input .ti-item.ti-selected-item {
  background: #ebde6e;
  color: #283944;
}

/* style the placeholders color across all browser */
.vue-tags-input ::-webkit-input-placeholder {
  color: #a4b1b6;
}

.vue-tags-input ::-moz-placeholder {
  color: #a4b1b6;
}

.vue-tags-input :-ms-input-placeholder {
  color: #a4b1b6;
}

.vue-tags-input :-moz-placeholder {
  color: #a4b1b6;
}

/* default styles for all the tags */
.vue-tags-input .ti-tag {
  position: relative;
  background: #1e88e5;
  color: #ffffff;
}

/* we defined a custom css class in the data model, now we are using it to style the tag */
.vue-tags-input .ti-tag.custom-class {
  background: transparent;
  border: 1px solid #ebde6e;
  color: #ebde6e;
  margin-right: 4px;
  border-radius: 0px;
  font-size: 13px;
}

/* the styles if a tag is invalid */
.vue-tags-input .ti-tag.ti-invalid {
  background-color: #e88a74;
}

/* if the user input is invalid, the input color should be red */
.vue-tags-input .ti-new-tag-input.ti-invalid {
  color: #e88a74;
}

/* if a tag or the user input is a duplicate, it should be crossed out */
.vue-tags-input .ti-duplicate span,
.vue-tags-input .ti-new-tag-input.ti-duplicate {
  text-decoration: line-through;
}

/* if the user presses backspace, the complete tag should be crossed out, to mark it for deletion */
.vue-tags-input .ti-tag:after {
  transition: transform 0.2s;
  position: absolute;
  content: "";
  height: 2px;
  width: 108%;
  left: -4%;
  top: calc(50% - 1px);
  background-color: #000;
  transform: scaleX(0);
}

.vue-tags-input .ti-deletion-mark:after {
  transform: scaleX(1);
}
</style>
