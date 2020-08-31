CREATE TABLE users (
  user_id SERIAL PRIMARY KEY,
  name    varchar(255) not null,
  email   varchar(255) not null,
  password   varchar(255) not null,
  salt   varchar(255) not null,
  status   int not null,
  created_at timestamp with time zone,
  updated_at timestamp with time zone
);