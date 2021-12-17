/*
Navicat MySQL Data Transfer

Source Server         : 数据库
Source Server Version : 50736
Source Host           : 127.0.0.1:3306
Source Database       : cmdb

Target Server Type    : MYSQL
Target Server Version : 50736
File Encoding         : 65001

Date: 2021-12-17 10:03:46
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `password` varchar(100) NOT NULL,
  `username` varchar(20) DEFAULT NULL,
  `head_url` varchar(200) DEFAULT NULL,
  `birthday` date DEFAULT NULL,
  `address` varchar(200) DEFAULT NULL,
  `desc` text,
  `gender` varchar(6) DEFAULT NULL,
  `role` int(11) NOT NULL,
  `mobile` varchar(12) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=124 DEFAULT CHARSET=utf8mb4;
