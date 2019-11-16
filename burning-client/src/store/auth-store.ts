import axios from 'axios';
import { make, commit } from 'vuex-pathify';
import { IUser } from '../types/User';
import _svc from '@/services/AuthService';
import { Store } from 'vuex';
import { ISignInRequest } from 'http';

class AuthStore {
    user: IUser;
}

const state = new AuthStore();

const mutations = make.mutations(state);
mutations['SET_USER'] = function(state: AuthStore, user: IUser) {
    state.user = user;
}

const actions = {
    async signIn({ commit }: Store<AuthStore>, userToSign: ISignInRequest) {
        const loggedInUser = await _svc.signIn(userToSign);
        if (!loggedInUser) return;
        commit('SET_USER', loggedInUser);
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