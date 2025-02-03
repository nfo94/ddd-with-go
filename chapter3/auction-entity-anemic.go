package chapter3

import (
	"time"

	"github.com/Rhymond/go-money"
)

type AnemicAuction struct {
	ID             int
	startingtPrice money.Money
	sellerID       int
	createdAt      time.Time
	auctionStart   time.Time
	auctionEnd     time.Time
}

func (a *AnemicAuction) GetID() int {
	return a.ID
}

func (a *AnemicAuction) StartingPrice() money.Money {
	return a.startingtPrice
}

func (a *AnemicAuction) SetStartingPrice(startingtPrice money.Money) {
	a.startingtPrice = startingtPrice
}

func (a *AnemicAuction) GetSellerID() int {
	return a.sellerID
}

func (a *AnemicAuction) SetSellerID(sellerID int) {
	a.sellerID = sellerID
}

func (a *AnemicAuction) GetCreatedAt() time.Time {
	return a.createdAt
}

func (a *AnemicAuction) SetCreatedAt(createdAt time.Time) {
	a.createdAt = createdAt
}

func (a *AnemicAuction) GetAuctionStart() time.Time {
	return a.auctionStart
}

func (a *AnemicAuction) SetAuctionStart(auctionStart time.Time) {
	a.auctionStart = auctionStart
}

func (a *AnemicAuction) GetAuctionEnd() time.Time {
	return a.auctionEnd
}

func (a *AnemicAuction) SetAuctionEnd(auctionEnd time.Time) {
	a.auctionEnd = auctionEnd
}
