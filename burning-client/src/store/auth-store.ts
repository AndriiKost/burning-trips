import axios from 'axios';
import { make } from 'vuex-pathify';
import { IUserToSign, IUser } from '../types/User';
import { Store } from 'vuex';

class AuthStore {
    user: IUser;
}

const state = new AuthStore();

const mutations = make.mutations(state);

const actions = {
    async signIn(_: Store<AuthStore>, userToSign: IUserToSign) {
        return await axios.post(`http://localhost:8080/login`, userToSign, { crossdomain: true, headers: {'Access-Control-Allow-Origin': '*', 'Content-Type': 'application/json'} });
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