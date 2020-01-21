import { IUser } from "./User";
import { IStoryVote } from "./Vote";

export interface IStory {
    authorId: number;
    content: string;
    id: number;
    imageUrl: string;
    title: string;
    author?: IUser;
    votes: IStoryVote[];
}