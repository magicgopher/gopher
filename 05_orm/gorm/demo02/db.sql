CREATE TABLE tb_user (
                         id BIGSERIAL PRIMARY KEY ,
                         uid VARCHAR(64) NOT NULL UNIQUE ,
                         name VARCHAR(50) NOT NULL DEFAULT '',
                         email VARCHAR(100) NOT NULL DEFAULT '',
                         created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
                         updated_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

INSERT INTO tb_user (uid, name, email, created_at, updated_at)  VALUES
                                                                    ('901241512345600001', '张三丰', 'zhangsanfeng@example.com', '2024-01-15 10:23:11+08', '2025-11-10 09:12:33+08'),
                                                                    ('901241512345600002', '李四喜', 'lisixi@example.com', '2024-02-20 14:30:22+08', '2025-10-25 18:45:10+08'),
                                                                    ('901241512345600003', '王五桐', 'wangwutong@example.com', '2024-03-08 09:15:47+08', '2025-11-01 11:11:11+08');