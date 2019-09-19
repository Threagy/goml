import axios from 'axios';
const Cookie = process.client ? require('js-cookie') : undefined

export default ({ app, store }, inject) => {
  function header() {
    var token = store.state.auth;
    var h = {
      headers: {
        authorization: token
      }
    };
    return h;
  }

  var gomlapi = {
    address: function () {
      return window.location.hostname + ":7777";
    },
    login: async function (body) {
      var res = await axios.post(`http://${this.address()}/v1/auth/login`, body);
      var token = res.data.token

      store.commit('setAuth', token);
      Cookie.set('auth', token);
      axios.defaults.headers.common['authorization'] = token;
    },
    logout: function () {
      store.commit('setAuth', '');
      Cookie.set('auth', '');
      axios.defaults.headers.common['authorization'] = '';
    },
    getUsers: function () {
      return axios.get(`http://${this.address()}/v1/users`, header());
    },
    getNodes: function () {
      return axios.get(`http://${this.address()}/v1/api/nodes`, header());
    },
    getNode: function (nodeAddress) {
      return axios.get(`http://${this.address()}/v1/api/node/${nodeAddress}`, header());
    },
    getSystem: function () {
      return axios.get(`http://${this.address()}/v1/api/system`, header());
    },
    setConfig: function (body) {
      return axios.put(`http://${this.address()}/v1/api/config`, body, header());
    },
    createContainer: function (body) {
      return axios.post(`http://${this.address()}/v1/api/container`, body, header());
    },
    startContainer: function (nodeID, containerID) {
      return axios.put(`http://${this.address()}/v1/api/container/start/${nodeID}/${containerID}`, {}, header());
    },
    stopContainer: function (nodeID, containerID) {
      return axios.put(`http://${this.address()}/v1/api/container/stop/${nodeID}/${containerID}`, {}, header());
    },
    deleteContainer: function (nodeID, containerID) {
      return axios.delete(`http://${this.address()}/v1/api/container/${nodeID}/${containerID}`, header());
    },
  }

  inject('gomlapi', gomlapi)
}
