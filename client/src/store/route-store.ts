import routeService from '@/services/RouteService';
import voteService from '@/services/VoteService';
import { IRoute } from '@/types/Route';
import { IRouteVote } from '@/types/Vote';
import { Store } from 'vuex';
import { make } from 'vuex-pathify';

class RouteStore {
    routes: Array<IRoute> = [];
}

const state = new RouteStore();

const mutations = make.mutations(state);

mutations['NEW_MUTATIONS'] = function(state: RouteStore, arg: any) {
}

const actions = {

    async createRoute({ commit, state }: Store<RouteStore>, route: IRoute) {
        return await routeService.createRoute(route);
    },

    async getAllRoutes({ commit, state }: Store<RouteStore>) {
        const routes = await routeService.getAllRoutes();
        if (!routes) return [];
        commit('SET_ROUTES', routes);
        return routes;
    },

    async getRoute({ commit, state }: Store<RouteStore>, id: number) {
        return await routeService.getRoute(id);
    },

    async updateRouteVote({ commit }: Store<RouteStore>, vote: IRouteVote) {
        return await voteService.updateRouteVote(vote);
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