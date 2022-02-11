CREATE TABLE `sp_config`(
    id int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `create_by` int(11) unsigned DEFAULT NULL,
    `update_by` int(11) unsigned DEFAULT NULL,
    user_id bigint unsigned NOT NULL COMMENT 'userid',
    marketplace_ids  varchar(200) NULL COMMENT '市场IDs',
    client_id varchar(100) NULL COMMENT 'SP-API客户端id',
    client_secret varchar(100) NULL COMMENT 'SP-API客户端密钥',
    refresh_token varchar(500) NULL COMMENT '刷新token',
    access_key_id varchar(100) NULL COMMENT 'AWS IAM User Access Key Id',
    secret_key varchar(100) NULL COMMENT 'AWS IAM User Secret Key',
    region varchar(100) NULL COMMENT 'AWS Region',
    role_arn varchar(100) NULL COMMENT 'AWS IAM Role ARN',
    PRIMARY KEY (`id`) USING BTREE,
    CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES sys_users(id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='批次成本合计';