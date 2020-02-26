export interface ISearchQuery {
    longitude: number;
    latitude: number;
}

export interface ISearchResult<T> {
    result: Array<T>;
}