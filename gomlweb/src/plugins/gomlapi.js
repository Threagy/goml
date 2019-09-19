import axios from 'axios';

export default {
  install(Vue) {
    Vue.prototype.gomlapi = {
      address: function () {
        return window.location.hostname + ":7777";
      },
      getNodes: function () {
        return axios.get(`http://${this.address()}/v1/api/nodes`);
      },
      getNode: function (nodeAddress) {
        return axios.get(`http://${this.address()}/v1/api/node/${nodeAddress}`);
      },
      getSystem: function () {
        return axios.get(`http://${this.address()}/v1/api/system`);
      },
      setConfig: function (body) {
        return axios.put(`http://${this.address()}/v1/api/config`, body);
      },
      createContainer: function (body) {
        return axios.post(`http://${this.address()}/v1/api/container`, body);
      },
      startContainer: function (nodeID, containerID) {
        return axios.put(`http://${this.address()}/v1/api/container/start/${nodeID}/${containerID}`);
      },
      stopContainer: function (nodeID, containerID) {
        return axios.put(`http://${this.address()}/v1/api/container/stop/${nodeID}/${containerID}`);
      },
      deleteContainer: function (nodeID, containerID) {
        return axios.delete(`http://${this.address()}/v1/api/container/${nodeID}/${containerID}`);
      },
    }
  }
}
