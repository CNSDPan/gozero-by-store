CREATE TABLE `store_member` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`store_member_id` bigint(20) NOT NULL COMMENT '会员ID',
`store_id` bigint(20) NOT NULL COMMENT '店铺ID',
`user_id` bigint(20) NOT NULL COMMENT '用户ID',
`created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
`updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
PRIMARY KEY (`id`) USING BTREE,
UNIQUE KEY `unx_smid` (`store_member_id`) USING BTREE,
KEY `idx_s_u_id` (`store_id`,`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='店铺会员表';