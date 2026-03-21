-- +goose Up
-- +goose StatementBegin
DO $$ BEGIN
    CREATE TYPE ledger_entry_type AS ENUM ('DEBIT', 'CREDIT');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;
-- +goose StatementEnd

-- +goose Down
DROP TYPE IF EXISTS ledger_entry_type;

