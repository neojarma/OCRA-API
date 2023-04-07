/*
 Navicat Premium Data Transfer

 Source Server         : neojarma
 Source Server Type    : MySQL
 Source Server Version : 100612
 Source Host           : neojarma.com:3306
 Source Schema         : neojarma_ocra_db

 Target Server Type    : MySQL
 Target Server Version : 100612
 File Encoding         : 65001

 Date: 08/04/2023 01:00:31
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for channels
-- ----------------------------
DROP TABLE IF EXISTS `channels`;
CREATE TABLE `channels`  (
  `channel_id` varchar(24) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `name` varchar(50) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `profile_image` text CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL,
  `banner_image` text CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL,
  `user_id` varchar(22) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `created_at` bigint NOT NULL,
  `updated_at` bigint NOT NULL,
  `subscriber` int NULL DEFAULT 0,
  PRIMARY KEY (`channel_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  CONSTRAINT `channels_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `comment_id` int NOT NULL AUTO_INCREMENT,
  `video_id` varchar(11) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `channel_id` varchar(22) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `comment` text CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `created_at` bigint NOT NULL,
  PRIMARY KEY (`comment_id`) USING BTREE,
  INDEX `video_id`(`video_id`) USING BTREE,
  INDEX `comments_ibfk_1`(`channel_id`) USING BTREE,
  CONSTRAINT `comments_ibfk_1` FOREIGN KEY (`channel_id`) REFERENCES `channels` (`channel_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `comments_ibfk_2` FOREIGN KEY (`video_id`) REFERENCES `videos` (`video_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 28 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for dislikes
-- ----------------------------
DROP TABLE IF EXISTS `dislikes`;
CREATE TABLE `dislikes`  (
  `dislike_id` int NOT NULL AUTO_INCREMENT,
  `video_id` varchar(11) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `user_id` varchar(22) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  PRIMARY KEY (`dislike_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `video_id`(`video_id`) USING BTREE,
  CONSTRAINT `dislikes_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `dislikes_ibfk_2` FOREIGN KEY (`video_id`) REFERENCES `videos` (`video_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 200 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for histories
-- ----------------------------
DROP TABLE IF EXISTS `histories`;
CREATE TABLE `histories`  (
  `history_id` int NOT NULL AUTO_INCREMENT,
  `video_id` varchar(11) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `channel_id` varchar(24) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `user_id` varchar(22) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  PRIMARY KEY (`history_id`) USING BTREE,
  INDEX `channel_id`(`channel_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `video_id`(`video_id`) USING BTREE,
  CONSTRAINT `histories_ibfk_1` FOREIGN KEY (`channel_id`) REFERENCES `channels` (`channel_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `histories_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `histories_ibfk_3` FOREIGN KEY (`video_id`) REFERENCES `videos` (`video_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for likes
-- ----------------------------
DROP TABLE IF EXISTS `likes`;
CREATE TABLE `likes`  (
  `like_id` int NOT NULL AUTO_INCREMENT,
  `video_id` varchar(11) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `user_id` varchar(22) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  PRIMARY KEY (`like_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `video_id`(`video_id`) USING BTREE,
  CONSTRAINT `likes_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `likes_ibfk_2` FOREIGN KEY (`video_id`) REFERENCES `videos` (`video_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 219 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sessions
-- ----------------------------
DROP TABLE IF EXISTS `sessions`;
CREATE TABLE `sessions`  (
  `session_id` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `user_id` varchar(22) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `created_at` bigint NOT NULL,
  `expires_at` bigint NOT NULL,
  PRIMARY KEY (`session_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  CONSTRAINT `sessions_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for subscribes
-- ----------------------------
DROP TABLE IF EXISTS `subscribes`;
CREATE TABLE `subscribes`  (
  `subs_id` int NOT NULL AUTO_INCREMENT,
  `channel_id` varchar(24) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `user_id` varchar(22) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  PRIMARY KEY (`subs_id`) USING BTREE,
  INDEX `channel_id`(`channel_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  CONSTRAINT `subscribes_ibfk_1` FOREIGN KEY (`channel_id`) REFERENCES `channels` (`channel_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `subscribes_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 79 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `user_id` varchar(22) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `full_name` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `profile_image` text CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL,
  `email` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `password` text CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `is_verified` tinyint(1) NOT NULL DEFAULT 0,
  `created_at` bigint NOT NULL,
  `updated_at` bigint NOT NULL,
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for verifications
-- ----------------------------
DROP TABLE IF EXISTS `verifications`;
CREATE TABLE `verifications`  (
  `verif_id` int NOT NULL AUTO_INCREMENT,
  `email` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `token` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `expires_at` bigint NOT NULL,
  PRIMARY KEY (`verif_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos`  (
  `video_id` varchar(11) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `channel_id` varchar(24) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `thumbnail` text CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `video` text CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `title` varchar(50) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `description` text CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `tags` text CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL,
  `created_at` bigint NOT NULL,
  `views_count` int NOT NULL DEFAULT 0,
  `likes_count` int NOT NULL DEFAULT 0,
  `dislikes_count` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`video_id`) USING BTREE,
  INDEX `channel_id`(`channel_id`) USING BTREE,
  CONSTRAINT `videos_ibfk_1` FOREIGN KEY (`channel_id`) REFERENCES `channels` (`channel_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for watch_laters
-- ----------------------------
DROP TABLE IF EXISTS `watch_laters`;
CREATE TABLE `watch_laters`  (
  `watch_id` int NOT NULL AUTO_INCREMENT,
  `video_id` varchar(11) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `channel_id` varchar(24) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `user_id` varchar(22) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  PRIMARY KEY (`watch_id`) USING BTREE,
  INDEX `video_id`(`video_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `channel_id`(`channel_id`) USING BTREE,
  CONSTRAINT `watch_laters_ibfk_1` FOREIGN KEY (`video_id`) REFERENCES `videos` (`video_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `watch_laters_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `watch_laters_ibfk_3` FOREIGN KEY (`channel_id`) REFERENCES `channels` (`channel_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Triggers structure for table dislikes
-- ----------------------------
DROP TRIGGER IF EXISTS `auto_increment_dislikes`;
delimiter ;;
CREATE TRIGGER `auto_increment_dislikes` AFTER INSERT ON `dislikes` FOR EACH ROW UPDATE
    videos
SET
    videos.dislikes_count = videos.dislikes_count + 1
WHERE
    videos.video_id = NEW.video_id
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table dislikes
-- ----------------------------
DROP TRIGGER IF EXISTS `auto_decrement_dislikes`;
delimiter ;;
CREATE TRIGGER `auto_decrement_dislikes` BEFORE DELETE ON `dislikes` FOR EACH ROW UPDATE
    videos
SET
    videos.dislikes_count = videos.dislikes_count - 1
WHERE
    videos.video_id = OLD.video_id
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table likes
-- ----------------------------
DROP TRIGGER IF EXISTS `auto_increment_likes`;
delimiter ;;
CREATE TRIGGER `auto_increment_likes` AFTER INSERT ON `likes` FOR EACH ROW UPDATE videos SET videos.likes_count = videos.likes_count + 1 WHERE videos.video_id = NEW.video_id
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table likes
-- ----------------------------
DROP TRIGGER IF EXISTS `auto_decrement_likes`;
delimiter ;;
CREATE TRIGGER `auto_decrement_likes` BEFORE DELETE ON `likes` FOR EACH ROW UPDATE videos SET videos.likes_count = videos.likes_count - 1 WHERE videos.video_id = OLD.video_id
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table subscribes
-- ----------------------------
DROP TRIGGER IF EXISTS `auto_increment_subscriber`;
delimiter ;;
CREATE TRIGGER `auto_increment_subscriber` AFTER INSERT ON `subscribes` FOR EACH ROW UPDATE
    channels
SET
    channels.subscriber = channels.subscriber + 1
WHERE
    channels.channel_id = NEW.channel_id
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table subscribes
-- ----------------------------
DROP TRIGGER IF EXISTS `auto_decrement_subscriber`;
delimiter ;;
CREATE TRIGGER `auto_decrement_subscriber` BEFORE DELETE ON `subscribes` FOR EACH ROW UPDATE
    channels
SET
    channels.subscriber = channels.subscriber - 1
WHERE
    channels.channel_id = OLD.channel_id
;;
delimiter ;

SET FOREIGN_KEY_CHECKS = 1;
