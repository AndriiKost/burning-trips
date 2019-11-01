import Vuex from 'vuex';

const regionStore = {

    namespaced: true,

    state: {
        regions: []
    },

    mutations: {

    },

    actions: {
        async getBurningRegions({ commit, store }, coords) {
            return [];
        }
    },

    getters: {

    }
}

export default regionStore;