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