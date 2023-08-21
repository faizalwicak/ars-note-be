-- Create "users" table
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `email` (`email`),
  INDEX `idx_users_deleted_at` (`deleted_at`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "books" table
CREATE TABLE `books` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `name` longtext NULL,
  `user_id` bigint unsigned NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_users_books` (`user_id`),
  INDEX `idx_books_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_users_books` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "categories" table
CREATE TABLE `categories` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `name` longtext NULL,
  `book_id` bigint unsigned NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_books_categories` (`book_id`),
  INDEX `idx_categories_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_books_categories` FOREIGN KEY (`book_id`) REFERENCES `books` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "locations" table
CREATE TABLE `locations` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `name` longtext NULL,
  `book_id` bigint unsigned NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_books_locations` (`book_id`),
  INDEX `idx_locations_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_books_locations` FOREIGN KEY (`book_id`) REFERENCES `books` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "transactions" table
CREATE TABLE `transactions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `date` datetime(3) NULL,
  `value` bigint NULL,
  `description` text NULL,
  `book_id` bigint unsigned NULL,
  `location_id` bigint unsigned NULL,
  `category_id` bigint unsigned NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_books_transactions` (`book_id`),
  INDEX `fk_categories_transactions` (`category_id`),
  INDEX `fk_locations_transactions` (`location_id`),
  INDEX `idx_transactions_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_books_transactions` FOREIGN KEY (`book_id`) REFERENCES `books` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `fk_categories_transactions` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `fk_locations_transactions` FOREIGN KEY (`location_id`) REFERENCES `locations` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
