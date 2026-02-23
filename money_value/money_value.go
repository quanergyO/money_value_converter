package moneyvalue

import (
	"errors"

	investapi "opensource.tbank.ru/invest/invest-go/proto"
)

const (
	RUB = "rub"
)

type Money struct {
	Units int64
}

func New(dto *investapi.MoneyValue) (Money, error) {
	const mult = 1000000000
	result := Money{Units: 0}
	if dto == nil {
		return result, errors.New("empty dto")
	}

	if dto.Currency != RUB {
		return result, errors.New("unsuported currency")
	}

	result.Units += dto.Units*mult + int64(dto.Nano)

	return result, nil
}

func (m *Money) Add(rhs *Money) {
	if rhs == nil {
		return
	}

	m.Units += rhs.Units
}

func (m *Money) Sub(rhs *Money) {
	if rhs == nil {
		return
	}

	m.Units -= rhs.Units
}
