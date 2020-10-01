package models

type SearchQuery struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Keyword   string  `json:"keyword"`
}

func getLatMinMax(query SearchQuery) (float32, float32) {
	min_lat := query.Latitude - 0.1
	max_lat := query.Latitude + 0.1
	return min_lat, max_lat
}
func getLngMinMax(query SearchQuery) (float32, float32) {
	min_lng := query.Longitude - 0.1
	max_lng := query.Longitude + 0.1
	return min_lng, max_lng
}
