/*
 * @Author: Joey
 * @Date: 2022-02-05 12:02:13
 * @LastEditors: Joey
 * @LastEditTime: 2022-02-06 21:53:02
 * @FilePath: \minamazon\listing\listingclient.go
 */
package listing

import (
	"context"
	"fmt"
	"minamazon/common"

	"github.com/JoeyFrancisTribbiani/selling-partner-api-sdk/listings"
)

func GetListings() {
	listingClient, err := listings.NewClientWithResponses(common.Endpoint,
		listings.WithRequestBefore(common.FnRequestBefore),
		listings.WithResponseAfter(common.FnResponseAfter),
	)

	if err != nil {
		fmt.Println(err, "err is not nil")
	}

	ctx := context.Background()
	// // var size int = 10
	// // var types []string = []string{"GET_MERCHANT_LISTINGS_ALL_DATA"}
	// // var status []string = []string{"DONE"}
	var marketplaceIds []string = []string{"ATVPDKIKX0DER"}
	// // m, _ := time.ParseDuration("-1M")
	// var since = time.Now().AddDate(0, -12, -21).UTC()
	// var until = time.Now().AddDate(0, -10, 0).UTC()

	// var nextToken = ""
	var sku = "6D-5N0Z-CDTM"
	var sellerId = "ASVGM7S7XZ793"

	listingClient.GetListingsItemWithResponse(ctx, sellerId, sku, &listings.GetListingsItemParams{
		MarketplaceIds: marketplaceIds,
	})

	if err != nil {
		// panic(err)
		fmt.Println(err, "err is not nil")
	}

}
