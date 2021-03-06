CREATE TABLE `erp_listing_detail`(
    id int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `create_by` int(11) unsigned DEFAULT NULL,
    `update_by` int(11) unsigned DEFAULT NULL,
    puid bigint(11) NULL COMMENT '2231',
    shop_id bigint(11) NULL COMMENT '4462',
    marketplace_id varchar(100) NULL COMMENT 'ATVPDKIKX0DER',
    full_cid varchar(100) NULL COMMENT '',
    parent_id bigint(11) NULL COMMENT '4762272',
    dxm_state varchar(100) NULL COMMENT 'online_amazon',
    dxm_publish_state varchar(100) NULL COMMENT 'null',
    err_msg varchar(100) NULL COMMENT 'null',
    asin varchar(100) NULL COMMENT 'B08X4MVXS7',
    listing_id varchar(100) NULL COMMENT '0222Y6BRSTD',
    parent_asin varchar(100) NULL COMMENT 'null',
    child_asins varchar(100) NULL COMMENT 'null',
    is_variation bigint(11) NULL COMMENT '0',
    sku varchar(100) NULL COMMENT 'VQ-C5VT-Y5GH',
    standard_product_type varchar(100) NULL COMMENT 'null',
    standard_product_id varchar(100) NULL COMMENT 'null',
    title varchar(100) NULL COMMENT 'Outdoor Faucet Cover ',
    brand varchar(100) NULL COMMENT 'null',
    manufacturer varchar(100) NULL COMMENT 'null',
    description varchar(100) NULL COMMENT 'null',
    bullet_point varchar(100) NULL COMMENT 'null',
    condition_type varchar(100) NULL COMMENT 'null',
    condition_value varchar(100) NULL COMMENT 'null',
    standard_price Double NULL COMMENT '8.99',
    listing_pricing Double NULL COMMENT '8.99',
    switch_fulfillment_to varchar(100) NULL COMMENT 'AFN',
    fulfillment_channel varchar(100) NULL COMMENT 'null',
    part_number varchar(100) NULL COMMENT 'null',
    main_image varchar(100) NULL COMMENT 'https://m.media-amazon.com/images/I/51QOuKgk22L._SL75_.jpg',
    sale_start_date varchar(100) NULL COMMENT 'null',
    sale_end_date varchar(100) NULL COMMENT 'null',
    sale_sale_price varchar(100) NULL COMMENT 'null',
    quantity bigint(11) NULL COMMENT '34',
    open_date varchar(100) NULL COMMENT '2021-02-22 04:20:10',
    item_dimensions varchar(100) NULL COMMENT 'null',
    package_dimensions varchar(100) NULL COMMENT 'null',
    standard_price_currency varchar(100) NULL COMMENT 'USD',
    specifics varchar(100) NULL COMMENT 'null',
    variation_child_str varchar(100) NULL COMMENT 'null',
    online_status varchar(100) NULL COMMENT 'Active',
    sale_num bigint(11) NULL COMMENT '0',
    fnsku varchar(100) NULL COMMENT 'X002TADZZ7',
    warehousing bigint(11) NULL COMMENT '0',
    unsellable bigint(11) NULL COMMENT '1',
    reserved_qty bigint(11) NULL COMMENT '0',
    purchase_cost bigint(11) NULL COMMENT '12',
    head_trip_cost bigint(11) NULL COMMENT '999',
    ship_cost varchar(100) NULL COMMENT 'null',
    price_feed_status varchar(100) NULL COMMENT 'null',
    inventory_feed_status varchar(100) NULL COMMENT 'null',
    update_standard_price varchar(100) NULL COMMENT 'null',
    update_quantity varchar(100) NULL COMMENT 'null',
    rating varchar(100) NULL COMMENT 'null',
    rating_count varchar(100) NULL COMMENT 'null',
    `rank` varchar(100) NULL COMMENT 'null',
    bsr_rank varchar(100) NULL COMMENT 'null',
    commodity_id bigint(11) NULL COMMENT '470356',
    math_commodity_time varchar(100) NULL COMMENT 'null',
    max_shipment_qty varchar(100) NULL COMMENT 'null',
    dev_id varchar(100) NULL COMMENT '9207',
    create_id varchar(100) NULL COMMENT 'null',
    update_id varchar(100) NULL COMMENT 'null',
    lprice_update_time varchar(100) NULL COMMENT 'null',
    lprice_status varchar(100) NULL COMMENT 'null',
    child_list varchar(100) NULL COMMENT 'null',
    sale_prices bigint(11) NULL COMMENT '0',
    ad_costs bigint(11) NULL COMMENT '0',
    ad_costs_sd bigint(11) NULL COMMENT '0',
    ad_costs_sb bigint(11) NULL COMMENT '0',
    ad_costs_sbv bigint(11) NULL COMMENT '0',
    ad_costs_sum bigint(11) NULL COMMENT '0',
    currency varchar(100) NULL COMMENT 'USD',
    cost_currency varchar(100) NULL COMMENT 'CNY',
    commodity_sku varchar(100) NULL COMMENT 'testsku123123',
    commodity_name varchar(100) NULL COMMENT '????????????',
    shop_name varchar(100) NULL COMMENT 'A3LR0SXN8QVIE0-na-US',
    site_name varchar(100) NULL COMMENT '??????',
    selling_partner_id varchar(100) NULL COMMENT 'null',
    domain varchar(100) NULL COMMENT 'www.amazon.com',
    child_num varchar(100) NULL COMMENT 'null',
    dev_name varchar(100) NULL COMMENT '?????????1',
    variation_str_map varchar(100) NULL COMMENT 'null',
    listing_price Double NULL COMMENT '8.99',
    shipping_price bigint(11) NULL COMMENT '0',
    buy_box_winner varchar(100) NULL COMMENT 'true',
    buy_box_currency varchar(100) NULL COMMENT 'USD',
    total_fee Double NULL COMMENT '4.82',
    referral_fee Double NULL COMMENT '1.35',
    variable_closing_fee bigint(11) NULL COMMENT '0',
    per_item_fee bigint(11) NULL COMMENT '0',
    fba_fees Double NULL COMMENT '3.47',
    fee_currency varchar(100) NULL COMMENT 'USD',
    is_put_ash varchar(100) NULL COMMENT 'null',
    bsr_rank_str varchar(100) NULL COMMENT 'null',
    vartation_str varchar(100) NULL COMMENT '',
    PRIMARY KEY (`id`) USING BTREE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='amazon listing detail';
