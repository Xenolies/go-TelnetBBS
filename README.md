# go-TelnetBBS
用Go语言写的Telnet BBS系统,使用了一部分Zinx框架的思路


 创建用户表
```mysql
create table users
(
    ID      int auto_increment
        primary key,
    level   int           not null,
    name    varchar(2000) not null,
    loginID varchar(2000) not null,
    pwd     varchar(2000) not null
);
```
