-- 创建库
CREATE DATABASE IF NOT EXISTS teaching_evaluation CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

-- 切换库
USE teaching_evaluation;

-- 班级表
CREATE TABLE IF NOT EXISTS student_class
(
    id              BIGINT AUTO_INCREMENT COMMENT 'id' PRIMARY KEY,
    class_number    varchar(128)      not null comment '班级编号',
    created_at      BIGINT  DEFAULT 0 NOT NULL COMMENT '创建时间',
    created_op_name VARCHAR(128)      NOT NULL COMMENT '创建者',
    created_op_id   VARCHAR(128)      NOT NULL COMMENT '创建者 id',
    updated_at      BIGINT  DEFAULT 0 NOT NULL COMMENT '更新时间',
    updated_op_name VARCHAR(128)      NOT NULL COMMENT '更新者',
    updated_op_id   VARCHAR(128)      NOT NULL COMMENT '更新者 id',
    deleted_at      TINYINT DEFAULT 0 NOT NULL COMMENT '删除时间',
    UNIQUE KEY uk_class_number (class_number)
) COMMENT '班级表' CHARSET = utf8mb4
                   COLLATE = utf8mb4_unicode_ci;
