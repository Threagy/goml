<template>
  <div></div>
</template>

<script>
/*
import { mapMutations } from "vuex";
import { addAdditionalProp, filterContainers } from "@/utils/utils-function";

export default {
  name: "NodeObserver",
  data: () => {
    return {
      timer: null
    };
  },
  created: async function() {
    console.log("created observer");
    await this.fetchNodeList();
  },
  methods: {
    ...mapMutations({ setNodes: "SET_NODES" }),
    ...mapMutations({ setContainers: "SET_CONTAINERS" }),
    async fetchNodeList(filters) {
      try {
        var res = await this.gomlapi.getNodes();

        var nodes = res.data.nodes;

        addAdditionalProp(nodes);

        var pickContainers = this._.map(nodes, "containers");
        var containers = this._.flatten(pickContainers);

        var filterdContainers = containers;

        if (this._.isNil(filters) === false)
          filterdContainers = filterContainers(containers);

        this.setNodes(nodes);
        this.setContainers(filterdContainers);
      } catch (err) {
        console.error(err);
      } finally {
        //this.timer = setTimeout(this.fetchNodeList, 10000);
      }
    },
  }
};
*/
</script>
