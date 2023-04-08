create extension if not exists "uuid-ossp";

create table subjects
(
    id              uuid      default uuid_generate_v4() not null
        constraint subjects_pk
            primary key,
    name                 varchar                              not null,
    description          varchar                              not null
);

create table tasks
(
    id              uuid      default uuid_generate_v4() not null
        constraint tasks_pk
            primary key,
    subject_id           uuid                              not null,
    description          varchar                              not null
);