
<template>
  <v-data-table
    :headers="headers"
    :items="nodes"
    class="elevation-1"
    hide-default-footer
  >
    <template v-slot:item="props">
      <tr>
        <td>{{ props.item.alias }}</td>
        <td class="text-xs-left" v-bind:key="index" v-for="index in 8">
          <v-tooltip bottom v-if="props.item.gpu_statuses.length > index - 1">
            <template v-slot:activator="{ on }">
              <v-progress-linear
                v-on="on"
                :size="36"
                :height="20"
                :value="props.item.gpu_statuses[index - 1].mem_used_percent"
                color="primary"
              />
            </template>
            <span>{{ props.item.gpu_statuses[index - 1].tooltip }}</span>
          </v-tooltip>
        </td>
      </tr>
    </template>
  </v-data-table>
</template>

<script>
export default {
  name: "NodeTable",
  inheritAttrs: false,
  props: {
    nodes: {
      type: [Array],
      default: () => []
    }
  },
  data: function() {
    return {
      pagination: {
        sortBy: "nodeAlias",
        rowsPerPage: 10
      },
      headers: [
          { text: "Node", align: "left", sortable: false },
          { text: "GPU 0", align: "center", sortable: false },
          { text: "GPU 1", align: "center", sortable: false },
          { text: "GPU 2", align: "center", sortable: false },
          { text: "GPU 3", align: "center", sortable: false },
          { text: "GPU 4", align: "center", sortable: false },
          { text: "GPU 5", align: "center", sortable: false },
          { text: "GPU 6", align: "center", sortable: false },
          { text: "GPU 7", align: "center", sortable: false }
      ],
    };
  },
};
</script>
