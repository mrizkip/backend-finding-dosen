-- MySQL dump 10.13  Distrib 8.0.12, for Win64 (x86_64)
--
-- Host: localhost    Database: findingdosen
-- ------------------------------------------------------
-- Server version	5.7.20-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8mb4 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `findingdosen`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `findingdosen` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `findingdosen`;

--
-- Table structure for table `access_point`
--

DROP TABLE IF EXISTS `access_point`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `access_point` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `bssid` varchar(45) NOT NULL,
  `ssid` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `bssid_UNIQUE` (`bssid`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `access_point`
--

LOCK TABLES `access_point` WRITE;
/*!40000 ALTER TABLE `access_point` DISABLE KEYS */;
INSERT INTO `access_point` VALUES (1,'a0:3d:6f:85:c0:c1','FILKOM.X'),(2,'a0:3d:6f:85:c3:01','FILKOM.X'),(3,'a0:e0:af:57:26:81','FILKOM.X'),(4,'a0:3d:6f:8b:11:61','FILKOM.X'),(5,'a0:3d:6f:8f:96:01','FILKOM.X'),(6,'a0:3d:6f:8b:0f:21','FILKOM.X'),(7,'a0:3d:6f:8b:0b:c1','FILKOM.X'),(8,'a0:3d:6f:89:22:c1','FILKOM.X'),(9,'a0:3d:6f:89:22:e1','FILKOM.X'),(10,'a0:3d:6f:8b:0d:81','FILKOM.X'),(11,'a0:e0:af:9a:12:01','FILKOM.X'),(12,'a0:e0:af:99:e4:e1','FILKOM.X'),(13,'a0:3d:6f:8b:0d:c1','FILKOM.X'),(14,'a0:3d:6f:89:20:e1','FILKOM.X'),(15,'a0:3d:6f:8f:52:81','FILKOM.X'),(16,'a0:3d:6f:85:c1:41','FILKOM.X'),(17,'a0:3d:6f:8f:92:21','FILKOM.X'),(18,'a0:3d:6f:5c:8a:a1','FILKOM.X'),(19,'a0:3d:6f:8f:94:61','FILKOM.X'),(20,'a0:3d:6f:76:b8:c1','FILKOM.X'),(21,'a0:3d:6f:85:c1:81','FILKOM.X');
/*!40000 ALTER TABLE `access_point` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `data_rgb`
--

DROP TABLE IF EXISTS `data_rgb`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `data_rgb` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `gedung` varchar(45) DEFAULT NULL,
  `ruang` varchar(45) DEFAULT NULL,
  `level_r` int(11) DEFAULT NULL,
  `level_g` int(11) DEFAULT NULL,
  `level_b` int(11) DEFAULT NULL,
  `ap1` int(11) DEFAULT NULL,
  `ap2` int(11) DEFAULT NULL,
  `ap3` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_ap1_idx` (`ap1`),
  KEY `fk_ap2_idx` (`ap2`),
  KEY `fk_ap3_idx` (`ap3`),
  CONSTRAINT `fk_ap1` FOREIGN KEY (`ap1`) REFERENCES `access_point` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_ap2` FOREIGN KEY (`ap2`) REFERENCES `access_point` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_ap3` FOREIGN KEY (`ap3`) REFERENCES `access_point` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=273 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `data_rgb`
--

LOCK TABLES `data_rgb` WRITE;
/*!40000 ALTER TABLE `data_rgb` DISABLE KEYS */;
INSERT INTO `data_rgb` VALUES (1,'A1','A1.1',57,89,54,1,2,3),(2,'A1','A1.1.1',61,54,26,1,2,3),(3,'A1','A1.1.2',51,61,38,1,2,3),(4,'A1','A1.1.3',29,61,45,1,2,3),(5,'A1','A1.2 A',80,16,19,1,2,3),(6,'A1','A1.2 B',86,16,32,1,2,3),(7,'A1','A1.3 A',96,255,255,1,2,3),(8,'A1','A1.3 B',80,255,207,1,2,3),(9,'A1','A1.3 C',99,255,255,1,2,3),(10,'A1','A1.3 D',105,255,255,1,2,3),(11,'A1','A1.4',112,255,255,1,2,3),(12,'A1','A1.4.1',105,255,255,1,2,3),(13,'A1','A1.4.2',86,255,255,1,2,3),(14,'A1','A1.4.3',64,255,255,1,2,3),(15,'A1','A1.5 A',80,102,255,1,2,3),(16,'A1','A1.5 B',77,131,255,1,2,3),(17,'A1','A1.6',48,67,77,1,2,3),(18,'A1','A1.6.1',41,51,26,1,2,3),(19,'A1','A1.6.2',51,61,19,1,2,3),(20,'A1','A1.6.3',57,64,140,1,2,3),(21,'A1','A1.7 A',140,89,57,1,2,3),(22,'A1','A1.7 B',204,67,51,1,2,3),(23,'A1','A1.8 A',137,108,86,1,2,3),(24,'A1','A1.8 B',102,48,89,1,2,3),(25,'A1','A1.9 A',255,0,105,1,2,3),(26,'A1','A1.9 B',255,64,67,1,2,3),(27,'A1','A1.9 C',255,57,64,1,2,3),(28,'A1','A1.9 D',255,137,108,1,2,3),(29,'A1','A1.10 A',255,0,64,1,2,3),(30,'A1','A1.10 B',255,57,73,1,2,3),(31,'A1','A1.10 C',255,99,51,1,2,3),(32,'A1','A1.10 D',255,172,57,1,2,3),(33,'A1','LA1.4',112,22,112,1,2,3),(34,'A1','LA1.5',147,48,96,1,2,3),(35,'A1','LA1.6 A',115,96,153,1,2,3),(36,'A1','LA1.6 B',89,77,80,1,2,3),(37,'A1','LA1.8',92,86,147,1,2,3),(38,'A1','LA1.10',57,41,112,1,2,3),(39,'A1','LA1T A',137,102,45,1,2,3),(40,'A1','LA1T B',124,102,108,1,2,3),(41,'A2','A2.16 A',98,158,179,4,5,6),(42,'A2','A2.16 B',64,77,207,4,5,6),(43,'A2','A2.16 C',86,22,10,4,5,6),(44,'A2','A2.16 D',45,16,230,4,5,6),(45,'A2','A2.17 A',83,29,134,4,5,6),(46,'A2','A2.17 B',67,45,10,4,5,6),(47,'A2','A2.17 C',118,48,29,4,5,6),(48,'A2','A2.17 D',112,57,89,4,5,6),(49,'A2','A2.18 A',70,86,92,4,5,6),(50,'A2','A2.18 B',57,13,45,4,5,6),(51,'A2','A2.18 C',61,80,70,4,5,6),(52,'A2','A2.18 D',32,38,48,4,5,6),(53,'A2','A2.19 A',137,121,108,4,5,6),(54,'A2','A2.19 B',10,16,96,4,5,6),(55,'A2','A2.19 C',38,51,92,4,5,6),(56,'A2','A2.19 D',13,26,137,4,5,6),(57,'A2','A2.20 A',124,32,61,4,5,6),(58,'A2','A2.20 B',230,73,80,4,5,6),(59,'A2','A2.20 C',0,3,51,4,5,6),(60,'A2','A2.20 D',153,166,51,4,5,6),(61,'A2','A2.21',26,16,38,4,5,6),(62,'A2','A2.22 A',48,64,255,4,5,6),(63,'A2','A2.22 B',29,16,128,4,5,6),(64,'A2','A2.22 C',80,105,233,4,5,6),(65,'A2','A2.22 D',61,115,233,4,5,6),(66,'A2','A2.23 A',89,156,255,4,5,6),(67,'A2','A2.23 B',102,41,22,4,5,6),(68,'A2','A2.24 A',35,96,48,4,5,6),(69,'A2','A2.24 B',29,92,22,4,5,6),(70,'A2','A2.24 C',45,115,29,4,5,6),(71,'A2','A2.24 D',16,124,32,4,5,6),(72,'A2','A2.25 A',255,22,26,4,5,6),(73,'A2','A2.25 B',38,89,64,4,5,6),(74,'A2','A2.25 C',150,45,29,4,5,6),(75,'A2','A2.25 D',57,54,22,4,5,6),(76,'A2','LA2.17',96,38,10,4,5,6),(77,'A2','LA2.18',147,131,115,4,5,6),(78,'A2','LA2.19',102,92,143,4,5,6),(79,'A2','LA2T A',64,83,96,4,5,6),(80,'A2','LA2T B',77,67,57,4,5,6),(81,'E1','E1.1 A',255,255,130,7,8,9),(82,'E1','E1.1 B',255,255,130,7,8,9),(83,'E1','E1.1 C',255,255,147,7,8,9),(84,'E1','E1.1 D',255,255,124,7,8,9),(85,'E1','E1.2 A',255,255,98,7,8,9),(86,'E1','E1.2 B',255,204,66,7,8,9),(87,'E1','E1.2 C',255,255,15,7,8,9),(88,'E1','E1.2 D',255,255,38,7,8,9),(89,'E1','E1.3 A',175,73,79,7,8,9),(90,'E1','E1.3 B',153,178,86,7,8,9),(91,'E1','E1.3 C',51,44,82,7,8,9),(92,'E1','E1.3 D',153,255,82,7,8,9),(93,'E1','E1.4 A',172,178,255,7,8,9),(94,'E1','E1.4 B',82,95,143,7,8,9),(95,'E1','E1.4 C',12,54,111,7,8,9),(96,'E1','E1.4 D',31,86,79,7,8,9),(97,'E1','E1.5 A',140,165,255,7,8,9),(98,'E1','E1.5 B',156,121,255,7,8,9),(99,'E1','E1.5 C',35,25,255,7,8,9),(100,'E1','E1.5 D',60,12,255,7,8,9),(101,'E1','E1.6 A',121,51,255,7,8,9),(102,'E1','E1.6 B',44,44,255,7,8,9),(103,'E1','E1.7 A',89,255,255,7,8,9),(104,'E1','E1.7 B',98,188,255,7,8,9),(105,'E1','E1.7 C',63,255,255,7,8,9),(106,'E1','E1.7 D',89,255,255,7,8,9),(107,'E1','E1.8 A',89,178,255,7,8,9),(108,'E1','E1.8 B',89,204,255,7,8,9),(109,'E1','E1.9 A',156,22,38,7,8,9),(110,'E1','E1.9 B',66,41,121,7,8,9),(111,'E1','E1.10 A',165,95,255,7,8,9),(112,'E1','E1.10 B',25,28,255,7,8,9),(113,'E1','LE1.1',82,207,82,7,8,9),(114,'E1','LE1.3',70,22,54,7,8,9),(115,'E1','LE1.4',86,63,9,7,8,9),(116,'E1','LE1.5',102,98,105,7,8,9),(117,'E1','LE1.7',143,191,229,7,8,9),(118,'E1','LE1T A',63,54,25,7,8,9),(119,'E1','LE1T B',92,98,66,7,8,9),(120,'E2','E2.1 A',255,63,117,10,11,12),(121,'E2','E2.1 B',229,117,184,10,11,12),(122,'E2','E2.1 C',200,95,98,10,11,12),(123,'E2','E2.1 D',22,76,54,10,11,12),(124,'E2','E2.2 A',184,175,127,10,11,12),(125,'E2','E2.2 B',210,197,117,10,11,12),(126,'E2','E2.2 C',255,153,127,10,11,12),(127,'E2','E2.2 D',255,146,124,10,11,12),(128,'E2','E2.3 A',159,105,162,10,11,12),(129,'E2','E2.3 B',130,92,207,10,11,12),(130,'E2','E2.3 C',181,124,70,10,11,12),(131,'E2','E2.3 D',28,98,19,10,11,12),(132,'E2','E2.4 A',44,38,0,10,11,12),(133,'E2','E2.4 B',25,38,0,10,11,12),(134,'E2','E2.4 C',35,35,0,10,11,12),(135,'E2','E2.4 D',38,28,0,10,11,12),(136,'E2','E2.5 A',38,63,15,10,11,12),(137,'E2','E2.5 B',31,57,15,10,11,12),(138,'E2','E2.5 C',70,79,210,10,11,12),(139,'E2','E2.5 D',118,149,143,10,11,12),(140,'E2','E2.6 A',76,31,178,10,11,12),(141,'E2','E2.6 B',86,143,184,10,11,12),(142,'E2','E2.7 A',144,102,204,10,11,12),(143,'E2','E2.7 B',63,73,149,10,11,12),(144,'E2','E2.7 C',140,255,229,10,11,12),(145,'E2','E2.7 D',111,60,255,10,11,12),(146,'E2','E2.8 A',54,255,255,10,11,12),(147,'E2','E2.8 B',89,178,204,10,11,12),(148,'E2','E2.8 C',86,229,255,10,11,12),(149,'E2','E2.8 D',76,255,255,10,11,12),(150,'E2','E2.9 A',63,0,255,10,11,12),(151,'E2','E2.9 B',82,0,255,10,11,12),(152,'E2','E2.9 C',70,9,255,10,11,12),(153,'E2','E2.9 D',76,35,255,10,11,12),(154,'E2','LE2.1',51,57,70,10,11,12),(155,'E2','LE2.3',38,73,44,10,11,12),(156,'E2','LE2.5',76,41,54,10,11,12),(157,'E2','LE2.7',137,156,95,10,11,12),(158,'E2','LE2.9',121,162,255,10,11,12),(159,'E2','LE2T A',51,156,165,10,11,12),(160,'E2','LE2T B',76,73,89,10,11,12),(161,'F2','F2.1 A',232,133,95,13,14,15),(162,'F2','F2.1 B',57,232,95,13,14,15),(163,'F2','F2.1 C',149,255,102,13,14,15),(164,'F2','F2.1 D',47,159,105,13,14,15),(165,'F2','F2.2 A',255,191,124,13,14,15),(166,'F2','F2.2 B',255,255,121,13,14,15),(167,'F2','F2.2 C',184,232,133,13,14,15),(168,'F2','F2.2 D',162,255,105,13,14,15),(169,'F2','F2.4 A',232,111,188,13,14,15),(170,'F2','F2.4 B',255,92,15,13,14,15),(171,'F2','F2.4 C',255,86,95,13,14,15),(172,'F2','F2.4 D',15,79,47,13,14,15),(173,'F2','F2.5 A',70,130,79,13,14,15),(174,'F2','F2.5 B',65,125,70,13,14,15),(175,'F2','F2.5 C',55,120,81,13,14,15),(176,'F2','F2.5 D',59,128,82,13,14,15),(177,'F2','F2.6 A',54,86,82,13,14,15),(178,'F2','F2.6 B',41,79,105,13,14,15),(179,'F2','F2.6 C',159,130,28,13,14,15),(180,'F2','F2.6 D',50,80,85,13,14,15),(181,'F2','F2.8 A',108,51,229,13,14,15),(182,'F2','F2.8 B',95,12,19,13,14,15),(183,'F2','F2.8 C',70,81,153,13,14,15),(184,'F2','F2.8 D',86,9,255,13,14,15),(185,'F2','F2.9 A',137,200,70,13,14,15),(186,'F2','F2.9 B',140,255,79,13,14,15),(187,'F2','F2.9 C',117,207,229,13,14,15),(188,'F2','F2.9 D',121,105,12,13,14,15),(189,'F2','F2.10',255,76,255,13,14,15),(190,'F2','F2 Hall A',79,73,66,13,14,15),(191,'F2','F2 Hall B',127,92,105,13,14,15),(192,'F2','F2 Hall C',124,153,38,13,14,15),(193,'F2','F2 Hall D',89,153,156,13,14,15),(194,'F2','F2 Hall E',108,191,127,13,14,15),(195,'F2','F2 Hall F',255,89,105,13,14,15),(196,'F2','F2 Hall G',86,153,82,13,14,15),(197,'F2','F2 Hall H',127,121,95,13,14,15),(198,'F2','F2 Hall I',188,89,146,13,14,15),(199,'F2','F2 Hall J',76,19,38,13,14,15),(200,'F2','F2 Hall K',255,162,197,13,14,15),(201,'F3','F3.1 A',255,31,51,16,17,18),(202,'F3','F3.1 B',255,31,51,16,17,18),(203,'F3','F3.1 C',255,38,4,16,17,18),(204,'F3','F3.1 D',255,98,44,16,17,18),(205,'F3','F3.2 A',25,255,210,16,17,18),(206,'F3','F3.2 B',25,124,11,16,17,18),(207,'F3','F3.3 A',255,86,89,16,17,18),(208,'F3','F3.3 B',255,204,63,16,17,18),(209,'F3','F3.3 C',255,54,102,16,17,18),(210,'F3','F3.3 D',255,38,60,16,17,18),(211,'F3','F3.4 A',188,188,82,16,17,18),(212,'F3','F3.4 B',207,255,114,16,17,18),(213,'F3','F3.5 A',255,255,184,16,17,18),(214,'F3','F3.5 B',111,255,66,16,17,18),(215,'F3','F3.5 C',175,255,66,16,17,18),(216,'F3','F3.6 A',137,117,95,16,17,18),(217,'F3','F3.6 B',137,117,95,16,17,18),(218,'F3','F3.7 A',213,255,255,16,17,18),(219,'F3','F3.7 B',108,255,255,16,17,18),(220,'F3','F3.7 C',92,255,79,16,17,18),(221,'F3','F3.8 A',162,143,168,16,17,18),(222,'F3','F3.8 B',92,105,70,16,17,18),(223,'F3','F3.9 A',6,38,255,16,17,18),(224,'F3','F3.9 B',15,60,76,16,17,18),(225,'F3','F3.9 C',0,44,255,16,17,18),(226,'F3','F3.9 D',15,54,73,16,17,18),(227,'F3','F3.10 A',95,146,255,16,17,18),(228,'F3','F3.10 B',63,60,51,16,17,18),(229,'F3','F3.11 A',10,70,255,16,17,18),(230,'F3','F3.11 B',25,82,255,16,17,18),(231,'F3','F3.11 C',9,81,255,16,17,18),(232,'F3','F3.11 D',20,77,255,16,17,18),(233,'F3','F3.12 A',137,143,156,16,17,18),(234,'F3','F3.12 B',22,95,117,16,17,18),(235,'F3','F3 Hall A',63,98,229,16,17,18),(236,'F3','F3 Hall B',89,57,54,16,17,18),(237,'F3','F3 Hall C',79,184,28,16,17,18),(238,'F3','F3 Hall D',89,89,255,16,17,18),(239,'F3','F3 Hall E',121,229,66,16,17,18),(240,'F3','F3 Hall F',82,255,178,16,17,18),(241,'F3','F3 Hall G',178,255,76,16,17,18),(242,'F3','F3 Hall H',41,156,133,16,17,18),(243,'F3','F3 Hall I',35,76,92,16,17,18),(244,'F3','F3 Hall J',181,54,86,16,17,18),(245,'F3','F3 Hall K',255,111,63,16,17,18),(246,'F3','F3 Hall L',204,117,130,16,17,18),(247,'F4','F4.1',60,35,92,19,20,21),(248,'F4','F4.2 A',172,172,95,19,20,21),(249,'F4','F4.2 B',79,41,89,19,20,21),(250,'F4','F4.3 A',57,47,73,19,20,21),(251,'F4','F4.3 B',63,70,95,19,20,21),(252,'F4','F4.4 A',63,70,66,19,20,21),(253,'F4','F4.4 B',60,66,70,19,20,21),(254,'F4','F4.5 A',70,57,60,19,20,21),(255,'F4','F4.5 B',63,98,63,19,20,21),(256,'F4','F4.6 A',79,51,66,19,20,21),(257,'F4','F4.6 B',54,66,51,19,20,21),(258,'F4','F4.7 A',114,143,86,19,20,21),(259,'F4','F4.7 B',95,54,79,19,20,21),(260,'F4','F4.8',70,127,54,19,20,21),(261,'F4','F4 Hall A',130,21,77,19,20,21),(262,'F4','F4 Hall B',149,31,70,19,20,21),(263,'F4','F4 Hall C',124,76,73,19,20,21),(264,'F4','F4 Hall D',73,165,168,19,20,21),(265,'F4','F4 Hall E',57,105,44,19,20,21),(266,'F4','F4 Hall F',117,108,102,19,20,21),(267,'F4','F4 Hall G',210,239,223,19,20,21),(268,'F4','F4 Hall H',47,255,102,19,20,21),(269,'F4','F4 Hall I',45,105,123,19,20,21),(270,'F4','F4 Hall J',70,92,146,19,20,21),(271,'F4','F4 Hall K',44,20,102,19,20,21),(272,'F4','F4 Hall L',63,28,60,19,20,21);
/*!40000 ALTER TABLE `data_rgb` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `status`
--

DROP TABLE IF EXISTS `status`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `status` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `desc_status` varchar(255) DEFAULT NULL,
  `posisi` varchar(255) DEFAULT NULL,
  `ket_status` varchar(255) DEFAULT NULL,
  `last_update` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UserIDIndex` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `status`
--

LOCK TABLES `status` WRITE;
/*!40000 ALTER TABLE `status` DISABLE KEYS */;
INSERT INTO `status` VALUES (1,3,'Tidak Aktif','','','','2018-04-08 09:19:00'),(2,4,'Tidak Aktif','','A1.8 A','','2018-07-18 12:45:00'),(3,5,'Tidak Aktif','','','','2018-04-08 09:19:00'),(4,6,'Tidak Aktif','','','','2018-04-08 09:19:00'),(5,7,'Aktif','Di Ruangan','A1.10','A1.10','2018-04-08 09:19:00'),(6,9,'Tidak Aktif','','','','2018-04-26 05:06:00');
/*!40000 ALTER TABLE `status` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `nama` varchar(255) DEFAULT NULL,
  `jenis_identitas` varchar(255) DEFAULT NULL,
  `no_identitas` varchar(255) DEFAULT NULL,
  `no_telpon` varchar(255) DEFAULT NULL,
  `role` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `EmailIdentitasIndex` (`email`,`no_identitas`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'test@mail.com','$2a$10$y.SZN7UDWSv6IKxWmwkhdO42mopKemOUSibnleUTw0buTEv93o4YK','Test Account','NIM','145110101111111','081234567890','mahasiswa'),(2,'test1@mail.com','$2a$10$B6Vgsh3Ig3UuEGT1kUnbx.mkqLJQS3qdqWErlvMcoHVWOH1XrL1Du','Test Account 1','NIM','145110101111111','081234567891','mahasiswa'),(3,'dosen@mail.com','$2a$10$iGU/Xk6H0G4S9/uChJgKb..Tr9ivNMRaYcbS7L3oRdDU77oYHzK9G','Test Account Dosen','NIP','8912830812387192837','081234567892','dosen'),(4,'dosen1@mail.com','$2a$10$87N3s163VDhuJRbpQAPD1ekk5NHnc1WVNmpupnRsF9Dk3HjL8StPW','Test Account Dosen 1','NIP','8912830812387192838','081234567893','dosen'),(5,'dosen2@mail.com','$2a$10$mkGsAfFrWpS1/xbl6z7Yb.DHZot3aV2cHs.NF/jKHufZIsTc1sJOK','Test Account Dosen 2','NIP','8912830812387192839','081234567894','dosen'),(6,'dosen3@mail.com','$2a$10$chWtqcnpTp5C0qYXKEgW8u9WC7nFQaCXRdaH4Up0WjiCGRke.pwFG','Test Account Dosen 3','NIP','8912830812387192840','081234567895','dosen'),(7,'dosen4@mail.com','$2a$10$Lmpo55mkzefIN9SWlNUOhuY6txxrqd8DtWS9FgzACD5tKh58q9oAe','Test Account Dosen 4','NIP','8912830812387192841','081234567896','dosen'),(8,'dosen@mail.com','$2a$10$W41nyZ.2b7Ke0KbSTabNFug/Q7QLn59h0zkU2upuhRyiLyXHCGaQ6','Dosen 5','NIP','123445161231412','085698745214','dosen'),(9,'dosen5@mail.com','$2a$10$Am34Kj1B7Zi2Wt.W85cyxu62FIJCl4Zow212rWiucxfsHbvIA41tG','Dosen 5','NIP','123445161231412','085698745214','dosen');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-11-04 15:02:00
