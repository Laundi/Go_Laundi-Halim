CREATE TABLE `product_types` (
  `id` INT(11) PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(255),
  `created_at` TIMESTAMP,
  `updated_at` TIMESTAMP
);

CREATE TABLE `operators` (
  `id` INT(11) PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(255),
  `created_at` TIMESTAMP,
  `updated_at` TIMESTAMP
);

CREATE TABLE `payment_methods` (
  `id` INT(11) PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(255),
  `status` SMALLINT,
  `created_at` TIMESTAMP,
  `updated_at` TIMESTAMP
);

CREATE TABLE `users` (
  `id` INT(11) PRIMARY KEY AUTO_INCREMENT,
  `status` SMALLINT,
  `dob` DATE,
  `gender` CHAR(1),
  `created_at` TIMESTAMP,
  `updated_at` TIMESTAMP
);

CREATE TABLE `products` (
  `id` INT(11) PRIMARY KEY AUTO_INCREMENT,
  `product_type_id` INT(11),
  `operator_id` INT(11),
  `code` VARCHAR(50),
  `name` VARCHAR(100),
  `status` SMALLINT,
  `created_at` TIMESTAMP,
  `updated_at` TIMESTAMP
);

CREATE TABLE `product_descriptions` (
  `id` INT(11) PRIMARY KEY AUTO_INCREMENT,
  `product_id` INT(11),
  `description` TEXT,
  `created_at` TIMESTAMP,
  `updated_at` TIMESTAMP
);

CREATE TABLE `transactions` (
  `id` INT(11) PRIMARY KEY AUTO_INCREMENT,
  `user_id` INT(11),
  `payment_method_id` INT(11),
  `status` VARCHAR(10),
  `total_qty` INT(11),
  `total_price` NUMERIC(25,2),
  `created_at` TIMESTAMP,
  `updated_at` TIMESTAMP
);

CREATE TABLE `transaction_details` (
  `transaction_id` INT(11),
  `product_id` INT(11),
  `status` VARCHAR(10),
  `qty` INT(11),
  `price` NUMERIC(25,2),
  `created_at` TIMESTAMP,
  `updated_at` TIMESTAMP
);

ALTER TABLE `products` ADD FOREIGN KEY (`product_type_id`) REFERENCES `product_types` (`id`);

ALTER TABLE `products` ADD FOREIGN KEY (`operator_id`) REFERENCES `operators` (`id`);

ALTER TABLE `product_descriptions` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);

ALTER TABLE `transactions` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `transactions` ADD FOREIGN KEY (`payment_method_id`) REFERENCES `payment_methods` (`id`);

ALTER TABLE `transaction_details` ADD FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`);

ALTER TABLE `transaction_details` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);
