CREATE TABLE user_mfas (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id uuid not null,
  mfa_type varchar(50) not null,
  secret varchar(255),
  is_default boolean not null DEFAULT false,
  created_at timestamp with time zone not null DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp with time zone not null DEFAULT CURRENT_TIMESTAMP,
  deleted_at timestamp with time zone
);