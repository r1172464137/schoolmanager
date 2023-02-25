create table if not exists s_student
(
    id       int                    not null comment '学号'
        primary key,
    username varchar(15)            not null comment '用户名',
    password varchar(20)            not null comment '密码',
    name     varchar(5)             not null comment '学生姓名',
    college  varchar(10) default '' null comment '所在学院',
    course   varchar(15) default '' null comment '专业',
    status   tinyint(1)  default 0  not null comment '学生状态',
    constraint s_student_username_uindex
        unique (username)
)
    comment '学生数据库';

create table if not exists s_teacher
(
    id       int         not null comment '学工号'
        primary key,
    name     varchar(5)  not null comment '姓名',
    username varchar(15) not null comment '登录名',
    password varchar(15) not null comment '密码',
    college  varchar(20) not null comment '所属学院',
    constraint s_teacher_username_uindex
        unique (username)
)
    comment '老师表';

create table if not exists s_leave
(
    id     int auto_increment comment '假条id'
        primary key,
    ask    int                    not null comment '申请id',
    review int                    not null comment '审核id',
    reason varchar(50) default '' null comment '申请理由',
    status tinyint(1)             not null comment '0通过1未通过',
    creat  date                   not null comment '申请时间',
    time   varchar(8)             null comment '时常',
    constraint s_leave_s_student_null_fk
        foreign key (ask) references s_student (id),
    constraint s_leave_s_teacher_null_fk
        foreign key (review) references s_teacher (id)
)
    comment '请假';

create table if not exists s_scan
(
    id   int         not null comment '假条id'
        primary key,
    base varchar(50) not null,
    hash varchar(50) not null,
    last varchar(50) not null,
    constraint s_scan_s_leave_null_fk
        foreign key (id) references s_leave (id)
)
    comment '出入证明';

