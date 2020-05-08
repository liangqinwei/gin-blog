
-- 标签表
CREATE TABLE `blog_tag` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(32) NOT NULL COMMENT '标签名称',
  `created_by` varchar(32) NOT NULL DEFAULT '' COMMENT '创建人',
  `modified_by` varchar(32) NOT NULL DEFAULT '' COMMENT '修改人',
  `summary` varchar(256) NOT NULL DEFAULT '' COMMENT '摘要',
  `state` tinyint(5) NOT NULL DEFAULT '0' COMMENT 'tag状态,0为禁用1为启用',
  `deleted_time` datetime DEFAULT NULL COMMENT '删除时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `index_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章表';

-- 文章表
CREATE TABLE `article` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `article_id` varchar(32) NOT NULL COMMENT '文章编号',
  `tag_id` varchar(32) NOT NULL COMMENT '节点id',
  `title` varchar(256) NOT NULL COMMENT '文章标题',
  `summary` varchar(256) NOT NULL DEFAULT '' COMMENT '摘要',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT '文章内容',
  `status` tinyint(5) NOT NULL DEFAULT '0' COMMENT '文章状态,0为禁用1为启用',
  `create_by` varchar(32) NOT NULL DEFAULT '' COMMENT '创建者',
  `update_by` varchar(32) NOT NULL DEFAULT '' COMMENT '更新者',
  `deleted_time` datetime DEFAULT NULL COMMENT '删除时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_articleid` (`article_id`),
  UNIQUE KEY `title` (`title`),
  KEY `index_title` (`title`),
  KEY `index_summary` (`summary`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章表';

-- 用户表
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `username` varchar(20) NOT NULL COMMENT '名字',
  `nickname` varchar(20) NOT NULL DEFAULT '' COMMENT '昵称',
  `password` varchar(555) NOT NULL  COMMENT '密码',
  `userface` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `telephone` varchar(13) NOT NULL DEFAULT '' COMMENT '电话',
  `mail` varchar(20) NOT NULL DEFAULT '' COMMENT '邮箱',
  `isDelete` int(10) NOT NULL DEFAULT '0' COMMENT '是否删除,0:不删除，1:删除',
  `deleted_time` datetime DEFAULT NULL COMMENT '删除时间',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录修改时间',
  PRIMARY KEY (`id`),
  KEY `index_name` (`username`),
  KEY `index_phone` (`telephone`)
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- 用户角色映射表
CREATE TABLE `user_role` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `uid` bigint(20) NOT NULL COMMENT '用户id',
  `rid` bigint(20) NOT NULL COMMENT '角色id',
  `is_delete` int(5) NOT NULL DEFAULT '0' COMMENT '角色状态，0:正常，1删除',
  `deleted_time` datetime DEFAULT NULL COMMENT '删除时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户角色映射表';

-- 用户角色
CREATE TABLE `role` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `role` varchar(20) NOT NULL COMMENT '名称',
  `is_delete` int(10) NOT NULL DEFAULT '0' COMMENT '是否删除,0:不删除，1:删除',
  `description` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `deleted_time` datetime DEFAULT NULL COMMENT '删除时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

