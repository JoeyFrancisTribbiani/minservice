// 自动生成模板SpConfig
package autocode

import (
	"github.com/JoeyFrancisTribbiani/minerpserver/global"
)

// SpConfig 结构体
// 如果含有time.Time 请自行import time包
type SpConfig struct {
      global.GVA_MODEL
      CreateBy  *int `json:"createBy" form:"createBy" gorm:"column:create_by;comment:;size:10;"`
      UpdateBy  *int `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:;size:10;"`
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:userid;size:20;"`
      MarketplaceIds  string `json:"marketplaceIds" form:"marketplaceIds" gorm:"column:marketplace_ids;comment:市场IDs;size:200;"`
      ClientId  string `json:"clientId" form:"clientId" gorm:"column:client_id;comment:SP-API客户端id;size:100;"`
      ClientSecret  string `json:"clientSecret" form:"clientSecret" gorm:"column:client_secret;comment:SP-API客户端密钥;size:100;"`
      RefreshToken  string `json:"refreshToken" form:"refreshToken" gorm:"column:refresh_token;comment:刷新token;size:500;"`
      AccessKeyId  string `json:"accessKeyId" form:"accessKeyId" gorm:"column:access_key_id;comment:AWS IAM User Access Key Id;size:100;"`
      SecretKey  string `json:"secretKey" form:"secretKey" gorm:"column:secret_key;comment:AWS IAM User Secret Key;size:100;"`
      Region  string `json:"region" form:"region" gorm:"column:region;comment:AWS Region;size:100;"`
      RoleArn  string `json:"roleArn" form:"roleArn" gorm:"column:role_arn;comment:AWS IAM Role ARN;size:100;"`
}


// TableName SpConfig 表名
func (SpConfig) TableName() string {
  return "sp_config"
}

