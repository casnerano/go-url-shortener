create table if not exists short_url (
  id serial not null,
  primary key (id),
  code varchar(64) unique not null,
  original varchar(255) not null,
  user_uuid varchar(64) default null,
  created_at timestamp default now(),
  unique (original, user_uuid)
);
