import { IStory } from '@/types/Story';
import { IStoryVote } from '@/types/Vote';

export default class Story implements IStory {
    id: number = 0;
    title: string = '';
    authorId: number = 0;
    content: string = '';
    votes: IStoryVote[] = [];
    imageUrl: string = '';
    
    constructor(authorId: number) {
        this.authorId = authorId;
    }
}