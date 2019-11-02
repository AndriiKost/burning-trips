import { IVotable } from './Vote';
import { IAddress, ICoordinates } from './Geo';

export interface IStop extends IVotable {
    address: IAddress;
    authorID: string;
    description: string;
    uid: string;
    imageUrl: string;
    name: string;
    coords: ICoordinates;
}