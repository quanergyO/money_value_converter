package moneyvalue

import (
	"errors"

	investapi "opensource.tbank.ru/invest/invest-go/proto"
)

const (
	RUB = "rub"
)

type Money struct {
	units int64
}

func New(dto *investapi.MoneyValue) (Money, error) {
	const mult = 1000000000
	result := Money{units: 0}
	if dto == nil {
		return result, errors.New("empty dto")
	}

	if dto.Currency != RUB {
		return result, errors.New("unsuported currency")
	}

	result.units += dto.Units*mult + int64(dto.Nano)

	return result, nil
}
