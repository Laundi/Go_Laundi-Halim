ALTER TABLE `operators` MODIFY COLUMN `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE `operators` MODIFY COLUMN `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE `payment_methods` MODIFY COLUMN `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE `payment_methods` MODIFY COLUMN `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE `product_descriptions` MODIFY COLUMN `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE `product_descriptions` MODIFY COLUMN `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE `product_types` MODIFY COLUMN `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE `product_types` MODIFY COLUMN `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE `products` MODIFY COLUMN `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE `products` MODIFY COLUMN `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE `transaction_details` MODIFY COLUMN `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE `transaction_details` MODIFY COLUMN `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE `transactions` MODIFY COLUMN `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE `transactions` MODIFY COLUMN `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE `users` MODIFY COLUMN `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE `users` MODIFY COLUMN `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

INSERT INTO `operators`(`name`) VALUES('a');
INSERT INTO `operators`(`name`) VALUES('b');
INSERT INTO `operators`(`name`) VALUES('c');
INSERT INTO `operators`(`name`) VALUES('d');
INSERT INTO `operators`(`name`) VALUES('e');

INSERT INTO `product_types`(`name`) VALUES('a');
INSERT INTO `product_types`(`name`) VALUES('b');
INSERT INTO `product_types`(`name`) VALUES('c');
INSERT INTO `product_types`(`name`) VALUES('d');

ALTER TABLE products AUTO_INCREMENT = 9;
ALTER TABLE operators AUTO_INCREMENT=20;

INSERT INTO `operators`(`name`) VALUES('z');
SELECT `id`, `name` FROM `operators` WHERE `name` = 'z'; -- 20, z

-- product_type_id,operator_id 1,3
INSERT INTO `products`(id,product_type_id,operator_id,code,name,status) VALUES(1, 1, 3, 'a', 'a', 0);
INSERT INTRO `products`(id,product_type_id,operator_id,code,name,status) VALUES(2, 1, 3, 'b', 'b', 0);

-- product_type_id,operator_id 2,1
INSERT INTO `products`(id,product_type_id,operator_id,code,name,status) VALUES(3, 2, 1, 'a', 'a', 0);
INSERT INTO `products`(id,product_type_id,operator_id,code,name,status) VALUES(4, 2, 1, 'b', 'b', 0);
INSERT INTO `products`(id,product_type_id,operator_id,code,name,status) VALUES(5, 2, 1, 'c', 'c', 0);

-- product_type_id,operator_id 3,4
INSERT INTO `products`(id,product_type_id,operator_id,code,name,status) VALUES(6, 3, 4, 'a', 'a', 0);
INSERT INTO `products`(id,product_type_id,operator_id,code,name,status) VALUES(7, 3, 4, 'b', 'b', 0);
INSERT INTO `products`(id,product_type_id,operator_id,code,name,status) VALUES(8, 3, 4, 'b', 'b', 0);

INSERT INTO `product_descriptions`(`product_id`, `description`) VALUES(1, 'a');
INSERT INTO `product_descriptions`(`product_id`, `description`) VALUES(2, 'b');
INSERT INTO `product_descriptions`(`product_id`, `description`) VALUES(3, 'c');
INSERT INTO `product_descriptions`(`product_id`, `description`) VALUES(4, 'd');
INSERT INTO `product_descriptions`(`product_id`, `description`) VALUES(5, 'e');
INSERT INTO `product_descriptions`(`product_id`, `description`) VALUES(6, 'f');
INSERT INTO `product_descriptions`(`product_id`, `description`) VALUES(7, 'g');
INSERT INTO `product_descriptions`(`product_id`, `description`) VALUES(8, 'h');

INSERT INTO `payment_methods`(`name`,`status`) VALUES('a', 0);
INSERT INTO `payment_methods`(`name`,`status`) VALUES('b', 0);
INSERT INTO `payment_methods`(`name`,`status`) VALUES('c', 0);

ALTER TABLE `users` ADD COLUMN `name` VARCHAR(50) NOT NULL;

INSERT INTO `users`(`name`, `status`,`dob`,`gender`) VALUES('a', 0, '1970-01-01 00:00:00', 'L');
INSERT INTO `users`(`name`, `status`,`dob`,`gender`) VALUES('b', 0, '1970-01-01 00:00:00', 'P');
INSERT INTO `users`(`name`, `status`,`dob`,`gender`) VALUES('c', 0, '1970-01-01 00:00:00', 'L');
INSERT INTO `users`(`name`, `status`,`dob`,`gender`) VALUES('d', 0, '1970-01-01 00:00:00', 'P');
INSERT INTO `users`(`name`, `status`,`dob`,`gender`) VALUES('e', 0, '1970-01-01 00:00:00', 'L');

INSERT INTO `transactions`(`user_id`, `payment_method_id`, `status`, `total_qty`, `total_price`) VALUES(1, 1, 0, 0, 0.00);
INSERT INTO `transactions`(`user_id`, `payment_method_id`, `status`, `total_qty`, `total_price`) VALUES(1, 1, 0, 1, 1.00);
INSERT INTO `transactions`(`user_id`, `payment_method_id`, `status`, `total_qty`, `total_price`) VALUES(1, 1, 0, 2, 2.00);

INSERT INTO `transactions`(`user_id`, `payment_method_id`, `status`, `total_qty`, `total_price`) VALUES(2, 2, 0, 0, 0.00);
INSERT INTO `transactions`(`user_id`, `payment_method_id`, `status`, `total_qty`, `total_price`) VALUES(2, 2, 0, 1, 1.00);
INSERT INTO `transactions`(`user_id`, `payment_method_id`, `status`, `total_qty`, `total_price`) VALUES(2, 2, 0, 2, 2.00);

INSERT INTO `transactions`(`user_id`, `payment_method_id`, `status`, `total_qty`, `total_price`) VALUES(3, 3, 0, 0, 0.00);
INSERT INTO `transactions`(`user_id`, `payment_method_id`, `status`, `total_qty`, `total_price`) VALUES(3, 3, 0, 1, 1.00);
INSERT INTO `transactions`(`user_id`, `payment_method_id`, `status`, `total_qty`, `total_price`) VALUES(3, 3, 0, 2, 2.00);

INSERT INTO `transactions`(`user_id`, `payment_method_id`, `status`, `total_qty`, `total_price`) VALUES(4, 1, 0, 0, 0.00);
INSERT INTO `transactions`(`user_id`, `payment_method_id`, `status`, `total_qty`, `total_price`) VALUES(4, 3, 0, 1, 1.00);
INSERT INTO `transactions`(`user_id`, `payment_method_id`, `status`, `total_qty`, `total_price`) VALUES(4, 3, 0, 2, 2.00);

INSERT INTO `transactions`(`user_id`, `payment_method_id`, `status`, `total_qty`, `total_price`) VALUES(5, 2, 0, 0, 0.00);
INSERT INTO `transactions`(`user_id`, `payment_method_id`, `status`, `total_qty`, `total_price`) VALUES(5, 2, 0, 1, 1.00);
INSERT INTO `transactions`(`user_id`, `payment_method_id`, `status`, `total_qty`, `total_price`) VALUES(5, 2, 0, 2, 2.00);

INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(1, 5, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(1, 4, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(1, 3, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(2, 5, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(2, 4, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(2, 3, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(3, 5, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(3, 4, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(3, 3, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(4, 5, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(4, 4, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(4, 3, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(5, 5, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(5, 4, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(5, 3, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(6, 5, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(6, 4, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(6, 3, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(7, 5, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(7, 4, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(7, 3, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(8, 5, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(8, 4, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(8, 3, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(9, 5, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(9, 4, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(9, 3, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(10, 5, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(10, 4, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(10, 3, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(11, 5, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(11, 4, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(11, 3, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(12, 5, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(12, 4, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(12, 3, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(13, 5, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(13, 4, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(13, 3, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(14, 5, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(14, 4, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(14, 3, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(15, 5, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(15, 4, 0, 1, 5.00);
INSERT INTO `transaction_details`(`transaction_id`, `product_id`, `status`, `qty`, `price`) VALUES(15, 3, 0, 1, 5.00);

SELECT `name` FROM `users` WHERE `gender` = 'L';
SELECT `name` FROM `products` WHERE `id` = 3;
-- Number(new Date('2023-03-19 00:00:00')) // 1679158800000 0e-3 ~ 10
-- Number(new Date('2023-03-12 00:00:00')) // 1678554000000 0e-3 ~ 10
SELECT * FROM `users` WHERE `name` LIKE '%a%' AND `created_at` BETWEEN '2023-03-12 00:00:00' AND '2023-03-19 00:00:00';
SELECT * FROM `users` WHERE `name` LIKE '%a%' AND `created_at` BETWEEN FROM_UNIXTIME(1678554000) AND FROM_UNIXTIME(1679158800);

SELECT COUNT(*) FROM `users` WHERE `gender` = 'P';

SELECT * FROM `products` LIMIT 5;

UPDATE `products` SET `name` = 'dummy' WHERE `id` = 1;
UPDATE `transactions` SET `total_qty` = 3 WHERE `id` = 3;

DELETE FROM `product_descriptions` WHERE `product_id` = 1;
DELETE FROM `products` WHERE `id` = 1;

DELETE FROM `product_descriptions` WHERE `product_id` IN (SELECT `id` FROM `products` WHERE `product_type_id` = 1);
DELETE FROM `products` WHERE `product_type_id` = 1;

-- ALTER TABLE `product_descriptions` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE CASCADE;

SELECT * FROM `transactions` WHERE `user_id` = 1 UNION SELECT * FROM `transactions` WHERE `user_id` = 2;

SELECT COUNT(total_price) FROM `transactions` WHERE `user_id` = 1 GROUP BY `user_id`;
SELECT SUM(total_price) FROM `transactions` WHERE `user_id` = 1 GROUP BY `user_id`;

SELECT COUNT(*) FROM `transactions` AS `t` INNER JOIN `transaction_details` AS d ON `t`.`id` = `d`.`transaction_id` INNER JOIN `products` AS `p` ON `d`.`product_id` = `p`.`id` WHERE `p`.`product_type_id` = 2 GROUP BY `p`.`product_type_id`;

SELECT * FROM `products` AS `p` INNER JOIN `product_types` AS `t` ON `p`.`product_type_id` = `t`.`id`;

SELECT `t`.*, `p`.*, `u`.* FROM `transactions` AS `t` INNER JOIN `transaction_details` AS `d` ON `t`.`id` = `d`.`transaction_id` INNER JOIN `products` AS `p` ON `d`.`product_id` = `p`.`id` INNER JOIN `users` AS `u` ON `t`.`user_id` = `u`.`id`;

SET GLOBAL log_bin_trust_function_creators = 1; -- root user

DELIMITER $
CREATE TRIGGER `sf_transaction_cascade_bf_e` BEFORE DELETE ON `transactions`
FOR EACH ROW
BEGIN
DELETE FROM `transaction_details` WHERE `transaction_id` = OLD.`id`;
END;
$
DELIMITER ;

CREATE TRIGGER `sf_transaction_detail_update_qty_bf_e` AFTER DELETE ON `transaction_details`
FOR EACH ROW SET `total_qty` = OLD.`total_qty`;

SELECT `name` FROM `products` WHERE `id` NOT IN (SELECT `product_id` FROM `transaction_details`);
