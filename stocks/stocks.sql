CREATE TABLE stock_symbols (                                                                     
    id SERIAL PRIMARY KEY,
	symbol TEXT NOT NULL,
	name TEXT NOT NULL DEFAULT ''
);


