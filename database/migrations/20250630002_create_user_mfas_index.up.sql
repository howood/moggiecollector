CREATE INDEX user_mfas_user_id_index ON user_mfas (user_id);
CREATE INDEX user_mfas_mfa_type_index ON user_mfas (mfa_type);
CREATE INDEX user_mfas_deleted_at_index ON user_mfas (deleted_at);