<template>
  <div>
    <v-app-bar
      dense
      app
      tabs
      class="elevation-2"
    >
      <v-toolbar-title>Users</v-toolbar-title>

      <v-spacer></v-spacer>

    </v-app-bar>
    <v-card>
      <v-card-text>
        <no-ssr>
          <v-data-table
            :headers="headers"
            :items="users"
            class="elevation-1"
          >
          </v-data-table>
        </no-ssr>
      </v-card-text>
    </v-card>
  </div>
</template>

<script>
export default {
  layout: 'dashboard',
  middleware: 'authenticated',
  data: () => ({
    users: [],
    selected: [],
    snackBar: false,
    snackMessage: "",
    headers: [
      { text: "ID", align: "left", value: "user_id" },
      { text: "Name", align: "left", value: "name" },
      { text: "Role", align: "left", value: "role" },
    ],
  }),
  created: async function () {
    if (_.isNil(this.$gomlapi) === true) return;
    var res = await this.$gomlapi.getUsers();
    var users = res.data.users;

    this.users = users;
  },
  asyncData(context) {
    return { project: 'nuxt' }
  }
}
</script>
