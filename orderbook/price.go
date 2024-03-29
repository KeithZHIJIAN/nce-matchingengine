package orderbook

import (
	"github.com/shopspring/decimal"
)

type Price struct {
	price decimal.Decimal
	isBuy bool
}

func NewPrice(price decimal.Decimal, isBuy bool) *Price {
	return &Price{price: price, isBuy: isBuy}
}

func (p Price) Cmp(rhs Price) int {
	if p.price.Equal(rhs.price) {
		return 0
	}
	if p.isBuy {
		return rhs.price.Cmp(p.price)
	} else {
		if p.price.Equal(decimal.Zero) {
			return -1
		} else if rhs.price.Equal(decimal.Zero) {
			return 1
		} else {
			return p.price.Cmp(rhs.price)
		}
	}
}

func (p Price) Match(rhs decimal.Decimal) bool {
	// Match when buy side is greater than or equal to sell side
	if p.price.Equal(decimal.Zero) {
		return true
	}
	if rhs.Equal(decimal.Zero) {
		return true
	}
	if p.isBuy {
		return p.price.GreaterThanOrEqual(rhs)
	} else {
		return p.price.LessThanOrEqual(rhs)
	}
}

func (p Price) String() string {
	return p.price.String()
}
