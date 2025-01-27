package chapter3

import (
	"time"

	"github.com/Rhymond/go-money"
)

type Auction struct {
	ID             int
	startingtPrice money.Money
	sellerID       int
	createdAt      time.Time
	auctionStart   time.Time
	auctionEnd     time.Time
}
