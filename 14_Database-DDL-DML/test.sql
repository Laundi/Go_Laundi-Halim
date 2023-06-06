CREATE DATABASE alta_online_shop;

-- start

CREATE TABLE `users` (
  `id` VARCHAR(33) PRIMARY KEY,
  `name` VARCHAR(51),
  `address` VARCHAR(255),
  `birthday` DATE,
  `status` VARCHAR(21),
  `gender` VARCHAR(1),
  `admin` BOOLEAN,
  `created_at` DATE,
  `updated_at` DATE
);

CREATE TABLE `products` (
  `id` VARCHAR(33) PRIMARY KEY,
  `name` VARCHAR(51),
  `type` VARCHAR(21),
  `description` VARCHAR(255),
  `price` BIGINT,
  `payment_methods` TEXT, -- CSV
  `created_at` DATE,
  `updated_at` DATE,
  `deleted_at` DATE
);

CREATE TABLE `invoices` (
  `id` VARCHAR(33) PRIMARY KEY,
  `user_id` VARCHAR(33),
  `payment_code` VARCHAR(31),
  `payment_method` VARCHAR(21),
  `created_at` DATE
);

CREATE TABLE `transactions` (
  `id` VARCHAR(33) PRIMARY KEY,
  `user_id` VARCHAR(33),
  `product_id` VARCHAR(33),
  `invoice_id` VARCHAR(33)
);

ALTER TABLE `invoices` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
ALTER TABLE `transactions` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
ALTER TABLE `transactions` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);
ALTER TABLE `transactions` ADD FOREIGN KEY (`invoice_id`) REFERENCES `invoices` (`id`);

-- end

CREATE TABLE `kurir`(
  `id` VARCHAR(33) PRIMARY KEY,
  `name` VARCHAR(21) UNIQUE NOT NULL,
  `created_at` DATE NOT NULL,
  `updated_at` DATE NOT NULL
);

ALTER TABLE `kurir` ADD COLUMN `ongkos_dasar` VARCHAR(33);
ALTER TABLE `kurir` RENAME `shipping`;
DROP TABLE `shipping`

-- payment_method_description
-- user address
-- user_payment_method_detail
