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
        <el-form-item label="userid:">
          <el-input v-model.number="formData.userId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="市场IDs:">
          <el-input v-model="formData.marketplaceIds" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="SP-API客户端id:">
          <el-input v-model="formData.clientId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="SP-API客户端密钥:">
          <el-input v-model="formData.clientSecret" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="刷新token:">
          <el-input v-model="formData.refreshToken" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="AWS IAM User Access Key Id:">
          <el-input v-model="formData.accessKeyId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="AWS IAM User Secret Key:">
          <el-input v-model="formData.secretKey" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="AWS Region:">
          <el-input v-model="formData.region" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="AWS IAM Role ARN:">
          <el-input v-model="formData.roleArn" clearable placeholder="请输入" />
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
  name: 'SpConfig'
}
</script>

<script setup>
import {
  createSpConfig,
  updateSpConfig,
  findSpConfig
} from '@/api/spConfig'

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
        userId: 0,
        marketplaceIds: '',
        clientId: '',
        clientSecret: '',
        refreshToken: '',
        accessKeyId: '',
        secretKey: '',
        region: '',
        roleArn: '',
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findSpConfig({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.respConfig
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
          res = await createSpConfig(formData.value)
          break
        case 'update':
          res = await updateSpConfig(formData.value)
          break
        default:
          res = await createSpConfig(formData.value)
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
