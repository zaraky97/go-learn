CREATE TABLE products (
    id SERIAL,
    name VARCHAR(32) NOT NULL,
    price INT(11) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO products (name,price) VALUES ('PC',100000);