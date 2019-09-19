import { SET_NODES, SET_CONTAINERS } from './mutation-types';
import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    accessToken: '',
    nickName: '',
    nodes: [],
    containers: [],
  },
  getters: {
    getAccessToken: state => {
      return state.accessToken;
    }
  },
  mutations: {
    [SET_NODES] (state, payload) {
      state.nodes = payload;
    },
    [SET_CONTAINERS] (state, payload) {
      state.containers = payload;
    },
    // [LOGOUT] (state) {
    //   state.accessToken = '';
    //   state.nickName = '';
    // }
  },
  actions: {
    // [SETNICKNAME] ...
  }
});
