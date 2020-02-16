import { IUser } from '@/types/User';
import { make } from 'vuex-pathify';

class GlobalState {
    user: IUser = null;
}

const state = new GlobalState();

const mutations = make.mutations(state);

const actions = {
    
}

const getters = {

}

export default {
    state,
    mutations,
    actions,
    getters
}