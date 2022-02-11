// 自动生成模板ErpFinanceAsinSummary
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ErpFinanceAsinSummary 结构体
// 如果含有time.Time 请自行import time包
type ErpFinanceAsinSummary struct {
      global.GVA_MODEL
      AdjCompensation  string `json:"adjCompensation" form:"adjCompensation" gorm:"column:adj_compensation;comment:库存调整赔偿;size:100;"`
      AdjCompensationNum  *int `json:"adjCompensationNum" form:"adjCompensationNum" gorm:"column:adj_compensation_num;comment:库存调整数量;size:19;"`
      AdjCompensationPcFee  string `json:"adjCompensationPcFee" form:"adjCompensationPcFee" gorm:"column:adj_compensation_pc_fee;comment:FBA库存赔偿;size:100;"`
      Asin  string `json:"asin" form:"asin" gorm:"column:asin;comment:ASIN;size:100;"`
      AsinList  string `json:"asinList" form:"asinList" gorm:"column:asin_list;comment:ASIN列表;size:100;"`
      AsinOtherFee  string `json:"asinOtherFee" form:"asinOtherFee" gorm:"column:asin_other_fee;comment:ASIN其他费;size:100;"`
      AsinOtherFees  string `json:"asinOtherFees" form:"asinOtherFees" gorm:"column:asin_other_fees;comment:ASIN其他费;size:100;"`
      BatchPurchaseApcFee  string `json:"batchPurchaseApcFee" form:"batchPurchaseApcFee" gorm:"column:batch_purchase_apc_fee;comment:批次采购成本（赔偿）;size:100;"`
      BatchPurchaseCrpcFee  string `json:"batchPurchaseCrpcFee" form:"batchPurchaseCrpcFee" gorm:"column:batch_purchase_crpc_fee;comment:批次采购成本（退货）;size:100;"`
      BatchPurchaseDfpcFee  string `json:"batchPurchaseDfpcFee" form:"batchPurchaseDfpcFee" gorm:"column:batch_purchase_dfpc_fee;comment:批次采购成本（差异）;size:100;"`
      BatchPurchasePcFee  string `json:"batchPurchasePcFee" form:"batchPurchasePcFee" gorm:"column:batch_purchase_pc_fee;comment:批次采购成本;size:100;"`
      BatchPurchaseVrpcFee  string `json:"batchPurchaseVrpcFee" form:"batchPurchaseVrpcFee" gorm:"column:batch_purchase_vrpc_fee;comment:批次采购成本（供应商退货）;size:100;"`
      BatchTripApcFee  string `json:"batchTripApcFee" form:"batchTripApcFee" gorm:"column:batch_trip_apc_fee;comment:批次头程费用（赔偿）;size:100;"`
      BatchTripCrpcFee  string `json:"batchTripCrpcFee" form:"batchTripCrpcFee" gorm:"column:batch_trip_crpc_fee;comment:批次头程费用（退货）;size:100;"`
      BatchTripDfpcFee  string `json:"batchTripDfpcFee" form:"batchTripDfpcFee" gorm:"column:batch_trip_dfpc_fee;comment:批次头程费用（差异）;size:100;"`
      BatchTripPcFee  string `json:"batchTripPcFee" form:"batchTripPcFee" gorm:"column:batch_trip_pc_fee;comment:批次头程费用;size:100;"`
      BatchTripVrpcFee  string `json:"batchTripVrpcFee" form:"batchTripVrpcFee" gorm:"column:batch_trip_vrpc_fee;comment:批次头程费用（供应商退货）;size:100;"`
      CommodityList  string `json:"commodityList" form:"commodityList" gorm:"column:commodity_list;comment:商品列表;size:100;"`
      CommodityName  string `json:"commodityName" form:"commodityName" gorm:"column:commodity_name;comment:商品名称;size:100;"`
      CommoditySku  string `json:"commoditySku" form:"commoditySku" gorm:"column:commodity_sku;comment:商品SKU;size:100;"`
      CouponsPromoteFee  string `json:"couponsPromoteFee" form:"couponsPromoteFee" gorm:"column:coupons_promote_fee;comment:优惠劵;size:100;"`
      CpcCostDiff  string `json:"cpcCostDiff" form:"cpcCostDiff" gorm:"column:cpc_cost_diff;comment:CPC广告费差异;size:100;"`
      CpcSales  string `json:"cpcSales" form:"cpcSales" gorm:"column:cpc_sales;comment:CPC销售额;size:100;"`
      CpcSalesNum  *int `json:"cpcSalesNum" form:"cpcSalesNum" gorm:"column:cpc_sales_num;comment:CPC销量;size:19;"`
      CpcSdCost  string `json:"cpcSdCost" form:"cpcSdCost" gorm:"column:cpc_sd_cost;comment:CPC广告成本;size:100;"`
      CpcSdSalesNum  *int `json:"cpcSdSalesNum" form:"cpcSdSalesNum" gorm:"column:cpc_sd_sales_num;comment:CPC广告销量;size:19;"`
      CpcSpCost  string `json:"cpcSpCost" form:"cpcSpCost" gorm:"column:cpc_sp_cost;comment:CPC推广费;size:100;"`
      CpcSpSalesNum  *int `json:"cpcSpSalesNum" form:"cpcSpSalesNum" gorm:"column:cpc_sp_sales_num;comment:CPC推广销量;size:19;"`
      CreateBy  *int `json:"createBy" form:"createBy" gorm:"column:create_by;comment:;size:10;"`
      Currency  string `json:"currency" form:"currency" gorm:"column:currency;comment:币种;size:100;"`
      Domain  string `json:"domain" form:"domain" gorm:"column:domain;comment:区域域名;size:100;"`
      ErpPromoteFee  string `json:"erpPromoteFee" form:"erpPromoteFee" gorm:"column:erp_promote_fee;comment:早期评论人计划;size:100;"`
      EvaluationCapital  string `json:"evaluationCapital" form:"evaluationCapital" gorm:"column:evaluation_capital;comment:测评本金;size:100;"`
      EvaluationCommission  string `json:"evaluationCommission" form:"evaluationCommission" gorm:"column:evaluation_commission;comment:测评佣金;size:100;"`
      EvaluationFee  string `json:"evaluationFee" form:"evaluationFee" gorm:"column:evaluation_fee;comment:测评费;size:100;"`
      FbaAdjustmentDiffFee  string `json:"fbaAdjustmentDiffFee" form:"fbaAdjustmentDiffFee" gorm:"column:fba_adjustment_diff_fee;comment:FBA库存调整费差异;size:100;"`
      FbaAdjustmentFee  string `json:"fbaAdjustmentFee" form:"fbaAdjustmentFee" gorm:"column:fba_adjustment_fee;comment:库存调整费;size:100;"`
      FbaDisposalFee  string `json:"fbaDisposalFee" form:"fbaDisposalFee" gorm:"column:fba_disposal_fee;comment:FBA销毁费;size:100;"`
      FbaDisposalNum  *int `json:"fbaDisposalNum" form:"fbaDisposalNum" gorm:"column:fba_disposal_num;comment:FBA销毁量;size:19;"`
      FbaInventoryPlacementFee  string `json:"fbaInventoryPlacementFee" form:"fbaInventoryPlacementFee" gorm:"column:fba_inventory_placement_fee;comment:库存配置费;size:100;"`
      FbaLongStorageFee  string `json:"fbaLongStorageFee" form:"fbaLongStorageFee" gorm:"column:fba_long_storage_fee;comment:FBA长期仓储费;size:100;"`
      FbaRefund  string `json:"fbaRefund" form:"fbaRefund" gorm:"column:fba_refund;comment:FBA退款;size:100;"`
      FbaRefundNum  *int `json:"fbaRefundNum" form:"fbaRefundNum" gorm:"column:fba_refund_num;comment:FBA退款量;size:19;"`
      FbaReissueNum  *int `json:"fbaReissueNum" form:"fbaReissueNum" gorm:"column:fba_reissue_num;comment:FBA补发量;size:19;"`
      FbaRemovalFee  string `json:"fbaRemovalFee" form:"fbaRemovalFee" gorm:"column:fba_removal_fee;comment:FBA移除费;size:100;"`
      FbaRemovalNum  *int `json:"fbaRemovalNum" form:"fbaRemovalNum" gorm:"column:fba_removal_num;comment:FBA移除量;size:19;"`
      FbaReturnFee  string `json:"fbaReturnFee" form:"fbaReturnFee" gorm:"column:fba_return_fee;comment:退货入仓费;size:100;"`
      FbaSales  string `json:"fbaSales" form:"fbaSales" gorm:"column:fba_sales;comment:FBA销售额;size:100;"`
      FbaSalesNum  *int `json:"fbaSalesNum" form:"fbaSalesNum" gorm:"column:fba_sales_num;comment:亚马逊FBA销量;size:19;"`
      FbaShipFee  string `json:"fbaShipFee" form:"fbaShipFee" gorm:"column:fba_ship_fee;comment:亚马逊合作承运费;size:100;"`
      FbaShippingCredits  string `json:"fbaShippingCredits" form:"fbaShippingCredits" gorm:"column:fba_shipping_credits;comment:FBA运输信用分;size:100;"`
      FbaStorageFee  string `json:"fbaStorageFee" form:"fbaStorageFee" gorm:"column:fba_storage_fee;comment:FBA月仓储费;size:100;"`
      FbaStorageFeeDiff  string `json:"fbaStorageFeeDiff" form:"fbaStorageFeeDiff" gorm:"column:fba_storage_fee_diff;comment:FBA月仓储费差;size:100;"`
      FbaStorageOtherFee  string `json:"fbaStorageOtherFee" form:"fbaStorageOtherFee" gorm:"column:fba_storage_other_fee;comment:其他费;size:100;"`
      FbmRefund  string `json:"fbmRefund" form:"fbmRefund" gorm:"column:fbm_refund;comment:FBM退款;size:100;"`
      FbmRefundNum  *int `json:"fbmRefundNum" form:"fbmRefundNum" gorm:"column:fbm_refund_num;comment:FBM退款量;size:19;"`
      FbmSales  string `json:"fbmSales" form:"fbmSales" gorm:"column:fbm_sales;comment:FBM销售额;size:100;"`
      FbmSalesNum  *int `json:"fbmSalesNum" form:"fbmSalesNum" gorm:"column:fbm_sales_num;comment:亚马逊FBM销量;size:19;"`
      FbmShippingCredits  string `json:"fbmShippingCredits" form:"fbmShippingCredits" gorm:"column:fbm_shipping_credits;comment:FBM运输信用分;size:100;"`
      FixedFee  string `json:"fixedFee" form:"fixedFee" gorm:"column:fixed_fee;comment:固定费用;size:100;"`
      GrossProfit  string `json:"grossProfit" form:"grossProfit" gorm:"column:gross_profit;comment:毛利润;size:100;"`
      GrossProfitRate  string `json:"grossProfitRate" form:"grossProfitRate" gorm:"column:gross_profit_rate;comment:毛利率;size:100;"`
      HeadTripAcpcFee  string `json:"headTripAcpcFee" form:"headTripAcpcFee" gorm:"column:head_trip_acpc_fee;comment:头程费用（赔偿）;size:100;"`
      HeadTripDpcFee  string `json:"headTripDpcFee" form:"headTripDpcFee" gorm:"column:head_trip_dpc_fee;comment:头程费用（销毁）;size:100;"`
      HeadTripMcpcFee  string `json:"headTripMcpcFee" form:"headTripMcpcFee" gorm:"column:head_trip_mcpc_fee;comment:头程费用;size:100;"`
      HeadTripPcFee  string `json:"headTripPcFee" form:"headTripPcFee" gorm:"column:head_trip_pc_fee;comment:头程费用;size:100;"`
      HeadTripRpcFee  string `json:"headTripRpcFee" form:"headTripRpcFee" gorm:"column:head_trip_rpc_fee;comment:头程费用;size:100;"`
      HeadTripRtpcFee  string `json:"headTripRtpcFee" form:"headTripRtpcFee" gorm:"column:head_trip_rtpc_fee;comment:头程费用（退货）;size:100;"`
      LdPromoteFee  string `json:"ldPromoteFee" form:"ldPromoteFee" gorm:"column:ld_promote_fee;comment:LD费;size:100;"`
      MainImg  string `json:"mainImg" form:"mainImg" gorm:"column:main_img;comment:主图;size:100;"`
      MarketplaceId  string `json:"marketplaceId" form:"marketplaceId" gorm:"column:marketplace_id;comment:市场id;size:100;"`
      MarketplaceName  string `json:"marketplaceName" form:"marketplaceName" gorm:"column:marketplace_name;comment:市场名称;size:100;"`
      McFbaShipFee  string `json:"mcFbaShipFee" form:"mcFbaShipFee" gorm:"column:mc_fba_ship_fee;comment:FBA发货费;size:100;"`
      Month  string `json:"month" form:"month" gorm:"column:month;comment:月份;size:100;"`
      MultiChannelNum  *int `json:"multiChannelNum" form:"multiChannelNum" gorm:"column:multi_channel_num;comment:多渠道销量;size:19;"`
      MultiChannelOrderFee  string `json:"multiChannelOrderFee" form:"multiChannelOrderFee" gorm:"column:multi_channel_order_fee;comment:多渠道订单;size:100;"`
      OtherFee  string `json:"otherFee" form:"otherFee" gorm:"column:other_fee;comment:其他费;size:100;"`
      ParentAsin  string `json:"parentAsin" form:"parentAsin" gorm:"column:parent_asin;comment:父ASIN;size:100;"`
      ProductSales  string `json:"productSales" form:"productSales" gorm:"column:product_sales;comment:销售额;size:100;"`
      PromotionalRebates  string `json:"promotionalRebates" form:"promotionalRebates" gorm:"column:promotional_rebates;comment:促销折扣;size:100;"`
      PurchaseAcpcFee  string `json:"purchaseAcpcFee" form:"purchaseAcpcFee" gorm:"column:purchase_acpc_fee;comment:采购成本（赔偿）;size:100;"`
      PurchaseDpcFee  string `json:"purchaseDpcFee" form:"purchaseDpcFee" gorm:"column:purchase_dpc_fee;comment:采购成本（销毁）;size:100;"`
      PurchaseMcpcFee  string `json:"purchaseMcpcFee" form:"purchaseMcpcFee" gorm:"column:purchase_mcpc_fee;comment:采购成本;size:100;"`
      PurchasePcFee  string `json:"purchasePcFee" form:"purchasePcFee" gorm:"column:purchase_pc_fee;comment:采购成本;size:100;"`
      PurchaseRpcFee  string `json:"purchaseRpcFee" form:"purchaseRpcFee" gorm:"column:purchase_rpc_fee;comment:采购成本;size:100;"`
      PurchaseRtpcFee  string `json:"purchaseRtpcFee" form:"purchaseRtpcFee" gorm:"column:purchase_rtpc_fee;comment:采购成本（退货）;size:100;"`
      Refund  string `json:"refund" form:"refund" gorm:"column:refund;comment:退款;size:100;"`
      RefundFbaFee  string `json:"refundFbaFee" form:"refundFbaFee" gorm:"column:refund_fba_fee;comment:FBA发货费;size:100;"`
      RefundNum  *int `json:"refundNum" form:"refundNum" gorm:"column:refund_num;comment:退款量;size:19;"`
      RefundOrderFees  string `json:"refundOrderFees" form:"refundOrderFees" gorm:"column:refund_order_fees;comment:退款订单;size:100;"`
      RefundOtherFee  string `json:"refundOtherFee" form:"refundOtherFee" gorm:"column:refund_other_fee;comment:其他费;size:99;"`
      RefundRate  string `json:"refundRate" form:"refundRate" gorm:"column:refund_rate;comment:退款率;size:100;"`
      RefundSellingFee  string `json:"refundSellingFee" form:"refundSellingFee" gorm:"column:refund_selling_fee;comment:平台费;size:100;"`
      ReturnNum  *int `json:"returnNum" form:"returnNum" gorm:"column:return_num;comment:退货量;size:19;"`
      ReturnRate  string `json:"returnRate" form:"returnRate" gorm:"column:return_rate;comment:退货比例;size:100;"`
      ReturnSellableNum  *int `json:"returnSellableNum" form:"returnSellableNum" gorm:"column:return_sellable_num;comment:退货可售商品数;size:19;"`
      ReturnUnsellableNum  *int `json:"returnUnsellableNum" form:"returnUnsellableNum" gorm:"column:return_unsellable_num;comment:退货不可售商品数;size:19;"`
      Roi  string `json:"roi" form:"roi" gorm:"column:roi;comment:ROI;size:100;"`
      SaleFbaFee  string `json:"saleFbaFee" form:"saleFbaFee" gorm:"column:sale_fba_fee;comment:FBA发货费;size:100;"`
      SaleFbmFee  string `json:"saleFbmFee" form:"saleFbmFee" gorm:"column:sale_fbm_fee;comment:FBM发货费;size:100;"`
      SaleOrderFee  string `json:"saleOrderFee" form:"saleOrderFee" gorm:"column:sale_order_fee;comment:销售订单;size:100;"`
      SaleOtherFee  string `json:"saleOtherFee" form:"saleOtherFee" gorm:"column:sale_other_fee;comment:其他费;size:100;"`
      SaleRatioByAsin  string `json:"saleRatioByAsin" form:"saleRatioByAsin" gorm:"column:sale_ratio_by_asin;comment:ASIN销售率;size:100;"`
      SaleRatioByShop  string `json:"saleRatioByShop" form:"saleRatioByShop" gorm:"column:sale_ratio_by_shop;comment:店铺销售率;size:100;"`
      SaleSellingFee  string `json:"saleSellingFee" form:"saleSellingFee" gorm:"column:sale_selling_fee;comment:销售成本;size:100;"`
      ShopId  string `json:"shopId" form:"shopId" gorm:"column:shop_id;comment:店铺id;size:100;"`
      ShopName  string `json:"shopName" form:"shopName" gorm:"column:shop_name;comment:店铺名称;size:100;"`
      ShopOtherFee  string `json:"shopOtherFee" form:"shopOtherFee" gorm:"column:shop_other_fee;comment:店铺其他费;size:100;"`
      ShopOtherFees  string `json:"shopOtherFees" form:"shopOtherFees" gorm:"column:shop_other_fees;comment:店铺其他费用;size:100;"`
      Sku  string `json:"sku" form:"sku" gorm:"column:sku;comment:SKU;size:100;"`
      SubscriptionFee  string `json:"subscriptionFee" form:"subscriptionFee" gorm:"column:subscription_fee;comment:订阅费;size:100;"`
      SumBatchHeadTripFee  string `json:"sumBatchHeadTripFee" form:"sumBatchHeadTripFee" gorm:"column:sum_batch_head_trip_fee;comment:批次头程费用合计;size:100;"`
      SumBatchHeadTripOtherFee  string `json:"sumBatchHeadTripOtherFee" form:"sumBatchHeadTripOtherFee" gorm:"column:sum_batch_head_trip_other_fee;comment:批次头程其他费用合计;size:100;"`
      SumBatchPurchaseFee  string `json:"sumBatchPurchaseFee" form:"sumBatchPurchaseFee" gorm:"column:sum_batch_purchase_fee;comment:批次采购成本合计;size:100;"`
      SumBatchPurchaseOtherFee  string `json:"sumBatchPurchaseOtherFee" form:"sumBatchPurchaseOtherFee" gorm:"column:sum_batch_purchase_other_fee;comment:批次采购其他成本合计;size:100;"`
      SumCpcCost  string `json:"sumCpcCost" form:"sumCpcCost" gorm:"column:sum_cpc_cost;comment:广告费;size:100;"`
      SumCustomizeFee  string `json:"sumCustomizeFee" form:"sumCustomizeFee" gorm:"column:sum_customize_fee;comment:自定义费用;size:100;"`
      SumFbaAdjustmentFee  string `json:"sumFbaAdjustmentFee" form:"sumFbaAdjustmentFee" gorm:"column:sum_fba_adjustment_fee;comment:FBA仓储调整费用合计;size:100;"`
      SumFbaStorageFee  string `json:"sumFbaStorageFee" form:"sumFbaStorageFee" gorm:"column:sum_fba_storage_fee;comment:仓储费用;size:100;"`
      SumHeadTripFee  string `json:"sumHeadTripFee" form:"sumHeadTripFee" gorm:"column:sum_head_trip_fee;comment:总头程费用;size:100;"`
      SumOrderFee  string `json:"sumOrderFee" form:"sumOrderFee" gorm:"column:sum_order_fee;comment:订单费用;size:100;"`
      SumOtherFee  string `json:"sumOtherFee" form:"sumOtherFee" gorm:"column:sum_other_fee;comment:其他费合计;size:100;"`
      SumPromoteFee  string `json:"sumPromoteFee" form:"sumPromoteFee" gorm:"column:sum_promote_fee;comment:推广费用;size:100;"`
      SumPurchaseFee  string `json:"sumPurchaseFee" form:"sumPurchaseFee" gorm:"column:sum_purchase_fee;comment:总采购成本;size:100;"`
      SumSalesNum  *int `json:"sumSalesNum" form:"sumSalesNum" gorm:"column:sum_sales_num;comment:销售数量;size:19;"`
      SumStorageOtherFee  string `json:"sumStorageOtherFee" form:"sumStorageOtherFee" gorm:"column:sum_storage_other_fee;comment:其他仓储费用;size:100;"`
      Tax  string `json:"tax" form:"tax" gorm:"column:tax;comment:税费;size:100;"`
      UpdateBy  *int `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:;size:10;"`
}


// TableName ErpFinanceAsinSummary 表名
func (ErpFinanceAsinSummary) TableName() string {
  return "erp_finance_asin_summary"
}

