package models

import (
	"fmt"
)

type BoolInt bool

func (b *BoolInt) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "1", "true":
		*b = true
	case "0", "false":
		*b = false
	default:
		return fmt.Errorf("invalid value for BoolInt: %s", data)
	}
	return nil
}
