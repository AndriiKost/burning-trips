import exploreService from '@/services/ExploreService';
import { ISearchQuery } from '@/types/Explore';
import { IStop } from '@/types/Stop';
import { Store } from 'vuex';
import { make } from 'vuex-pathify';

class ExploreState {
    stops: IStop[] = [];
}

const state = new ExploreState();

const mutations = make.mutations(state);

const actions = {

    async searchStops({ commit }: Store<ExploreState>, query: ISearchQuery): Promise<IStop[]> {
        const stops = await exploreService.searchStops(query);
        if (stops && stops.length > 0) {
            commit('SET_STOPS', stops);
            return stops;
        }
    },

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