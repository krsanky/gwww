CREATE TABLE registration (                                                                     
	account INT references account(id),
	activation_key text,
	ts TIMESTAMP, 
	tstz TIMESTAMPTZ
);
