import storyService from '@/services/StoryService';
import voteService from '@/services/VoteService';
import { IStop } from '@/types/Stop';
import { IStory } from '@/types/Story';
import { IStoryVote } from '@/types/Vote';
import { Store } from 'vuex';
import { make } from 'vuex-pathify';

class StoryStore {
    stories: Array<IStop> = [];
}

const state = new StoryStore();

const mutations = make.mutations(state);

mutations['NEW_MUTATIONS'] = function(state: StoryStore, arg: any) {
}

const actions = {

    async createStory(_: Store<StoryStore>, story: IStory) {
        return await storyService.createStory(story);
    },

    async getAllStories({ commit }: Store<StoryStore>) {
        const stops = await storyService.getAllStories();
        if (!stops) return [];
        commit('SET_STORIES', stops);
        return stops;
    },

    async getStory(_, id: number) {
        return await storyService.getStory(id);
    },

    async updateStoryVote(_: Store<StoryStore>, vote: IStoryVote) {
        return await voteService.updateStoryVote(vote);
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