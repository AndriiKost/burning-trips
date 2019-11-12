import { IStop } from '@/types/Stop';
import { IVote } from '@/types/Vote';
import { IAddress, ICoordinates } from '@/types/Geo';
import { IUser } from '@/types/User';

export default class Stop implements IStop {
    trending: boolean = false;
    votes: IVote[] = [];
    id: number = 0;
    content: string = '';
    imageUrl: string = '';
    name: string = '';
    authorID: number = 0;
    address: string = '';

    // IAddress = {
    //     address: '',
    //     address2: '',
    //     city: '',
    //     zipcode: '',
    //     country: '',
    //     state: ''
    // };

    // coords: ICoordinates = {
    //     lat: 0,
    //     lng: 0
    // };

}