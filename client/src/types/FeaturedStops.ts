import { IRoute } from './Route';
import { IStop } from './Stop';
import { IStory } from './Story';

export interface IFeaturedStops {
	nearestStops: IStop[];
	topRankedStops: IStop[];
	stories: IStory[];
	routes: IRoute[];
}