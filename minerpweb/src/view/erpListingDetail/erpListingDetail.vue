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
        <el-table-column align="left" label="createBy字段" prop="createBy" width="120" />
        <el-table-column align="left" label="updateBy字段" prop="updateBy" width="120" />
        <el-table-column align="left" label="2231" prop="puid" width="120" />
        <el-table-column align="left" label="4462" prop="shopId" width="120" />
        <el-table-column align="left" label="ATVPDKIKX0DER" prop="marketplaceId" width="120" />
        <el-table-column align="left" label="fullCid字段" prop="fullCid" width="120" />
        <el-table-column align="left" label="4762272" prop="parentId" width="120" />
        <el-table-column align="left" label="online_amazon" prop="dxmState" width="120" />
        <el-table-column align="left" label="null" prop="dxmPublishState" width="120" />
        <el-table-column align="left" label="null" prop="errMsg" width="120" />
        <el-table-column align="left" label="B08X4MVXS7" prop="asin" width="120" />
        <el-table-column align="left" label="0222Y6BRSTD" prop="listingId" width="120" />
        <el-table-column align="left" label="null" prop="parentAsin" width="120" />
        <el-table-column align="left" label="null" prop="childAsins" width="120" />
        <el-table-column align="left" label="0" prop="isVariation" width="120" />
        <el-table-column align="left" label="VQ-C5VT-Y5GH" prop="sku" width="120" />
        <el-table-column align="left" label="null" prop="standardProductType" width="120" />
        <el-table-column align="left" label="null" prop="standardProductId" width="120" />
        <el-table-column align="left" label="Outdoor Faucet Cover" prop="title" width="120" />
        <el-table-column align="left" label="null" prop="brand" width="120" />
        <el-table-column align="left" label="null" prop="manufacturer" width="120" />
        <el-table-column align="left" label="null" prop="description" width="120" />
        <el-table-column align="left" label="null" prop="bulletPoint" width="120" />
        <el-table-column align="left" label="null" prop="conditionType" width="120" />
        <el-table-column align="left" label="null" prop="conditionValue" width="120" />
        <el-table-column align="left" label="8.99" prop="standardPrice" width="120" />
        <el-table-column align="left" label="8.99" prop="listingPricing" width="120" />
        <el-table-column align="left" label="AFN" prop="switchFulfillmentTo" width="120" />
        <el-table-column align="left" label="null" prop="fulfillmentChannel" width="120" />
        <el-table-column align="left" label="null" prop="partNumber" width="120" />
        <el-table-column align="left" label="https://m.media-amazon.com/images/I/51QOuKgk22L._SL75_.jpg" prop="mainImage" width="120" />
        <el-table-column align="left" label="null" prop="saleStartDate" width="120" />
        <el-table-column align="left" label="null" prop="saleEndDate" width="120" />
        <el-table-column align="left" label="null" prop="saleSalePrice" width="120" />
        <el-table-column align="left" label="34" prop="quantity" width="120" />
        <el-table-column align="left" label="2021-02-22 04:20:10" prop="openDate" width="120" />
        <el-table-column align="left" label="null" prop="itemDimensions" width="120" />
        <el-table-column align="left" label="null" prop="packageDimensions" width="120" />
        <el-table-column align="left" label="USD" prop="standardPriceCurrency" width="120" />
        <el-table-column align="left" label="null" prop="specifics" width="120" />
        <el-table-column align="left" label="null" prop="variationChildStr" width="120" />
        <el-table-column align="left" label="Active" prop="onlineStatus" width="120" />
        <el-table-column align="left" label="0" prop="saleNum" width="120" />
        <el-table-column align="left" label="X002TADZZ7" prop="fnsku" width="120" />
        <el-table-column align="left" label="0" prop="warehousing" width="120" />
        <el-table-column align="left" label="1" prop="unsellable" width="120" />
        <el-table-column align="left" label="0" prop="reservedQty" width="120" />
        <el-table-column align="left" label="12" prop="purchaseCost" width="120" />
        <el-table-column align="left" label="999" prop="headTripCost" width="120" />
        <el-table-column align="left" label="null" prop="shipCost" width="120" />
        <el-table-column align="left" label="null" prop="priceFeedStatus" width="120" />
        <el-table-column align="left" label="null" prop="inventoryFeedStatus" width="120" />
        <el-table-column align="left" label="null" prop="updateStandardPrice" width="120" />
        <el-table-column align="left" label="null" prop="updateQuantity" width="120" />
        <el-table-column align="left" label="null" prop="rating" width="120" />
        <el-table-column align="left" label="null" prop="ratingCount" width="120" />
        <el-table-column align="left" label="null" prop="rank" width="120" />
        <el-table-column align="left" label="null" prop="bsrRank" width="120" />
        <el-table-column align="left" label="470356" prop="commodityId" width="120" />
        <el-table-column align="left" label="null" prop="mathCommodityTime" width="120" />
        <el-table-column align="left" label="null" prop="maxShipmentQty" width="120" />
        <el-table-column align="left" label="9207" prop="devId" width="120" />
        <el-table-column align="left" label="null" prop="createId" width="120" />
        <el-table-column align="left" label="null" prop="updateId" width="120" />
        <el-table-column align="left" label="null" prop="lpriceUpdateTime" width="120" />
        <el-table-column align="left" label="null" prop="lpriceStatus" width="120" />
        <el-table-column align="left" label="null" prop="childList" width="120" />
        <el-table-column align="left" label="0" prop="salePrices" width="120" />
        <el-table-column align="left" label="0" prop="adCosts" width="120" />
        <el-table-column align="left" label="0" prop="adCostsSd" width="120" />
        <el-table-column align="left" label="0" prop="adCostsSb" width="120" />
        <el-table-column align="left" label="0" prop="adCostsSbv" width="120" />
        <el-table-column align="left" label="0" prop="adCostsSum" width="120" />
        <el-table-column align="left" label="USD" prop="currency" width="120" />
        <el-table-column align="left" label="CNY" prop="costCurrency" width="120" />
        <el-table-column align="left" label="testsku123123" prop="commoditySku" width="120" />
        <el-table-column align="left" label="测试商品" prop="commodityName" width="120" />
        <el-table-column align="left" label="A3LR0SXN8QVIE0-na-US" prop="shopName" width="120" />
        <el-table-column align="left" label="美国" prop="siteName" width="120" />
        <el-table-column align="left" label="null" prop="sellingPartnerId" width="120" />
        <el-table-column align="left" label="www.amazon.com" prop="domain" width="120" />
        <el-table-column align="left" label="null" prop="childNum" width="120" />
        <el-table-column align="left" label="业务员1" prop="devName" width="120" />
        <el-table-column align="left" label="null" prop="variationStrMap" width="120" />
        <el-table-column align="left" label="8.99" prop="listingPrice" width="120" />
        <el-table-column align="left" label="0" prop="shippingPrice" width="120" />
        <el-table-column align="left" label="true" prop="buyBoxWinner" width="120" />
        <el-table-column align="left" label="USD" prop="buyBoxCurrency" width="120" />
        <el-table-column align="left" label="4.82" prop="totalFee" width="120" />
        <el-table-column align="left" label="1.35" prop="referralFee" width="120" />
        <el-table-column align="left" label="0" prop="variableClosingFee" width="120" />
        <el-table-column align="left" label="0" prop="perItemFee" width="120" />
        <el-table-column align="left" label="3.47" prop="fbaFees" width="120" />
        <el-table-column align="left" label="USD" prop="feeCurrency" width="120" />
        <el-table-column align="left" label="null" prop="isPutAsh" width="120" />
        <el-table-column align="left" label="null" prop="bsrRankStr" width="120" />
        <el-table-column align="left" label="vartationStr字段" prop="vartationStr" width="120" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateErpListingDetailFunc(scope.row)">变更</el-button>
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
        <el-form-item label="createBy字段:">
          <el-input v-model.number="formData.createBy" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="updateBy字段:">
          <el-input v-model.number="formData.updateBy" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="2231:">
          <el-input v-model.number="formData.puid" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="4462:">
          <el-input v-model.number="formData.shopId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="ATVPDKIKX0DER:">
          <el-input v-model="formData.marketplaceId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="fullCid字段:">
          <el-input v-model="formData.fullCid" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="4762272:">
          <el-input v-model.number="formData.parentId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="online_amazon:">
          <el-input v-model="formData.dxmState" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.dxmPublishState" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.errMsg" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="B08X4MVXS7:">
          <el-input v-model="formData.asin" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="0222Y6BRSTD:">
          <el-input v-model="formData.listingId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.parentAsin" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.childAsins" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="0:">
          <el-input v-model.number="formData.isVariation" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="VQ-C5VT-Y5GH:">
          <el-input v-model="formData.sku" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.standardProductType" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.standardProductId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Outdoor Faucet Cover:">
          <el-input v-model="formData.title" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.brand" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.manufacturer" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.description" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.bulletPoint" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.conditionType" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.conditionValue" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="8.99:">
          <el-input-number v-model="formData.standardPrice"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="8.99:">
          <el-input-number v-model="formData.listingPricing"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="AFN:">
          <el-input v-model="formData.switchFulfillmentTo" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.fulfillmentChannel" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.partNumber" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="https://m.media-amazon.com/images/I/51QOuKgk22L._SL75_.jpg:">
          <el-input v-model="formData.mainImage" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.saleStartDate" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.saleEndDate" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.saleSalePrice" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="34:">
          <el-input v-model.number="formData.quantity" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="2021-02-22 04:20:10:">
          <el-input v-model="formData.openDate" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.itemDimensions" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.packageDimensions" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="USD:">
          <el-input v-model="formData.standardPriceCurrency" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.specifics" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.variationChildStr" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Active:">
          <el-input v-model="formData.onlineStatus" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="0:">
          <el-input v-model.number="formData.saleNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="X002TADZZ7:">
          <el-input v-model="formData.fnsku" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="0:">
          <el-input v-model.number="formData.warehousing" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="1:">
          <el-input v-model.number="formData.unsellable" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="0:">
          <el-input v-model.number="formData.reservedQty" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="12:">
          <el-input v-model.number="formData.purchaseCost" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="999:">
          <el-input v-model.number="formData.headTripCost" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.shipCost" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.priceFeedStatus" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.inventoryFeedStatus" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.updateStandardPrice" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.updateQuantity" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.rating" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.ratingCount" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.rank" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.bsrRank" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="470356:">
          <el-input v-model.number="formData.commodityId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.mathCommodityTime" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.maxShipmentQty" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="9207:">
          <el-input v-model="formData.devId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.createId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.updateId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.lpriceUpdateTime" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.lpriceStatus" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.childList" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="0:">
          <el-input v-model.number="formData.salePrices" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="0:">
          <el-input v-model.number="formData.adCosts" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="0:">
          <el-input v-model.number="formData.adCostsSd" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="0:">
          <el-input v-model.number="formData.adCostsSb" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="0:">
          <el-input v-model.number="formData.adCostsSbv" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="0:">
          <el-input v-model.number="formData.adCostsSum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="USD:">
          <el-input v-model="formData.currency" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="CNY:">
          <el-input v-model="formData.costCurrency" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="testsku123123:">
          <el-input v-model="formData.commoditySku" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="测试商品:">
          <el-input v-model="formData.commodityName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="A3LR0SXN8QVIE0-na-US:">
          <el-input v-model="formData.shopName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="美国:">
          <el-input v-model="formData.siteName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.sellingPartnerId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="www.amazon.com:">
          <el-input v-model="formData.domain" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.childNum" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="业务员1:">
          <el-input v-model="formData.devName" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.variationStrMap" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="8.99:">
          <el-input-number v-model="formData.listingPrice"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="0:">
          <el-input v-model.number="formData.shippingPrice" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="true:">
          <el-input v-model="formData.buyBoxWinner" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="USD:">
          <el-input v-model="formData.buyBoxCurrency" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="4.82:">
          <el-input-number v-model="formData.totalFee"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="1.35:">
          <el-input-number v-model="formData.referralFee"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="0:">
          <el-input v-model.number="formData.variableClosingFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="0:">
          <el-input v-model.number="formData.perItemFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="3.47:">
          <el-input-number v-model="formData.fbaFees"  style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="USD:">
          <el-input v-model="formData.feeCurrency" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.isPutAsh" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="null:">
          <el-input v-model="formData.bsrRankStr" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="vartationStr字段:">
          <el-input v-model="formData.vartationStr" clearable placeholder="请输入" />
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
  name: 'ErpListingDetail'
}
</script>

<script setup>
import {
  createErpListingDetail,
  deleteErpListingDetail,
  deleteErpListingDetailByIds,
  updateErpListingDetail,
  findErpListingDetail,
  getErpListingDetailList
} from '@/api/erpListingDetail'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        createBy: 0,
        updateBy: 0,
        puid: 0,
        shopId: 0,
        marketplaceId: '',
        fullCid: '',
        parentId: 0,
        dxmState: '',
        dxmPublishState: '',
        errMsg: '',
        asin: '',
        listingId: '',
        parentAsin: '',
        childAsins: '',
        isVariation: 0,
        sku: '',
        standardProductType: '',
        standardProductId: '',
        title: '',
        brand: '',
        manufacturer: '',
        description: '',
        bulletPoint: '',
        conditionType: '',
        conditionValue: '',
        standardPrice: 0,
        listingPricing: 0,
        switchFulfillmentTo: '',
        fulfillmentChannel: '',
        partNumber: '',
        mainImage: '',
        saleStartDate: '',
        saleEndDate: '',
        saleSalePrice: '',
        quantity: 0,
        openDate: '',
        itemDimensions: '',
        packageDimensions: '',
        standardPriceCurrency: '',
        specifics: '',
        variationChildStr: '',
        onlineStatus: '',
        saleNum: 0,
        fnsku: '',
        warehousing: 0,
        unsellable: 0,
        reservedQty: 0,
        purchaseCost: 0,
        headTripCost: 0,
        shipCost: '',
        priceFeedStatus: '',
        inventoryFeedStatus: '',
        updateStandardPrice: '',
        updateQuantity: '',
        rating: '',
        ratingCount: '',
        rank: '',
        bsrRank: '',
        commodityId: 0,
        mathCommodityTime: '',
        maxShipmentQty: '',
        devId: '',
        createId: '',
        updateId: '',
        lpriceUpdateTime: '',
        lpriceStatus: '',
        childList: '',
        salePrices: 0,
        adCosts: 0,
        adCostsSd: 0,
        adCostsSb: 0,
        adCostsSbv: 0,
        adCostsSum: 0,
        currency: '',
        costCurrency: '',
        commoditySku: '',
        commodityName: '',
        shopName: '',
        siteName: '',
        sellingPartnerId: '',
        domain: '',
        childNum: '',
        devName: '',
        variationStrMap: '',
        listingPrice: 0,
        shippingPrice: 0,
        buyBoxWinner: '',
        buyBoxCurrency: '',
        totalFee: 0,
        referralFee: 0,
        variableClosingFee: 0,
        perItemFee: 0,
        fbaFees: 0,
        feeCurrency: '',
        isPutAsh: '',
        bsrRankStr: '',
        vartationStr: '',
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
  const table = await getErpListingDetailList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteErpListingDetailFunc(row)
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
      const res = await deleteErpListingDetailByIds({ ids })
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
const updateErpListingDetailFunc = async(row) => {
    const res = await findErpListingDetail({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reerpListingDetail
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteErpListingDetailFunc = async (row) => {
    const res = await deleteErpListingDetail({ ID: row.ID })
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
        createBy: 0,
        updateBy: 0,
        puid: 0,
        shopId: 0,
        marketplaceId: '',
        fullCid: '',
        parentId: 0,
        dxmState: '',
        dxmPublishState: '',
        errMsg: '',
        asin: '',
        listingId: '',
        parentAsin: '',
        childAsins: '',
        isVariation: 0,
        sku: '',
        standardProductType: '',
        standardProductId: '',
        title: '',
        brand: '',
        manufacturer: '',
        description: '',
        bulletPoint: '',
        conditionType: '',
        conditionValue: '',
        standardPrice: 0,
        listingPricing: 0,
        switchFulfillmentTo: '',
        fulfillmentChannel: '',
        partNumber: '',
        mainImage: '',
        saleStartDate: '',
        saleEndDate: '',
        saleSalePrice: '',
        quantity: 0,
        openDate: '',
        itemDimensions: '',
        packageDimensions: '',
        standardPriceCurrency: '',
        specifics: '',
        variationChildStr: '',
        onlineStatus: '',
        saleNum: 0,
        fnsku: '',
        warehousing: 0,
        unsellable: 0,
        reservedQty: 0,
        purchaseCost: 0,
        headTripCost: 0,
        shipCost: '',
        priceFeedStatus: '',
        inventoryFeedStatus: '',
        updateStandardPrice: '',
        updateQuantity: '',
        rating: '',
        ratingCount: '',
        rank: '',
        bsrRank: '',
        commodityId: 0,
        mathCommodityTime: '',
        maxShipmentQty: '',
        devId: '',
        createId: '',
        updateId: '',
        lpriceUpdateTime: '',
        lpriceStatus: '',
        childList: '',
        salePrices: 0,
        adCosts: 0,
        adCostsSd: 0,
        adCostsSb: 0,
        adCostsSbv: 0,
        adCostsSum: 0,
        currency: '',
        costCurrency: '',
        commoditySku: '',
        commodityName: '',
        shopName: '',
        siteName: '',
        sellingPartnerId: '',
        domain: '',
        childNum: '',
        devName: '',
        variationStrMap: '',
        listingPrice: 0,
        shippingPrice: 0,
        buyBoxWinner: '',
        buyBoxCurrency: '',
        totalFee: 0,
        referralFee: 0,
        variableClosingFee: 0,
        perItemFee: 0,
        fbaFees: 0,
        feeCurrency: '',
        isPutAsh: '',
        bsrRankStr: '',
        vartationStr: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
      let res
      switch (type.value) {
        case 'create':
          res = await createErpListingDetail(formData.value)
          break
        case 'update':
          res = await updateErpListingDetail(formData.value)
          break
        default:
          res = await createErpListingDetail(formData.value)
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
