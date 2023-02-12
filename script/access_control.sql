CREATE DATABASE access_control;

CREATE TABLE `access_control`.`user` (
  `id` INT(11) NOT NULL,
  `uid` INT(11) NOT NULL,
  `role_id` INT(11) NULL,
  `create_time` VARCHAR(45) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uid_UNIQUE` (`uid` ASC));

CREATE TABLE `access_control`.`role` (
  `id` INT(11) NOT NULL,
  `name` VARCHAR(64) NULL,
  `status` TINYINT(2) NULL,
  `create_time` TIMESTAMP NULL,
  `update_time` TIMESTAMP NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `access_control`.`permission` (
  `id` INT(11) NOT NULL,
  `title` VARCHAR(64) NULL,
  `action` VARCHAR(64) NULL,
  `status` TINYINT(1) NULL,
  `create_time` TIMESTAMP NULL,
  `update_time` TIMESTAMP NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `access_control`.`role_permission` (
  `role_id` INT(11) NOT NULL,
  `permission_id` INT(11) NULL,
  `create_time` TIMESTAMP NULL,
  PRIMARY KEY (`role_id`));

CREATE TABLE `access_control`.`user_role` (
  `id` INT(11) NOT NULL,
  `uid` INT(11) NULL,
  `role_id` INT(11) NULL,
  `create_time` TIMESTAMP NULL,
  PRIMARY KEY (`id`));

