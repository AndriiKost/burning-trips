import { make } from 'vuex-pathify';
import { IStop } from '@/types/Stop';

class StopStore {
    stops: Array<IStop> = [
        {
            uid: 'sajkdajknsdknjsa',
            name: "Stop Name",
            authorID: "asjdkjasdaklsjdklj12",
            description: "Such a cool place omg lol omg lorem ispum ups dolor set isjask!",
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
            description: "Super duper awesome place!",
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
        }
    ];
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