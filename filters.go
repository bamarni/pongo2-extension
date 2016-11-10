package pongo2extension

import (
	"fmt"

	"github.com/flosch/pongo2"
)

func init() {
	pongo2.RegisterFilter("price", filterPrice)
}

func filterPrice(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	val := fmt.Sprintf("%.2f", in.Float()/100.0)

	return pongo2.AsValue(val), nil
}
