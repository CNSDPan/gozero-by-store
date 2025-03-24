CREATE TABLE `chat_log` (
`user_id` bigint(20) NOT NULL COMMENT '用户ID',
`store_id` bigint(20) DEFAULT '1' COMMENT '店铺ID',
`message` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '消息',
`timestamp` bigint(20) DEFAULT '0' COMMENT '记录时间;微秒',
`created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
`updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
KEY `idx_uid` (`user_id`) USING BTREE,
KEY `idx_storeid` (`store_id`),
KEY `idx_time` (`timestamp`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='聊天记录表';