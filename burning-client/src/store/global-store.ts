import { make } from 'vuex-pathify';
import { User } from '@/types/User';
import { Store } from 'vuex';

class GlobalState {
    user: User = {
        firstName: 'Andrii',
        uid: 'Andrii1293812-213213',
        username: 'andriikost',
        email: 'andriikost@yahoo.com'
    };
}

const state = new GlobalState();

const mutations = make.mutations(state);

mutations['NEW_MUTATIONS'] = function(state: GlobalState, arg: any) {
}

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