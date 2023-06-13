-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Jun 13, 2023 at 06:56 PM
-- Server version: 8.0.30
-- PHP Version: 8.1.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `dbs-ginprojeact`
--

-- --------------------------------------------------------

--
-- Table structure for table `categories`
--

CREATE TABLE `categories` (
  `id` bigint UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` longtext COLLATE utf8mb4_general_ci,
  `slug` longtext COLLATE utf8mb4_general_ci,
  `parint_id` bigint UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `categories`
--

INSERT INTO `categories` (`id`, `created_at`, `updated_at`, `deleted_at`, `title`, `slug`, `parint_id`) VALUES
(1, '2023-06-05 09:01:02.023', '2023-06-05 09:01:02.023', NULL, 'sad', 'sad', 0),
(2, '2023-06-05 09:07:25.321', '2023-06-05 16:43:56.454', NULL, 'دسته بندی شماره 2', 'دسته_بندی_جدید', 0),
(3, '2023-06-05 09:07:39.331', '2023-06-05 09:07:39.331', NULL, 'sad', 'sad', 0),
(4, '2023-06-05 09:09:00.366', '2023-06-05 09:09:00.366', NULL, '', '', 0),
(5, '2023-06-05 09:10:16.694', '2023-06-05 17:03:58.160', '2023-06-05 17:04:07.853', 'دسته بندی جدید ', 'دسته_بندی_جدید', 0);

-- --------------------------------------------------------

--
-- Table structure for table `halls`
--

CREATE TABLE `halls` (
  `id` bigint UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` longtext COLLATE utf8mb4_general_ci,
  `body` longtext COLLATE utf8mb4_general_ci,
  `span` bigint UNSIGNED DEFAULT NULL,
  `active` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `movies`
--

CREATE TABLE `movies` (
  `id` bigint UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` longtext COLLATE utf8mb4_general_ci,
  `synopsis` longtext COLLATE utf8mb4_general_ci,
  `body` longtext COLLATE utf8mb4_general_ci,
  `category_id` bigint UNSIGNED DEFAULT NULL,
  `price` longtext COLLATE utf8mb4_general_ci,
  `year` longtext COLLATE utf8mb4_general_ci,
  `view` bigint UNSIGNED DEFAULT NULL,
  `image` longtext COLLATE utf8mb4_general_ci,
  `actived` tinyint(1) DEFAULT NULL,
  `minutes` bigint UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `movies`
--

INSERT INTO `movies` (`id`, `created_at`, `updated_at`, `deleted_at`, `title`, `synopsis`, `body`, `category_id`, `price`, `year`, `view`, `image`, `actived`, `minutes`) VALUES
(1, '2023-06-06 18:24:09.551', '2023-06-06 18:24:09.551', NULL, 'دسته بندی جدید ', '', '', 0, '', '', 1, '', 0, NULL),
(2, '2023-06-06 18:28:13.484', '2023-06-06 18:39:27.759', NULL, 'filme  ', 'so in filme is verry butefful', 'verey butefule filme', 0, '', '', 1, '', 0, NULL),
(3, '2023-06-06 18:30:59.875', '2023-06-06 19:03:25.927', NULL, 'filme  ', 'so in filme is verry butefful', 'verey butefule filme', 0, '', '', 1, '', 1, NULL),
(4, '2023-06-06 18:32:29.843', '2023-06-06 19:03:41.785', NULL, 'filme  ', '', 'verey butefule filme', 0, '', '', 1, '', 1, NULL),
(5, '2023-06-06 18:33:44.660', '2023-06-06 19:04:06.604', NULL, 'filme  ', '', 'verey butefule filme', 0, '', '', 1, '', 1, NULL),
(6, '2023-06-06 18:34:41.420', '2023-06-06 18:34:41.420', NULL, 'filme  ', '', 'verey butefule filme', 0, '', '', 1, '', 0, NULL);

-- --------------------------------------------------------

--
-- Table structure for table `movie_halls`
--

CREATE TABLE `movie_halls` (
  `id` bigint UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `hall_id` bigint UNSIGNED NOT NULL,
  `movie_id` bigint UNSIGNED NOT NULL,
  `year_start` longtext COLLATE utf8mb4_general_ci,
  `month_start` longtext COLLATE utf8mb4_general_ci,
  `day_start` longtext COLLATE utf8mb4_general_ci,
  `hour_start` longtext COLLATE utf8mb4_general_ci,
  `minutes_start` longtext COLLATE utf8mb4_general_ci,
  `span` bigint UNSIGNED DEFAULT NULL,
  `year_end` longtext COLLATE utf8mb4_general_ci,
  `month_end` longtext COLLATE utf8mb4_general_ci,
  `day_end` longtext COLLATE utf8mb4_general_ci,
  `hour_end` longtext COLLATE utf8mb4_general_ci,
  `minutes_end` longtext COLLATE utf8mb4_general_ci,
  `number_sales` bigint UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` varchar(191) COLLATE utf8mb4_general_ci NOT NULL,
  `password` longtext COLLATE utf8mb4_general_ci NOT NULL,
  `email` varchar(191) COLLATE utf8mb4_general_ci NOT NULL,
  `fullname` longtext COLLATE utf8mb4_general_ci NOT NULL,
  `active` bigint UNSIGNED DEFAULT '0',
  `status` bigint UNSIGNED DEFAULT '1',
  `code_active` bigint UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `categories`
--
ALTER TABLE `categories`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_categories_deleted_at` (`deleted_at`);

--
-- Indexes for table `halls`
--
ALTER TABLE `halls`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_halls_deleted_at` (`deleted_at`);

--
-- Indexes for table `movies`
--
ALTER TABLE `movies`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_movies_deleted_at` (`deleted_at`);

--
-- Indexes for table `movie_halls`
--
ALTER TABLE `movie_halls`
  ADD PRIMARY KEY (`id`,`hall_id`,`movie_id`),
  ADD KEY `idx_movie_halls_deleted_at` (`deleted_at`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`),
  ADD UNIQUE KEY `email` (`email`),
  ADD KEY `idx_users_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `categories`
--
ALTER TABLE `categories`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `halls`
--
ALTER TABLE `halls`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `movies`
--
ALTER TABLE `movies`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `movie_halls`
--
ALTER TABLE `movie_halls`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
