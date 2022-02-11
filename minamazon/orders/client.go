/*
 * @Author: Joey
 * @Date: 2022-01-01 23:48:38
 * @LastEditors: Joey
 * @LastEditTime: 2022-02-06 20:48:13
 * @FilePath: \minamazon\orders\client.go
 */
package orders

import (
	"context"
	"fmt"
	"time"

	"minamazon/common"

	"github.com/JoeyFrancisTribbiani/selling-partner-api-sdk/ordersV0"
)

func GetOrders() error {
	order, err := ordersV0.NewClientWithResponses(common.Endpoint,
		ordersV0.WithRequestBefore(common.FnRequestBefore),

		ordersV0.WithResponseAfter(common.FnResponseAfter),
	)

	if err != nil {
		// panic(err)
		fmt.Println(err, "err is not nil")
	}

	ctx := context.Background()
	var size int = 10
	// var types []string = []string{"GET_MERCHANT_LISTINGS_ALL_DATA"}
	// var status []string = []string{"DONE"}
	var marketplaceIds []string = []string{"ATVPDKIKX0DER"}
	// m, _ := time.ParseDuration("-1M")
	var since = time.Now().AddDate(0, -3, -21).UTC().Format(time.RFC3339)
	var until = time.Now().AddDate(0, -1, 0).UTC().Format(time.RFC3339)
	var fuck = false

	// var nextToken = ""

	order.GetOrdersWithResponse(ctx, &ordersV0.GetOrdersParams{
		&since,
		&until,
		nil,
		nil,
		nil,
		marketplaceIds,
		nil,
		nil,
		nil,
		nil,
		&size,
		nil,
		nil,
		nil,
		nil,
		&fuck,
		nil,
	})

	if err != nil {
		// panic(err)
		fmt.Println(err, "err is not nil")
	}

	return nil
}
