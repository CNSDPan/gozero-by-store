CREATE TABLE `users` (
 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
 `user_id` bigint(20) NOT NULL COMMENT '用户ID',
 `status` tinyint(1) DEFAULT '1' COMMENT '1-启用、2-禁用',
 `mobile` int(11) DEFAULT '0' COMMENT '手机号',
 `password` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '密码',
 `name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '昵称',
 `avatar` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '头像',
 `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
 `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
 PRIMARY KEY (`id`) USING BTREE,
 UNIQUE KEY `unx_uid` (`user_id`) USING BTREE,
 UNIQUE KEY `unx_mobile` (`mobile`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';