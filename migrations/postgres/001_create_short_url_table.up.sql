create table if not exists short_url (
  id serial not null,
  primary key (id),
  code varchar(64) unique not null,
  original varchar(255) not null,
  created_at timestamp DEFAULT now(),
  deleted_at timestamp not null
);
