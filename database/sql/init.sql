use test_database;

create table products (
product_key VARCHAR(32) primary key,
product_value VARCHAR(32),
created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
deleted_at datetime NULL DEFAULT NULL
) engine = MyISAM default charset = utf8;

INSERT INTO products (product_key,product_value) VALUES ('401','unauthorized');