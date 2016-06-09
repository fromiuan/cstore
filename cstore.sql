/*
Navicat MySQL Data Transfer

Source Server         : local
Source Server Version : 50620
Source Host           : localhost:3306
Source Database       : cstore

Target Server Type    : MYSQL
Target Server Version : 50620
File Encoding         : 65001

Date: 2015-02-25 10:34:30
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `group`
-- ----------------------------
DROP TABLE IF EXISTS `group`;
CREATE TABLE `group` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `title` varchar(100) NOT NULL,
  `status` int(11) NOT NULL,
  `sort` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of group
-- ----------------------------
INSERT INTO `group` VALUES ('1', '审核组', 'System', '2', '1');
INSERT INTO `group` VALUES ('2', '客服组', '客服', '2', '1');

-- ----------------------------
-- Table structure for `member`
-- ----------------------------
DROP TABLE IF EXISTS `member`;
CREATE TABLE `member` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uname` varchar(64) NOT NULL,
  `name` varchar(32) DEFAULT NULL,
  `pwd` varchar(64) NOT NULL,
  `status` varchar(64) NOT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `phone` varchar(64) NOT NULL,
  `address` varchar(128) NOT NULL,
  `email` varchar(64) DEFAULT NULL,
  `birthday` varchar(64) DEFAULT NULL,
  `experience` varchar(32) DEFAULT NULL,
  `grade` varchar(32) DEFAULT NULL,
  `client_type` varchar(32) DEFAULT NULL,
  `logintime` datetime DEFAULT NULL,
  `ctime` datetime NOT NULL,
  `utime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of member
-- ----------------------------

-- ----------------------------
-- Table structure for `resource`
-- ----------------------------
DROP TABLE IF EXISTS `resource`;
CREATE TABLE `resource` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL,
  `pid` int(11) NOT NULL DEFAULT '0',
  `key` varchar(64) NOT NULL,
  `type` varchar(10) DEFAULT NULL,
  `url` varchar(64) DEFAULT NULL,
  `level` int(11) DEFAULT NULL,
  `description` varchar(200) DEFAULT NULL,
  `group_id` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of resource
-- ----------------------------
INSERT INTO `resource` VALUES ('1', '用户管理', '0', 'user', '0', 'user', '1', '用户管理', '0');
INSERT INTO `resource` VALUES ('2', '全部用户', '1', 'user_list', '1', 'cstore/user/list', '2', '查看所有的用户，以及相应的操作', '0');
INSERT INTO `resource` VALUES ('3', '新增用户', '1', 'user_add', '1', 'cstore/user/add', '2', '添加相关的成员', '0');
INSERT INTO `resource` VALUES ('4', '角色管理', '0', 'role', '0', 'role', '1', '角色管理', '0');
INSERT INTO `resource` VALUES ('5', '角色列表', '4', 'role_list', '1', 'cstore/role/list', '2', '展示所有的角色名称', '0');
INSERT INTO `resource` VALUES ('6', '新增角色', '4', 'role_add', '1', 'cstore/role/add', '2', '角色添加', '0');
INSERT INTO `resource` VALUES ('7', '资源管理', '0', 'resource', '0', 'resource', '1', '资源管理', '0');
INSERT INTO `resource` VALUES ('8', '资源列表', '7', 'resource_list', '1', 'cstore/resource/list', '2', '所有权限', '0');
INSERT INTO `resource` VALUES ('9', '分配角色', '1000', 'user_allocation_role', '2', 'cstore/user/role', '2', '给相应的用户分配角色', '0');
INSERT INTO `resource` VALUES ('10', '分配权限', '1000', 'role_allocation_permission', '2', 'cstore/role/resource', '2', '给相应的角色分配相应的权限', '0');
INSERT INTO `resource` VALUES ('11', '编辑用户', '1000', 'edit_user', '2', 'cstore/user/edit', '2', '编辑用户信息，包括密码', '0');
INSERT INTO `resource` VALUES ('12', '删除用户', '1000', 'del_user', '2', 'cstore/user/delete', '2', '删除相应的用户', '0');
INSERT INTO `resource` VALUES ('13', '显示角色信息', '1000', 'show_role', '2', 'cstore/role/show', '2', '显示角色的详细信息', '0');
INSERT INTO `resource` VALUES ('14', '编辑角色', '1000', 'edit_role', '2', 'cstore/role/edit', '2', '编辑角色信息', '0');
INSERT INTO `resource` VALUES ('15', '删除角色', '1000', 'del_role', '2', 'cstore/role/delete', '2', '删除角色信息', '0');
INSERT INTO `resource` VALUES ('16', '显示资源信息', '1000', 'show_res', '2', 'cstore/resource/show', '2', '显示资源信息', '0');
INSERT INTO `resource` VALUES ('17', '编辑权限资源', '1000', 'edit_res', '2', 'cstore/resource/edit', '2', '编辑权限资源', '0');
INSERT INTO `resource` VALUES ('18', '删除权限资源', '1000', 'del_res', '2', 'cstore/resource/delete', '2', '删除资源', '0');
INSERT INTO `resource` VALUES ('19', '资源添加', '7', 'resource_add', '1', 'cstore/resource/add', '2', '资源添加', '0');
INSERT INTO `resource` VALUES ('20', ' 用户信息', '1000', 'user_show', '2', 'cstore/user/show  ', '2', '显示单个用户信息', '0');

-- ----------------------------
-- Table structure for `resource_roles`
-- ----------------------------
DROP TABLE IF EXISTS `resource_roles`;
CREATE TABLE `resource_roles` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `resource_id` bigint(20) NOT NULL,
  `role_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of resource_roles
-- ----------------------------
INSERT INTO `resource_roles` VALUES ('13', '1', '1');
INSERT INTO `resource_roles` VALUES ('14', '2', '1');
INSERT INTO `resource_roles` VALUES ('15', '3', '1');
INSERT INTO `resource_roles` VALUES ('16', '4', '1');
INSERT INTO `resource_roles` VALUES ('17', '5', '1');
INSERT INTO `resource_roles` VALUES ('18', '6', '1');
INSERT INTO `resource_roles` VALUES ('19', '7', '1');
INSERT INTO `resource_roles` VALUES ('20', '8', '1');
INSERT INTO `resource_roles` VALUES ('21', '19', '1');
INSERT INTO `resource_roles` VALUES ('22', '9', '1');
INSERT INTO `resource_roles` VALUES ('23', '1000', '1');
INSERT INTO `resource_roles` VALUES ('24', '10', '1');
INSERT INTO `resource_roles` VALUES ('32', '2', '2');
INSERT INTO `resource_roles` VALUES ('33', '1', '2');
INSERT INTO `resource_roles` VALUES ('34', '4', '2');
INSERT INTO `resource_roles` VALUES ('35', '5', '2');
INSERT INTO `resource_roles` VALUES ('36', '19', '2');
INSERT INTO `resource_roles` VALUES ('37', '7', '2');

-- ----------------------------
-- Table structure for `role`
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL,
  `key` varchar(64) NOT NULL,
  `description` varchar(200) DEFAULT NULL,
  `status` int(2) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES ('1', '管理员', 'admin', '这是一个管理员最大权限', '2');
INSERT INTO `role` VALUES ('2', '           admin', ' admin          ', '试的管理员权限', '1');
INSERT INTO `role` VALUES ('3', '被的想法', 'vd', 'V的', '1');

-- ----------------------------
-- Table structure for `storer`
-- ----------------------------
DROP TABLE IF EXISTS `storer`;
CREATE TABLE `storer` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uname` varchar(64) NOT NULL,
  `pwd` varchar(64) NOT NULL,
  `name` varchar(64) NOT NULL,
  `logo` varchar(64) DEFAULT NULL,
  `expires` datetime NOT NULL,
  `tel` varchar(32) DEFAULT NULL,
  `phone` varchar(32) DEFAULT NULL,
  `address` varchar(64) NOT NULL,
  `ctime` datetime NOT NULL,
  `utime` datetime DEFAULT NULL,
  `logintime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of storer
-- ----------------------------

-- ----------------------------
-- Table structure for `user`
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uname` varchar(64) NOT NULL,
  `pwd` varchar(64) NOT NULL,
  `nickname` varchar(64) NOT NULL,
  `email` varchar(64) NOT NULL,
  `phone` varchar(64) NOT NULL,
  `status` int(11) NOT NULL,
  `remark` varchar(200) DEFAULT NULL,
  `logintime` datetime DEFAULT NULL,
  `ctime` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', 'admin', '54dea6d17eba63690c1f8ec4b4888e99', 'Admin', 'qi19901212@163.com', '18510970061', '2', '孙凤齐哈哈哈', '2015-01-18 21:26:33', '2015-01-16 09:58:08');
INSERT INTO `user` VALUES ('4', 'sunqi', 'c42b11d1e88ea4d0e9f25fd9c0c9eb3e', '孙齐', 'qi19901212@163.com', '18311376490', '1', '这是一个go程序员', null, '2015-01-27 08:02:58');

-- ----------------------------
-- Table structure for `user_roles`
-- ----------------------------
DROP TABLE IF EXISTS `user_roles`;
CREATE TABLE `user_roles` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL DEFAULT '0',
  `role_id` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user_roles
-- ----------------------------
INSERT INTO `user_roles` VALUES ('1', '1', '1');
INSERT INTO `user_roles` VALUES ('9', '4', '2');
