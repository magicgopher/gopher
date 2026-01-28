#####################
## 一对多查询示例SQL
#####################
# 删除订单表
DROP TABLE IF EXISTS orders;
# 删除用户表
DROP TABLE IF EXISTS users;
# 创建用户表
CREATE TABLE users
(
    id    BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    name  VARCHAR(50)  NOT NULL COMMENT '姓名',
    email VARCHAR(100) NOT NULL COMMENT '邮件'
) COMMENT '用户表';
# 创建订单表
CREATE TABLE orders
(
    id       BIGINT AUTO_INCREMENT PRIMARY KEY,
    order_no VARCHAR(50) NOT NULL,
    amount   DECIMAL(10, 2),
    user_id  BIGINT,
    FOREIGN KEY (user_id) REFERENCES users (id)
) COMMENT '订单表';
# 用户表数据
INSERT INTO users (id, name, email)
VALUES (1, '张三', 'zhangsan@example.com');
INSERT INTO users (id, name, email)
VALUES (2, '李四', 'lisi@example.com');
# 订单表数据
INSERT INTO orders (order_no, amount, user_id)
VALUES ('ORD-1001', 99.50, 1);
INSERT INTO orders (order_no, amount, user_id)
VALUES ('ORD-1002', 199.00, 1);
INSERT INTO orders (order_no, amount, user_id)
VALUES ('ORD-2001', 50.00, 2);
INSERT INTO orders (order_no, amount, user_id)
VALUES ('ORD-2002', 230.50, 2);
INSERT INTO orders (order_no, amount, user_id)
VALUES ('ORD-2004', 155.50, 2);

#####################
## 多对多查询示例SQL
#####################
DROP TABLE IF EXISTS student_courses;
DROP TABLE IF EXISTS courses;
DROP TABLE IF EXISTS students;
# 学生表
CREATE TABLE students
(
    id   BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '技术主键 (内部关联用)',
    sid  VARCHAR(20) NOT NULL COMMENT '学号 (业务主键)',
    name VARCHAR(50) NOT NULL COMMENT '姓名',
    -- 必须保证学号唯一，否则逻辑会有问题
    UNIQUE KEY uk_sid (sid)
) COMMENT '学生表';
#   课程表
CREATE TABLE courses
(
    id      BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '技术主键',
    cid     VARCHAR(20)  NOT NULL COMMENT '课程编号 (业务主键)',
    title   VARCHAR(100) NOT NULL COMMENT '课程名称',
    credits INT COMMENT '学分',
    UNIQUE KEY uk_cid (cid)
) COMMENT '课程表';
-- 中间关联表
CREATE TABLE student_courses
(
    student_id BIGINT NOT NULL COMMENT '对应 students.id',
    course_id  BIGINT NOT NULL COMMENT '对应 courses.id',
    PRIMARY KEY (student_id, course_id),
    -- 外键指向两张表的主键 id (BIGINT 对 BIGINT)
    CONSTRAINT fk_sc_student FOREIGN KEY (student_id) REFERENCES students (id),
    CONSTRAINT fk_sc_course FOREIGN KEY (course_id) REFERENCES courses (id)
) COMMENT '学生选课中间表';
-- 模拟真实的学号格式：S + 年份 + 序号
INSERT INTO students (sid, name)
VALUES ('S2023001', '张三'); -- id=1
INSERT INTO students (sid, name)
VALUES ('S2023002', '李四'); -- id=2
INSERT INTO students (sid, name)
VALUES ('S2023003', '王五'); -- id=3 (学霸)
INSERT INTO students (sid, name)
VALUES ('S2023004', '赵六');
-- id=4 (无课)
-- 3.2 插入课程
-- 模拟课程编号：CS=计算机, MA=数学, EN=英语
INSERT INTO courses (cid, title, credits)
VALUES ('CS-101', 'Go语言编程', 4); -- id=1
INSERT INTO courses (cid, title, credits)
VALUES ('CS-102', 'MySQL进阶', 3); -- id=2
INSERT INTO courses (cid, title, credits)
VALUES ('MA-001', '高等数学', 5); -- id=3
INSERT INTO courses (cid, title, credits)
VALUES ('EN-001', '专业英语', 2);
-- id=4
-- 3.3 建立关联 (使用 id 关联)
-- 张三(id=1): 选了 Go(id=1) 和 数学(id=3)
INSERT INTO student_courses (student_id, course_id)
VALUES (1, 1);
INSERT INTO student_courses (student_id, course_id)
VALUES (1, 3);
-- 李四(id=2): 选了 MySQL(id=2)
INSERT INTO student_courses (student_id, course_id)
VALUES (2, 2);
-- 王五(id=3): 全选 (学霸测试)
INSERT INTO student_courses (student_id, course_id)
VALUES (3, 1);
INSERT INTO student_courses (student_id, course_id)
VALUES (3, 2);
INSERT INTO student_courses (student_id, course_id)
VALUES (3, 3);
INSERT INTO student_courses (student_id, course_id)
VALUES (3, 4);