/*
SQLyog Ultimate v12.14 (64 bit)
MySQL - 10.4.17-MariaDB : Database - sitedb
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`sitedb` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `sitedb`;

/*Table structure for table `accses` */

DROP TABLE IF EXISTS `accses`;

CREATE TABLE `accses` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL,
  `id_handler` int(11) NOT NULL,
  PRIMARY KEY (`id`,`role_id`,`id_handler`),
  KEY `id_role` (`role_id`),
  KEY `id_handler` (`id_handler`),
  CONSTRAINT `accses_ibfk_4` FOREIGN KEY (`id_handler`) REFERENCES `handlers` (`id`),
  CONSTRAINT `accses_ibfk_5` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8;

/*Data for the table `accses` */

insert  into `accses`(`id`,`role_id`,`id_handler`) values 
(1,1,1),
(2,1,2),
(3,1,3),
(4,1,4),
(5,1,5),
(6,1,6),
(7,1,7),
(8,1,8),
(9,1,9),
(10,1,10),
(11,1,11),
(12,1,12),
(13,1,13),
(14,2,1),
(15,2,2),
(16,2,3),
(17,2,4),
(18,2,5),
(19,2,9),
(20,2,11),
(21,2,12),
(22,2,13);

/*Table structure for table `directs` */

DROP TABLE IF EXISTS `directs`;

CREATE TABLE `directs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `direct` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

/*Data for the table `directs` */

insert  into `directs`(`id`,`direct`) values 
(1,'D:\\projekts\\tests');

/*Table structure for table `handlers` */

DROP TABLE IF EXISTS `handlers`;

CREATE TABLE `handlers` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `method` varchar(10) NOT NULL,
  `info` varchar(40) DEFAULT NULL,
  `status` tinyint(1) NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;

/*Data for the table `handlers` */

insert  into `handlers`(`id`,`name`,`method`,`info`,`status`) values 
(1,'files.PostFile','POST',NULL,1),
(2,'files.GetUserFile','GET',NULL,1),
(3,'files.PutFile','PUT',NULL,1),
(4,'files.DeleteFile','DELETE',NULL,1),
(5,'users.GetUser','GET',NULL,1),
(6,'users.GetUsers','GET',NULL,1),
(7,'users.PostUser','POST',NULL,1),
(8,'users.DeleteUser','DELETE',NULL,1),
(9,'users.PutUser','PUT',NULL,1),
(10,'files.GetFile','GET',NULL,1),
(11,'users.UserChangePassword','PUT',NULL,1),
(12,'users.chenckLogin','GET',NULL,1),
(13,'users.permission','GET',NULL,1);

/*Table structure for table `log_errors` */

DROP TABLE IF EXISTS `log_errors`;

CREATE TABLE `log_errors` (
  `date` varchar(20) DEFAULT NULL,
  `error_type` varbinary(255) DEFAULT NULL,
  `error` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `log_errors` */

/*Table structure for table `musics` */

DROP TABLE IF EXISTS `musics`;

CREATE TABLE `musics` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `atist` varchar(20) DEFAULT NULL,
  `name` varchar(30) DEFAULT NULL,
  `name_orig` varchar(60) DEFAULT NULL,
  `size` varchar(10) DEFAULT NULL,
  `duration` varchar(8) DEFAULT NULL,
  `downloaded` int(11) DEFAULT NULL,
  `date_upload` varchar(11) DEFAULT NULL,
  `id_user` varchar(30) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `id_user` (`id_user`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8;

/*Data for the table `musics` */

insert  into `musics`(`id`,`atist`,`name`,`name_orig`,`size`,`duration`,`downloaded`,`date_upload`,`id_user`) values 
(20,'i am','no','i_am_-_no.mp3','3.13 mb','01:00',NULL,'11-19-2021 ','32418eb3-02e1-4b85-8770-dedc75'),
(21,'i am2','no','i_am2_-_no.mp3','3.13 mb','01:00',NULL,'11-24-2021 ','41b5a24d-5925-45e5-92d8-9e6ffa'),
(22,'i amwe','no','i_amwe_-_no.mp3','3.13 mb','01:00',NULL,'11-25-2021 ','41b5a24d-5925-45e5-92d8-9e6ffa');

/*Table structure for table `role` */

DROP TABLE IF EXISTS `role`;

CREATE TABLE `role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

/*Data for the table `role` */

insert  into `role`(`id`,`name`) values 
(1,'admin'),
(2,'user'),
(3,NULL);

/*Table structure for table `users` */

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `uid` varchar(30) NOT NULL,
  `frist_name` varchar(20) DEFAULT NULL,
  `last_name` varchar(20) DEFAULT NULL,
  `data_regist` varchar(50) DEFAULT NULL,
  `last_visit` varchar(50) NOT NULL DEFAULT '00-00-0000 00:00',
  `login` varchar(20) DEFAULT NULL,
  `password` varchar(40) DEFAULT NULL,
  `token` varchar(100) NOT NULL,
  `id_role` int(11) DEFAULT NULL,
  `create_at` varchar(50) DEFAULT NULL,
  `update_at` varchar(50) NOT NULL,
  UNIQUE KEY `id` (`uid`),
  KEY `id_role` (`id_role`),
  CONSTRAINT `users_ibfk_1` FOREIGN KEY (`id_role`) REFERENCES `role` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `users` */

insert  into `users`(`uid`,`frist_name`,`last_name`,`data_regist`,`last_visit`,`login`,`password`,`token`,`id_role`,`create_at`,`update_at`) values 
('32418eb3-02e1-4b85-8770-dedc75','Makhsudov','Mukhammed','09-06-2021','27-01-2022 12:50:28','Alone','202cb962ac59075b964b07152d234b70','d83f56dfad0fa1806cbbd861aa8a79ededb989931548fea07fbe7af64096c0da',1,'09-06-2021','27-01-2022 12:50:28'),
('41b5a24d-5925-45e5-92d8-9e6ff1','USER2','USER1',NULL,'04-12-2021 20:56:32','USER1','3b712de48137572f3849aabd5666a4e3','d83f56dfad0fa1806cbbd861aa8a79ededb989931548fea07fbe7af64096c0d',2,NULL,'04-12-2021 20:56:32');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
