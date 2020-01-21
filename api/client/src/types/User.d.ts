import { IBookmark } from "./Bookmark";

export interface IUser {
    firstName?: string;
    lastName?: string;
    username: string;
    id: number;
    email: string;
    bookmarks?: IBookmark;
}

export interface IUserToSign {
    email: string;
    password: string;
}