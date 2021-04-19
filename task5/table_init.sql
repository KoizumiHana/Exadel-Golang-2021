create table task
(
    id          int auto_increment,
    description varchar(225)                                                             not null,
    duedate     datetime                                                                 not null,
    status      enum ('NEW', 'IN_PROGRESS', 'CANCELED', 'DONE', 'EXPIRED') default 'NEW' not null,
    name        varchar(100)                                                             not null,
    constraint task_id_uindex
        unique (id)
);

alter table task
    add primary key (id);

