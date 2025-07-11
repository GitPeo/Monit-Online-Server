-- 创建数据库
CREATE DATABASE IF NOT EXISTS `monit` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

-- 切换到数据库
USE `monit`;

-- 创建表
create table todo
(
    id          bigint unsigned auto_increment comment 'id'
        primary key,
    tid         bigint                                   not null comment 'todo id',
    title       varchar(256)                             not null comment '标题',
    checked     tinyint(1)  default 0                    not null comment '是否完成',
    `order`     double                                   not null comment '排序',
    version     bigint      default 0                    not null comment '版本号',
    deleted     tinyint(1)  default 0                    not null comment '是否删除',
    create_time datetime(3) default CURRENT_TIMESTAMP(3) not null comment '创建时间',
    update_time datetime(3) default CURRENT_TIMESTAMP(3) not null on update CURRENT_TIMESTAMP(3) comment '修改时间'
)
    comment 'todo表' row_format = DYNAMIC;