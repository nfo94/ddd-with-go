package recommendation

import (
	"context"
	"time"

	"github.com/Rhymond/go-money"
)

type Recommendation struct {
	TripStart time.Time
	TripEnd   time.Time
	HotelName string
	Location  string
	TripPrice money.Money
}

type Option struct {
	HotelName     string
	Location      string
	PricePernight money.Money
}

type AvailabilityGetter interface {
	GetAvailability(ctx context.Context, tripStart time.Time, tripStart time.Time, location string) ([]Option, error)
}
