ALTER TABLE transactions
ALTER COLUMN time TYPE TIME USING time::time,
ALTER COLUMN date TYPE DATE USING date::date;