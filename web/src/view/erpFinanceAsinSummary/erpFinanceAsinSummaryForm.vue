<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="库存调整赔偿:">
          <el-input v-model="formData.adjCompensation" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="库存调整数量:">
          <el-input v-model.number="formData.adjCompensationNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="FBA库存赔偿:">
          <el-input v-model="formData.adjCompensationPcFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="ASIN:">
          <el-input v-model="formData.asin" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="ASIN列表:">
          <el-input v-model="formData.asinList" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="ASIN其他费:">
          <el-input v-model="formData.asinOtherFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="ASIN其他费:">
          <el-input v-model="formData.asinOtherFees" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="批次采购成本（赔偿）:">
          <el-input v-model="formData.batchPurchaseApcFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="批次采购成本（退货）:">
          <el-input v-model="formData.batchPurchaseCrpcFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="批次采购成本（差异）:">
          <el-input v-model="formData.batchPurchaseDfpcFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="批次采购成本:">
          <el-input v-model="formData.batchPurchasePcFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="批次采购成本（供应商退货）:">
          <el-input v-model="formData.batchPurchaseVrpcFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="批次头程费用（赔偿）:">
          <el-input v-model="formData.batchTripApcFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="批次头程费用（退货）:">
          <el-input v-model="formData.batchTripCrpcFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="批次头程费用（差异）:">
          <el-input v-model="formData.batchTripDfpcFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="批次头程费用:">
          <el-input v-model="formData.batchTripPcFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="批次头程费用（供应商退货）:">
          <el-input v-model="formData.batchTripVrpcFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品列表:">
          <el-input v-model="formData.commodityList" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品名称:">
          <el-input v-model="formData.commodityName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品SKU:">
          <el-input v-model="formData.commoditySku" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="优惠劵:">
          <el-input v-model="formData.couponsPromoteFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="CPC广告费差异:">
          <el-input v-model="formData.cpcCostDiff" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="CPC销售额:">
          <el-input v-model="formData.cpcSales" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="CPC销量:">
          <el-input v-model.number="formData.cpcSalesNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="CPC广告成本:">
          <el-input v-model="formData.cpcSdCost" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="CPC广告销量:">
          <el-input v-model.number="formData.cpcSdSalesNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="CPC推广费:">
          <el-input v-model="formData.cpcSpCost" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="CPC推广销量:">
          <el-input v-model.number="formData.cpcSpSalesNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="createBy字段:">
          <el-input v-model.number="formData.createBy" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="币种:">
          <el-input v-model="formData.currency" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="区域域名:">
          <el-input v-model="formData.domain" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="早期评论人计划:">
          <el-input v-model="formData.erpPromoteFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="测评本金:">
          <el-input v-model="formData.evaluationCapital" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="测评佣金:">
          <el-input v-model="formData.evaluationCommission" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="测评费:">
          <el-input v-model="formData.evaluationFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="FBA库存调整费差异:">
          <el-input v-model="formData.fbaAdjustmentDiffFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="库存调整费:">
          <el-input v-model="formData.fbaAdjustmentFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="FBA销毁费:">
          <el-input v-model="formData.fbaDisposalFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="FBA销毁量:">
          <el-input v-model.number="formData.fbaDisposalNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="库存配置费:">
          <el-input v-model="formData.fbaInventoryPlacementFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="FBA长期仓储费:">
          <el-input v-model="formData.fbaLongStorageFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="FBA退款:">
          <el-input v-model="formData.fbaRefund" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="FBA退款量:">
          <el-input v-model.number="formData.fbaRefundNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="FBA补发量:">
          <el-input v-model.number="formData.fbaReissueNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="FBA移除费:">
          <el-input v-model="formData.fbaRemovalFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="FBA移除量:">
          <el-input v-model.number="formData.fbaRemovalNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="退货入仓费:">
          <el-input v-model="formData.fbaReturnFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="FBA销售额:">
          <el-input v-model="formData.fbaSales" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="亚马逊FBA销量:">
          <el-input v-model.number="formData.fbaSalesNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="亚马逊合作承运费:">
          <el-input v-model="formData.fbaShipFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="FBA运输信用分:">
          <el-input v-model="formData.fbaShippingCredits" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="FBA月仓储费:">
          <el-input v-model="formData.fbaStorageFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="FBA月仓储费差:">
          <el-input v-model="formData.fbaStorageFeeDiff" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="其他费:">
          <el-input v-model="formData.fbaStorageOtherFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="FBM退款:">
          <el-input v-model="formData.fbmRefund" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="FBM退款量:">
          <el-input v-model.number="formData.fbmRefundNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="FBM销售额:">
          <el-input v-model="formData.fbmSales" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="亚马逊FBM销量:">
          <el-input v-model.number="formData.fbmSalesNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="FBM运输信用分:">
          <el-input v-model="formData.fbmShippingCredits" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="固定费用:">
          <el-input v-model="formData.fixedFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="毛利润:">
          <el-input v-model="formData.grossProfit" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="毛利率:">
          <el-input v-model="formData.grossProfitRate" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="头程费用（赔偿）:">
          <el-input v-model="formData.headTripAcpcFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="头程费用（销毁）:">
          <el-input v-model="formData.headTripDpcFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="头程费用:">
          <el-input v-model="formData.headTripMcpcFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="头程费用:">
          <el-input v-model="formData.headTripPcFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="头程费用:">
          <el-input v-model="formData.headTripRpcFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="头程费用（退货）:">
          <el-input v-model="formData.headTripRtpcFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="LD费:">
          <el-input v-model="formData.ldPromoteFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="主图:">
          <el-input v-model="formData.mainImg" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="市场id:">
          <el-input v-model="formData.marketplaceId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="市场名称:">
          <el-input v-model="formData.marketplaceName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="FBA发货费:">
          <el-input v-model="formData.mcFbaShipFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="月份:">
          <el-input v-model="formData.month" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="多渠道销量:">
          <el-input v-model.number="formData.multiChannelNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="多渠道订单:">
          <el-input v-model="formData.multiChannelOrderFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="其他费:">
          <el-input v-model="formData.otherFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="父ASIN:">
          <el-input v-model="formData.parentAsin" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="销售额:">
          <el-input v-model="formData.productSales" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="促销折扣:">
          <el-input v-model="formData.promotionalRebates" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="采购成本（赔偿）:">
          <el-input v-model="formData.purchaseAcpcFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="采购成本（销毁）:">
          <el-input v-model="formData.purchaseDpcFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="采购成本:">
          <el-input v-model="formData.purchaseMcpcFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="采购成本:">
          <el-input v-model="formData.purchasePcFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="采购成本:">
          <el-input v-model="formData.purchaseRpcFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="采购成本（退货）:">
          <el-input v-model="formData.purchaseRtpcFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="退款:">
          <el-input v-model="formData.refund" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="FBA发货费:">
          <el-input v-model="formData.refundFbaFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="退款量:">
          <el-input v-model.number="formData.refundNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="退款订单:">
          <el-input v-model="formData.refundOrderFees" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="其他费:">
          <el-input v-model="formData.refundOtherFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="退款率:">
          <el-input v-model="formData.refundRate" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="平台费:">
          <el-input v-model="formData.refundSellingFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="退货量:">
          <el-input v-model.number="formData.returnNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="退货比例:">
          <el-input v-model="formData.returnRate" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="退货可售商品数:">
          <el-input v-model.number="formData.returnSellableNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="退货不可售商品数:">
          <el-input v-model.number="formData.returnUnsellableNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="ROI:">
          <el-input v-model="formData.roi" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="FBA发货费:">
          <el-input v-model="formData.saleFbaFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="FBM发货费:">
          <el-input v-model="formData.saleFbmFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="销售订单:">
          <el-input v-model="formData.saleOrderFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="其他费:">
          <el-input v-model="formData.saleOtherFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="ASIN销售率:">
          <el-input v-model="formData.saleRatioByAsin" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="店铺销售率:">
          <el-input v-model="formData.saleRatioByShop" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="销售成本:">
          <el-input v-model="formData.saleSellingFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="店铺id:">
          <el-input v-model="formData.shopId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="店铺名称:">
          <el-input v-model="formData.shopName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="店铺其他费:">
          <el-input v-model="formData.shopOtherFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="店铺其他费用:">
          <el-input v-model="formData.shopOtherFees" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="SKU:">
          <el-input v-model="formData.sku" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="订阅费:">
          <el-input v-model="formData.subscriptionFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="批次头程费用合计:">
          <el-input v-model="formData.sumBatchHeadTripFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="批次头程其他费用合计:">
          <el-input v-model="formData.sumBatchHeadTripOtherFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="批次采购成本合计:">
          <el-input v-model="formData.sumBatchPurchaseFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="批次采购其他成本合计:">
          <el-input v-model="formData.sumBatchPurchaseOtherFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="广告费:">
          <el-input v-model="formData.sumCpcCost" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="自定义费用:">
          <el-input v-model="formData.sumCustomizeFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="FBA仓储调整费用合计:">
          <el-input v-model="formData.sumFbaAdjustmentFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="仓储费用:">
          <el-input v-model="formData.sumFbaStorageFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="总头程费用:">
          <el-input v-model="formData.sumHeadTripFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="订单费用:">
          <el-input v-model="formData.sumOrderFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="其他费合计:">
          <el-input v-model="formData.sumOtherFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="推广费用:">
          <el-input v-model="formData.sumPromoteFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="总采购成本:">
          <el-input v-model="formData.sumPurchaseFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="销售数量:">
          <el-input v-model.number="formData.sumSalesNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="其他仓储费用:">
          <el-input v-model="formData.sumStorageOtherFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="税费:">
          <el-input v-model="formData.tax" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="updateBy字段:">
          <el-input v-model.number="formData.updateBy" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ErpFinanceAsinSummary'
}
</script>

<script setup>
import {
  createErpFinanceAsinSummary,
  updateErpFinanceAsinSummary,
  findErpFinanceAsinSummary
} from '@/api/erpFinanceAsinSummary'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
        adjCompensation: '',
        adjCompensationNum: 0,
        adjCompensationPcFee: '',
        asin: '',
        asinList: '',
        asinOtherFee: '',
        asinOtherFees: '',
        batchPurchaseApcFee: '',
        batchPurchaseCrpcFee: '',
        batchPurchaseDfpcFee: '',
        batchPurchasePcFee: '',
        batchPurchaseVrpcFee: '',
        batchTripApcFee: '',
        batchTripCrpcFee: '',
        batchTripDfpcFee: '',
        batchTripPcFee: '',
        batchTripVrpcFee: '',
        commodityList: '',
        commodityName: '',
        commoditySku: '',
        couponsPromoteFee: '',
        cpcCostDiff: '',
        cpcSales: '',
        cpcSalesNum: 0,
        cpcSdCost: '',
        cpcSdSalesNum: 0,
        cpcSpCost: '',
        cpcSpSalesNum: 0,
        createBy: 0,
        currency: '',
        domain: '',
        erpPromoteFee: '',
        evaluationCapital: '',
        evaluationCommission: '',
        evaluationFee: '',
        fbaAdjustmentDiffFee: '',
        fbaAdjustmentFee: '',
        fbaDisposalFee: '',
        fbaDisposalNum: 0,
        fbaInventoryPlacementFee: '',
        fbaLongStorageFee: '',
        fbaRefund: '',
        fbaRefundNum: 0,
        fbaReissueNum: 0,
        fbaRemovalFee: '',
        fbaRemovalNum: 0,
        fbaReturnFee: '',
        fbaSales: '',
        fbaSalesNum: 0,
        fbaShipFee: '',
        fbaShippingCredits: '',
        fbaStorageFee: '',
        fbaStorageFeeDiff: '',
        fbaStorageOtherFee: '',
        fbmRefund: '',
        fbmRefundNum: 0,
        fbmSales: '',
        fbmSalesNum: 0,
        fbmShippingCredits: '',
        fixedFee: '',
        grossProfit: '',
        grossProfitRate: '',
        headTripAcpcFee: '',
        headTripDpcFee: '',
        headTripMcpcFee: '',
        headTripPcFee: '',
        headTripRpcFee: '',
        headTripRtpcFee: '',
        ldPromoteFee: '',
        mainImg: '',
        marketplaceId: '',
        marketplaceName: '',
        mcFbaShipFee: '',
        month: '',
        multiChannelNum: 0,
        multiChannelOrderFee: '',
        otherFee: '',
        parentAsin: '',
        productSales: '',
        promotionalRebates: '',
        purchaseAcpcFee: '',
        purchaseDpcFee: '',
        purchaseMcpcFee: '',
        purchasePcFee: '',
        purchaseRpcFee: '',
        purchaseRtpcFee: '',
        refund: '',
        refundFbaFee: '',
        refundNum: 0,
        refundOrderFees: '',
        refundOtherFee: '',
        refundRate: '',
        refundSellingFee: '',
        returnNum: 0,
        returnRate: '',
        returnSellableNum: 0,
        returnUnsellableNum: 0,
        roi: '',
        saleFbaFee: '',
        saleFbmFee: '',
        saleOrderFee: '',
        saleOtherFee: '',
        saleRatioByAsin: '',
        saleRatioByShop: '',
        saleSellingFee: '',
        shopId: '',
        shopName: '',
        shopOtherFee: '',
        shopOtherFees: '',
        sku: '',
        subscriptionFee: '',
        sumBatchHeadTripFee: '',
        sumBatchHeadTripOtherFee: '',
        sumBatchPurchaseFee: '',
        sumBatchPurchaseOtherFee: '',
        sumCpcCost: '',
        sumCustomizeFee: '',
        sumFbaAdjustmentFee: '',
        sumFbaStorageFee: '',
        sumHeadTripFee: '',
        sumOrderFee: '',
        sumOtherFee: '',
        sumPromoteFee: '',
        sumPurchaseFee: '',
        sumSalesNum: 0,
        sumStorageOtherFee: '',
        tax: '',
        updateBy: 0,
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findErpFinanceAsinSummary({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reerpFinanceAsinSummary
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createErpFinanceAsinSummary(formData.value)
          break
        case 'update':
          res = await updateErpFinanceAsinSummary(formData.value)
          break
        default:
          res = await createErpFinanceAsinSummary(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
