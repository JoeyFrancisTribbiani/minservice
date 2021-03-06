/*
 * @Author: Joey
 * @Date: 2022-02-05 12:02:13
 * @LastEditors: Joey
 * @LastEditTime: 2022-02-06 21:53:02
 * @FilePath: \minamazon\listing\listingclient.go
 */

package finance

import (
	"context"
	"fmt"
	"minamazon/common"
	"time"

	"github.com/JoeyFrancisTribbiani/minerpserver/model/autocode"
	service "github.com/JoeyFrancisTribbiani/minerpserver/service/autocode"
	"github.com/JoeyFrancisTribbiani/selling-partner-api-sdk/finances"
)

func GetFinanceEvents() {

	finance, err := finances.NewClientWithResponses(common.Endpoint,
		finances.WithRequestBefore(common.FnRequestBefore),
		finances.WithResponseAfter(common.FnResponseAfter))
	if err != nil {
		fmt.Println(err, "err is not nil")
	}
	ctx := context.Background()

	// // var size int = 10
	// // var types []string = []string{"GET_MERCHANT_LISTINGS_ALL_DATA"}
	// // var status []string = []string{"DONE"}
	// var marketplaceIds []string = []string{"ATVPDKIKX0DER"}
	// // m, _ := time.ParseDuration("-1M")
	// 不得晚于当前时间前两分钟
	var since = time.Now().AddDate(0, -3, 0).UTC()
	var until = time.Now().Add(-time.Minute * 5).UTC()

	// var nextToken = ""
	// var sku = "6D-5N0Z-CDTM"
	// var sellerId = "ASVGM7S7XZ793"
	// var includedData = []string{"attributes", "issues", "offers", "fulfillmentAvailability"}

	resp, err := finance.ListFinancialEventGroupsWithResponse(ctx, &finances.ListFinancialEventGroupsParams{
		FinancialEventGroupStartedBefore: &until,
		FinancialEventGroupStartedAfter:  &since,
	})

	if err != nil {
		fmt.Println(err, "err is not nil")
		fmt.Println(resp.Body)
	}

	jsonmodel := resp.JSON200
	eventGroupList := *jsonmodel.Payload.FinancialEventGroupList
	for _, group := range eventGroupList {
		fmt.Print(group)
	}

}

func SaveListing(listing *autocode.ErpListingDetail) {
	service := &service.ErpListingDetailService{}
	err := service.CreateErpListingDetail(*listing)
	if err != nil {
		panic("error when create listing")
	}
}
