create table public.posts (
    id serial not null constraint posts_pk primary key,
    created_at timestamp not null default now(),
    deleted_at timestamp,
    title text not null,
    description text not null,
    content text not null
);
create unique index posts_id_uindex on public.posts (id)
where deleted_at is null;
create index posts_created_at_index on public.posts (created_at);