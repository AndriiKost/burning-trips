import VoteSection from '@/components/global/VoteSection.vue';
import CKEditor from '@ckeditor/ckeditor5-vue';
import ElementUI from 'element-ui';
import Vue from "vue";
import AppContent from './components/layout/AppContent.vue';
import './registerServiceWorker';
import router from './router';
import './scss/global.scss';
import store from './store';

Vue.component('vote-section', VoteSection);
Vue.use(CKEditor);

Vue.use(ElementUI);

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: h => h(AppContent)
}).$mount("#app");
