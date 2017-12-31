# tsdump
[![Build Status](https://travis-ci.org/voidint/tsdump.svg?branch=master)](https://travis-ci.org/voidint/tsdump)
[![Coverage Status](https://coveralls.io/repos/github/voidint/tsdump/badge.svg?branch=master)](https://coveralls.io/github/voidint/tsdump?branch=master)

## 目录
- [特性](#特性)
- [安装](#安装)
- [基本使用](#基本使用)
- [Changelog](#changelog)

## 特性
- 支持将数据库(当前仅支持`MySQL`)及其表结构的元数据以`text`、`markdown`、`json`、`csv`形式输出。


## 安装
```shell
$ go get -u github.com/voidint/tsdump
```

## 基本使用
- 全局选项
```shell
GLOBAL OPTIONS:
   -D, --debug               enable debug mode
   -H value, --host value    connect to host (default: "127.0.0.1")
   -P value, --port value    port number to use for connection (default: 3306)
   -u value, --user value    user for login if not current user (default: "voidint")
   -V value, --viewer value  output viewer. Optional values: txt|csv|json|md (default: "txt")
   -o value, --output value  write to a file, instead of STDOUT
   --help, -h                show help
   --version, -v             print the version
```

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

- 将目标数据库及其所有表的表结构数据以表格形式输出到console
    ```shell
    $ tsdump -H 127.0.0.1 -P 3307 -u root mydb
    Enter Password:
    |----------|---------------|--------------------|
    | DATABASE | CHARACTER SET |     COLLATION      |
    |----------|---------------|--------------------|
    | mydb     | utf8mb4       | utf8mb4_general_ci |
    |----------|---------------|--------------------|

    TABLE:	student	学生信息表
    |----------|--------------|----------|-----|---------|---------------|--------------------|----------|
    |  COLUMN  |  DATA TYPE   | NULLABLE | KEY | DEFAULT | CHARACTER SET |     COLLATION      | COMMENT  |
    |----------|--------------|----------|-----|---------|---------------|--------------------|----------|
    | sno      | char(8)      | NO       | PRI |         | utf8mb4       | utf8mb4_general_ci | 学号     |
    | sname    | varchar(255) | NO       |     |         | utf8mb4       | utf8mb4_general_ci | 姓名     |
    | gender   | char(2)      | YES      |     |         | utf8mb4       | utf8mb4_general_ci | 性别     |
    | native   | char(20)     | YES      |     |         | utf8mb4       | utf8mb4_general_ci | 籍贯     |
    | birthday | datetime     | YES      |     |         |               |                    | 出生日期 |
    | dno      | char(6)      | YES      |     |         | utf8mb4       | utf8mb4_general_ci | 所在院系 |
    | spno     | char(8)      | YES      |     |         | utf8mb4       | utf8mb4_general_ci | 专业代码 |
    | classno  | char(4)      | YES      |     |         | utf8mb4       | utf8mb4_general_ci | 班级号   |
    | entime   | date         | YES      |     |         |               |                    | 入校时间 |
    | home     | varchar(40)  | YES      |     |         | utf8mb4       | utf8mb4_general_ci | 家庭住址 |
    | tell     | varchar(40)  | YES      |     |         | utf8mb4       | utf8mb4_general_ci | 联系电话 |
    |----------|--------------|----------|-----|---------|---------------|--------------------|----------|
    ```

- 将目标数据库下目标表的表结构数据输出到markdown文件
    ```shell
    $ tsdump -H 127.0.0.1 -P 3307 -u root -V md -o ./student.md mydb student
    ```

    output: 

    | DATABASE | CHARACTER SET |     COLLATION      |
    |----------|---------------|--------------------|
    | mydb     | utf8mb4       | utf8mb4_general_ci |

    ### `student`
    学生信息表

    |  COLUMN  |  DATA TYPE   | NULLABLE | KEY | DEFAULT | CHARACTER SET |     COLLATION      | COMMENT  |
    |----------|--------------|----------|-----|---------|---------------|--------------------|----------|
    | sno      | char(8)      | NO       | PRI |         | utf8mb4       | utf8mb4_general_ci | 学号     |
    | sname    | varchar(255) | NO       |     |         | utf8mb4       | utf8mb4_general_ci | 姓名     |
    | gender   | char(2)      | YES      |     |         | utf8mb4       | utf8mb4_general_ci | 性别     |
    | native   | char(20)     | YES      |     |         | utf8mb4       | utf8mb4_general_ci | 籍贯     |
    | birthday | datetime     | YES      |     |         |               |                    | 出生日期 |
    | dno      | char(6)      | YES      |     |         | utf8mb4       | utf8mb4_general_ci | 所在院系 |
    | spno     | char(8)      | YES      |     |         | utf8mb4       | utf8mb4_general_ci | 专业代码 |
    | classno  | char(4)      | YES      |     |         | utf8mb4       | utf8mb4_general_ci | 班级号   |
    | entime   | date         | YES      |     |         |               |                    | 入校时间 |
    | home     | varchar(40)  | YES      |     |         | utf8mb4       | utf8mb4_general_ci | 家庭住址 |
    | tell     | varchar(40)  | YES      |     |         | utf8mb4       | utf8mb4_general_ci | 联系电话 |

- 将用户权限范围内数据库及其表结构数据输出到csv文件
    ```shell
    $ tsdump -H 127.0.0.1 -P 3307 -u root -V csv -o ./mydb.csv
    ```

- 将目标数据库及其所有表的表结构数据输出到JSON文件
    ```shell
    $ tsdump -H 127.0.0.1 -P 3307 -u root -V json -o mydb.json mydb
    ```

## Changelog
### 0.1.0 - 2017/12/31
- 支持以`csv`视图方式导出表结构数据。[#1](https://github.com/voidint/tsdump/issues/1)
- 支持以`markdown`视图方式导出表结构数据。[#2](https://github.com/voidint/tsdump/issues/2)
- 支持以`text`视图方式导出表结构数据。[#3](https://github.com/voidint/tsdump/issues/3)
- 支持以`json`视图方式导出表结构数据。[#4](https://github.com/voidint/tsdump/issues/4)

### 0.2.0 - 
- 支持从stdin console中读取数据库登录密码。[#5](https://github.com/voidint/tsdump/issues/5)
- 支持命令行参数指定目标数据库和表。[#12](https://github.com/voidint/tsdump/issues/12)