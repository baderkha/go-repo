package gorepo

import (
	"fmt"
	"strings"
)

type FieldInput map[string]any

func (i FieldInput) With(field string, val any) *FieldInput {
	i[field] = val
	return &i
}

func (i FieldInput) MustHave(keys ...string) error {
	var keysNotFound []string
	for _, k := range keys {
		_, ok := i[k]
		if !ok {
			keysNotFound = append(keysNotFound, k)
		}
	}

	return fmt.Errorf(
		"gorepo index error , you're missing the following keys (%s)",
		strings.Join(keysNotFound, ","),
	)
}

func FLD() *FieldInput {
	var in FieldInput = make(FieldInput)
	return &in
}
