import { IRouteVote, IStopVote, IStoryVote } from '@/types/Vote';
import BaseHttpService from './BaseHttpService';

class VoteService extends BaseHttpService {

    async updateStopVote(vote: IStopVote) {
        const res = await this.post<IStopVote>('/votes/stop-votes', vote);
        if (!res.ok || !res) return null
        return res.result;
    }

    async updateRouteVote(vote: IRouteVote) {
        const res = await this.post<IRouteVote>('/votes/route-votes', vote);
        if (!res.ok || !res) return null
        return res.result;
    }

    async updateStoryVote(vote: IStoryVote) {
        const res = await this.post<IStoryVote>('/votes/story-votes', vote);
        if (!res.ok || !res) return null
        return res.result;
    }
    
}

export default new VoteService();