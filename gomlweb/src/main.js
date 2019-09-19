import Vue from 'vue'
import VueRouter from 'vue-router'
import VueLodash from 'vue-lodash'
import App from './App.vue'
import Gomlapi from './plugins/gomlapi'
import routes from './routes';
import store from './store';
import vuetify from './plugins/vuetify';
import VuePrism from 'vue-prism'
import 'prismjs/themes/prism-okaidia.css'

Vue.use(VueRouter);
Vue.use(VueLodash);
Vue.use(VuePrism);
Vue.use(Gomlapi);

const router = new VueRouter({ mode: 'history', routes });

new Vue({
  router,
  vuetify,
  store,
  render: h => h(App),
}).$mount('#app');
