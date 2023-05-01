ALTER TABLE costs ALTER COLUMN money TYPE real USING money::real;

ALTER TABLE debts ALTER COLUMN money TYPE real USING money::real;