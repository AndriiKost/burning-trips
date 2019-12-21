import Vue from "vue";
import App from "./App.vue";
import "./registerServiceWorker";
import router from "./router";
import store from './store';

import ElementUI from 'element-ui';
import './scss/global.scss';
import config from './config';

Vue.use(ElementUI);

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: h => h(App),
  created() {
    console.log('config.API_URL', config.API_URL)
  }
}).$mount("#app");
