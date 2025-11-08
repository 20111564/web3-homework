CREATE TABLE `employee` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '姓名',
  `department` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '部门',
  `salary` decimal(10,2) DEFAULT NULL COMMENT '薪水',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `books` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `title` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '书名',
  `author` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '作者',
  `price` decimal(10,2) DEFAULT NULL COMMENT '价格',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;