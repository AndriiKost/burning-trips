import { IRoute } from '../types/Route';
import BaseHttpService from './BaseHttpService';

class RouteService extends BaseHttpService {

    async createRoute(route: IRoute) {
        const res = await this.post<IRoute>('/routes', route);
        if (!res.ok || !res) return null
        return res.result;
    }

    async getAllRoutes() {
        const res = await this.get<IRoute[]>('/routes');
        if (!res.ok || !res) return [];
        return res.result;
    }

    async getRoute(id: number) {
        const res = await this.get<IRoute>(`/routes/${id}`);
        if (!res.ok || !res) return null;
        return res.result;
    }

}

export default new RouteService();