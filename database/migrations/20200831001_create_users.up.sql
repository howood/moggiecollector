CREATE TABLE users (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  name    varchar(255) not null,
  email   varchar(255) not null,
  password   varchar(255) not null,
  salt   varchar(255) not null,
  status   int not null,
  created_at timestamp with time zone not null DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp with time zone not null DEFAULT CURRENT_TIMESTAMP,
  deleted_at timestamp with time zone
);