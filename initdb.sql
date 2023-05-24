-- lets_study.newtable definition
-- 生成用户表
CREATE TABLE `user_info` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户名称',
  `age` tinyint DEFAULT NULL COMMENT '用户年龄',
  `username` varchar(100) NOT NULL COMMENT '用户名，由英文字母，数字，下划线组成',
  `password` varchar(100) NOT NULL COMMENT '用户密码，加密',
  `power` tinyint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_info_un` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
INSERT INTO lets_study.user_info (name,age,username,password,power) VALUES ('tututuf',0,'root','root',5) 


