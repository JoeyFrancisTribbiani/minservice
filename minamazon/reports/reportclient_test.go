/*
 * @Author: Joey
 * @Date: 2022-01-01 23:48:38
 * @LastEditors: Joey
 * @LastEditTime: 2022-02-06 21:27:05
 * @FilePath: \minamazon\reports\client_test.go
 */
package reports

import (
	"testing"
)

func TestGetReports(t *testing.T) {
	GetReports()
}

func TestCreateReports(t *testing.T) {
	CreateReports()
}

func TestGetReportDocument(t *testing.T) {
	GetReportDocument()
}

func TestParseReportDocument(t *testing.T) {
	// ParseReportDocument("3_amzn1.spdoc.1.3.5663339c-6bc8-49a1-86b2-b8c04009d9a9.T3CXLA26QNMVF5.47700_GET_MERCHANT_LISTINGS_ALL_DATA_report.txt")
}
