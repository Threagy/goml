<template>
  <v-content>
    <v-container fill-height fluid>
      <v-layout align-center justify-center>
        <v-flex xs12 sm8 md4>
          <v-card class="elevation-12">
            <v-toolbar color="general">
              <v-toolbar-title>Login</v-toolbar-title>
              <v-spacer />
            </v-toolbar>
            <v-card-text>
              <v-form>
                <v-text-field
                  ref="username"
                  v-model="username"
                  :rules="[() => !!username || 'This field is required']"
                  prepend-icon="mdi-account"
                  label="Login"
                  placeholder="TotallyNotThanos"
                  required
                />
                <v-text-field
                  ref="password"
                  v-model="password"
                  :rules="[() => !!password || 'This field is required']"
                  :append-icon="showPassword ? 'mdi-eye-off' : 'mdi-eye'"
                  :type="showPassword ? 'text' : 'password'"
                  prepend-icon="mdi-lock"
                  label="Password"
                  placeholder="*********"
                  counter
                  required
                  @keydown.enter="login"
                  @click:append="showPassword = !showPassword"
                />
              </v-form>
            </v-card-text>
            <v-divider class="mt-5" />
            <v-card-actions>
              <v-spacer />
              <v-btn align-center justify-center color="general" @click="login"
                >Login
              </v-btn>
            </v-card-actions>
            <v-snackbar v-model="snackbar" :color="color" :top="true">
              {{ errorMessages }}
              <v-btn dark flat @click="snackbar = false">
                Close
              </v-btn>
            </v-snackbar>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
  </v-content>
  <!-- <div class="container">
    <h1>Sign in to access the secret page</h1>
    <div>
      <p>Username: <input v-model="formUsername" type="text" name="username"></p>
      <p>Password: <input v-model="formPassword" type="password" name="password"></p>
      <button @click="login">
        login
      </button>
      <p>The credentials are not verified for the example purpose.</p>
    </div>
  </div> -->
</template>

<script>
const Cookie = process.client ? require('js-cookie') : undefined

export default {
  layout: 'login',
  middleware: 'notAuthenticated',
  data() {
    return {
      username: '',
      password: '',
      errorMessages: 'Incorrect login info',
      snackbar: false,
      color: 'general',
      showPassword: false
      // formUsername: '',
      // formPassword: ''
    }
  },
  methods: {
    async login() {
      let username = this.username
      let password = this.password

      try {
        await this.$gomlapi.login({ user_id: username, password: password });
        this.$router.push('/')
      }
      catch (err) {
        console.error(err);
      }
    }
    // async login () {
    //   await this.$store.dispatch('login', {
    //       username: this.formUsername,
    //       password: this.formPassword
    //     })
    //   // setTimeout(() => { // we simulate the async request with timeout.
    //   //   const auth = {
    //   //     accessToken: 'someStringGotFromApiServiceWithAjax'
    //   //   }
    //   //   this.$store.commit('setAuth', auth) // mutating to store for client rendering
    //     Cookie.set('auth', this.$store.state.auth) // saving token in cookie for server rendering
    //     this.$router.push('/')
    //   // }, 1000)
    // }
  }
}
</script>
