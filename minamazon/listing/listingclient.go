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
	"log"
	"minamazon/common"
	"net/http"
	"net/http/httputil"
	"strconv"

	"github.com/JoeyFrancisTribbiani/minerpserver/model/autocode"
	"github.com/JoeyFrancisTribbiani/selling-partner-api-sdk/listings"
	"github.com/pkg/errors"
)

func GetListings(model *autocode.ErpListingDetail) *autocode.ErpListingDetail {
	listingClient, err := listings.NewClientWithResponses(common.Endpoint,
		listings.WithRequestBefore(common.FnRequestBefore),
		listings.WithResponseAfter(SaveListing),
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
	var includedData = []string{"attributes", "issues", "offers", "fulfillmentAvailability"}

	resp, err := listingClient.GetListingsItemWithResponse(ctx, sellerId, sku, &listings.GetListingsItemParams{
		MarketplaceIds: marketplaceIds,
		IncludedData:   &includedData,
	})

	if err != nil {
		fmt.Println(err, "err is not nil")
	}

	jsonmodel := resp.JSON200
	model.Asin = jsonmodel.Attributes.AdditionalProperties[""].(string)
	model.Currency = (*jsonmodel.Offers)[0].Price.CurrencyCode
	price, err := strconv.ParseFloat(string((*jsonmodel.Offers)[0].Price.Amount), 64)
	if err != nil {
		fmt.Println(err, "err is not nil")
	}
	model.ListingPrice = &price
	model.Quantity = (*jsonmodel.FulfillmentAvailability)[0].Quantity
	model.FulfillmentChannel = (*jsonmodel.FulfillmentAvailability)[0].FulfillmentChannelCode

	return model
}

func SaveListing(ctx context.Context, rsp *http.Response) error {
	dump, err := httputil.DumpResponse(rsp, true)
	if err != nil {
		return errors.Wrap(err, "DumpResponse Error")
	}
	log.Printf("DumpResponse = %s", dump)

	return nil
}
