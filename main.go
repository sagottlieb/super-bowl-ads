package main

import (
	"context"
	"fmt"
	"os"

	"github.com/MontFerret/ferret/pkg/compiler"
	"github.com/MontFerret/ferret/pkg/drivers"
	"github.com/MontFerret/ferret/pkg/drivers/cdp"
	"github.com/MontFerret/ferret/pkg/drivers/http"
)

var query = `
LET doc = DOCUMENT("https://admeter.usatoday.com/results/2021")

FOR ad IN ELEMENTS(doc, '#post-')
    LET link = ELEMENT(ad, 'a')
    RETURN {link: link.attributes.href}
`

func main() {
	comp := compiler.New()

	program, err := comp.Compile(query)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ctx := context.Background()
	ctx = drivers.WithContext(ctx, cdp.NewDriver())
	// site returns 404 with content of interest
	ctx = drivers.WithContext(ctx, http.NewDriver(http.WithAllowedHTTPCode(404)), drivers.AsDefault())

	_, err = program.Run(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}