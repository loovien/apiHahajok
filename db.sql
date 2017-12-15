use hahajok;
-- 用户表
create table if not exists user (
    id int unsigned auto_increment,
    openId varchar(64) not null default "" comment "用户OPENID",
    unionId varchar(64) not null default "" comment "用户对应开放平台的唯一ID",
    nickname varchar(32) not null default "" comment "用户昵称",
    avatar varchar(255) not null default "" comment "用户头像",
    gender tinyint unsigned not null default "0" comment "用户性别",
    country varchar(32) not null default "" comment "用户所在国家",
    province varchar(32) not null default "" comment "用户所在省份",
    city varchar(32) not null default "" comment "用户所在城市",
    lang varchar(16) not null default "" comment "en",
    createdAt int unsigned not null default 0 comment "用户平台创建时间",
    updatedAt int unsigned not null default 0 comment "用户平台更新时间",

    primary key (id),
    unique key (openId) using hash,
    index idx_nickname (nickname) using btree,
    index idx_gender (gender),
    index idx_country (country),
    index idx_province (province),
    index idx_city (city)
) engine=innodb charset=utf8 comment "用户信息表";

