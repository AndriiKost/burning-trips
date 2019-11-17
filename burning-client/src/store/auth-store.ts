import Cookies from 'js-cookie';
import { make, commit } from 'vuex-pathify';
import { IUser } from '../types/User';
import _svc from '@/services/AuthService';
import { Store } from 'vuex';
import { ISignInRequest } from 'http';
import config from '@/config';
import authService from '@/services/AuthService';

class AuthStore {
    loggedInUser: IUser;
}

const state = new AuthStore();

const mutations = make.mutations(state);
mutations['SET_LOGGED_IN_USER'] = function(state: AuthStore, user: IUser) {
    state.loggedInUser = user;
}

const actions = {
    async signIn({ commit }: Store<AuthStore>, userToSign: ISignInRequest) {
        const loggedInUser = await _svc.signIn(userToSign);
        if (!loggedInUser) return false;
        commit('SET_LOGGED_IN_USER', loggedInUser);
        return true;
    },

    async initSession({ commit, state }: Store<AuthStore>): Promise<Boolean> {
        if (state.loggedInUser) return true;
        const token = Cookies.get(config.JWT_NAME);
        if (!token) return false;
        const user = await authService.getUserData();
        if (!user) return false;
        commit('SET_LOGGED_IN_USER', user);
        return true;
    }
}

const getters = {
    dummy(state) {
        return 'dummy';
    }
}

export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
}