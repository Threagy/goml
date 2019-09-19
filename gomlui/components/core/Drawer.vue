<template>
  <v-navigation-drawer
    id="app-drawer"
    v-model="inputValue"
    app
    overflow
    dark
    mobile-break-point="991"
    width="260"
  >
    <v-list nav dense height="100%">
      <v-list-item>
        <v-list-item-content class="title pl-4">
          Metallica
        </v-list-item-content>
      </v-list-item>
      <v-divider />
      <v-list-item two-line v-if="$store.state.auth !== ''">
        <v-list-item-avatar>
          <no-ssr>
          <avatar :username="username" :size="40"></avatar>
          </no-ssr>
          <!-- <img src="https://randomuser.me/api/portraits/women/81.jpg" /> -->
        </v-list-item-avatar>

        <v-list-item-content>
          <v-list-item-title>{{username}}</v-list-item-title>
          <v-list-item-subtitle>Logged In</v-list-item-subtitle>
        </v-list-item-content>
        <v-btn @click="logout" class="ma-2" small icon tile>
          <v-icon>mdi-power</v-icon>
        </v-btn>
      </v-list-item>

      <v-list-item v-if="$store.state.auth === ''">
        <v-list-item-content>
          <v-btn @click="login">
            Login
            <v-icon right class="px-4">mdi-account</v-icon>
          </v-btn>
        </v-list-item-content>
      </v-list-item>

      <v-divider />

      <v-list-item
        active-class="grey--text"
        v-for="item in links"
        :key="item.text"
        :to="item.to"
      >
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
const Cookie = process.client ? require('js-cookie') : undefined
const jwt = process.client ? require('jsonwebtoken') : undefined

// Utilities
import {
  mapMutations,
  mapState
} from 'vuex'

export default {
  data: () => ({
    logo: require('~/assets/img/redditicon.png'),
    username: '',
    links: [
      {
        to: '/',
        icon: 'dashboard',
        text: 'Dashboard'
      },
      {
        to: '/containers',
        icon: 'grid_on',
        text: 'Containers'
      },
      {
        to: '/images',
        icon: 'filter_3',
        text: 'Images'
      },
      {
        to: '/settings',
        icon: 'settings',
        text: 'Settings'
      },
      {
        to: '/users',
        icon: 'people',
        text: 'Users'
      },
    ],
    responsive: false
  }),
  computed: {
    ...mapState('app', ['image', 'color']),
    inputValue: {
      get() {
        return this.$store.state.app.drawer
      },
      set(val) {
        this.setDrawer(val)
      }
    },
    items() {
      return this.$t('Layout.View.items')
    }
  },
  mounted() {
    this.onResponsiveInverted()
    window.addEventListener('resize', this.onResponsiveInverted)
  },
  beforeDestroy() {
    window.removeEventListener('resize', this.onResponsiveInverted)
  },
  created: async function () {
    try {
      if (_.isNil(this.$gomlapi) === true || _.isNil(jwt) === true) return;
      var decoded = jwt.decode(this.$store.state.auth);
      this.username = decoded.Name;
    } catch (err) {
      console.error(err);
    }
  },
  methods: {
    ...mapMutations('app', ['setDrawer', 'toggleDrawer']),
    onResponsiveInverted() {
      if (window.innerWidth < 991) {
        this.responsive = true
      } else {
        this.responsive = false
      }
    },
    login() {
      this.$router.push('/login')
    },
    async logout() {
      await this.$gomlapi.logout();
      this.$router.push({path:'/login'});
    }
  }
}
</script>

<style lang="scss" scoped>
// @import "@/assets/material/_color.scss";
// @import "@/assets/material/_variables.scss";

// #app-drawer {
//   background: $sidebar-bg;
// }
</style>
