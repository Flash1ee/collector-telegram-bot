ALTER TABLE costs ALTER COLUMN money TYPE bigint USING money::integer;

ALTER TABLE debts ALTER COLUMN money TYPE bigint USING money::integer;
