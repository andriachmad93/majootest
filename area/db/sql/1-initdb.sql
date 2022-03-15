USE majooapptest;

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `area_value` bigint(20) DEFAULT NULL,
  `type` varchar(45) DEFAULT NULL
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;
