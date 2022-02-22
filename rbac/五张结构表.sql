/*
 Navicat Premium Data Transfer

 Source Server         : mysql57
 Source Server Type    : MySQL
 Source Server Version : 50721
 Source Host           : localhost:3307
 Source Schema         : mytest

 Target Server Type    : MySQL
 Target Server Version : 50721
 File Encoding         : 65001

 Date: 28/09/2020 19:35:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles`  (
                          `role_id` int(11) NOT NULL AUTO_INCREMENT,
                          `role_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
                          `role_pid` int(11) NULL DEFAULT 0,
                          `role_comment` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
                          PRIMARY KEY (`role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of roles
-- ----------------------------
INSERT INTO `roles` VALUES (1, 'guest', 0, '游客');
INSERT INTO `roles` VALUES (2, 'sysadmin', 0, '系统管理员');
INSERT INTO `roles` VALUES (3, 'deptadmin', 2, '部门管理员');
INSERT INTO `roles` VALUES (4, 'useradmin', 2, '会员总管理员');
INSERT INTO `roles` VALUES (5, 'frontadmin', 4, '前台会员管理员');
INSERT INTO `roles` VALUES (6, 'backendadmin', 4, '后台会员管理员');

-- ----------------------------
-- Table structure for router_roles
-- ----------------------------
DROP TABLE IF EXISTS `router_roles`;
CREATE TABLE `router_roles`  (
                                 `id` int(11) NOT NULL AUTO_INCREMENT,
                                 `router_id` int(11) NULL DEFAULT NULL,
                                 `role_id` int(11) NULL DEFAULT NULL,
                                 PRIMARY KEY (`id`) USING BTREE,
                                 UNIQUE INDEX `router_id`(`router_id`, `role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of router_roles
-- ----------------------------
INSERT INTO `router_roles` VALUES (1, 1, 3);
INSERT INTO `router_roles` VALUES (2, 1, 4);
INSERT INTO `router_roles` VALUES (5, 3, 3);
INSERT INTO `router_roles` VALUES (3, 5, 3);
INSERT INTO `router_roles` VALUES (4, 5, 4);
INSERT INTO `router_roles` VALUES (6, 7, 4);

-- ----------------------------
-- Table structure for routers
-- ----------------------------
DROP TABLE IF EXISTS `routers`;
CREATE TABLE `routers`  (
                            `r_id` int(11) NOT NULL AUTO_INCREMENT,
                            `r_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
                            `r_uri` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
                            `r_method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
                            `r_status` tinyint(4) NULL DEFAULT NULL,
                            PRIMARY KEY (`r_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of routers
-- ----------------------------
INSERT INTO `routers` VALUES (1, '部门列表', '/detps', 'GET', 1);
INSERT INTO `routers` VALUES (2, '部门详细', '/depts/:id', 'GET', 1);
INSERT INTO `routers` VALUES (3, '新增部门', '/depts', 'POST', 1);
INSERT INTO `routers` VALUES (5, '会员列表', '/users', 'GET', 1);
INSERT INTO `routers` VALUES (6, '会员详细', '/users/:id', 'GET', 1);
INSERT INTO `routers` VALUES (7, '新增会员', '/users', 'POST', 1);

-- ----------------------------
-- Table structure for user_roles
-- ----------------------------
DROP TABLE IF EXISTS `user_roles`;
CREATE TABLE `user_roles`  (
                               `id` int(11) NOT NULL AUTO_INCREMENT,
                               `user_id` int(11) NOT NULL,
                               `role_id` int(11) NOT NULL,
                               PRIMARY KEY (`id`) USING BTREE,
                               UNIQUE INDEX `user_id`(`user_id`, `role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_roles
-- ----------------------------
INSERT INTO `user_roles` VALUES (1, 1, 2);
INSERT INTO `user_roles` VALUES (2, 2, 3);
INSERT INTO `user_roles` VALUES (3, 3, 4);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
                          `user_id` int(11) NOT NULL AUTO_INCREMENT,
                          `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
                          PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'shenyi');
INSERT INTO `users` VALUES (1, 'alinx');
INSERT INTO `users` VALUES (1, 'lixiong');
INSERT INTO `users` VALUES (2, 'lisi');
INSERT INTO `users` VALUES (3, 'zhangsan');

SET FOREIGN_KEY_CHECKS = 1;
