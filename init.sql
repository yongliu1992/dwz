CREATE DATABASE `dwz` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */;
use dwz;
CREATE TABLE `wzs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sub` varchar(60) COLLATE utf8mb4_unicode_ci NOT NULL,
  `origin` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL,
  `add_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `wz` (`sub`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci