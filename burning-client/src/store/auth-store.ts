import axios from 'axios';
import { make } from 'vuex-pathify';
import { IUserToSign, IUser } from '../types/User';
import _svc from '@/services/AuthService';
import { Store } from 'vuex';

class AuthStore {
    user: IUser;
}

const state = new AuthStore();

const mutations = make.mutations(state);

const actions = {
    async signIn(_: Store<AuthStore>, userToSign: IUserToSign) {
        return await _svc.signIn(userToSign);
        // return await axios.post(`http://192.168.1.219:80/api/login`, userToSign);
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