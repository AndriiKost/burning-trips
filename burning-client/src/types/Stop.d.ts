import { IVotable } from './Vote';
import { IAddress, ICoordinates } from './Geo';
import { IUser } from './User';

export interface IStop extends IVotable {
    address?: string;
    authorID: number;
    content: string;
    id: number;
    imageUrl: string;
    name: string;
    coords?: ICoordinates;
    author?: IUser;
    longtitude?: number;
    latitude?: number;
}