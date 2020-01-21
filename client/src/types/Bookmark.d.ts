import { ResourceType } from './Resource';

export interface IBookmark {
    uid: string;
    userID: string;
    resourceType: ResourceType;
}