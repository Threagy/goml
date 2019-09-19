<template>
  <v-app-bar app dense :elevation="0">
    <v-spacer></v-spacer>

    <NuxtLink to="/login">
      <v-btn icon>
          <v-icon>mdi-account</v-icon>
      </v-btn>
    </NuxtLink>

    <v-btn icon @click="logout">
        <v-icon>mdi-power</v-icon>
    </v-btn>

  </v-app-bar>
</template>

<script>
import { mapMutations, mapGetters } from "vuex";

export default {
  data: () => ({
    notifications: [
      "Mike, Thanos is coming",
      "5 new avengers joined the team",
      "You're now friends with Capt",
      "Another Notification",
      "Another One - Dj Khalid voice"
    ],
    title: "I got a digital dash -Future Hendrixx",
    responsive: false,
    responsiveInput: false
  }),

  computed: {
    ...mapGetters(["authorized"])
  },

  watch: {
    $route(val) {
      this.title = val.meta.name;
    }
  },

  mounted() {
    this.onResponsiveInverted();
    window.addEventListener("resize", this.onResponsiveInverted);
  },
  beforeDestroy() {
    window.removeEventListener("resize", this.onResponsiveInverted);
  },

  methods: {
    ...mapMutations("app", ["setDrawer", "toggleDrawer"]),
    onClickBtn() {
      this.setDrawer(!this.$store.state.app.drawer);
    },
    onClick() {
      //
    },
    onResponsiveInverted() {
      if (window.innerWidth < 991) {
        this.responsive = true;
        this.responsiveInput = false;
      } else {
        this.responsive = false;
        this.responsiveInput = true;
      }
    },
    async logout() {
      // alert('logout')
      await this.$store.dispatch("logout");
      this.$router.push("");
    }
  }
};
</script>

<style>
#core-toolbar a {
  text-decoration: none;
}
</style>
