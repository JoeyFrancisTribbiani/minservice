/*
 * @Author: Joey
 * @Date: 2022-01-01 23:48:38
 * @LastEditors: Joey
 * @LastEditTime: 2022-02-06 21:26:37
 * @FilePath: \minamazon\reports\reportclient.go
 */
package reports

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"minamazon/common"

	"github.com/JoeyFrancisTribbiani/selling-partner-api-sdk/reports"
)

func CreateReports() error {
	report, err := reports.NewClientWithResponses(common.Endpoint,
		reports.WithRequestBefore(common.FnRequestBefore),

		reports.WithResponseAfter(common.FnResponseAfter),
	)

	if err != nil {
		// panic(err)
		fmt.Println(err, "err is not nil")
	}

	ctx := context.Background()
	// var size int = 10
	// var types []string = []string{"GET_MERCHANT_LISTINGS_ALL_DATA"}
	// var status []string = []string{"DONE"}
	var marketplaceIds []string = []string{"ATVPDKIKX0DER"}
	// m, _ := time.ParseDuration("-1M")
	var since = time.Now().AddDate(0, -12, -21).UTC()
	var until = time.Now().AddDate(0, -10, 0).UTC()

	// var nextToken = ""

	report.CreateReportWithResponse(ctx, reports.CreateReportJSONRequestBody{
		DataEndTime:    &until,
		DataStartTime:  &since,
		MarketplaceIds: marketplaceIds,
		ReportType:     "GET_MERCHANT_LISTINGS_ALL_DATA",
	})

	if err != nil {
		// panic(err)
		fmt.Println(err, "err is not nil")
	}

	return nil
}

func GetReports() error {
	report, err := reports.NewClientWithResponses(common.Endpoint,
		reports.WithRequestBefore(common.FnRequestBefore),

		reports.WithResponseAfter(common.FnResponseAfter),
	)

	if err != nil {
		// panic(err)
		fmt.Println(err, "err is not nil")
	}

	ctx := context.Background()
	// var size int = 10
	var types []string = []string{"GET_MERCHANT_LISTINGS_ALL_DATA"}
	// var status []string = []string{"DONE"}
	var marketplaceIds []string = []string{"ATVPDKIKX0DER"}
	// m, _ := time.ParseDuration("-1M")
	// var since = time.Now().AddDate(0, -3, -21).UTC().Format(time.RFC3339)
	// var until = time.Now().AddDate(0, -1, 0).UTC().Format(time.RFC3339)

	// var nextToken = ""

	report.GetReports(ctx, &reports.GetReportsParams{
		&types,
		nil,
		&marketplaceIds,
		nil,
		nil,
		nil,
		nil,
	})

	if err != nil {
		// panic(err)
		fmt.Println(err, "err is not nil")
	}

	return nil
}
func GetReportDocument() error {
	report, err := reports.NewClientWithResponses(common.Endpoint,
		reports.WithRequestBefore(common.FnRequestBefore),

		reports.WithResponseAfter(common.FnResponseAfter),
	)

	if err != nil {
		// panic(err)
		fmt.Println(err, "err is not nil")
	}

	ctx := context.Background()
	// var size int = 10
	// var types []string = []string{"GET_MERCHANT_LISTINGS_ALL_DATA"}
	// var status []string = []string{"DONE"}
	// var marketplaceIds []string = []string{"ATVPDKIKX0DER"}
	// m, _ := time.ParseDuration("-1M")
	// var since = time.Now().AddDate(0, -3, -21).UTC().Format(time.RFC3339)
	// var until = time.Now().AddDate(0, -1, 0).UTC().Format(time.RFC3339)

	// var nextToken = ""
	// var reportType = "GET_MERCHANT_LISTINGS_ALL_DATA"
	var reportDocumentId = "amzn1.spdoc.1.3.5663339c-6bc8-49a1-86b2-b8c04009d9a9.T3CXLA26QNMVF5.47700"
	resp, err := report.GetReportDocument(ctx, reportDocumentId)
	if err != nil {
		log.Fatal("Failed to make GetReportDocument request: ", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.Fatal("Service returned a status that isn't 2xx: ", resp.StatusCode)
	}
	defer resp.Body.Close()
	var bodyData []byte
	resp.Body.Read(bodyData)
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read GetReportDocument response body: ", err)
	}
	var reportDocumentResp reports.GetReportDocumentResp
	err = json.Unmarshal(content, &reportDocumentResp.JSON200)
	if err != nil {
		log.Fatal("Failed to unmarshall GetReportDocument response: ", err)
	}
	// if reportDocumentResp.StatusCode() != 200 {
	// 	log.Fatal("Got errors in the GetReportDocument response: ", reportDocumentResp.StatusCode())
	// }
	respWithDocContent, err := http.Get(reportDocumentResp.JSON200.Url)
	if err != nil {
		log.Fatal("failed to make request to "+reportDocumentResp.JSON200.Url+" : ", err)
	}
	if respWithDocContent.StatusCode < 200 || respWithDocContent.StatusCode > 299 {
		log.Fatal("Service returned a status that isn't 2xx: ", respWithDocContent.StatusCode)
	}
	defer respWithDocContent.Body.Close()

	// reportContentBytes, err := ioutil.ReadAll(respWithDocContent.Body)
	// if err != nil {
	// 	log.Fatal("Failed to read Report content body: ", err)
	// }
	// ioutil.WriteFile("./3_encrypted_"+reportDocumentId+"_"+reportType+"_report.txt", reportContentBytes, 0777)
	// decryptedFilecontent, err := decryption.Decrypt(*reportDocumentResp.JSON200.CompressionAlgorithm.EncryptionDetails.Key, reportDocumentResp.Payload.EncryptionDetails.InitializationVector, reportContentBytes)
	// if err != nil {
	// 	log.Fatal("Failed to decrypt file content: ", err)
	// }
	// documentFile := "./3_" + reportDocumentId + "_" + reportType + "_report.txt"
	// ioutil.WriteFile(documentFile, decryptedFilecontent, 0777)
	// if err != nil {
	// 	// panic(err)
	// 	fmt.Println(err, "err is not nil")
	// }
	ParseReportDocument(respWithDocContent.Body)

	return nil
}

func ParseReportDocument(content io.ReadCloser) error {
	// csvfile, err := os.Open(filepath)
	// if err != nil {
	// 	log.Fatalln("Couldn't open the csv file", err)
	// }
	// defer csvfile.Close()

	// Parse the file
	r := csv.NewReader(content)
	r.LazyQuotes = true
	r.Comma = '\t'

	// Iterate through the records
	header, _ := r.Read()
	headers := strings.Split(header[0], "\t")
	fmt.Printf("header has %d columns:%s\n", len(headers), headers)
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Record has %d columns.\n", len(record))
		cols := strings.Split(record[0], "\t")
		fmt.Printf("line has %d columns:%s\n", len(cols), cols)
	}
	return nil
}
