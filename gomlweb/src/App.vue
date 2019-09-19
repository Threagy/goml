<template>
  <v-app>
    <v-app>
      <gm-app-navigation :drawer="drawer"></gm-app-navigation>
      <gm-app-bar :drawer.sync="drawer"></gm-app-bar>
      <gm-node-observer ref="nodeObserver"></gm-node-observer>

      <v-content>
        <router-view />
      </v-content>
    </v-app>
  </v-app>
</template>

<script>
import AppBar from "@/components/layout/AppBar.vue";
import AppNavigation from "@/components/layout/AppNavigation.vue";
import NodeObserver from "@/components/common/NodeObserver.vue";

export default {
  name: "App",
  components: {
    "gm-app-bar": AppBar,
    "gm-app-navigation": AppNavigation,
    "gm-node-observer": NodeObserver,
  },
  data() {
    return {
      drawer: true
    };
  }
};
</script>
