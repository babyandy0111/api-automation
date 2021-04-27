CREATE TABLE `point_earn` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL COMMENT 'ref: user.id',
  `point` int(11) unsigned NOT NULL COMMENT '點數',
  `description` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '點數相關描述',
  `expired_at` datetime NOT NULL COMMENT '點數有效期限',
  `is_expired` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否已過期',
  `operator_id` int(11) NOT NULL DEFAULT '0' COMMENT 'ref: sys_account.id',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_point_earn_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `point_use` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL COMMENT 'ref: user.id',
  `point` int(11) unsigned NOT NULL COMMENT '點數',
  `description` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '點數相關描述',
  `use_type` int(11) NOT NULL COMMENT '使用類型 1: gift shop, 2: donation',
  `relate_id` int(11) NOT NULL COMMENT '關聯 ID',
  `operator_id` int(11) NOT NULL DEFAULT '0' COMMENT 'ref: sys_account.id',
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_point_use_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


CREATE TABLE `user_point` (
  `user_id` int(11) NOT NULL COMMENT 'ref: user.id',
  `point` int(11) unsigned NOT NULL COMMENT '總點數',
  `soon_expired_point` int(11) unsigned NOT NULL COMMENT '即將過期的點數',
  `version` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '樂觀鎖',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `voucher` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(150) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '商品名稱',
  `image_url` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '商品圖片連結',
  `point` int(11) unsigned NOT NULL COMMENT '需要多少點數兌換',
  `exchange_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '兌換種類 1: 寄送, 2: 親領, 3: Redeem code',
  `note` varchar(350) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '說明欄位 (使用步驟須知)',
  `term_condition` varchar(350) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '注意事項',
  `remark` varchar(350) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '兌換說明',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '1: 上架, 2: 下架',
  `expired_dt` datetime NOT NULL COMMENT '使用期限',
  `is_expired` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否已過期',
  `country` varchar(4) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '國家 (ISO 3166)',
  `sort` int(11) NOT NULL COMMENT '排序',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `voucher_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `voucher_id` int(11) NOT NULL COMMENT 'ref: voucher.id',
  `user_id` int(11) NOT NULL COMMENT 'ref: user.id',
  `serial_num` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '票券序號',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0: 未使用, 1: 已使用',
  `expired_dt` datetime NOT NULL COMMENT '票券有效期限',
  `is_expired` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否已過期',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_voucher_user_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;