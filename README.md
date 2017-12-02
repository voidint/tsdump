# tsdump
[![Build Status](https://travis-ci.org/voidint/tsdump.svg?branch=master)](https://travis-ci.org/voidint/tsdump)

## 目录
- [特性](#特性)
- [安装](#安装)
- [基本使用](#基本使用)
- [Changelog](#Changelog)

## 特性
- 支持将数据库(当前仅支持`MySQL`)及其表结构的元数据以`text`、`markdown`、`json`、`csv`形式输出。


## 安装
```shell
$ go get -u github.com/voidint/tsdump
```

## 基本使用
- 使用`root`用户创建一个名为`mydb`的数据库实例，以及一张`student`的表。
    ```SQL
    CREATE DATABASE IF NOT EXISTS `mydb` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

    USE `mydb`;

    CREATE TABLE `student` (
    `sno` char(8) NOT NULL COMMENT '学号',
    `sname` varchar(255) NOT NULL COMMENT '姓名',
    `gender` char(2) DEFAULT NULL COMMENT '性别',
    `native` char(20) DEFAULT NULL COMMENT '籍贯',
    `birthday` datetime DEFAULT NULL COMMENT '出生日期',
    `dno` char(6) DEFAULT NULL COMMENT '所在院系',
    `spno` char(8) DEFAULT NULL COMMENT '专业代码',
    `classno` char(4) DEFAULT NULL COMMENT '班级号',
    `entime` date DEFAULT NULL COMMENT '入校时间',
    `home` varchar(40) DEFAULT NULL COMMENT '家庭住址',
    `tell` varchar(40) DEFAULT NULL COMMENT '联系电话',
    PRIMARY KEY (`sno`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='学生信息表';
    ```

- 将数据库及其表结构数据以表格形式输出到console
    ```shell
    $ tsdump -H 127.0.0.1 -P 3307 -u root -p "mypassword" --db mydb
    |----------|---------------|--------------------|
    | DATABASE | CHARACTER SET |     COLLATION      |
    |----------|---------------|--------------------|
    | mydb     | utf8mb4       | utf8mb4_general_ci |
    |----------|---------------|--------------------|

    TABLE:	student	学生信息表
    |----------|----------|--------------|---------------|--------------------|----------|
    |  COLUMN  | NULLABLE |  DATA TYPE   | CHARACTER SET |     COLLATION      | COMMENT  |
    |----------|----------|--------------|---------------|--------------------|----------|
    | sno      | NO       | char(8)      | utf8mb4       | utf8mb4_general_ci | 学号     |
    | sname    | NO       | varchar(255) | utf8mb4       | utf8mb4_general_ci | 姓名     |
    | gender   | YES      | char(2)      | utf8mb4       | utf8mb4_general_ci | 性别     |
    | native   | YES      | char(20)     | utf8mb4       | utf8mb4_general_ci | 籍贯     |
    | birthday | YES      | datetime     |               |                    | 出生日期 |
    | dno      | YES      | char(6)      | utf8mb4       | utf8mb4_general_ci | 所在院系 |
    | spno     | YES      | char(8)      | utf8mb4       | utf8mb4_general_ci | 专业代码 |
    | classno  | YES      | char(4)      | utf8mb4       | utf8mb4_general_ci | 班级号   |
    | entime   | YES      | date         |               |                    | 入校时间 |
    | home     | YES      | varchar(40)  | utf8mb4       | utf8mb4_general_ci | 家庭住址 |
    | tell     | YES      | varchar(40)  | utf8mb4       | utf8mb4_general_ci | 联系电话 |
    |----------|----------|--------------|---------------|--------------------|----------|
    ```

- 将数据库及其表结构数据输出到markdown文件
    ```shell
    $ tsdump -H 127.0.0.1 -P 3307 -u root -p "mypassword" --db mydb -V md > ./mydb.md
    ```

    output: 
    ### `student`
    学生信息表

    |  COLUMN  | NULLABLE |  DATA TYPE   | CHARACTER SET |     COLLATION      | COMMENT  |
    |----------|----------|--------------|---------------|--------------------|----------|
    | sno      | NO       | char(8)      | utf8mb4       | utf8mb4_general_ci | 学号     |
    | sname    | NO       | varchar(255) | utf8mb4       | utf8mb4_general_ci | 姓名     |
    | gender   | YES      | char(2)      | utf8mb4       | utf8mb4_general_ci | 性别     |
    | native   | YES      | char(20)     | utf8mb4       | utf8mb4_general_ci | 籍贯     |
    | birthday | YES      | datetime     |               |                    | 出生日期 |
    | dno      | YES      | char(6)      | utf8mb4       | utf8mb4_general_ci | 所在院系 |
    | spno     | YES      | char(8)      | utf8mb4       | utf8mb4_general_ci | 专业代码 |
    | classno  | YES      | char(4)      | utf8mb4       | utf8mb4_general_ci | 班级号   |
    | entime   | YES      | date         |               |                    | 入校时间 |
    | home     | YES      | varchar(40)  | utf8mb4       | utf8mb4_general_ci | 家庭住址 |
    | tell     | YES      | varchar(40)  | utf8mb4       | utf8mb4_general_ci | 联系电话 |

- 将数据库及其表结构数据输出到csv文件
    ```shell
    $ tsdump -H 127.0.0.1 -P 3307 -u root -p "mypassword" --db mydb -V csv -o ./mydb.csv
    ```

- 将数据库及其表结构数据输出到JSON文件
    ```shell
    $ tsdump -H 127.0.0.1 -P 3307 -u root -p "mypassword" --db mydb -V json -o ./mydb.json
    ```

## Changelog