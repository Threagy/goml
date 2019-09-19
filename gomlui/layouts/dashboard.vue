<template>
  <v-content>
    <v-app id="dashboard-view" dark>
      <main>
        <!-- <core-filter /> -->

        <!-- <core-toolbar /> -->

        <core-drawer />

        <v-container fluid>
          <v-fade-transition mode="out-in">
            <router-view />
          </v-fade-transition>
        </v-container>
        <!-- <core-footer /> -->
      </main>
    </v-app>
  </v-content>
</template>

<script>
export default {
  name: 'dashboard',
  metaInfo() {
    return {
      title: 'dashboard page'
    }
  }
}
</script>

<style lang="scss">
@import "@/assets/index.scss";

/* Remove in 1.2 */
.v-datatable thead th.column.sortable i {
  vertical-align: unset;
}
</style>
