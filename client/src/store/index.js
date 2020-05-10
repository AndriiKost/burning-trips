import Vue from "vue";
import Vuex from "vuex";
import pathify from 'vuex-pathify';
import stopStore from './stop-store';
import routeStore from './route-store';
import globalStore from './global-store';
import authStore from './auth-store';
import storyStore from './story-store';
import uploadStore from './upload-store';
import exploreStore from './explore-store';

Vue.use(Vuex);

export default new Vuex.Store({
	...globalStore,
	modules: {
		stop: stopStore,
		auth: authStore,
		upload: uploadStore,
		route: routeStore,
		explore: exploreStore,
		story: storyStore
	},
	plugins: [pathify.plugin]
});
