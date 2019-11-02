export interface IVote {
    uid?: string;
    userID: string;
    userVotes: number;
}

export interface IVotable {
    trending: boolean;
    votes: IVote[];
}