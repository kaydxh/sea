# init database for seatlet

DELIMITER ;;
CREATE DATABASE /*!32312 IF NOT EXISTS*/ `sealet` /*!40100 DEFAULT CHARACTER SET utf8*/;
USE `sealet`;

CREATE TABLE IF NOT EXISTS `task`
(
    # 必备字段
    `id`                  BIGINT UNSIGNED  NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `created_at`          DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`          DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最近更新时间',
    # 软删除
    `is_deleted`          TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'soft delete, 0 for not deleted, 1 for deleted',
    `deleted_at`          DATETIME                  DEFAULT NULL COMMENT '最近删除时间',
    # 版本控制
    `version`             INT              NOT NULL DEFAULT 0 COMMENT '记录版本',

    `task_name`           VARCHAR(200)     NOT NULL DEFAULT '' COMMENT '任务名称',
    `task_id`             VARCHAR(200)     NOT NULL DEFAULT '' COMMENT '任务ID',
    `task_type`           INT UNSIGNED     NOT NULL DEFAULT 0  COMMENT '任务类型 1 任务a  2 任务b',
    `task_status`         INT UNSIGNED     NOT NULL DEFAULT 0  COMMENT '任务状态 1 未开始 2 处理中 3 成功 4 失败',

    PRIMARY KEY `pk_id` (`id`),
    UNIQUE KEY `idx_task_id` (`task_id`),
    KEY `idx_task_type` (`task_type`),
    KEY `idx_task_name` (`task_name`),
    KEY `idx_task_status` (`task_status`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COMMENT = 'task table';

SHOW TABLES;
