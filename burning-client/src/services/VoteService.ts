import BaseHttpService from './BaseHttpService';
import { IStop } from '../types/Stop';
import { IStopVote } from '@/types/Vote';

class VoteService extends BaseHttpService {

    async updateStopVote(vote: IStopVote) {
        const res = await this.post<IStop>('/votes/stop-votes', vote);
        if (!res.ok || !res) return null
        return res.result;
    }
    
}

export default new VoteService();