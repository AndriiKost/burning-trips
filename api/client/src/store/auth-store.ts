import config from '@/config';
import { default as authService, default as _svc } from '@/services/AuthService';
import { ISignInRequest } from 'http';
import Cookies from 'js-cookie';
import { Store } from 'vuex';
import { make } from 'vuex-pathify';
import { IUser } from '../types/User';

class AuthStore {
    loggedInUser: IUser = null;
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