SELECT
	'CREATE DATABASE trvl'
WHERE NOT EXISTS (
	SELECT FROM pg_database WHERE datname = 'trvl'
);
