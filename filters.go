package pongo2extension

import (
	"fmt"
	"regexp"

	"github.com/flosch/pongo2"
)

var (
	DefaultLocale = "fr_FR"

	phoneRgx = map[string]*regexp.Regexp{
		"fr_FR": regexp.MustCompile("^\\+33([1-9])([0-9]{2})([0-9]{2})([0-9]{2})([0-9]{2})$"),
	}

	phoneFmt = map[string]string{
		"fr_FR": "0$1 $2 $3 $4 $5",
	}
)

func init() {
	pongo2.RegisterFilter("phone", filterPhone)
	pongo2.RegisterFilter("price", filterPrice)
}

func filterPhone(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	// TODO: allow to override the default locale with a filter param
	val := phoneRgx[DefaultLocale].ReplaceAllString(
		in.String(),
		phoneFmt[DefaultLocale],
	)

	return pongo2.AsValue(val), nil
}

func filterPrice(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	val := fmt.Sprintf("%.2f", in.Float()/100.0)

	return pongo2.AsValue(val), nil
}
