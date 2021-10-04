package ipapi

import "github.com/avtara/travair-api/businesses/iplocator"

type Response struct {
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
}

func (resp *Response) toDomain() *iplocator.Domain {
	return &iplocator.Domain{
		Latitude: resp.Latitude,
		Longitude: resp.Longitude,
	}
}