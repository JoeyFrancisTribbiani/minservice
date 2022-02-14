<template>
  <div>
    <div class="gva-form-box">
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
          <el-input-number v-model="formData.standardPrice" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="8.99:">
          <el-input-number v-model="formData.listingPricing" :precision="2" clearable></el-input-number>
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
          <el-input-number v-model="formData.listingPrice" :precision="2" clearable></el-input-number>
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
          <el-input-number v-model="formData.totalFee" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="1.35:">
          <el-input-number v-model="formData.referralFee" :precision="2" clearable></el-input-number>
        </el-form-item>
        <el-form-item label="0:">
          <el-input v-model.number="formData.variableClosingFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="0:">
          <el-input v-model.number="formData.perItemFee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="3.47:">
          <el-input-number v-model="formData.fbaFees" :precision="2" clearable></el-input-number>
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
  name: 'ErpListingDetail'
}
</script>

<script setup>
import {
  createErpListingDetail,
  updateErpListingDetail,
  findErpListingDetail
} from '@/api/erpListingDetail'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findErpListingDetail({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reerpListingDetail
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
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
