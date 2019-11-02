export interface Stop {
    id: String;
    name: string;
    authorID: string;
    votes: Array<Vote>;
    description: string;
    imageUrl: string;
    isBurning: boolean;
}

export interface Vote {
    uid: string;
    votes: number;
}