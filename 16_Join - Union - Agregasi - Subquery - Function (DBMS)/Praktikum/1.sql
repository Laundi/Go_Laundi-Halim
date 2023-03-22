CREATE TABLE operators(
    -> id INT PRIMARY KEY AUTO_INCREMENT,
    -> name VARCHAR(255) NOT NULL,
    -> created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    -> updated_at DATETIME DEFAULT CURRENT_TIMESTAMP);

CREATE TABLE product_types (
    -> id INT PRIMARY KEY AUTO_INCREMENT
    -> name VARCHAR(255),
    -> created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    -> updated_at DATETIME DEFAULT CURRENT_TIMESTAMP);

CREATE TABLE product_descriptions (
    -> id INT PRIMARY KEY AUTO_INCREMENT,
    -> description TEXT,
    -> created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    -> updated_at DATETIME DEFAULT CURRENT_TIMESTAMP);

CREATE TABLE payment_methods (
    -> id INT PRIMARY KEY AUTO_INCREMENT,
    -> name VARCHAR(255),
    -> status SMALLINT,
    -> created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    -> updated_at DATETIME DEFAULT CURRENT_TIMESTAMP);

CREATE TABLE users (
    -> id INT PRIMARY KEY AUTO_INCREMENT
    -> status SMALLINT,
    -> dob DATE,
    -> gender CHAR(1),
    -> created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    -> updated_at DATETIME DEFAULT CURRENT_TIMESTAMP);

CREATE TABLE transaction_details (
    -> transaction_id INT AUTO_INCREMENT,
    -> product_id INT,
    -> status VARCHAR(10),
    -> qty INT,
    -> price NUMERIC(25,2),
    -> created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    -> updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    -> PRIMARY KEY (transaction_id, product_id)
    -> );

CREATE TABLE products (
    -> id INT PRIMARY KEY AUTO_INCREMENT,
    -> product_type_id INT,
    -> operator_id INT,
    -> code VARCHAR(50),
    -> name VARCHAR(100),
    -> status SMALLINT,
    -> created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    -> updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    -> FOREIGN KEY (product_type_id) REFERENCES product_type(id),
    -> FOREIGN KEY (operator_id) REFERENCES operators(id));

CREATE TABLE transaction (
    -> id INT PRIMARY KEY AUTO_INCREMENT,
    -> user_id INT,
    -> payment_method_id INT,
    -> status VARCHAR(10),
    -> total_qty INT,
    -> total_price NUMERIC(25,2),
    -> created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    -> updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    -> FOREIGN KEY (user_id) REFERENCES users(id),
    -> FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id));

1a. INSERT INTO operators (name) values ("A");
    INSERT INTO operators (name) values ("B");
    INSERT INTO operators (name) values ("C");
    INSERT INTO operators (name) values ("D");
    INSERT INTO operators (name) values ("E");

1b. INSERT INTO product_types (name) values ("Elektronik");
    INSERT INTO product_types (name) values ("Pakaian"); 
    INSERT INTO product_types (name) values ("Kecantikan");

1c-e.   INSERT INTO products (product_type_id, operator_id, code, name, status) values (2, 1, "AAB", "T-Shirt A, 1);
        INSERT INTO products (product_type_id, operator_id, code, name, status) values (3, 4, "AAC", "Skincare A, 1);
        INSERT INTO products (product_type_id, operator_id, code, name, status) values (1, 2, "AA1", "Acer Predator, 1);
        INSERT INTO products (product_type_id, operator_id, code, name, status) values (1, 1, "AA2", "Lenovo Legion, 1);
        INSERT INTO products (product_type_id, operator_id, code, name, status) values (2, 2, "AB1", "Kenzo, 1);
        INSERT INTO products (product_type_id, operator_id, code, name, status) values (2, 3, "AB2", "ADLV, 1);
        INSERT INTO products (product_type_id, operator_id, code, name, status) values (3, 2, "AC1", "Somethinc, 1);
        INSERT INTO products (product_type_id, operator_id, code, name, status) values (3, 1, "AC2", "Scarlett, 1);

1f. INSERT INTO product_descriptions (description) values ("T-shirt Kenzo");
    INSERT INTO product_descriptions (description) values ("Skincare Somethinc");
    INSERT INTO product_descriptions (description) values ("Laptop Acer");
    INSERT INTO product_descriptions (description) values ("Laptop Lenovo");
    INSERT INTO product_descriptions (description) values ("Baju Kenzo 2.0");
    INSERT INTO product_descriptions (description) values ("Baju ADLV");
    INSERT INTO product_descriptions (description) values ("Skincare Somethinc 2.0");
    INSERT INTO product_descriptions (description) values ("Skincare Scarlett");

1g. INSERT INTO payment_descriptions (description) values ("Gopay,1);
    INSERT INTO payment_descriptions (description) values ("OVO,1);
    INSERT INTO payment_descriptions (description) values ("Mobile Banking,1);

1h. INSERT INTO user (status, dob, gender), values (1, "2001-03-04", "M");
    INSERT INTO user (status, dob, gender), values (1, "2003-07-05", "F");
    INSERT INTO user (status, dob, gender), values (1, "2003-07-06", "F");
    INSERT INTO user (status, dob, gender), values (1, "2008-12-06", "F");
    INSERT INTO user (status, dob, gender), values (1, "1998-02-22", "M");

2a. select * from users where gender = "M"

2b. select * from product where id =3; 

2d. select count(*) from users where gender = "F";

2f. select * from products limit 5;

3a. UPDATE products set name = "product dummy" where id = 1;

3b. UPDATE transaction_details set qty = 3 where product_id = 1;

4a. DELETE FROM products where id = 1;

4b. DELETE FROM products where product_type_id

5a. select * from transaction where user_id = 1 or user_id = 2

5b. select sum(total_price) from transaction where user_id = 1;

5c. select count(*) from transaction_details a left join product b on b.id = a.product_id where b.product_type_id

5d. SELECT * FROM product INNER JOIN product_types ON product.product_type_id = product_type.id;

5e. deliminter $$
    -> create trigger delete_transaction_detail
    -> AFTER DELETE ON transactions
    -> FOR EACH ROM
    -> BEGIN
    -> DELETE FROM TRANSACTION_DETAILS
    -> WHERE transaction_id = OLD.id;
    -> END;
    -> $$

5f. delimiter $$
    > create trigger update_total_qty
    -> AFTER DELETE ON transaction_details
    -> FOR EACH ROM
    -> BEGIN
    -> UPDATE transaction a
    -> set a.total_qty = (SELECT SUM(b.qty) from transaction_details where b.transaction_id = a.id)
    -> WHERE b.id = OLD.transaction_id;
    -> END;
    -> $$

5g. select * from product where id not in (select product_id from transaction details);