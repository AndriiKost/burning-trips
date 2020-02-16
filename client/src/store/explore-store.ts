import exploreService from '@/services/ExploreService';
import { ISearchQuery, ISearchResult } from '@/types/Explore';
import { IStop } from '@/types/Stop';
import { Store } from 'vuex';
import { commit, make } from 'vuex-pathify';

class ExploreStore {
    searchResult: any = [];
}

const state = new ExploreStore();

const mutations = make.mutations(state);

mutations['NEW_MUTATIONS'] = function(state: ExploreStore, arg: any) {
}

const actions = {

    async searchStops(_: Store<ExploreStore>, query: ISearchQuery): Promise<ISearchResult<IStop>> {
        const searchResult = await exploreService.searchStops(query);
        if (searchResult.result && searchResult.result.length > 0) {
            commit('searchResult', searchResult);
            return searchResult;
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