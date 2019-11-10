import { IStop } from '@/types/Stop';
import { IVote } from '@/types/Vote';
import { IAddress, ICoordinates } from '@/types/Geo';

export default class Stop implements IStop {
    trending: boolean = false;
    votes: IVote[] = [];
    uid: string = '';
    description: string = '';
    imageUrl: string = '';
    name: string = '';
    authorID: string = '';

    address: IAddress = {
        address: '',
        address2: '',
        city: '',
        zipcode: '',
        country: '',
        state: ''
    };

    coords: ICoordinates = {
        lat: 0,
        lng: 0
    };

}