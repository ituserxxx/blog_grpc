-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- 主机： blog_mysql
-- 生成日期： 2023-02-23 05:18:51
-- 服务器版本： 8.0.31
-- PHP 版本： 8.0.19

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `gin_vue_blog`
--
CREATE DATABASE gin_vue_blog;
use gin_vue_blog;
-- --------------------------------------------------------

--
-- 表的结构 `blog_article`
--

CREATE TABLE `blog_article` (
  `id` int NOT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `status` tinyint DEFAULT '2' COMMENT '1-发布，2-草稿'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

--
-- 转存表中的数据 `blog_article`
--
         
--
-- 表的结构 `blog_comment`
--

CREATE TABLE `blog_comment` (
  `id` int UNSIGNED NOT NULL,
  `username` varchar(16) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `is_author` tinyint(1) NOT NULL DEFAULT '0',
  `parent_id` int UNSIGNED DEFAULT NULL,
  `root_id` int UNSIGNED DEFAULT NULL,
  `content` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `created_time` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `article_id` int DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;

-- --------------------------------------------------------

--
-- 表的结构 `blog_image`
--

CREATE TABLE `blog_image` (
  `id` int NOT NULL,
  `img_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '图片地址',
  `create_time` datetime NOT NULL,
  `type` tinyint DEFAULT NULL,
  `status` tinyint DEFAULT '1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='图片素材库' ROW_FORMAT=DYNAMIC;

-- --------------------------------------------------------

--
-- 表的结构 `blog_tag`
--

CREATE TABLE `blog_tag` (
  `id` int UNSIGNED NOT NULL,
  `tag_name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `article_sum` int DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

--
-- 转存表中的数据 `blog_tag`
--
 

-- --------------------------------------------------------

--
-- 表的结构 `blog_tag_article`
--

CREATE TABLE `blog_tag_article` (
  `id` int UNSIGNED NOT NULL,
  `tag_id` int UNSIGNED NOT NULL,
  `article_id` int UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

--
-- 转存表中的数据 `blog_tag_article`
--
    

-- --------------------------------------------------------

--
-- 表的结构 `blog_user`
--

CREATE TABLE `blog_user` (
  `id` int UNSIGNED NOT NULL,
  `username` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `introduction` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `nickname` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `about` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

--
-- 转存表中的数据 `blog_user`
--
 
-- --------------------------------------------------------

--
-- 表的结构 `blog_visitor`
--

CREATE TABLE `blog_visitor` (
  `id` bigint UNSIGNED NOT NULL,
  `uri` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '访问路径',
  `referer` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
  `ua` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'user_agent',
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `visit_time` datetime NOT NULL COMMENT '访问时间',
  `ip_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '地址'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

--  
-- 转储表的索引
--

--
-- 表的索引 `blog_article`
--
ALTER TABLE `blog_article`
  ADD PRIMARY KEY (`id`) USING BTREE;

--
-- 表的索引 `blog_comment`
--
ALTER TABLE `blog_comment`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `parent` (`parent_id`) USING BTREE,
  ADD KEY `root` (`root_id`) USING BTREE;

--
-- 表的索引 `blog_image`
--
ALTER TABLE `blog_image`
  ADD PRIMARY KEY (`id`) USING BTREE;

--
-- 表的索引 `blog_tag`
--
ALTER TABLE `blog_tag`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD UNIQUE KEY `uni_name` (`tag_name`) USING BTREE;

--
-- 表的索引 `blog_tag_article`
--
ALTER TABLE `blog_tag_article`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `tag_id` (`tag_id`) USING BTREE,
  ADD KEY `article_id` (`article_id`) USING BTREE;

--
-- 表的索引 `blog_user`
--
ALTER TABLE `blog_user`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD UNIQUE KEY `uni_nam` (`username`) USING BTREE;

--
-- 表的索引 `blog_visitor`
--
ALTER TABLE `blog_visitor`
  ADD PRIMARY KEY (`id`) USING BTREE;

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `blog_article`
--
ALTER TABLE `blog_article`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=94;

--
-- 使用表AUTO_INCREMENT `blog_comment`
--
ALTER TABLE `blog_comment`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `blog_image`
--
ALTER TABLE `blog_image`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `blog_tag`
--
ALTER TABLE `blog_tag`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;

--
-- 使用表AUTO_INCREMENT `blog_tag_article`
--
ALTER TABLE `blog_tag_article`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=105;

--
-- 使用表AUTO_INCREMENT `blog_user`
--
ALTER TABLE `blog_user`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `blog_visitor`
--
ALTER TABLE `blog_visitor`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12532;

--
-- 限制导出的表
--

--
-- 限制表 `blog_comment`
--
ALTER TABLE `blog_comment`
  ADD CONSTRAINT `parent` FOREIGN KEY (`parent_id`) REFERENCES `blog_comment` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `root` FOREIGN KEY (`root_id`) REFERENCES `blog_comment` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
