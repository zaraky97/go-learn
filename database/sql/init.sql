use test_database;

create table products (
ID int  AUTO_INCREMENT NOT NULL primary key,
ProductKey VARCHAR(32),
ProductValue VARCHAR(32),
created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
deleted_at datetime NULL DEFAULT NULL
) engine = MyISAM default charset = utf8;

INSERT INTO products (ProductKey,ProductValue) VALUES ('401','unauthorized');