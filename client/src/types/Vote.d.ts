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

export interface IRouteVote extends IVote {
    routeId: number;
}

export interface IStoryVote extends IVote {
    storyId: number;
}