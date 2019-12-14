import Vue from "vue";
import Vuex from "vuex";
import pathify from 'vuex-pathify';
import regionStore from './region-store';
import stopStore from './stop-store';
import globalStore from './global-store';
import authStore from './auth-store';
import uploadStore from './upload-store';

Vue.use(Vuex);

export default new Vuex.Store({
	...globalStore,
	modules: {
		region: regionStore,
		stop: stopStore,
		auth: authStore,
		upload: uploadStore
	},
	plugins: [pathify.plugin]
});
