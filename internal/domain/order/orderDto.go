package order

import (
	"encoding/json"

	"factory-pattern/internal/errors"
)

type OrderDto struct {
	Id        string `json:"id"`
	Pizza     string `json:"pizza"`
	OrderedAt string `json:"orderedAt"`
}

func NewOrderDtoFromBytes(b []byte) (*OrderDto, error) {
	var orderDto *OrderDto

	err := json.Unmarshal(b, &orderDto)
	if err != nil {
		return nil, errors.WithStack("failed to initialize order from bytes", err)
	}

	return orderDto, nil
}
