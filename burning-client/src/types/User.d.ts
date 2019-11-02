import { IBookmark } from "./Bookmark";

export interface IUser {
    firstName?: string;
    lastName?: string;
    username: string;
    uid: string;
    email: string;
    bookmarks: IBookmark;
}