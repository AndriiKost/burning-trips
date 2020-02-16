import { ISearchQuery, ISearchResult } from '@/types/Explore';
import { IStop } from '@/types/Stop';
import BaseHttpService from './BaseHttpService';

class ExploreService extends BaseHttpService {

    async searchStops(query: ISearchQuery): Promise<ISearchResult<IStop>> {
        const res = await this.post<ISearchResult<IStop>>('/search-stops', query);
        if (!res.ok || !res) return null;
        return res.result;
    }

}

export default new ExploreService();