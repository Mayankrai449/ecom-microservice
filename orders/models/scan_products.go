package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type ProductItem struct {
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

type ProductItems []ProductItem

func (p *ProductItems) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to unmarshal JSONB value")
	}

	var items []ProductItem
	err := json.Unmarshal(bytes, &items)
	*p = ProductItems(items)
	return err
}

func (p ProductItems) Value() (driver.Value, error) {
	return json.Marshal(p)
}
