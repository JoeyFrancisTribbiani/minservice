<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button size="mini" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="mini" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="mini" type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
                <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" size="mini" style="margin-left: 10px;" :disabled="!multipleSelection.length">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="库存调整赔偿" prop="adjCompensation" width="120" />
        <el-table-column align="left" label="库存调整数量" prop="adjCompensationNum" width="120" />
        <el-table-column align="left" label="FBA库存赔偿" prop="adjCompensationPcFee" width="120" />
        <el-table-column align="left" label="ASIN" prop="asin" width="120" />
        <el-table-column align="left" label="ASIN列表" prop="asinList" width="120" />
        <el-table-column align="left" label="ASIN其他费" prop="asinOtherFee" width="120" />
        <el-table-column align="left" label="ASIN其他费" prop="asinOtherFees" width="120" />
        <el-table-column align="left" label="批次采购成本（赔偿）" prop="batchPurchaseApcFee" width="120" />
        <el-table-column align="left" label="批次采购成本（退货）" prop="batchPurchaseCrpcFee" width="120" />
        <el-table-column align="left" label="批次采购成本（差异）" prop="batchPurchaseDfpcFee" width="120" />
        <el-table-column align="left" label="批次采购成本" prop="batchPurchasePcFee" width="120" />
        <el-table-column align="left" label="批次采购成本（供应商退货）" prop="batchPurchaseVrpcFee" width="120" />
        <el-table-column align="left" label="批次头程费用（赔偿）" prop="batchTripApcFee" width="120" />
        <el-table-column align="left" label="批次头程费用（退货）" prop="batchTripCrpcFee" width="120" />
        <el-table-column align="left" label="批次头程费用（差异）" prop="batchTripDfpcFee" width="120" />
        <el-table-column align="left" label="批次头程费用" prop="batchTripPcFee" width="120" />
        <el-table-column align="left" label="批次头程费用（供应商退货）" prop="batchTripVrpcFee" width="120" />
        <el-table-column align="left" label="商品列表" prop="commodityList" width="120" />
        <el-table-column align="left" label="商品名称" prop="commodityName" width="120" />
        <el-table-column align="left" label="商品SKU" prop="commoditySku" width="120" />
        <el-table-column align="left" label="优惠劵" prop="couponsPromoteFee" width="120" />
        <el-table-column align="left" label="CPC广告费差异" prop="cpcCostDiff" width="120" />
        <el-table-column align="left" label="CPC销售额" prop="cpcSales" width="120" />
        <el-table-column align="left" label="CPC销量" prop="cpcSalesNum" width="120" />
        <el-table-column align="left" label="CPC广告成本" prop="cpcSdCost" width="120" />
        <el-table-column align="left" label="CPC广告销量" prop="cpcSdSalesNum" width="120" />
        <el-table-column align="left" label="CPC推广费" prop="cpcSpCost" width="120" />
        <el-table-column align="left" label="CPC推广销量" prop="cpcSpSalesNum" width="120" />
        <el-table-column align="left" label="createBy字段" prop="createBy" width="120" />
        <el-table-column align="left" label="币种" prop="currency" width="120" />
        <el-table-column align="left" label="区域域名" prop="domain" width="120" />
        <el-table-column align="left" label="早期评论人计划" prop="erpPromoteFee" width="120" />
        <el-table-column align="left" label="测评本金" prop="evaluationCapital" width="120" />
        <el-table-column align="left" label="测评佣金" prop="evaluationCommission" width="120" />
        <el-table-column align="left" label="测评费" prop="evaluationFee" width="120" />
        <el-table-column align="left" label="FBA库存调整费差异" prop="fbaAdjustmentDiffFee" width="120" />
        <el-table-column align="left" label="库存调整费" prop="fbaAdjustmentFee" width="120" />
        <el-table-column align="left" label="FBA销毁费" prop="fbaDisposalFee" width="120" />
        <el-table-column align="left" label="FBA销毁量" prop="fbaDisposalNum" width="120" />
        <el-table-column align="left" label="库存配置费" prop="fbaInventoryPlacementFee" width="120" />
        <el-table-column align="left" label="FBA长期仓储费" prop="fbaLongStorageFee" width="120" />
        <el-table-column align="left" label="FBA退款" prop="fbaRefund" width="120" />
        <el-table-column align="left" label="FBA退款量" prop="fbaRefundNum" width="120" />
        <el-table-column align="left" label="FBA补发量" prop="fbaReissueNum" width="120" />
        <el-table-column align="left" label="FBA移除费" prop="fbaRemovalFee" width="120" />
        <el-table-column align="left" label="FBA移除量" prop="fbaRemovalNum" width="120" />
        <el-table-column align="left" label="退货入仓费" prop="fbaReturnFee" width="120" />
        <el-table-column align="left" label="FBA销售额" prop="fbaSales" width="120" />
        <el-table-column align="left" label="亚马逊FBA销量" prop="fbaSalesNum" width="120" />
        <el-table-column align="left" label="亚马逊合作承运费" prop="fbaShipFee" width="120" />
        <el-table-column align="left" label="FBA运输信用分" prop="fbaShippingCredits" width="120" />
        <el-table-column align="left" label="FBA月仓储费" prop="fbaStorageFee" width="120" />
        <el-table-column align="left" label="FBA月仓储费差" prop="fbaStorageFeeDiff" width="120" />
        <el-table-column align="left" label="其他费" prop="fbaStorageOtherFee" width="120" />
        <el-table-column align="left" label="FBM退款" prop="fbmRefund" width="120" />
        <el-table-column align="left" label="FBM退款量" prop="fbmRefundNum" width="120" />
        <el-table-column align="left" label="FBM销售额" prop="fbmSales" width="120" />
        <el-table-column align="left" label="亚马逊FBM销量" prop="fbmSalesNum" width="120" />
        <el-table-column align="left" label="FBM运输信用分" prop="fbmShippingCredits" width="120" />
        <el-table-column align="left" label="固定费用" prop="fixedFee" width="120" />
        <el-table-column align="left" label="毛利润" prop="grossProfit" width="120" />
        <el-table-column align="left" label="毛利率" prop="grossProfitRate" width="120" />
        <el-table-column align="left" label="头程费用（赔偿）" prop="headTripAcpcFee" width="120" />
        <el-table-column align="left" label="头程费用（销毁）" prop="headTripDpcFee" width="120" />
        <el-table-column align="left" label="头程费用" prop="headTripMcpcFee" width="120" />
        <el-table-column align="left" label="头程费用" prop="headTripPcFee" width="120" />
        <el-table-column align="left" label="头程费用" prop="headTripRpcFee" width="120" />
        <el-table-column align="left" label="头程费用（退货）" prop="headTripRtpcFee" width="120" />
        <el-table-column align="left" label="LD费" prop="ldPromoteFee" width="120" />
        <el-table-column align="left" label="主图" prop="mainImg" width="120" />
        <el-table-column align="left" label="市场id" prop="marketplaceId" width="120" />
        <el-table-column align="left" label="市场名称" prop="marketplaceName" width="120" />
        <el-table-column align="left" label="FBA发货费" prop="mcFbaShipFee" width="120" />
        <el-table-column align="left" label="月份" prop="month" width="120" />
        <el-table-column align="left" label="多渠道销量" prop="multiChannelNum" width="120" />
        <el-table-column align="left" label="多渠道订单" prop="multiChannelOrderFee" width="120" />
        <el-table-column align="left" label="其他费" prop="otherFee" width="120" />
        <el-table-column align="left" label="父ASIN" prop="parentAsin" width="120" />
        <el-table-column align="left" label="销售额" prop="productSales" width="120" />
        <el-table-column align="left" label="促销折扣" prop="promotionalRebates" width="120" />
        <el-table-column align="left" label="采购成本（赔偿）" prop="purchaseAcpcFee" width="120" />
        <el-table-column align="left" label="采购成本（销毁）" prop="purchaseDpcFee" width="120" />
        <el-table-column align="left" label="采购成本" prop="purchaseMcpcFee" width="120" />
        <el-table-column align="left" label="采购成本" prop="purchasePcFee" width="120" />
        <el-table-column align="left" label="采购成本" prop="purchaseRpcFee" width="120" />
        <el-table-column align="left" label="采购成本（退货）" prop="purchaseRtpcFee" width="120" />
        <el-table-column align="left" label="退款" prop="refund" width="120" />
        <el-table-column align="left" label="FBA发货费" prop="refundFbaFee" width="120" />
        <el-table-column align="left" label="退款量" prop="refundNum" width="120" />
        <el-table-column align="left" label="退款订单" prop="refundOrderFees" width="120" />
        <el-table-column align="left" label="其他费" prop="refundOtherFee" width="120" />
        <el-table-column align="left" label="退款率" prop="refundRate" width="120" />
        <el-table-column align="left" label="平台费" prop="refundSellingFee" width="120" />
        <el-table-column align="left" label="退货量" prop="returnNum" width="120" />
        <el-table-column align="left" label="退货比例" prop="returnRate" width="120" />
        <el-table-column align="left" label="退货可售商品数" prop="returnSellableNum" width="120" />
        <el-table-column align="left" label="退货不可售商品数" prop="returnUnsellableNum" width="120" />
        <el-table-column align="left" label="ROI" prop="roi" width="120" />
        <el-table-column align="left" label="FBA发货费" prop="saleFbaFee" width="120" />
        <el-table-column align="left" label="FBM发货费" prop="saleFbmFee" width="120" />
        <el-table-column align="left" label="销售订单" prop="saleOrderFee" width="120" />
        <el-table-column align="left" label="其他费" prop="saleOtherFee" width="120" />
        <el-table-column align="left" label="ASIN销售率" prop="saleRatioByAsin" width="120" />
        <el-table-column align="left" label="店铺销售率" prop="saleRatioByShop" width="120" />
        <el-table-column align="left" label="销售成本" prop="saleSellingFee" width="120" />
        <el-table-column align="left" label="店铺id" prop="shopId" width="120" />
        <el-table-column align="left" label="店铺名称" prop="shopName" width="120" />
        <el-table-column align="left" label="店铺其他费" prop="shopOtherFee" width="120" />
        <el-table-column align="left" label="店铺其他费用" prop="shopOtherFees" width="120" />
        <el-table-column align="left" label="SKU" prop="sku" width="120" />
        <el-table-column align="left" label="订阅费" prop="subscriptionFee" width="120" />
        <el-table-column align="left" label="批次头程费用合计" prop="sumBatchHeadTripFee" width="120" />
        <el-table-column align="left" label="批次头程其他费用合计" prop="sumBatchHeadTripOtherFee" width="120" />
        <el-table-column align="left" label="批次采购成本合计" prop="sumBatchPurchaseFee" width="120" />
        <el-table-column align="left" label="批次采购其他成本合计" prop="sumBatchPurchaseOtherFee" width="120" />
        <el-table-column align="left" label="广告费" prop="sumCpcCost" width="120" />
        <el-table-column align="left" label="自定义费用" prop="sumCustomizeFee" width="120" />
        <el-table-column align="left" label="FBA仓储调整费用合计" prop="sumFbaAdjustmentFee" width="120" />
        <el-table-column align="left" label="仓储费用" prop="sumFbaStorageFee" width="120" />
        <el-table-column align="left" label="总头程费用" prop="sumHeadTripFee" width="120" />
        <el-table-column align="left" label="订单费用" prop="sumOrderFee" width="120" />
        <el-table-column align="left" label="其他费合计" prop="sumOtherFee" width="120" />
        <el-table-column align="left" label="推广费用" prop="sumPromoteFee" width="120" />
        <el-table-column align="left" label="总采购成本" prop="sumPurchaseFee" width="120" />
        <el-table-column align="left" label="销售数量" prop="sumSalesNum" width="120" />
        <el-table-column align="left" label="其他仓储费用" prop="sumStorageOtherFee" width="120" />
        <el-table-column align="left" label="税费" prop="tax" width="120" />
        <el-table-column align="left" label="updateBy字段" prop="updateBy" width="120" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateErpFinanceAsinSummaryFunc(scope.row)">变更</el-button>
            <el-button type="text" icon="delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
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
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
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
  deleteErpFinanceAsinSummary,
  deleteErpFinanceAsinSummaryByIds,
  updateErpFinanceAsinSummary,
  findErpFinanceAsinSummary,
  getErpFinanceAsinSummaryList
} from '@/api/erpFinanceAsinSummary'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
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

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getErpFinanceAsinSummaryList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteErpFinanceAsinSummaryFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteErpFinanceAsinSummaryByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateErpFinanceAsinSummaryFunc = async(row) => {
    const res = await findErpFinanceAsinSummary({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reerpFinanceAsinSummary
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteErpFinanceAsinSummaryFunc = async (row) => {
    const res = await deleteErpFinanceAsinSummary({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
        closeDialog()
        getTableData()
      }
}
</script>

<style>
</style>
