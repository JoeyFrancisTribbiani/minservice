/*
 * @Author: Joey
 * @Date: 2022-01-01 23:48:32
 * @LastEditors: Joey
 * @LastEditTime: 2022-02-06 20:58:34
 * @FilePath: \minamazon\common\spconfig.go
 */
package common

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"

	sp "github.com/JoeyFrancisTribbiani/selling-partner-api-sdk/pkg/selling-partner"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	SellingPartner, _ = sp.NewSellingPartner(&sp.Config{
		ClientID:     "amzn1.application-oa2-client.0c73f7463a85441f9ea68bbcb59702af",
		ClientSecret: "2bc5c7bed95d3891bf7b7562c7eb8497c307f8beebecc628614bfcde0d81a3b5",
		// RefreshToken: "Atzr|IwEBIFgMrY3c7wGNme00P-GXaiH1LEIkKrjl5pYKACuKjsC7qRqQefKlnv5gaTv9sXfM_-eXpIQBIZiqgh8f7wer0cZVPxqcGyrpAJVPj5Zi_WU4PSp3vVjNks9C47HIxGES94YDGrOTVXev-KtKkEsE-laUBsswyv3oFAO02G017HUHVf2V_jQo_Vw_fibbS4XL0FfzTlvKUKZzOe34A28MxyL5FVuoZEiOftrPLMs0BMyqD3bp03Q3Ex8zG1JamggJJRH-kFAz5VfCBQnTR93ifsYRiE2KFCDquau3T4K_OYrAX5wW_BJBjU01IvjqE_2A2rk",
		RefreshToken: "Atzr|IwEBIDLdczknrBmE9Bh4TLUARxlMcdFCR1zW79YtDPt006uP7a9SXCTfZa8M80OXoxxGaIp0hn-BjN5oC09ZDlg-SweQ9RnsUOf1nqqsgrVPHFZYD1esdrfYd6ymMzkXv4XaCNI97C0MH_ZiuUV03FdjGxrR85gLNYB_cW8_K8AknXE-n1Z9fKAQnXeReMzFSgbTZV8GBT2IlcHUyOsJjRDSpkX_CZb_GdoHj1FdBZC6HtWrBJd36UVdgT-eTNFpmi9-dYFsXQtPlkGkElGaFgdm8viAXXaZoUA8FGBh-1a_3gEMk1L-AbgBTqjfV6grZI9t-5w",
		AccessKeyID:  "AKIAVCZHDCZBJL6PSYNE",
		SecretKey:    "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",
		Region:       "us-east-1",
		RoleArn:      "arn:aws:iam::349584954946:role/erp_role",
	})
	// SellingPartner, _ = sp.NewSellingPartner(&sp.Config{
	// ClientID:     "",
	// ClientSecret: "",
	// // RefreshToken: "Atzr|IwEBIFgMrY3c7wGNme00P-GXaiH1LEIkKrjl5pYKACuKjsC7qRqQefKlnv5gaTv9sXfM_-eXpIQBIZiqgh8f7wer0cZVPxqcGyrpAJVPj5Zi_WU4PSp3vVjNks9C47HIxGES94YDGrOTVXev-KtKkEsE-laUBsswyv3oFAO02G017HUHVf2V_jQo_Vw_fibbS4XL0FfzTlvKUKZzOe34A28MxyL5FVuoZEiOftrPLMs0BMyqD3bp03Q3Ex8zG1JamggJJRH-kFAz5VfCBQnTR93ifsYRiE2KFCDquau3T4K_OYrAX5wW_BJBjU01IvjqE_2A2rk",
	// RefreshToken: "",
	// AccessKeyID:  "",
	// SecretKey:    "",
	// Region:       "",
	// RoleArn:      "",
	// })

	Endpoint = "https://sellingpartnerapi-na.amazon.com"

	FnRequestBefore = func(ctx context.Context, req *http.Request) error {
		req.Header.Add("X-Amzn-Requestid", uuid.New().String())
		err := SellingPartner.SignRequest(req)
		if err != nil {
			return errors.Wrap(err, "sign error")
		}

		dump, err := httputil.DumpRequest(req, true)
		if err != nil {
			return errors.Wrap(err, "DumpRequest Error")
		}
		log.Printf("DumpRequest = %s", dump)
		return nil
	}

	FnResponseAfter = func(ctx context.Context, rsp *http.Response) error {
		dump, err := httputil.DumpResponse(rsp, true)
		if err != nil {
			return errors.Wrap(err, "DumpResponse Error")
		}
		log.Printf("DumpResponse = %s", dump)

		return nil
	}
)
