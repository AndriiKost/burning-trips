import { make } from 'vuex-pathify';
import { Stop } from '@/types/stops.type';

class StopStore {
    stops: Array<Stop> = [
        {
            id: 'sajkdajknsdknjsa',
            name: "Stop Name",
            authorID: "asjdkjasdaklsjdklj12",
            description: "Such a cool place omg lol omg lorem ispum ups dolor set isjask!",
            votes: [
                { uid: 'sakldjaskljdklsa', votes: 59 }
            ],
            imageUrl: 'https://images.unsplash.com/photo-1572295727871-7638149ea3d7?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=500&q=60'
        }
    ];
    voteLimit: number = 0;
}

const state = new StopStore();

const mutations = make.mutations(state);

mutations['NEW_MUTATIONS'] = function(state: StopStore, arg: any) {
    console.log('New mutations, with argument = ', arg);
}

const actions = {

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