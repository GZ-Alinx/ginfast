/*
Navicat MySQL Data Transfer

Source Server         : 本地数据库API
Source Server Version : 50736
Source Host           : 10.1.20.130:3306
Source Database       : proinfo

Target Server Type    : MYSQL
Target Server Version : 50736
File Encoding         : 65001

Date: 2021-12-22 17:11:50
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for pro_master
-- ----------------------------
DROP TABLE IF EXISTS `pro_userall`;
CREATE TABLE `pro_userall` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `mobile` varchar(255) NOT NULL,
  `email` varchar(255) DEFAULT NULL,
  `userid` varchar(255) DEFAULT NULL,
  `dept` varchar(255) DEFAULT '',
  `updatetime` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `mobile` (`mobile`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=206532 DEFAULT CHARSET=utf8mb4;
