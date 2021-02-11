package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/MontFerret/ferret/pkg/compiler"
	"github.com/MontFerret/ferret/pkg/drivers"
	"github.com/MontFerret/ferret/pkg/drivers/cdp"
	"github.com/MontFerret/ferret/pkg/drivers/http"
)

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

	out, err := program.Run(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var results []*ad
	err = json.Unmarshal(out, &results)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for idx, ad := range results {
		fmt.Printf("#%d, %s, %s, %s, %s, %s\n", idx+1, ad.Brand, ad.Quarter, ad.Title, ad.Score, ad.Link)
	}
}
