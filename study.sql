CREATE TABLE `bp_peiqi_version`
(
    `id`          int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `version`     varchar(100) NOT NULL DEFAULT '' COMMENT '配奇版本',
    `type`        varchar(50)  NOT NULL COMMENT '版本类型，默认空字符串，线上online，灰度gray',
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY           `version_index` (`version`)
) ENGINE=InnoDB AUTO_INCREMENT=66 DEFAULT CHARSET=utf8 COMMENT='配奇版本管理'