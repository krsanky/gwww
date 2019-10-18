CREATE TABLE phrase (                                                                     
    id SERIAL PRIMARY KEY,
	phrase TEXT NOT NULL,
	tags TEXT NOT NULL DEFAULT '',
	path TEXT NOT NULL DEFAULT '',
	order_ INT NOT NULL DEFAULT 0,
	UNIQUE(tags , path, order_) 
);
