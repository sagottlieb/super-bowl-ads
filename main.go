package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

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

	file, err := os.Create("2021.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	csvwriter := csv.NewWriter(file)

	for idx, ad := range results {
		csvwriter.Write([]string{"2021", ad.Brand, ad.Title, strconv.Itoa(idx + 1), ad.Score, ad.Quarter, ad.Link})
	}

	csvwriter.Flush()
	file.Close()
}
