-- MySQL dump 10.13  Distrib 8.0.23, for Win64 (x86_64)
--
-- Host: localhost    Database: mydroid
-- ------------------------------------------------------
-- Server version	8.0.23

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `mydroid`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `mydroid` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `mydroid`;

--
-- Table structure for table `account`
--

DROP TABLE IF EXISTS `account`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `account` (
  `id` int NOT NULL AUTO_INCREMENT,
  `type` int NOT NULL,
  `date_joined` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `password` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `account`
--

LOCK TABLES `account` WRITE;
/*!40000 ALTER TABLE `account` DISABLE KEYS */;
INSERT INTO `account` VALUES (2,3,'2021-02-26 10:40:58','Ramin','123123'),(3,3,'2021-02-26 10:45:01','GameDevC00L','123'),(4,3,'2021-02-26 10:45:15','AppDevl33t','123'),(16,3,'2021-03-04 09:08:52','rmclassic','123123');
/*!40000 ALTER TABLE `account` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `app`
--

DROP TABLE IF EXISTS `app`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `app` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `publisher_id` int NOT NULL,
  `download_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `thumb_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `date_modified` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_pid` (`publisher_id`),
  CONSTRAINT `fk_pid` FOREIGN KEY (`publisher_id`) REFERENCES `account` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `app`
--

LOCK TABLES `app` WRITE;
/*!40000 ALTER TABLE `app` DISABLE KEYS */;
INSERT INTO `app` VALUES (2,'Divar','Buy and sell stuff in this app easily and with a click!',4,'','','2021-02-26 10:47:54'),(3,'Snapp','With snapp! you can get your cab at your service at no cost!',4,'','','2021-02-26 10:48:57'),(4,'AP','Get charged! everywhere you are, with whomever you are!',4,'','','2021-02-26 10:49:57'),(6,'Big Cars','Go into the road with weird big cars',3,'','','2021-02-26 10:52:42'),(7,'Clash Of Clans','A game for people who love to kill',3,'','','2021-02-26 10:54:15'),(8,'Vikings','kill, kill and kill',2,'','','2021-02-27 19:05:45'),(9,'Auto Cards','Run! with the best cars, specially pride 111',3,'','','2021-02-27 19:06:47'),(10,'Shepard','Sth about sheeps',3,'','','2021-02-27 19:07:35'),(11,'Zombie Hunter 2','Kill zombies with all your might, die hard',3,'','','2021-02-27 19:08:16'),(12,'Zapya','Share your pics with ppl fast',4,'','','2021-02-27 19:15:14'),(13,'telewebion','tv on web, web in your hands',4,'','','2021-02-27 19:15:37'),(14,'Digikala','Buy stuff more expensive than market',4,'','','2021-02-27 19:16:07'),(15,'Mellat mobile bank','mobile banking in your hands',4,'','','2021-02-27 19:16:47'),(17,'Aparat','Youtube in a nutshell',3,'','','2021-02-27 19:17:19'),(33,'Pou','Your virtual pet',3,'','','2021-03-03 14:53:22'),(34,'SFG','Gun game',3,'','','2021-03-03 14:53:50'),(35,'Racing game','Race forever',3,'','','2021-03-03 14:54:38'),(45,'dwa','wda',2,'','','2021-03-05 12:06:52');
/*!40000 ALTER TABLE `app` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `app_category`
--

DROP TABLE IF EXISTS `app_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `app_category` (
  `app_id` int NOT NULL,
  `category_id` int NOT NULL,
  PRIMARY KEY (`app_id`,`category_id`),
  KEY `fk_cat_id` (`category_id`),
  CONSTRAINT `fk_app_id` FOREIGN KEY (`app_id`) REFERENCES `app` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_cat_id` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `app_category`
--

LOCK TABLES `app_category` WRITE;
/*!40000 ALTER TABLE `app_category` DISABLE KEYS */;
INSERT INTO `app_category` VALUES (4,3),(14,3),(8,4),(33,4),(7,5),(10,5),(34,5),(9,6),(6,7),(11,7),(35,7),(2,8),(3,8),(12,8),(15,8),(45,8),(13,9),(17,9);
/*!40000 ALTER TABLE `app_category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `app_management`
--

DROP TABLE IF EXISTS `app_management`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `app_management` (
  `app_id` int NOT NULL,
  `user_id` int NOT NULL,
  KEY `app_id` (`app_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `app_management_ibfk_1` FOREIGN KEY (`app_id`) REFERENCES `app` (`id`) ON DELETE CASCADE,
  CONSTRAINT `app_management_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `account` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `app_management`
--

LOCK TABLES `app_management` WRITE;
/*!40000 ALTER TABLE `app_management` DISABLE KEYS */;
/*!40000 ALTER TABLE `app_management` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `category`
--

DROP TABLE IF EXISTS `category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `category` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `parent_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_catpid` (`parent_id`),
  CONSTRAINT `fk_catpid` FOREIGN KEY (`parent_id`) REFERENCES `category` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category`
--

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
INSERT INTO `category` VALUES (3,'Apps',NULL),(4,'Games',NULL),(5,'Action Games',4),(6,'AR Games',4),(7,'RPG Games',5),(8,'Utility Apps',3),(9,'Multimedia',3);
/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comment`
--

DROP TABLE IF EXISTS `comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comment` (
  `app_id` int NOT NULL,
  `user_id` int NOT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  KEY `app_id` (`app_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `comment_ibfk_1` FOREIGN KEY (`app_id`) REFERENCES `app` (`id`) ON DELETE CASCADE,
  CONSTRAINT `comment_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `account` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment`
--

LOCK TABLES `comment` WRITE;
/*!40000 ALTER TABLE `comment` DISABLE KEYS */;
INSERT INTO `comment` VALUES (3,2,'This is a great app!'),(3,3,'illegal app'),(3,3,'go drink your kafir'),(4,3,'always crashes'),(4,16,'میشود خیلی راحت با این نرم افزار گوشی را شارژ کرد\nسریع است!'),(8,16,'بازی خیلی هیجان انگیزی است'),(8,16,'نمیتوانم آن را ران کنم'),(33,16,'میتوان معتاد این بازی شد ');
/*!40000 ALTER TABLE `comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `download`
--

DROP TABLE IF EXISTS `download`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `download` (
  `app_id` int NOT NULL,
  `user_id` int NOT NULL,
  KEY `app_id` (`app_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `download_ibfk_1` FOREIGN KEY (`app_id`) REFERENCES `app` (`id`) ON DELETE CASCADE,
  CONSTRAINT `download_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `account` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `download`
--

LOCK TABLES `download` WRITE;
/*!40000 ALTER TABLE `download` DISABLE KEYS */;
INSERT INTO `download` VALUES (2,2),(7,4),(6,2),(4,2),(12,3),(13,3),(12,3),(17,3),(17,3),(17,3),(17,3),(17,3),(17,3),(15,3),(15,3),(13,3),(11,3),(11,3),(11,3),(11,3),(11,3),(10,3),(10,3),(10,3),(9,3),(9,3),(9,3),(9,3),(8,3),(8,3),(7,3),(7,3),(7,3),(7,3),(7,3),(7,3),(7,3),(7,3),(7,3),(6,3),(33,2),(17,3),(34,4),(33,2),(17,3),(34,4),(33,2),(17,3),(34,4),(33,2),(17,3),(34,4),(33,2),(17,3),(34,4),(33,2),(17,3),(34,4),(33,2),(17,3),(34,4),(3,2),(10,3),(35,4),(35,3),(35,2);
/*!40000 ALTER TABLE `download` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `message`
--

DROP TABLE IF EXISTS `message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `message` (
  `sender_id` int NOT NULL,
  `recv_id` int NOT NULL,
  PRIMARY KEY (`sender_id`,`recv_id`),
  KEY `fk_recv` (`recv_id`),
  CONSTRAINT `fk_recv` FOREIGN KEY (`recv_id`) REFERENCES `account` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_sender` FOREIGN KEY (`sender_id`) REFERENCES `account` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `message`
--

LOCK TABLES `message` WRITE;
/*!40000 ALTER TABLE `message` DISABLE KEYS */;
/*!40000 ALTER TABLE `message` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `token`
--

DROP TABLE IF EXISTS `token`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `token` (
  `uid` int NOT NULL,
  `token` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`uid`),
  CONSTRAINT `token_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `account` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `token`
--

LOCK TABLES `token` WRITE;
/*!40000 ALTER TABLE `token` DISABLE KEYS */;
INSERT INTO `token` VALUES (2,'RBLP8Q7GCMCDG34CTJU');
/*!40000 ALTER TABLE `token` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_management`
--

DROP TABLE IF EXISTS `user_management`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_management` (
  `manager_id` int NOT NULL,
  `user_id` int NOT NULL,
  PRIMARY KEY (`manager_id`,`user_id`),
  KEY `fk_uid` (`user_id`),
  CONSTRAINT `fk_mid` FOREIGN KEY (`manager_id`) REFERENCES `account` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_uid` FOREIGN KEY (`user_id`) REFERENCES `account` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_management`
--

LOCK TABLES `user_management` WRITE;
/*!40000 ALTER TABLE `user_management` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_management` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-03-05 15:55:36
