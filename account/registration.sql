CREATE TABLE registration (                                                                     
	account INT references account(id),
	activaton_key text
);
