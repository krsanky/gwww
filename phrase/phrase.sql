CREATE TABLE phrase (                                                                     
	id INT GENERATED ALWAYS AS IDENTITY (MINVALUE 1),
	phrase TEXT NOT NULL DEFAULT '',
	tags TEXT NOT NULL DEFAULT '',
	path TEXT NOT NULL DEFAULT '',
	source TEXT NOT NULL DEFAULT '',
	order_ INT NOT NULL DEFAULT 0,
	UNIQUE(tags , path, order_) 
);

-- alter table phrase add column source text NOT NULL default '';
-- alter table phrase alter  column phrase set   default '';


