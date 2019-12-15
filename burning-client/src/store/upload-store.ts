import { make } from 'vuex-pathify';
import { Store } from 'vuex';
import uploadService from '@/services/UploadService';

class UploadStore {
    
}

const state = new UploadStore();

const mutations = make.mutations(state);

const actions = {

    async getUploadUrl(_) {
        return await uploadService.getPresignedUploadUrl();
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