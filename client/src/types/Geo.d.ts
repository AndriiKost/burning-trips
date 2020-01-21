export interface IAddress {
    address: string;
    address2?: string;
    city: string;
    zipcode: number | string;
    state?: string;
    country: string;
}

export interface ICoordinates {
    lat: number;
    lng: number;
}