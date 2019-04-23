CREATE TABLE account (                                                                     
    id SERIAL PRIMARY KEY,
    password character varying(512) NOT NULL,
    is_superuser boolean NOT NULL,                                       
    username character varying(150) NOT NULL UNIQUE,
    first_name character varying(30) NOT NULL,
    last_name character varying(30) NOT NULL,
    email character varying(254) NOT NULL,
    is_staff boolean NOT NULL,
    is_active boolean NOT NULL,
    timezone character varying(128) default 'America/New_York' 
);
