import { IStop } from '@/types/Stop';
import { IStopVote } from '@/types/Vote';
// import { IAddress, ICoordinates } from '@/types/Geo';
// import { IUser } from '@/types/User';

export default class Stop implements IStop {
    trending: boolean = false;
    votes: IStopVote[] = [];
    id: number = 0;
    content: string = '';
    imageUrl: string = '';
    name: string = '';
    authorID: number = 0;
    address: string = '';

    longtitude: any = null;
    latitude: any = null;

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