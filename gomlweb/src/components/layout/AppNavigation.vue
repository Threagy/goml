<template>
  <v-navigation-drawer v-model="visibleDrawer" app clipped>
    <v-list dense>
      <v-list-item v-for="item in items" :key="item.text" :to="item.link">
        <v-list-item-action>
          <v-icon>{{ item.icon }}</v-icon>
        </v-list-item-action>
        <v-list-item-content>
          <v-list-item-title>{{ item.text }}</v-list-item-title>
        </v-list-item-content>
      </v-list-item>
    </v-list>
  </v-navigation-drawer>
</template>

<script>
export default {
  name: "AppNavigation",
  inheritAttrs: false,
  props: {
    drawer: {
      type: Boolean
    }
  },
  computed: {
    visibleDrawer: {
      get() {
        return this.drawer;
      },
      set(visible) {
        this.$emit("update:drawer", visible);
      }
    }
  },
  data() {
    return {
      items: [
        // { icon: "home", text: "Home", link: "/" },
        { icon: "dashboard", text: "Dashboard", link: "/" },
        { icon: "settings", text: "Settings", link: "/settings" }
      ]
    };
  }
};
</script>
