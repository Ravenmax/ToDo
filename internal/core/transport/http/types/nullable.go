package core_http_types

import (
	"encoding/json"

	"github.com/Ravenmax/ToDo/internal/core/domain"
)

type Nullable[T any] struct {
	domain.Nullable[T]
}

/*
Nullable
-fields not provided {Set=false}
-fields provided: value{Set=true, Value != null}
-fileds proveded: null {Set=True, Value = null}
*/

func (n *Nullable[T]) UnmarshalJSON(b []byte) error {
	n.Set = true
	if string(b) == "null" {
		n.Value = nil
		return nil
	}
	var value T
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}
	n.Value = &value
	return nil
}

func (n *Nullable[T]) ToDomain() domain.Nullable[T] {
	return domain.Nullable[T]{
		Value: n.Value,
		Set:   n.Set,
	}
}
