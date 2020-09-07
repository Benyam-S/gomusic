DROP DATABASE gomusic;

CREATE DATABASE gomusic;

use gomusic;

CREATE TABLE customers (
    id INT PRIMARY KEY AUTO_INCREMENT,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    cc_customer_id VARCHAR(255),
    logged_in BOOLEAN,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME
);

CREATE TABLE products (
    id INT PRIMARY KEY AUTO_INCREMENT,
    img_src VARCHAR(255),
    img_alt VARCHAR(255),
    description VARCHAR(255),
    product_name VARCHAR(255),
    price FLOAT,
    promotion FLOAT,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME
);

CREATE TABLE orders (
    id INT PRIMARY KEY AUTO_INCREMENT,
    customer_id INT,
    product_id INT,
    price FLOAT,
    purchase_date DATETIME,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME
);