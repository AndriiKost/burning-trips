import { make } from 'vuex-pathify';
import { IStop } from '@/types/Stop';
import { Store } from 'vuex';
import stopService from '@/services/StopService';

class StopStore {
    stops: Array<IStop> = [];
}

const state = new StopStore();

const mutations = make.mutations(state);

mutations['NEW_MUTATIONS'] = function(state: StopStore, arg: any) {
}

const actions = {

    async createStop({ commit, state }: Store<StopStore>, newStop: IStop) {
        return await stopService.createStop(newStop);
    },

    async getAllStops({ commit, state }: Store<StopStore>) {
        const stops = await stopService.getAllStops();
        if (!stops) return [];
        commit('SET_STOPS', stops);
        return stops;
    },

    async getStop(_, stopID: number) {
        return await stopService.getStop(stopID);
    }
}

const getters = {

}

export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
}