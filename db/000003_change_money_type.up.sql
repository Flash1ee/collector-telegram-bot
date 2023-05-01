ALTER TABLE costs ALTER COLUMN money TYPE bigint USING not null;

ALTER TABLE debts ALTER COLUMN money TYPE bigint USING not null;
