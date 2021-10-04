package iplocator

import "context"

type Domain struct {
	Latitude  float64
	Longitude float64
}

type Repository interface {
	GetLocationByIP(ctx context.Context, ip string) (*Domain, error)
}
