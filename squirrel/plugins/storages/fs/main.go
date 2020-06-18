package main

import (
//	"fmt"
	"errors"
)

import (
	"github.com/mysheep/squirrel/types"
)

type LoadStore string

/*
type Loader interface {
	Load(s string) (*types.Cell, error)
}

type Storer interface {
	Store(s string, e *types.Cell) error
}
*/

func (ls LoadStore) Load(location string) (*types.Cell, error) {
	return nil, errors.New("Not implemented")
}

func (ls LoadStore) Store(location string, c *types.Cell) error {
	return errors.New("Not implemented")
}