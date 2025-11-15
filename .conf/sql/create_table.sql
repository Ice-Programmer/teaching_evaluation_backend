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

-- 管理员表
CREATE TABLE IF NOT EXISTS admin
(
    id        BIGINT AUTO_INCREMENT COMMENT 'id' PRIMARY KEY,
    username  VARCHAR(256)      NOT NULL COMMENT '账号',
    password  VARCHAR(256)      NOT NULL COMMENT '密码',
    create_at BIGINT  DEFAULT 0 NOT NULL COMMENT '创建时间',
    deleted_at TINYINT DEFAULT 0 NOT NULL COMMENT '删除时间',
    UNIQUE INDEX idx_username (username)
) COMMENT '管理员表' CHARSET = utf8mb4
                     COLLATE = utf8mb4_unicode_ci;

-- 学生表
CREATE TABLE IF NOT EXISTS student
(
    id             BIGINT AUTO_INCREMENT COMMENT 'id' PRIMARY KEY,
    student_number VARCHAR(128)      NOT NULL COMMENT '学生学号',
    student_name   VARCHAR(128)      NOT NULL COMMENT '学生姓名',
    class_id       BIGINT            NOT NULL COMMENT '学生班级id',
    major          TINYINT DEFAULT 0 NOT NULL COMMENT '专业（0 - 计算机 1 - 自动化）',
    grade          INT               NOT NULL COMMENT '学生年级',
    status         TINYINT DEFAULT 0 NOT NULL COMMENT '学生状态（0 - 正常使用 1 - 拒绝访问）',
    create_at      BIGINT  DEFAULT 0 NOT NULL COMMENT '创建时间',
    deleted_at     TINYINT DEFAULT 0 NOT NULL COMMENT '是否删除',
    INDEX idx_student_number (student_number),
    INDEX idx_class_id (class_id),
    INDEX idx_grade_major (grade, major)
) COMMENT '学生表' CHARSET = utf8mb4
                   COLLATE = utf8mb4_unicode_ci;
