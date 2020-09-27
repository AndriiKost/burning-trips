import stopService from '@/services/StopService';
import ViewModelService from '@/services/ViewModelService';
import voteService from '@/services/VoteService';
import { IFeaturedStops } from '@/types/FeaturedStops';
import { IStop } from '@/types/Stop';
import { IStopVote } from '@/types/Vote';
import { Store } from 'vuex';
import { make } from 'vuex-pathify';

const vmSvc = new ViewModelService();

class StopStore {
    stops: IStop[] = [];
    featuredLandmarks: IStop[] = [];
    upsellLandmarks: IStop[] = [];
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
    },

    async updateStopVote({ commit }: Store<StopStore>, vote: IStopVote) {
        return await voteService.updateStopVote(vote);
    },

    async getStopViewModel(_, stopId: number): Promise<IFeaturedStops> {
        return await vmSvc.getStopViewModel(stopId);
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