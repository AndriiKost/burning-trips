import { make } from 'vuex-pathify';
import { IStop } from '@/types/Stop';
import { Store } from 'vuex';
import stopService from '@/services/StopService';

class StopStore {
    stops: Array<IStop> = [
        {
            uid: 'sajkdajknsdknjsa',
            name: "Stop Name",
            authorID: "asjdkjasdaklsjdklj12",
            content: "Such a cool place omg lol omg lorem ispum ups dolor set isjask!",
            trending: true,
            address: {
                state: 'SA',
                address: '1 Capitol Square',
                zipcode: 53701,
                country: 'USA',
                city: 'Madison'
            },
            coords: {
                lat: 8912381297398127938721,
                lng: 128938912398781293871293
            },
            votes: [
                { uid: 'sakldjaskljdklsa', userVotes: 59, userID: 'ahsdjkhaskjdhjkashdjkhas' }
            ],
            imageUrl: 'https://images.unsplash.com/photo-1572295727871-7638149ea3d7?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=500&q=60'
        },
        {
            uid: 'sdasadasdasdass',
            name: "Madison Stop",
            authorID: "sheila shasl222",
            content: "Super duper awesome place!",
            trending: false,
            address: {
                state: 'SA',
                address: '1 Capitol Square',
                zipcode: 53701,
                country: 'USA',
                city: 'Madison'
            },
            coords: {
                lat: 8912381297398127938721,
                lng: 128938912398781293871293
            },
            votes: [
                { uid: 'sakldjaskljdklsa', userVotes: 19, userID: '213asdajksh' }
            ],
            imageUrl: 'https://images.unsplash.com/photo-1561953131-4b07835d56e8?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1351&q=80'
        },
    ];
}

const state = new StopStore();

const mutations = make.mutations(state);

mutations['NEW_MUTATIONS'] = function(state: StopStore, arg: any) {
}

const actions = {

    async createStop({ commit, state }: Store<StopStore>, newStop: IStop) {
        return await stopService.createStop(newStop);
    },

    async getAllStops({ commit, state }: Store<StopStore>) {
        const stops = await stopService.getAllStops();
        if (!stops) return [];
        commit('SET_STOPS', stops);
        return stops;
    },

    async getStop(_, stopID: number) {
        return await stopService.getStop(stopID);
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