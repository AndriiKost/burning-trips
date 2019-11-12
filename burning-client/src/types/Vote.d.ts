export interface IVote {
    id: number;
    userID: number;
    userVotes: number;
}

export interface IVotable {
    trending: boolean;
    votes: IVote[];
}