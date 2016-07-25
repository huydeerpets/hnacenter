/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 50617
Source Host           : localhost:3306
Source Database       : hnacenter

Target Server Type    : MYSQL
Target Server Version : 50617
File Encoding         : 65001

Date: 2016-07-25 17:16:55
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for company
-- ----------------------------
DROP TABLE IF EXISTS `company`;
CREATE TABLE `company` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `code` varchar(255) NOT NULL DEFAULT '',
  `name` varchar(255) NOT NULL DEFAULT '',
  `status` int(11) NOT NULL DEFAULT '1',
  `expiration_time` bigint(20) NOT NULL DEFAULT '0',
  `add_time` bigint(20) NOT NULL DEFAULT '0',
  `login_url` varchar(255) NOT NULL DEFAULT '',
  `update_url` varchar(255) NOT NULL DEFAULT '',
  `soft_version` bigint(20) NOT NULL DEFAULT '0',
  `soft_update_url` varchar(255) NOT NULL DEFAULT '',
  `soft_md5` varchar(255) NOT NULL DEFAULT '',
  `soft_describe` varchar(255) NOT NULL DEFAULT '',
  `soft_cleardata` int(11) NOT NULL DEFAULT '0',
  `rom_version` bigint(20) NOT NULL DEFAULT '0',
  `rom_update_url` varchar(255) NOT NULL DEFAULT '',
  `rom_md5` varchar(255) NOT NULL DEFAULT '',
  `rom_describe` varchar(255) NOT NULL DEFAULT '',
  `rom_cleardata` int(11) NOT NULL DEFAULT '0',
  `apk_version` bigint(20) NOT NULL DEFAULT '0',
  `apk_update_url` varchar(255) NOT NULL DEFAULT '',
  `apk_md5` varchar(255) NOT NULL DEFAULT '',
  `apk_describe` varchar(255) NOT NULL DEFAULT '',
  `apk_clerdata` int(11) NOT NULL DEFAULT '0',
  `customer_url` varchar(255) NOT NULL DEFAULT '',
  `images_url` varchar(255) NOT NULL DEFAULT '',
  `wechat_data_url` varchar(255) NOT NULL DEFAULT '',
  `user_max_count` bigint(20) NOT NULL DEFAULT '0',
  `total_max_count` bigint(20) NOT NULL DEFAULT '0',
  `func_customer` int(11) NOT NULL DEFAULT '1',
  `func_image` int(11) NOT NULL DEFAULT '1',
  `func_upload_wechat` int(11) NOT NULL DEFAULT '1',
  `func_limit_customer` int(11) NOT NULL DEFAULT '1',
  `func_auto_add` int(11) NOT NULL DEFAULT '1',
  `func_auto_search` int(11) NOT NULL DEFAULT '1',
  `func_location` int(11) NOT NULL DEFAULT '1',
  `func_shake` int(11) NOT NULL DEFAULT '1',
  `custtomer_push` bigint(20) NOT NULL DEFAULT '0',
  `customer_user` bigint(20) NOT NULL DEFAULT '888888',
  `customer_num` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of company
-- ----------------------------
INSERT INTO `company` VALUES ('1', '12587', '湖北雅倩集团', '1', '0', '0', '', '', '0', '', '', '', '0', '0', '', '', '', '0', '0', '', '', '', '0', '', '', '', '20', '852', '1', '1', '1', '1', '1', '1', '1', '1', '0', '888888', '0');
INSERT INTO `company` VALUES ('3', '85497', '湖北烽火科技', '1', '0', '0', '', '', '0', '', '', '', '0', '0', '', '', '', '0', '0', '', '', '', '0', '', '', '', '854', '99875', '1', '1', '1', '1', '1', '1', '1', '1', '0', '888888', '0');
INSERT INTO `company` VALUES ('4', '15487', '湖北野火集团', '0', '-62135596800', '-62135596800', '位lgujsdhlfkjg', '地方规划的法规【’', '0', 'fghjgh', 'kl', '减肥vghj', '0', '0', '', '估计', 'dfghdfgh\'', '0', '0', '估计', '泛海国际、', '大发光火', '0', 'fghdfgh', 'dfghdfg', '地方规划的法规【‘', '16', '8520', '1', '1', '1', '1', '1', '1', '1', '1', '0', '0', '0');

-- ----------------------------
-- Table structure for company_config
-- ----------------------------
DROP TABLE IF EXISTS `company_config`;
CREATE TABLE `company_config` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `code` varchar(255) NOT NULL DEFAULT '',
  `databse_info` varchar(255) NOT NULL DEFAULT '',
  `addr` varchar(255) NOT NULL DEFAULT '',
  `httpport` int(11) NOT NULL DEFAULT '0',
  `day_max_download_count` int(11) NOT NULL DEFAULT '150',
  `day_max_hid_download` int(11) NOT NULL DEFAULT '30',
  `clear_data_times` varchar(255) NOT NULL DEFAULT '0 28 17 * * *',
  `status` int(11) NOT NULL DEFAULT '1',
  `server_addr` varchar(255) NOT NULL DEFAULT '',
  `ormdebug` int(11) NOT NULL DEFAULT '0',
  `testdebug` int(11) NOT NULL DEFAULT '0',
  `maxuploadimages` int(11) NOT NULL DEFAULT '150',
  `adminuser` bigint(20) NOT NULL DEFAULT '51875511',
  `recycle` bigint(20) NOT NULL DEFAULT '999999',
  PRIMARY KEY (`id`),
  UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of company_config
-- ----------------------------

-- ----------------------------
-- Table structure for company_mobile_provinces
-- ----------------------------
DROP TABLE IF EXISTS `company_mobile_provinces`;
CREATE TABLE `company_mobile_provinces` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `company_id` bigint(20) NOT NULL,
  `mobile_province_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of company_mobile_provinces
-- ----------------------------
INSERT INTO `company_mobile_provinces` VALUES ('1', '1', '2');
INSERT INTO `company_mobile_provinces` VALUES ('2', '3', '2');
INSERT INTO `company_mobile_provinces` VALUES ('3', '4', '2');
INSERT INTO `company_mobile_provinces` VALUES ('4', '1', '4');
INSERT INTO `company_mobile_provinces` VALUES ('5', '1', '5');

-- ----------------------------
-- Table structure for device_info
-- ----------------------------
DROP TABLE IF EXISTS `device_info`;
CREATE TABLE `device_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `davice` varchar(255) NOT NULL DEFAULT '',
  `register_time` bigint(20) NOT NULL DEFAULT '0',
  `status` int(11) NOT NULL DEFAULT '0',
  `user_id` varchar(255) NOT NULL DEFAULT '',
  `serial` varchar(255) NOT NULL DEFAULT '',
  `deviceid` varchar(255) NOT NULL DEFAULT '',
  `mac_addr` varchar(255) NOT NULL DEFAULT '',
  `bass_addr` varchar(255) NOT NULL DEFAULT '',
  `ssid` varchar(255) NOT NULL DEFAULT '',
  `android` varchar(255) NOT NULL DEFAULT '',
  `imsi` varchar(255) NOT NULL DEFAULT '',
  `iccid` varchar(255) NOT NULL DEFAULT '',
  `phone_number` varchar(255) NOT NULL DEFAULT '',
  `company` varchar(255) NOT NULL DEFAULT '',
  `ip_address` varchar(255) NOT NULL DEFAULT '',
  `device_type` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `device_info_davice` (`davice`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of device_info
-- ----------------------------
INSERT INTO `device_info` VALUES ('1', '522', '125410545', '0', 'hjkhk', 'jhkjh', '452', '525', '4253', '53252', '5242', '4524', 'jnmk', '45656', '78645645', '456456', '0');

-- ----------------------------
-- Table structure for mobile_province
-- ----------------------------
DROP TABLE IF EXISTS `mobile_province`;
CREATE TABLE `mobile_province` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `provinceid` bigint(20) NOT NULL DEFAULT '0',
  `provincename` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `mobile_province_provinceid` (`provinceid`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of mobile_province
-- ----------------------------
INSERT INTO `mobile_province` VALUES ('1', '1', '湖北');
INSERT INTO `mobile_province` VALUES ('2', '2', '厦门');
INSERT INTO `mobile_province` VALUES ('3', '3', '澳门');
INSERT INTO `mobile_province` VALUES ('4', '4', '宁夏');
INSERT INTO `mobile_province` VALUES ('5', '5', '福建');

-- ----------------------------
-- Table structure for resource
-- ----------------------------
DROP TABLE IF EXISTS `resource`;
CREATE TABLE `resource` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `reskey` varchar(64) NOT NULL DEFAULT '',
  `level` bigint(20) NOT NULL DEFAULT '1',
  `fid` bigint(20) NOT NULL DEFAULT '0',
  `url` varchar(64) NOT NULL DEFAULT '',
  `status` bigint(20) NOT NULL DEFAULT '0',
  `sort` bigint(20) NOT NULL DEFAULT '0',
  `ico` varchar(64) NOT NULL DEFAULT '',
  `isfunction` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of resource
-- ----------------------------
INSERT INTO `resource` VALUES ('1', '用户管理', '1', '0', 'user', '0', '1', 'menu-icon fa fa-user', '0');
INSERT INTO `resource` VALUES ('2', '全部用户', '2', '1', '/hnacenter/user/list', '0', '2', '全部用户', '0');
INSERT INTO `resource` VALUES ('3', '新增用户', '2', '1', '/hnacenter/user/add', '0', '2', '新增用户', '0');
INSERT INTO `resource` VALUES ('4', '编辑用户', '1000', '-1', '/hnacenter/user/edit', '0', '2', '编辑用户', '0');
INSERT INTO `resource` VALUES ('5', '删除用户', '1000', '-1', '/hnacenter/user/delete', '0', '2', '删除用户', '0');
INSERT INTO `resource` VALUES ('7', '角色管理', '1', '0', 'role', '0', '1', 'menu-icon fa fa fa-users', '0');
INSERT INTO `resource` VALUES ('8', '全部角色', '2', '7', '/hnacenter/role/list', '0', '2', '全部角色', '0');
INSERT INTO `resource` VALUES ('9', '新增角色', '2', '7', '/hnacenter/role/add', '0', '2', '新增角色', '0');
INSERT INTO `resource` VALUES ('10', '编辑角色', '1000', '-1', '/hnacenter/role/edit', '0', '2', '编辑角色', '0');
INSERT INTO `resource` VALUES ('11', '删除角色', '1000', '-1', '/hnacenter/role/delete', '0', '2', '删除角色', '0');
INSERT INTO `resource` VALUES ('12', '设置权限', '1000', '-1', '/hnacenter/user/allocation', '0', '2', '给相应的角色分配相应的权限', '0');
INSERT INTO `resource` VALUES ('13', '资源管理', '1', '0', 'resource', '0', '1', 'menu-icon fa fa-folder', '0');
INSERT INTO `resource` VALUES ('14', '全部资源', '2', '13', '/hnacenter/resource/list', '0', '2', '全部资源', '0');
INSERT INTO `resource` VALUES ('15', '新增资源', '2', '13', '/hnacenter/resource/add', '0', '2', '新增资源', '0');
INSERT INTO `resource` VALUES ('16', '编辑资源', '1000', '-1', '/hnacenter/resource/edit', '0', '2', '编辑资源', '0');
INSERT INTO `resource` VALUES ('17', '删除资源', '1000', '-1', '/hnacenter/resource/delete', '0', '2', '删除资源', '0');
INSERT INTO `resource` VALUES ('18', '公司管理', '1', '0', 'company', '0', '1', 'menu-icon fa fa-sitemap', '0');
INSERT INTO `resource` VALUES ('19', '公司列表', '2', '18', '/hnacenter/company/list', '0', '2', '公司列表', '0');
INSERT INTO `resource` VALUES ('20', '添加公司', '2', '18', '/hnacenter/company/add', '0', '2', '添加公司', '0');
INSERT INTO `resource` VALUES ('21', '修改公司信息', '1000', '-1', '/hnacenter/company/edit', '0', '2', '修改公司信息', '0');
INSERT INTO `resource` VALUES ('22', '删除公司', '1000', '-1', '/hnacenter/company/delete', '0', '2', '删除公司', '0');
INSERT INTO `resource` VALUES ('23', '修改公司状态', '1000', '-1', '/hnacenter/company/changestatus', '0', '2', '修改公司状态', '0');
INSERT INTO `resource` VALUES ('24', '搜索公司', '1000', '-1', '/hnacenter/company/select', '0', '2', '搜索公司', '0');
INSERT INTO `resource` VALUES ('25', '设备管理', '1', '0', 'device', '0', '1', 'menu-icon fa fa-desktop', '0');
INSERT INTO `resource` VALUES ('26', '设备列表', '2', '25', '/hnacenter/device/list', '0', '2', '设备列表', '0');
INSERT INTO `resource` VALUES ('27', '删除设备', '1000', '-1', '/hnacenter/device/delete', '0', '2', '删除设备', '0');

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `rolekey` varchar(64) NOT NULL DEFAULT '',
  `name` varchar(64) NOT NULL DEFAULT '',
  `remark` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES ('1', '2e365bc5d8ad8ef8b', '超级管理员', '超级管理员');
INSERT INTO `role` VALUES ('2', '2e365bc5d8ad8ef8c', '研发经理', '研发部经理');
INSERT INTO `role` VALUES ('3', '2e365bc5d8ad8ef8e', '财务经理', '财务经理');
INSERT INTO `role` VALUES ('4', '2e365bc5d8ad8ef8e', '人事经理', '人事部经理');
INSERT INTO `role` VALUES ('5', '2e365bc5d8ad8ef8f', '研发组长', '研发部组长');
INSERT INTO `role` VALUES ('6', '2e365bc5d8ad8ef8g', 'WEB开发工程师', 'WEB开发工程师');
INSERT INTO `role` VALUES ('7', '2e365bc5d8ad8ef8h', '财务部会计', '财务部会计');
INSERT INTO `role` VALUES ('8', '2e365bc5d8ad8ef8r', '前台', '前台');
INSERT INTO `role` VALUES ('9', '2e365bc5d8ad8ef8r', 'HR', 'HR');
INSERT INTO `role` VALUES ('10', 'alsdfjlakdjfgajdflkasdjfhsd', 'java后台工程师', 'Java后台工程师');

-- ----------------------------
-- Table structure for role_resources
-- ----------------------------
DROP TABLE IF EXISTS `role_resources`;
CREATE TABLE `role_resources` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `role_id` bigint(20) NOT NULL,
  `resource_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of role_resources
-- ----------------------------
INSERT INTO `role_resources` VALUES ('19', '4', '1');
INSERT INTO `role_resources` VALUES ('20', '4', '2');
INSERT INTO `role_resources` VALUES ('21', '4', '3');
INSERT INTO `role_resources` VALUES ('22', '4', '7');
INSERT INTO `role_resources` VALUES ('23', '4', '8');
INSERT INTO `role_resources` VALUES ('24', '4', '9');
INSERT INTO `role_resources` VALUES ('25', '4', '13');
INSERT INTO `role_resources` VALUES ('26', '4', '14');
INSERT INTO `role_resources` VALUES ('27', '4', '15');
INSERT INTO `role_resources` VALUES ('28', '2', '1');
INSERT INTO `role_resources` VALUES ('29', '2', '2');
INSERT INTO `role_resources` VALUES ('30', '2', '3');
INSERT INTO `role_resources` VALUES ('31', '2', '7');
INSERT INTO `role_resources` VALUES ('32', '2', '8');
INSERT INTO `role_resources` VALUES ('33', '2', '9');
INSERT INTO `role_resources` VALUES ('34', '2', '13');
INSERT INTO `role_resources` VALUES ('35', '2', '14');
INSERT INTO `role_resources` VALUES ('36', '2', '15');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `usercode` varchar(255) NOT NULL DEFAULT '',
  `username` varchar(255) NOT NULL DEFAULT '',
  `pwd` varchar(64) NOT NULL DEFAULT '',
  `rid` bigint(20) NOT NULL DEFAULT '0',
  `status` bigint(20) NOT NULL DEFAULT '0',
  `remark` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `usercode` (`usercode`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', '100', 'admin', '21232f297a57a5a743894a0e4a801fc3', '1', '0', '超级管理员');
INSERT INTO `user` VALUES ('2', '101', 'zhang', '21232f297a57a5a743894a0e4a801fc3', '2', '0', 'admin111');
INSERT INTO `user` VALUES ('3', '125', 'wang', '21232f297a57a5a743894a0e4a801fc3', '8', '0', 'asdkljfa;jflkasjdfli阿法加快递费');
INSERT INTO `user` VALUES ('4', '103', '李华', '21232f297a57a5a743894a0e4a801fc3', '4', '0', 'akdjfhksdljfasdfasdfasdfasd');
INSERT INTO `user` VALUES ('5', '104', 'li', '21232f297a57a5a743894a0e4a801fc3', '5', '0', 'admin114');
INSERT INTO `user` VALUES ('7', '106', 'admin116', '21232f297a57a5a743894a0e4a801fc3', '7', '0', 'admin116');
INSERT INTO `user` VALUES ('10', '6987', '你好', '123', '9', '0', '前台');

-- ----------------------------
-- Table structure for user_roles
-- ----------------------------
DROP TABLE IF EXISTS `user_roles`;
CREATE TABLE `user_roles` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `role_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user_roles
-- ----------------------------
