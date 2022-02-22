/*
 * @Author: Joey
 * @Date: 2022-02-05 12:02:13
 * @LastEditors: Joey
 * @LastEditTime: 2022-02-06 21:53:02
 * @FilePath: \minamazon\listing\listingclient.go
 */

package finance

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"gopkg.me/selling-partner-api-sdk/finances"
	// "github.com/JoeyFrancisTribbiani/selling-partner-api-sdk/finances"
	// "github.com/JoeyFrancisTribbiani/selling-partner-api-sdk/finances"
)

func TestGetFinanceEvents(t *testing.T) {
	GetFinanceEvents()
}
func TestParseJson(t *testing.T) {
	// testtime := "2021-07-28T13:49:14Z"

	jsonbytes, err := ioutil.ReadFile("./testjson.json")
	if err != nil {
		panic(err)
	}

	var dest finances.ListFinancialEventGroupsResponse
	if err := json.Unmarshal(jsonbytes, &dest); err != nil {
		fmt.Println(err)
	}
}
