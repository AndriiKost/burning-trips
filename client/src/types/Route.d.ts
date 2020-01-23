import { IStop } from './Stop';
import { IUser } from './User';
import { IRouteVote } from './Vote';

export interface IRoute {
    authorID: number;
    content: string;
    id: number;
    imageUrl: string;
    name: string;
    author?: IUser;
    stops: IStop[];
    votes: IRouteVote[];
}