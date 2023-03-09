/*
 Navicat MySQL Data Transfer

 Source Server         : myLocalhost
 Source Server Type    : MySQL
 Source Server Version : 80032
 Source Host           : localhost:3306
 Source Schema         : datascs

 Target Server Type    : MySQL
 Target Server Version : 80032
 File Encoding         : 65001

 Date: 09/03/2023 16:12:31
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for apply
-- ----------------------------
DROP TABLE IF EXISTS `apply`;
CREATE TABLE `apply`  (
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `starttime` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `endtime` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `state` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of apply
-- ----------------------------
INSERT INTO `apply` VALUES ('德育分', '20210202', '20210303', 'allowed');
INSERT INTO `apply` VALUES ('智育分', '20210202', '20210303', 'allowed');
INSERT INTO `apply` VALUES ('体育分', '20210202', '20210204', 'disallowed');

-- ----------------------------
-- Table structure for coach
-- ----------------------------
DROP TABLE IF EXISTS `coach`;
CREATE TABLE `coach`  (
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `relativeid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of coach
-- ----------------------------
INSERT INTO `coach` VALUES ('金大坚', '202102222222', '1212');

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `id` int(10) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT,
  `created_At` datetime(0) NULL DEFAULT NULL,
  `updated_At` datetime(0) NULL DEFAULT NULL,
  `target` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `deleted_At` datetime(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  `time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of comment
-- ----------------------------
INSERT INTO `comment` VALUES ('eiusmod velit minim', 0000000001, '2023-02-18 17:51:24', '2023-02-18 17:51:24', '1', '2023-02-18 19:41:33', NULL);
INSERT INTO `comment` VALUES ('eiusmod velit minim', 0000000002, '2023-02-18 17:51:35', '2023-02-18 17:51:35', '1', NULL, NULL);
INSERT INTO `comment` VALUES ('ut quis', 0000000003, NULL, NULL, '2', NULL, '2023-02-21 19:55:21');
INSERT INTO `comment` VALUES ('ut quis', 0000000004, NULL, NULL, '2', NULL, '2023-02-22 19:32:39');
INSERT INTO `comment` VALUES ('ut quis', 0000000005, NULL, NULL, '16', NULL, '2023-02-22 19:33:13');
INSERT INTO `comment` VALUES ('ut quis', 0000000006, NULL, NULL, '16', NULL, '2023-02-22 19:42:22');
INSERT INTO `comment` VALUES ('ut quis', 0000000007, NULL, NULL, '16', NULL, '2023-02-22 19:44:14');
INSERT INTO `comment` VALUES ('ut quis', 0000000008, NULL, NULL, '16', NULL, '2023-02-22 19:59:20');
INSERT INTO `comment` VALUES ('ut quis', 0000000009, NULL, NULL, '16', NULL, '2023-02-22 22:44:17');
INSERT INTO `comment` VALUES ('ut quis', 0000000010, NULL, NULL, '16', NULL, '2023-02-22 22:44:24');

-- ----------------------------
-- Table structure for cookie
-- ----------------------------
DROP TABLE IF EXISTS `cookie`;
CREATE TABLE `cookie`  (
  `id` int NULL DEFAULT NULL,
  `token` char(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `role` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of cookie
-- ----------------------------
INSERT INTO `cookie` VALUES (2, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6IjIiLCJleHAiOjE2NzUxNzQyNzMsImlhdCI6MTY3NDU2OTQ3MywiaXNzIjoi5rC055eVIiwic3ViIjoidG9rZW4ifQ.jOrJCqeWIflathOyjeAx-jNNav-qz32QtHHaZMb_eZ4', NULL, NULL);
INSERT INTO `cookie` VALUES (3, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozLCJ1c2VybmFtZSI6IjEiLCJleHAiOjE2NzUxNzU3MTQsImlhdCI6MTY3NDU3MDkxNCwiaXNzIjoi5rC055eVIiwic3ViIjoidG9rZW4ifQ.TjWKKhcvl8DFRdixca0ARmdXLGrbExPteVJoBP4IJgE', NULL, NULL);
INSERT INTO `cookie` VALUES (3, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozLCJ1c2VybmFtZSI6IjEiLCJleHAiOjE2NzcyNDg0ODQsImlhdCI6MTY3NjY0MzY4NCwiaXNzIjoi5rC055eVIiwic3ViIjoidG9rZW4ifQ.59_3nw7vZnOsSLvWdbCJxaejhQUOTc7ZVkRCPh1wTWA', NULL, NULL);
INSERT INTO `cookie` VALUES (2, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6IjIiLCJleHAiOjE2NzcyNDg1NjMsImlhdCI6MTY3NjY0Mzc2MywiaXNzIjoi5rC055eVIiwic3ViIjoidG9rZW4ifQ.WTrYILA1QlB0v3GfKWRHv26A1DfHeVakch1GMXyCN5I', NULL, NULL);
INSERT INTO `cookie` VALUES (1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6IjMiLCJleHAiOjE2NzcyNDg1ODIsImlhdCI6MTY3NjY0Mzc4MiwiaXNzIjoi5rC055eVIiwic3ViIjoidG9rZW4ifQ.fsbInv6BrW5sFI_nBpbmfLfzBCOIzcIgu4su5bcYlFk', NULL, NULL);
INSERT INTO `cookie` VALUES (1, '123456', 'coach', '1');
INSERT INTO `cookie` VALUES (1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6IjMiLCJleHAiOjE2NzczMDAzMDQsImlhdCI6MTY3NjY5NTUwNCwiaXNzIjoi5rC055eVIiwic3ViIjoidG9rZW4ifQ.40fRsUd_9m8azX1aEmaM6AXkZYKzMWmM1DOfl85V4Kc', NULL, NULL);
INSERT INTO `cookie` VALUES (1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6IjMiLCJleHAiOjE2NzczMDAzMTEsImlhdCI6MTY3NjY5NTUxMSwiaXNzIjoi5rC055eVIiwic3ViIjoidG9rZW4ifQ.sHIxeiuotmrX56rJkVmcLHFPy7eXUrqWK7DJRLyWqnk', NULL, NULL);
INSERT INTO `cookie` VALUES (1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6IjMiLCJleHAiOjE2NzczMDA5MTYsImlhdCI6MTY3NjY5NjExNiwiaXNzIjoi5rC055eVIiwic3ViIjoidG9rZW4ifQ.fnk4vFRM2dtMyRPXdrbTOf3l2S3VqT8EM-7eeHF_L50', NULL, NULL);
INSERT INTO `cookie` VALUES (1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6IjMiLCJleHAiOjE2NzczMDA5MzIsImlhdCI6MTY3NjY5NjEzMiwiaXNzIjoi5rC055eVIiwic3ViIjoidG9rZW4ifQ._4LWDyfKwKbwLiCBLvqv6Ce_9-xCZ12gsxCGSf5fzDc', NULL, NULL);
INSERT INTO `cookie` VALUES (1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6IjMiLCJleHAiOjE2NzczMDA5NTAsImlhdCI6MTY3NjY5NjE1MCwiaXNzIjoi5rC055eVIiwic3ViIjoidG9rZW4ifQ.rN0Mq7lUAmun5owd795Tptj6xew6-7G1c0pNJxWZlN0', NULL, NULL);
INSERT INTO `cookie` VALUES (1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6IjMiLCJleHAiOjE2NzczMDA5NTEsImlhdCI6MTY3NjY5NjE1MSwiaXNzIjoi5rC055eVIiwic3ViIjoidG9rZW4ifQ.c8YOfwf9BId4Oc8L-Kc-ZzneNZInvskeIJ_H42XIovU', NULL, NULL);
INSERT INTO `cookie` VALUES (1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6IjMiLCJleHAiOjE2NzczMDEwOTksImlhdCI6MTY3NjY5NjI5OSwiaXNzIjoi5rC055eVIiwic3ViIjoidG9rZW4ifQ.duAYmRXLscbBdbi-dn3gMsbDb0C0VIwFnSpI8LImvVM', NULL, NULL);
INSERT INTO `cookie` VALUES (1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6IjMiLCJleHAiOjE2NzczMDE5MTQsImlhdCI6MTY3NjY5NzExNCwiaXNzIjoi5rC055eVIiwic3ViIjoidG9rZW4ifQ.1hhhDqCWa6Fkb2aLyGPFFcfztfXEOBeVHEfKXo-dAg0', NULL, NULL);
INSERT INTO `cookie` VALUES (1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6IjMiLCJleHAiOjE2NzczMDQzNjAsImlhdCI6MTY3NjY5OTU2MCwiaXNzIjoi5rC055eVIiwic3ViIjoidG9rZW4ifQ.lymD-8vwSySjGliJ5GVA7f0HIpmRWu9SuGV53mDxNrU', 'student', '3');
INSERT INTO `cookie` VALUES (1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6IjMiLCJleHAiOjE2NzczMDQ0NDAsImlhdCI6MTY3NjY5OTY0MCwiaXNzIjoi5rC055eVIiwic3ViIjoidG9rZW4ifQ.X_m3hGO9_MqUSceLo4Q5C4Qn6SU2SrZvFdcIgBTUx70', 'student', '3');
INSERT INTO `cookie` VALUES (1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6IjMiLCJleHAiOjE2Nzc0MjY1NjgsImlhdCI6MTY3NjgyMTc2OCwiaXNzIjoi5rC055eVIiwic3ViIjoidG9rZW4ifQ.GGdGBtfByj69yx7-SsMktKZkXE9iD6sQ9GFK1QVsv3Q', 'student', '3');

-- ----------------------------
-- Table structure for count
-- ----------------------------
DROP TABLE IF EXISTS `count`;
CREATE TABLE `count`  (
  `index` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `countid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `value` double NULL DEFAULT NULL,
  `content` json NULL,
  `list` json NULL,
  `updatedAt` datetime(0) NULL DEFAULT NULL,
  `createdAt` datetime(0) NULL DEFAULT NULL,
  `label` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `top` double NULL DEFAULT NULL,
  `unique` json NULL,
  `realindex` int NOT NULL,
  `id` int NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of count
-- ----------------------------
INSERT INTO `count` VALUES ('v', '202102222222', 12, NULL, NULL, NULL, NULL, '德育分', 30, 'true', 10, 1);
INSERT INTO `count` VALUES ('v-1', '202102222222', 11, NULL, NULL, NULL, NULL, '德育赋分', 20, 'false', 11, 2);
INSERT INTO `count` VALUES ('v', '2021022', 12, NULL, NULL, NULL, NULL, '德育分', 30, 'true', 10, 3);

-- ----------------------------
-- Table structure for countmodel
-- ----------------------------
DROP TABLE IF EXISTS `countmodel`;
CREATE TABLE `countmodel`  (
  `index` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `id` int NOT NULL AUTO_INCREMENT,
  `value` double NULL DEFAULT NULL,
  `content` json NULL,
  `list` json NULL,
  `updatedAt` datetime(0) NULL DEFAULT NULL,
  `createdAt` datetime(0) NULL DEFAULT NULL,
  `label` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `top` double NULL DEFAULT NULL,
  `unique` json NULL,
  `realindex` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of countmodel
-- ----------------------------
INSERT INTO `countmodel` VALUES ('v', 1, NULL, NULL, NULL, NULL, NULL, '德育分', 30, 'true', 10);
INSERT INTO `countmodel` VALUES ('v-1', 2, NULL, NULL, NULL, NULL, NULL, '志愿分', 10, 'false', 11);
INSERT INTO `countmodel` VALUES ('v-2', 3, NULL, NULL, NULL, NULL, NULL, '德育学习分', 20, 'false', 12);
INSERT INTO `countmodel` VALUES ('a', 4, NULL, NULL, NULL, NULL, NULL, '智育分', 60, 'true', 20);

-- ----------------------------
-- Table structure for post
-- ----------------------------
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post`  (
  `type` json NOT NULL,
  `id` int NOT NULL,
  `time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `userRole` json NOT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `commentList` json NOT NULL,
  `createdAt` datetime(0) NOT NULL,
  `updatedAt` datetime(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of post
-- ----------------------------

-- ----------------------------
-- Table structure for reason
-- ----------------------------
DROP TABLE IF EXISTS `reason`;
CREATE TABLE `reason`  (
  `reason` json NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of reason
-- ----------------------------
INSERT INTO `reason` VALUES ('[\"证据不充分\", \"申请内容格式有误\"]');
INSERT INTO `reason` VALUES ('[\"证据不充分\", \"申请内容格式有误\"]');
INSERT INTO `reason` VALUES ('[\"证据不充分\", \"申请内容格式有误\"]');

-- ----------------------------
-- Table structure for record
-- ----------------------------
DROP TABLE IF EXISTS `record`;
CREATE TABLE `record`  (
  `label` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `index` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'ͨ???㼶????????\n`v-2-2`??ʾ**????????-??ʵ?Ӽ???-???????μ?ʵ??**\n`v-1`??ʾ**????????-??????????**',
  `complain` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `id` int NOT NULL AUTO_INCREMENT,
  `value` int NULL DEFAULT NULL,
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `files` json NULL,
  `state` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `createdAt` datetime(0) NULL DEFAULT NULL,
  `updatedAt` datetime(0) NULL DEFAULT NULL,
  `countid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of record
-- ----------------------------
INSERT INTO `record` VALUES (NULL, NULL, NULL, 1, NULL, 'cupidatat aliquip laborum', '2023-02-21 21:22:14', NULL, '1', 'dolore reprehenderit officia ullamco', NULL, NULL, '202102222222');
INSERT INTO `record` VALUES (NULL, NULL, NULL, 2, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `record` VALUES (NULL, NULL, NULL, 3, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);

-- ----------------------------
-- Table structure for suggest
-- ----------------------------
DROP TABLE IF EXISTS `suggest`;
CREATE TABLE `suggest`  (
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `target` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `countid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of suggest
-- ----------------------------
INSERT INTO `suggest` VALUES ('velit id commodo aliquip exercitation', 'cupidatat ut eiusmod', '');
INSERT INTO `suggest` VALUES ('velit id commodo aliquip exercitation', 'cupidatat ut eiusmod', '');
INSERT INTO `suggest` VALUES ('velit id commodo aliquip exercitation', 'cupidatat ut eiusmod', '');
INSERT INTO `suggest` VALUES ('velit id commodo aliquip exercitation', 'cupidatat ut eiusmod', '');
INSERT INTO `suggest` VALUES ('velit id commodo aliquip exercitation', 'cupidatat ut eiusmod', '');
INSERT INTO `suggest` VALUES ('Excepteur eu occaecat esse', 'in eiusmod', '');

-- ----------------------------
-- Table structure for suncount
-- ----------------------------
DROP TABLE IF EXISTS `suncount`;
CREATE TABLE `suncount`  (
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `value` float NULL DEFAULT NULL,
  `content` json NULL,
  `index` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of suncount
-- ----------------------------
INSERT INTO `suncount` VALUES ('commodo voluptate', 11, '[\"eiusmod cupidatat\", \"cillum ut sunt proident\"]', NULL);
INSERT INTO `suncount` VALUES ('commodo voluptate', 11, '[\"eiusmod cupidatat\", \"cillum ut sunt proident\"]', NULL);
INSERT INTO `suncount` VALUES ('commodo voluptate', 11, '[\"eiusmod cupidatat\", \"cillum ut sunt proident\"]', NULL);
INSERT INTO `suncount` VALUES ('commodo voluptate', 11, '[\"eiusmod cupidatat\", \"cillum ut sunt proident\"]', NULL);

-- ----------------------------
-- Table structure for tip
-- ----------------------------
DROP TABLE IF EXISTS `tip`;
CREATE TABLE `tip`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `createdAt` datetime(0) NOT NULL,
  `updatedAt` datetime(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of tip
-- ----------------------------

-- ----------------------------
-- Table structure for topic
-- ----------------------------
DROP TABLE IF EXISTS `topic`;
CREATE TABLE `topic`  (
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `id` int(255) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT,
  `role` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `created_At` datetime(0) NOT NULL,
  `updated_At` datetime(0) NOT NULL,
  `deleted_At` datetime(0) NULL DEFAULT NULL,
  `commentList` json NULL,
  `time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of topic
-- ----------------------------
INSERT INTO `topic` VALUES ('ad laborum reprehenderit velit', 000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001, 'question', '强派划知写', '本习领己下比在式见白政近持此直果支很。水示因起府如反情家斗往被据九压少展。干白总即节并我统格据火已单结。约市多标新断术算通想而育用就己战。专料级等根保然此从受而提提式。员管素定需商程除就办采如划。改张共切并立快外走科容去。', '2023-02-18 17:19:14', '2023-02-18 17:19:14', NULL, NULL, NULL);
INSERT INTO `topic` VALUES ('laborum sint', 000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002, 'question', '人将做飞需', '路置记此三格生使可分想意难。石群红火没便得都数选权层度等看。安完院众个近火为个圆龙中入名写做。按国给后回少叫且共观响马科才带。战管度联专交次方号应美变育张走。前所体车东没等小见山毛变利且美料前属。', '2023-02-18 17:19:19', '2023-02-18 17:19:19', NULL, NULL, NULL);
INSERT INTO `topic` VALUES ('laborum sint', 000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003, 'question', '人将做飞需', '路置记此三格生使可分想意难。石群红火没便得都数选权层度等看。安完院众个近火为个圆龙中入名写做。按国给后回少叫且共观响马科才带。战管度联专交次方号应美变育张走。前所体车东没等小见山毛变利且美料前属。', '2023-02-18 17:51:57', '2023-02-18 17:51:57', NULL, NULL, NULL);
INSERT INTO `topic` VALUES ('laborum sint', 000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004, 'question', '人将做飞需', '路置记此三格生使可分想意难。石群红火没便得都数选权层度等看。安完院众个近火为个圆龙中入名写做。按国给后回少叫且共观响马科才带。战管度联专交次方号应美变育张走。前所体车东没等小见山毛变利且美料前属。', '2023-02-18 17:51:58', '2023-02-18 17:51:58', NULL, NULL, NULL);
INSERT INTO `topic` VALUES ('laborum sint', 000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005, 'question', '人将做飞需', '路置记此三格生使可分想意难。石群红火没便得都数选权层度等看。安完院众个近火为个圆龙中入名写做。按国给后回少叫且共观响马科才带。战管度联专交次方号应美变育张走。前所体车东没等小见山毛变利且美料前属。', '2023-02-19 20:13:49', '2023-02-19 20:13:49', NULL, NULL, '2023-02-19 20:13:48');
INSERT INTO `topic` VALUES ('laborum sint', 000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006, 'question', '人将做飞需', '路置记此三格生使可分想意难。石群红火没便得都数选权层度等看。安完院众个近火为个圆龙中入名写做。按国给后回少叫且共观响马科才带。战管度联专交次方号应美变育张走。前所体车东没等小见山毛变利且美料前属。', '2023-02-19 20:14:12', '2023-02-19 20:14:12', NULL, NULL, '2023-02-19 20:14:12');
INSERT INTO `topic` VALUES ('elit est Lorem dolore', 000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007, 'question', '能头值易从', '复划场红地经阶格相装百报火情。叫百起多收报素般下学走京那较。安出统深维究于音类京白状命确。较型身完后级又什九也给院出保准。地指代复际土干是积品龙容压派精比。', '2023-02-19 20:14:22', '2023-02-19 20:14:22', NULL, NULL, '2023-02-19 20:14:22');
INSERT INTO `topic` VALUES ('qui Duis ad est aliquip', 000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008, 'question', '明办结声来七', '派体行此行起干已完办象就金适属克们。构要车该从外单只标速接各最。委此细口须加状带期手时今期命。八构质总手质全走群十必者党原。些学导多条斯条由入管术设党。办先她三石难做化全多毛者量广。', '2023-02-19 20:14:23', '2023-02-19 20:14:23', NULL, NULL, '2023-02-19 20:14:23');
INSERT INTO `topic` VALUES ('occaecat sint deserunt', 000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000009, 'question', '算过应达', '也海劳没带何群活题院期放所十接性必。正长建物天外般合形前府备同步区学分。东红适面华号议委这打她构感身解。还由龙做过部美月活海石问只。层议极强石装节接力包选要该路治。传音研处八土三造标京打引论。支万直战命深车效九反向思工。', '2023-02-19 20:14:24', '2023-02-19 20:14:24', NULL, NULL, '2023-02-19 20:14:24');
INSERT INTO `topic` VALUES ('mollit in Ut nulla laboris', 000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010, 'question', '应管入并', '并铁精上变按般化看一是导听际接深里造。线指布小话放表算习交到而专十满深。由际用需子确中选求要情准专建。代周新者算是拉把飞增技团。织老张则本百从切达风已响华斗。证标高标六情果分此安流再。', '2023-02-19 20:14:25', '2023-02-19 20:14:25', NULL, NULL, '2023-02-19 20:14:24');
INSERT INTO `topic` VALUES ('ut irure non', 000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000011, 'question', '清意二素命', '照使者先科义由入达酸教却酸亲基。文属住什义目书界热众之资带党果。广再西和取科正音非应状习要最引素。千往往从意听才素法验构边开族准。', '2023-02-19 20:14:25', '2023-02-19 20:14:25', NULL, NULL, '2023-02-19 20:14:25');
INSERT INTO `topic` VALUES ('dolor enim ut', 000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000012, 'declare', '各民层斗非日军', '保斯率资造又题公酸相需真识解响级响风。半第边知土效斯际半十消非习好局史也。查科图构难所存省真比部统员史头证。县构许元究些时它是林确满内及动。周率进音什术率因半采济她及内几。气与一器南别作经七种往务提名快不。', '2023-02-19 20:14:26', '2023-02-19 20:14:26', NULL, NULL, '2023-02-19 20:14:26');
INSERT INTO `topic` VALUES ('enim velit consequat est', 000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000013, 'declare', '众路公响研', '特置口较意前前消少次力火交领想容。照克千装该万维这边流认九度般区率。关则离家等决历才由采活革全因条想。化热位给声委科点状划切维。会史体值阶深上先南面来众自。包油学业空验根合石解式把率金到。', '2023-02-19 20:14:27', '2023-02-19 20:14:27', NULL, NULL, '2023-02-19 20:14:27');
INSERT INTO `topic` VALUES ('pariatur eu deserunt amet occaecat', 000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000014, 'declare', '流体放前', '流前候积科八地去格重局线反院反将万。要都真性铁况定己维又验低千土六原。术识安具不世究九们车一元持济。离江五规军例在角值力取准此。斯律标究走六对其容领年越历半原。', '2023-02-19 20:44:53', '2023-02-19 20:44:53', NULL, NULL, '2023-02-19 20:44:52');
INSERT INTO `topic` VALUES ('dolore eu', 000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000015, 'question', '离没许影千美白', '能被际道类为总号青省品复产始义下求。细社法公线但情把究习回飞。构战界品水干代改术色七长年重平到。低心系她约单习其将除常老东应。命到选多从放酸对使出与技数与众角。', '2023-02-19 23:52:22', '2023-02-19 23:52:22', NULL, NULL, '2023-02-19 23:52:21');
INSERT INTO `topic` VALUES ('proident dolor qui eu ut', 000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000016, 'sint dolore commodo', '物众圆式表儿', '在门将三们提离要果向细省其。历西海毛众同计清号行青图物果。划之把青争列眼增酸做气感传上。思金此量管领利便期报消属石人。装去米什计眼我争龙目率因着老成放。', '2023-02-21 19:53:36', '2023-02-21 19:53:36', NULL, NULL, '2023-02-21 19:53:36');
INSERT INTO `topic` VALUES ('proident dolor qui eu ut', 000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000017, 'sint dolore commodo', '物众圆式表儿', '在门将三们提离要果向细省其。历西海毛众同计清号行青图物果。划之把青争列眼增酸做气感传上。思金此量管领利便期报消属石人。装去米什计眼我争龙目率因着老成放。', '2023-02-22 19:32:58', '2023-02-22 19:32:58', NULL, NULL, '2023-02-22 19:32:57');
INSERT INTO `topic` VALUES ('proident dolor qui eu ut', 000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000018, 'sint dolore commodo', '物众圆式表儿', '在门将三们提离要果向细省其。历西海毛众同计清号行青图物果。划之把青争列眼增酸做气感传上。思金此量管领利便期报消属石人。装去米什计眼我争龙目率因着老成放。', '2023-02-22 22:44:34', '2023-02-22 22:44:34', NULL, NULL, '2023-02-22 22:44:34');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int NOT NULL,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `role` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `deleted_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `coachname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `coachid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `countid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `session` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '文德嗣', '12345678', 'student', NULL, NULL, '', '202102222222', '202103333333', 'ajJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6IjMiLCJleHAiOjE2NzczMDQzNjAsImlhdCI6MTY3NjY5OTU2MCwiaXNzIjoi5rC055eVIiwic3ViIjoidG9rZW4ifQ.lymD-8vwSySjGliJ5GVA7f0HIpmRWu9SuGV53mDxNrU');
INSERT INTO `user` VALUES (2, '克里斯蒂娜', '45', 'student', NULL, '2023-02-22 22:43:49', '', '', '202102222222', 'cbdhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6IjMiLCJleHAiOjE2NzczMDQzNjAsImlhdCI6MTY3NjY5OTU2MCwiaXNzIjoi5rC055eVIiwic3ViIjoidG9rZW4ifQ.lymD-8vwSySjGliJ5GVA7f0HIpmRWu9SuGV53mDxNrU');
INSERT INTO `user` VALUES (3, '1', '1', 'coach', NULL, '2023-02-18 16:02:46', '', '202102222222', '1', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6IjMiLCJleHAiOjE2NzczMDQzNjAsImlhdCI6MTY3NjY5OTU2MCwiaXNzIjoi5rC055eVIiwic3ViIjoidG9rZW4ifQ.lymD-8vwSySjGliJ5GVA7f0HIpmRWu9SuGV53mDxNrU');

SET FOREIGN_KEY_CHECKS = 1;
