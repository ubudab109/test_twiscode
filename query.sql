DROP TABLE IF EXISTS `transactions`;

CREATE TABLE `transactions` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `order_at` datetime DEFAULT NULL,
  `status_paid` enum('pending','lunas','gagal') DEFAULT NULL,
  `paid-at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

/*Data for the table `transactions` */

insert  into `transactions`(`id`,`order_at`,`status_paid`,`paid-at`) values 
(1,'2022-05-08 16:31:13','lunas','2022-05-17 16:31:17'),
(2,'2022-05-09 16:31:21','pending',NULL);

/*Table structure for table `transactions_detail` */

DROP TABLE IF EXISTS `transactions_detail`;

CREATE TABLE `transactions_detail` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `transaction_id` bigint(20) unsigned NOT NULL,
  `price` double DEFAULT NULL,
  `qty` int(11) DEFAULT NULL,
  `sub_total` double DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `transaction_id` (`transaction_id`),
  CONSTRAINT `transactions_detail_ibfk_1` FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

/*Data for the table `transactions_detail` */

insert  into `transactions_detail`(`id`,`transaction_id`,`price`,`qty`,`sub_total`) values 
(1,1,10000,2,20000),
(2,2,6250,4,25000);
