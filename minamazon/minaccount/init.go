package minaccount

import (
	"minamazon/common"
	"minamazon/listing"

	"github.com/JoeyFrancisTribbiani/minerpserver/model/autocode"
	aotucode "github.com/JoeyFrancisTribbiani/minerpserver/model/autocode"
	sp "github.com/JoeyFrancisTribbiani/selling-partner-api-sdk/pkg/selling-partner"
)

func InitAccount(config *aotucode.SpConfig) error {
	spmodel, _ := sp.NewSellingPartner(&sp.Config{
		ClientID:     config.ClientId,
		ClientSecret: config.ClientSecret,
		// RefreshToken: "Atzr|IwEBIFgMrY3c7wGNme00P-GXaiH1LEIkKrjl5pYKACuKjsC7qRqQefKlnv5gaTv9sXfM_-eXpIQBIZiqgh8f7wer0cZVPxqcGyrpAJVPj5Zi_WU4PSp3vVjNks9C47HIxGES94YDGrOTVXev-KtKkEsE-laUBsswyv3oFAO02G017HUHVf2V_jQo_Vw_fibbS4XL0FfzTlvKUKZzOe34A28MxyL5FVuoZEiOftrPLMs0BMyqD3bp03Q3Ex8zG1JamggJJRH-kFAz5VfCBQnTR93ifsYRiE2KFCDquau3T4K_OYrAX5wW_BJBjU01IvjqE_2A2rk",
		RefreshToken: config.RefreshToken,
		AccessKeyID:  config.AccessKeyId,
		SecretKey:    config.SecretKey,
		Region:       config.Region,
		RoleArn:      config.RoleArn,
	})
	common.SellingPartner = spmodel
	listingModel := &autocode.ErpListingDetail{}
	listingModel = listing.GetListings(listingModel)
	return nil
}
