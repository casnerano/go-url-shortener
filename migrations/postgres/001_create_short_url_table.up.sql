create table if not exists short_url (
  id serial not null,
  primary key (id),
  code varchar(64) unique not null,
  original varchar(255) not null,
  user_id integer default null,
  created_at timestamp default now()
);
