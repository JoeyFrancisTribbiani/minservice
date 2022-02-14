// 自动生成模板ErpListingDetail
package autocode

import (
	"github.com/JoeyFrancisTribbiani/minerpserver/global"
)

// ErpListingDetail 结构体
// 如果含有time.Time 请自行import time包
type ErpListingDetail struct {
	global.GVA_MODEL
	CreateBy              *int     `json:"createBy" form:"createBy" gorm:"column:create_by;comment:;size:10;"`
	UpdateBy              *int     `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:;size:10;"`
	Puid                  *int     `json:"puid" form:"puid" gorm:"column:puid;comment:2231;size:19;"`
	ShopId                *int     `json:"shopId" form:"shopId" gorm:"column:shop_id;comment:4462;size:19;"`
	MarketplaceId         string   `json:"marketplaceId" form:"marketplaceId" gorm:"column:marketplace_id;comment:ATVPDKIKX0DER;size:100;"`
	FullCid               string   `json:"fullCid" form:"fullCid" gorm:"column:full_cid;comment:;size:100;"`
	ParentId              *int     `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:4762272;size:19;"`
	DxmState              string   `json:"dxmState" form:"dxmState" gorm:"column:dxm_state;comment:online_amazon;size:100;"`
	DxmPublishState       string   `json:"dxmPublishState" form:"dxmPublishState" gorm:"column:dxm_publish_state;comment:null;size:100;"`
	ErrMsg                string   `json:"errMsg" form:"errMsg" gorm:"column:err_msg;comment:null;size:100;"`
	Asin                  string   `json:"asin" form:"asin" gorm:"column:asin;comment:B08X4MVXS7;size:100;"`
	ListingId             string   `json:"listingId" form:"listingId" gorm:"column:listing_id;comment:0222Y6BRSTD;size:100;"`
	ParentAsin            string   `json:"parentAsin" form:"parentAsin" gorm:"column:parent_asin;comment:null;size:100;"`
	ChildAsins            string   `json:"childAsins" form:"childAsins" gorm:"column:child_asins;comment:null;size:100;"`
	IsVariation           *int     `json:"isVariation" form:"isVariation" gorm:"column:is_variation;comment:0;size:19;"`
	Sku                   string   `json:"sku" form:"sku" gorm:"column:sku;comment:VQ-C5VT-Y5GH;size:100;"`
	StandardProductType   string   `json:"standardProductType" form:"standardProductType" gorm:"column:standard_product_type;comment:null;size:100;"`
	StandardProductId     string   `json:"standardProductId" form:"standardProductId" gorm:"column:standard_product_id;comment:null;size:100;"`
	Title                 string   `json:"title" form:"title" gorm:"column:title;comment:Outdoor Faucet Cover;size:100;"`
	Brand                 string   `json:"brand" form:"brand" gorm:"column:brand;comment:null;size:100;"`
	Manufacturer          string   `json:"manufacturer" form:"manufacturer" gorm:"column:manufacturer;comment:null;size:100;"`
	Description           string   `json:"description" form:"description" gorm:"column:description;comment:null;size:100;"`
	BulletPoint           string   `json:"bulletPoint" form:"bulletPoint" gorm:"column:bullet_point;comment:null;size:100;"`
	ConditionType         string   `json:"conditionType" form:"conditionType" gorm:"column:condition_type;comment:null;size:100;"`
	ConditionValue        string   `json:"conditionValue" form:"conditionValue" gorm:"column:condition_value;comment:null;size:100;"`
	StandardPrice         *float64 `json:"standardPrice" form:"standardPrice" gorm:"column:standard_price;comment:8.99;size:22;"`
	ListingPricing        *float64 `json:"listingPricing" form:"listingPricing" gorm:"column:listing_pricing;comment:8.99;size:22;"`
	SwitchFulfillmentTo   string   `json:"switchFulfillmentTo" form:"switchFulfillmentTo" gorm:"column:switch_fulfillment_to;comment:AFN;size:100;"`
	FulfillmentChannel    string   `json:"fulfillmentChannel" form:"fulfillmentChannel" gorm:"column:fulfillment_channel;comment:null;size:100;"`
	PartNumber            string   `json:"partNumber" form:"partNumber" gorm:"column:part_number;comment:null;size:100;"`
	MainImage             string   `json:"mainImage" form:"mainImage" gorm:"column:main_image;comment:https://m.media-amazon.com/images/I/51QOuKgk22L._SL75_.jpg;size:100;"`
	SaleStartDate         string   `json:"saleStartDate" form:"saleStartDate" gorm:"column:sale_start_date;comment:null;size:100;"`
	SaleEndDate           string   `json:"saleEndDate" form:"saleEndDate" gorm:"column:sale_end_date;comment:null;size:100;"`
	SaleSalePrice         string   `json:"saleSalePrice" form:"saleSalePrice" gorm:"column:sale_sale_price;comment:null;size:100;"`
	Quantity              *int     `json:"quantity" form:"quantity" gorm:"column:quantity;comment:34;size:19;"`
	OpenDate              string   `json:"openDate" form:"openDate" gorm:"column:open_date;comment:2021-02-22 04:20:10;size:100;"`
	ItemDimensions        string   `json:"itemDimensions" form:"itemDimensions" gorm:"column:item_dimensions;comment:null;size:100;"`
	PackageDimensions     string   `json:"packageDimensions" form:"packageDimensions" gorm:"column:package_dimensions;comment:null;size:100;"`
	StandardPriceCurrency string   `json:"standardPriceCurrency" form:"standardPriceCurrency" gorm:"column:standard_price_currency;comment:USD;size:100;"`
	Specifics             string   `json:"specifics" form:"specifics" gorm:"column:specifics;comment:null;size:100;"`
	VariationChildStr     string   `json:"variationChildStr" form:"variationChildStr" gorm:"column:variation_child_str;comment:null;size:100;"`
	OnlineStatus          string   `json:"onlineStatus" form:"onlineStatus" gorm:"column:online_status;comment:Active;size:100;"`
	SaleNum               *int     `json:"saleNum" form:"saleNum" gorm:"column:sale_num;comment:0;size:19;"`
	Fnsku                 string   `json:"fnsku" form:"fnsku" gorm:"column:fnsku;comment:X002TADZZ7;size:100;"`
	Warehousing           *int     `json:"warehousing" form:"warehousing" gorm:"column:warehousing;comment:0;size:19;"`
	Unsellable            *int     `json:"unsellable" form:"unsellable" gorm:"column:unsellable;comment:1;size:19;"`
	ReservedQty           *int     `json:"reservedQty" form:"reservedQty" gorm:"column:reserved_qty;comment:0;size:19;"`
	PurchaseCost          *int     `json:"purchaseCost" form:"purchaseCost" gorm:"column:purchase_cost;comment:12;size:19;"`
	HeadTripCost          *int     `json:"headTripCost" form:"headTripCost" gorm:"column:head_trip_cost;comment:999;size:19;"`
	ShipCost              string   `json:"shipCost" form:"shipCost" gorm:"column:ship_cost;comment:null;size:100;"`
	PriceFeedStatus       string   `json:"priceFeedStatus" form:"priceFeedStatus" gorm:"column:price_feed_status;comment:null;size:100;"`
	InventoryFeedStatus   string   `json:"inventoryFeedStatus" form:"inventoryFeedStatus" gorm:"column:inventory_feed_status;comment:null;size:100;"`
	UpdateStandardPrice   string   `json:"updateStandardPrice" form:"updateStandardPrice" gorm:"column:update_standard_price;comment:null;size:100;"`
	UpdateQuantity        string   `json:"updateQuantity" form:"updateQuantity" gorm:"column:update_quantity;comment:null;size:100;"`
	Rating                string   `json:"rating" form:"rating" gorm:"column:rating;comment:null;size:100;"`
	RatingCount           string   `json:"ratingCount" form:"ratingCount" gorm:"column:rating_count;comment:null;size:100;"`
	Rank                  string   `json:"rank" form:"rank" gorm:"column:rank;comment:null;size:100;"`
	BsrRank               string   `json:"bsrRank" form:"bsrRank" gorm:"column:bsr_rank;comment:null;size:100;"`
	CommodityId           *int     `json:"commodityId" form:"commodityId" gorm:"column:commodity_id;comment:470356;size:19;"`
	MathCommodityTime     string   `json:"mathCommodityTime" form:"mathCommodityTime" gorm:"column:math_commodity_time;comment:null;size:100;"`
	MaxShipmentQty        string   `json:"maxShipmentQty" form:"maxShipmentQty" gorm:"column:max_shipment_qty;comment:null;size:100;"`
	DevId                 string   `json:"devId" form:"devId" gorm:"column:dev_id;comment:9207;size:100;"`
	CreateId              string   `json:"createId" form:"createId" gorm:"column:create_id;comment:null;size:100;"`
	UpdateId              string   `json:"updateId" form:"updateId" gorm:"column:update_id;comment:null;size:100;"`
	LpriceUpdateTime      string   `json:"lpriceUpdateTime" form:"lpriceUpdateTime" gorm:"column:lprice_update_time;comment:null;size:100;"`
	LpriceStatus          string   `json:"lpriceStatus" form:"lpriceStatus" gorm:"column:lprice_status;comment:null;size:100;"`
	ChildList             string   `json:"childList" form:"childList" gorm:"column:child_list;comment:null;size:100;"`
	SalePrices            *int     `json:"salePrices" form:"salePrices" gorm:"column:sale_prices;comment:0;size:19;"`
	AdCosts               *int     `json:"adCosts" form:"adCosts" gorm:"column:ad_costs;comment:0;size:19;"`
	AdCostsSd             *int     `json:"adCostsSd" form:"adCostsSd" gorm:"column:ad_costs_sd;comment:0;size:19;"`
	AdCostsSb             *int     `json:"adCostsSb" form:"adCostsSb" gorm:"column:ad_costs_sb;comment:0;size:19;"`
	AdCostsSbv            *int     `json:"adCostsSbv" form:"adCostsSbv" gorm:"column:ad_costs_sbv;comment:0;size:19;"`
	AdCostsSum            *int     `json:"adCostsSum" form:"adCostsSum" gorm:"column:ad_costs_sum;comment:0;size:19;"`
	Currency              string   `json:"currency" form:"currency" gorm:"column:currency;comment:USD;size:100;"`
	CostCurrency          string   `json:"costCurrency" form:"costCurrency" gorm:"column:cost_currency;comment:CNY;size:100;"`
	CommoditySku          string   `json:"commoditySku" form:"commoditySku" gorm:"column:commodity_sku;comment:testsku123123;size:100;"`
	CommodityName         string   `json:"commodityName" form:"commodityName" gorm:"column:commodity_name;comment:测试商品;size:100;"`
	ShopName              string   `json:"shopName" form:"shopName" gorm:"column:shop_name;comment:A3LR0SXN8QVIE0-na-US;size:100;"`
	SiteName              string   `json:"siteName" form:"siteName" gorm:"column:site_name;comment:美国;size:100;"`
	SellingPartnerId      string   `json:"sellingPartnerId" form:"sellingPartnerId" gorm:"column:selling_partner_id;comment:null;size:100;"`
	Domain                string   `json:"domain" form:"domain" gorm:"column:domain;comment:www.amazon.com;size:100;"`
	ChildNum              string   `json:"childNum" form:"childNum" gorm:"column:child_num;comment:null;size:100;"`
	DevName               string   `json:"devName" form:"devName" gorm:"column:dev_name;comment:业务员1;size:100;"`
	VariationStrMap       string   `json:"variationStrMap" form:"variationStrMap" gorm:"column:variation_str_map;comment:null;size:100;"`
	ListingPrice          *float64 `json:"listingPrice" form:"listingPrice" gorm:"column:listing_price;comment:8.99;size:22;"`
	ShippingPrice         *int     `json:"shippingPrice" form:"shippingPrice" gorm:"column:shipping_price;comment:0;size:19;"`
	BuyBoxWinner          string   `json:"buyBoxWinner" form:"buyBoxWinner" gorm:"column:buy_box_winner;comment:true;size:100;"`
	BuyBoxCurrency        string   `json:"buyBoxCurrency" form:"buyBoxCurrency" gorm:"column:buy_box_currency;comment:USD;size:100;"`
	TotalFee              *float64 `json:"totalFee" form:"totalFee" gorm:"column:total_fee;comment:4.82;size:22;"`
	ReferralFee           *float64 `json:"referralFee" form:"referralFee" gorm:"column:referral_fee;comment:1.35;size:22;"`
	VariableClosingFee    *int     `json:"variableClosingFee" form:"variableClosingFee" gorm:"column:variable_closing_fee;comment:0;size:19;"`
	PerItemFee            *int     `json:"perItemFee" form:"perItemFee" gorm:"column:per_item_fee;comment:0;size:19;"`
	FbaFees               *float64 `json:"fbaFees" form:"fbaFees" gorm:"column:fba_fees;comment:3.47;size:22;"`
	FeeCurrency           string   `json:"feeCurrency" form:"feeCurrency" gorm:"column:fee_currency;comment:USD;size:100;"`
	IsPutAsh              string   `json:"isPutAsh" form:"isPutAsh" gorm:"column:is_put_ash;comment:null;size:100;"`
	BsrRankStr            string   `json:"bsrRankStr" form:"bsrRankStr" gorm:"column:bsr_rank_str;comment:null;size:100;"`
	VartationStr          string   `json:"vartationStr" form:"vartationStr" gorm:"column:vartation_str;comment:;size:100;"`
}

// TableName ErpListingDetail 表名
func (ErpListingDetail) TableName() string {
	return "erp_listing_detail"
}
