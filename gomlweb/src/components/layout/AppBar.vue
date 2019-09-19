<template>
    <v-app-bar
      app
      clipped-left
      color="primary"
      dark
      dense
    >
      <v-app-bar-nav-icon @click.stop="toggleDrawer"></v-app-bar-nav-icon>
      <v-icon class="mx-4">fab fa-youtube</v-icon>
      <v-toolbar-title class="mr-12 align-center">
        <span class="title">Metallica</span>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-layout
        align-center
        style="max-width: 650px"
      >
      </v-layout>
    </v-app-bar>
</template>

<script>
export default {
  name: "AppBar",
  inheritAttrs: false,
  props: {
    drawer: {
      type: Boolean,
    }
  },
  methods: {
    toggleDrawer() {
      this.$emit('update:drawer', !this.drawer)
    }
  }
};
</script>
