SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '文章id',
  `category_id` int UNSIGNED NOT NULL COMMENT '分类id',
  `content` longtext  NOT NULL COMMENT '文章内容',
  `title` varchar(1024)  NOT NULL COMMENT '文章标题',
  `view_count` int UNSIGNED NOT NULL COMMENT '阅读次数',
  `comment_count` int UNSIGNED NOT NULL COMMENT '评论次数',
  `username` varchar(128)  NOT NULL COMMENT '作者',
  `status` int UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态',
  `summary` varchar(256) NOT NULL COMMENT '文章摘要',
  `created_at` datetime  NOT NULL DEFAULT CURRENT_TIMESTAMP  COMMENT '发布时间',
  `updated_at` datetime  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '更新时间',
  `deleted_at` datetime NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_view_count`(`view_count`) USING BTREE COMMENT '阅读次数索引',
  INDEX `idx_comment_count`(`comment_count`) USING BTREE COMMENT '评论数索引',
  INDEX `idx_category_id`(`category_id`) USING BTREE COMMENT '分类id索引'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分类',
  `category_name` varchar(255) NOT NULL COMMENT '分类名字',
  `category_no` int UNSIGNED NOT NULL COMMENT '分类排序',
  `created_at` datetime  NOT NULL DEFAULT CURRENT_TIMESTAMP  COMMENT '发布时间',
  `updated_at` datetime  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '更新时间',
  `deleted_at` datetime NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES (1, 'css/html', 1, '2018-08-12 10:55:45', '2018-08-12 10:59:00',NULL);
INSERT INTO `category` VALUES (2, '后端开发', 2, '2018-08-12 10:56:07', '2018-08-12 10:59:03',NULL);
INSERT INTO `category` VALUES (3, 'Java开发', 3, '2018-08-12 10:56:16', '2018-08-12 10:59:05',NULL);
INSERT INTO `category` VALUES (4, 'C++开发', 4, '2018-08-12 10:56:24', '2018-08-12 10:59:08',NULL);
INSERT INTO `category` VALUES (5, '架构剖析', 5, '2018-08-12 10:56:36', '2018-08-12 10:59:10',NULL);
INSERT INTO `category` VALUES (6, 'Golang开发', 6, '2018-08-12 10:56:45', '2018-08-12 10:59:14',NULL);

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `content` text  NOT NULL COMMENT '评论内容',
  `username` varchar(64)  NOT NULL COMMENT '评论作者',
  `created_at` datetime  NOT NULL DEFAULT CURRENT_TIMESTAMP  COMMENT '发布时间',
  `updated_at` datetime  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '更新时间',
  `deleted_at` datetime NULL COMMENT '删除时间',
  `status` int UNSIGNED NOT NULL COMMENT '评论状态: 0, 删除；1， 正常',
  `article_id` int UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for leave
-- ----------------------------
DROP TABLE IF EXISTS `leave`;
CREATE TABLE `leave`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '留言id',
  `username` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `created_at` datetime  NOT NULL DEFAULT CURRENT_TIMESTAMP  COMMENT '发布时间',
  `updated_at` datetime  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '更新时间',
  `deleted_at` datetime NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for leave
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '留言id',
  `created_at` datetime  NOT NULL DEFAULT CURRENT_TIMESTAMP  COMMENT '发布时间',
  `updated_at` datetime  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  COMMENT '更新时间',
  `deleted_at` datetime NULL COMMENT '删除时间',
  `user_id` int UNSIGNED NOT NULL,
  `username` varchar(16) NOT NULL,
  `password` varchar(64) NOT NULL,
  `telephone` int UNSIGNED NULL,
  `avatar` varchar(64) NOT NULL,
  `nick_name` varchar(16) NULL,
  `email` varchar(64) NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;



SET FOREIGN_KEY_CHECKS = 1;