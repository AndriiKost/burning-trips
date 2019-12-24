export interface IVote {
    id: number;
    userId: number;
    count: number;
}

export interface IVotable {
    trending: boolean;
    votes: IVote[];
}

export interface IStopVote extends IVote {
    stopId: number;
}