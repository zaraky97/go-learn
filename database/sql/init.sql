create table products (
id int  AUTO_INCREMENT NOT NULL primary key,
productKey VARCHAR(32),
productValue VARCHAR(32)
) engine = MyISAM default charset = utf8;

INSERT INTO products (productKey,productValue) VALUES ('404','not found');