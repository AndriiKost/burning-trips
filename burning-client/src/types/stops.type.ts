export interface Stop {
    id: String;
    name: string;
    authorID: string;
    votes: Array<Vote>;
    description: string;
    imageUrl: string;
}

export interface Vote {
    uid: string;
    votes: number;
}