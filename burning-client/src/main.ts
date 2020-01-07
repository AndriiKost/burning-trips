import ElementUI from 'element-ui';
import Vue from "vue";
import AppContent from './components/layout/AppContent.vue';
import config from './config';
import "./registerServiceWorker";
import router from "./router";
import './scss/global.scss';
import store from './store';


Vue.use(ElementUI);

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: h => h(AppContent),
  created() {
    console.log('config.API_URL', config.API_URL)
  }
}).$mount("#app");
