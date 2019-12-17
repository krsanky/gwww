CREATE TABLE account (                                                                     
	id INT GENERATED ALWAYS AS IDENTITY (MINVALUE 1),
    password character varying(512) NOT NULL,
    is_superuser boolean NOT NULL,                                       
    username character varying(150) UNIQUE,
    first_name character varying(30) NOT NULL,
    last_name character varying(30) NOT NULL,
    email character varying(254) NOT NULL UNIQUE,
    is_staff boolean NOT NULL,
    is_active boolean NOT NULL,
	UNIQUE(id)
);
--    timezone character varying(128) default 'America/New_York' 
