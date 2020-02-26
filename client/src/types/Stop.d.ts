import { IVotable, IStopVote } from './Vote';
import { IAddress, ICoordinates } from './Geo';
import { IUser } from './User';

export interface IStop {
    address?: string;
    authorID: number;
    content: string;
    id: number;
    imageUrl: string;
    name: string;
    coords?: ICoordinates;
    author?: IUser;
    longitude?: number;
    latitude?: number;
    votes: IStopVote[];
}