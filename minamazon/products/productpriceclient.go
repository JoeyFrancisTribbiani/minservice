/*
 * @Author: Joey
 * @Date: 2022-02-05 12:02:13
 * @LastEditors: Joey
 * @LastEditTime: 2022-02-06 21:53:02
 * @FilePath: \minamazon\listing\listingclient.go
 */
package products

import (
	"context"
	"fmt"
	"minamazon/common"

	"github.com/JoeyFrancisTribbiani/selling-partner-api-sdk/productPricing"
)

func GetProductsPrice() {
	productPricingClient, err := productPricing.NewClientWithResponses(common.Endpoint,
		productPricing.WithRequestBefore(common.FnRequestBefore),
		productPricing.WithResponseAfter(common.FnResponseAfter),
	)
	if err != nil {
		fmt.Println(err, "err is not nil")
	}

	ctx := context.Background()
	// // var size int = 10
	// // var types []string = []string{"GET_MERCHANT_LISTINGS_ALL_DATA"}
	// // var status []string = []string{"DONE"}
	var marketplaceIds []string = []string{"ATVPDKIKX0DER"}
	var asins []string = []string{"B08KGSXPZS", "B08PS5R74W", "B08VN9PS4G"}
	// // m, _ := time.ParseDuration("-1M")
	// var since = time.Now().AddDate(0, -12, -21).UTC()
	// var until = time.Now().AddDate(0, -10, 0).UTC()

	// var nextToken = ""
	// var skus []string = []string{"6D-5N0Z-CDTM", "8B-ZXJ8-GKUG", "CE-TM56-FB28", "UJ-H6NT-K02S"}
	// var sellerId = "ASVGM7S7XZ793"

	productPricingClient.GetPricingWithResponse(ctx, &productPricing.GetPricingParams{
		MarketplaceId: marketplaceIds[0],
		// A list of up to twenty Amazon Standard Identification Number (ASIN) values used to identify items in the given marketplace.
		Asins: &asins,
		// A list of up to twenty seller SKU values used to identify items in the given marketplace.
		// Skus: &skus,
		// Indicates whether ASIN values or seller SKU values are used to identify items. If you specify Asin, the information in the response will be dependent on the list of Asins you provide in the Asins parameter. If you specify Sku, the information in the response will be dependent on the list of Skus you provide in the Skus parameter.
		ItemType: "Asin",
		// Filters the offer listings based on item condition. Possible values: New, Used, Collectible, Refurbished, Club.
		ItemCondition: nil,
	})

	if err != nil {
		// panic(err)
		fmt.Println(err, "err is not nil")
	}

}
