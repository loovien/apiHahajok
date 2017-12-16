use hahajok;
-- 用户表
create table if not exists user (
    id int unsigned auto_increment,
    openId varchar(64) not null default '' comment '用户OPENID',
    unionId varchar(64) not null default '' comment '用户对应开放平台的唯一ID',
    nickname varchar(32) not null default '' comment '用户昵称',
    wallet decimal(8,2) not null default '0' comment '用户钱包',
    avatar varchar(255) not null default '' comment '用户头像',
    gender tinyint unsigned not null default '0' comment '用户性别',
    country varchar(32) not null default '' comment '用户所在国家',
    province varchar(32) not null default '' comment '用户所在省份',
    city varchar(32) not null default '' comment '用户所在城市',
    lang varchar(16) not null default '' comment 'en',
    issave tinyint not null default '0' comment '是否存储过',
    createdAt int unsigned not null default 0 comment '用户平台创建时间',
    updatedAt int unsigned not null default 0 comment '用户平台更新时间',

    primary key (id),
    unique key (openId) using hash,
    index idx_nickname (nickname) using btree,
    index idx_gender (gender),
    index idx_country (country),
    index idx_province (province),
    index idx_city (city)
) engine=innodb charset=utf8 comment '用户信息表';

-- 分类名称
create table if not exists classification (
    id int unsigned auto_increment,
    name varchar(64) not null default '' comment '分类名称',
    icon varchar(255) not null default '' comment '分类ICON',
    createdAt int unsigned not null default 0 comment '用户平台创建时间',
    updatedAt int unsigned not null default 0 comment '用户平台更新时间',

    primary key (id)
) engine=innodb charset=utf8 comment '笑话分类表';

-- 段子表
create table if not exists joker (
    id int unsigned auto_increment,
    uid int unsigned not null default '0' comment '所属用户UID',
    classId int unsigned not null default '0' comment '所属于分类ID',
    title varchar(64) not null default '' comment '段子标题',
    content text null default null comment '段子内容',
    imageList text null default null comment '段子图片, json',
    mediaUrl varchar(255) not null default '' comment '视频地址',
    replies int unsigned not null default '0' comment '回复数量',
    status tinyint not null default '0' comment '状态, [-1/删除, 0/审核中, 1/正常 ]',
    createdAt int unsigned not null default 0 comment '用户平台创建时间',
    passedAt int unsigned not null default 0 comment '审核通过时间',
    updatedAt int unsigned not null default 0 comment '用户平台更新时间',

    primary key (id),
    index idx_uid (uid),
    index idx_classid (classId),
    index idx_title (title),
    index idx_status (status),
    index idx_createdat (createdAt),
    index idx_updatedat (updatedAt)
) engine=innodb charset=utf8 comment '段子表';

-- 评论表
create table if not exists replies (
    id int unsigned auto_increment,
    jokerId int unsigned not null default '0' comment '段子ID',
    uid int unsigned not null default '0' comment '评论用户UID',
    content text null default null comment '评论内容',
    imageList text null default null comment '评论图片列表, json',
    status tinyint not null default '0' comment '状态, [-1/删除, 0/审核中, 1/正常 ]',
    createdAt int unsigned not null default 0 comment '用户平台创建时间',
    passedAt int unsigned not null default 0 comment '审核通过时间',
    updatedAt int unsigned not null default 0 comment '用户平台更新时间',

    primary key (id),
    index idx_jokerid (jokerId),
    index idx_uid (uid),
    index idx_status (status),
    index idx_createdat (createdAt),
    index idx_updatedat (updatedAt)
) engine=innodb charset=utf8 comment '用户评论表';
