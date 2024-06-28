
use shuyi;
drop table if exists content;
CREATE TABLE `content` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `content_id` bigint NOT NULL DEFAULT 0 COMMENT '内容ID',
    `text` longtext COMMENT '详细内容',
    `content_type` int NOT NULL DEFAULT 0 COMMENT '内容类型',
    `title` varchar(128) NOT NULL DEFAULT '' COMMENT '内容标题',
    `backup` text COMMENT '备份的信息',
    `extra` text COMMENT '一些额外的信息',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `created_by` varchar(128) NOT NULL DEFAULT '' COMMENT '创建人',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `updated_by` varchar(128) NOT NULL DEFAULT '' COMMENT '最后更新人',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_ci` (`content_id`)
) ENGINE = InnoDB CHARSET = utf8mb4 AUTO_INCREMENT = 2020 COMMENT '内容';

