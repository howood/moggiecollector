CREATE TABLE request_logs (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  x_request_id uuid not null,
  endpoint    varchar(255) not null,
  method    varchar(100) not null,
  http_type   varchar(100) not null,
  url_query   text,
  body   text,
  header   text,
  created_at timestamp with time zone not null DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp with time zone not null DEFAULT CURRENT_TIMESTAMP,
  deleted_at timestamp with time zone
);