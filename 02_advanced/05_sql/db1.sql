# 创建user用户表
CREATE TABLE `user`
(
    `id`   BIGINT AUTO_INCREMENT COMMENT '主键',
    `name` VARCHAR(30) NOT NULL COMMENT '用户名称',
    `age`  TINYINT(1)  NOT NULL COMMENT '年龄',
    `sex`  VARCHAR(30) NOT NULL COMMENT '性别',
    PRIMARY KEY (id)
) COMMENT '用户表';

# 添加user表的数据
INSERT INTO `user` (`name`, `age`, `sex`)
VALUES ('张伟', 28, '男'),
       ('李娜', 24, '女'),
       ('王磊', 35, '男'),
       ('刘洋', 19, '女'),
       ('陈静', 42, '女'),
       ('杨明', 31, '男'),
       ('赵雪', 22, '女'),
       ('孙浩', 27, '男'),
       ('周婷', 29, '女'),
       ('吴刚', 38, '男'),
       ('徐晓雯', 25, '女'),
       ('胡兵', 33, '男'),
       ('郭佳怡', 20, '女'),
       ('林峰', 45, '男'),
       ('何雨欣', 26, '女'),
       ('马强', 30, '男'),
       ('朱莉', 23, '女'),
       ('韩宇航', 37, '男'),
       ('沈梦琪', 21, '女'),
       ('郑凯文', 32, '男');

# 创建账户表
CREATE TABLE `account`
(
    `id`       BIGINT AUTO_INCREMENT COMMENT '主键',
    `user_id`  BIGINT         NOT NULL COMMENT '关联的用户ID',
    `balance`  DECIMAL(15, 2) NOT NULL DEFAULT 0.00 COMMENT '账户余额',
    `currency` CHAR(3)        NOT NULL DEFAULT 'USD' COMMENT '货币代码',
    PRIMARY KEY (`id`),
    INDEX idx_user_id (`user_id`)
) COMMENT '账户表';

-- 插入两个测试账户
INSERT INTO `account` (`user_id`, `balance`, `currency`)
VALUES (1, 1000.00, 'CNY'),
       (2, 500.00, 'CNY');