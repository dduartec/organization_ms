    CREATE TABLE logMovements
    (	
        move_id serial NOT NULL,
	owner character varying(500) NOT NULL,
        origen character varying(500) NOT NULL,
        destiny character varying(500) NOT NULL,
        date date,
        CONSTRAINT logMovimientos_pkey PRIMARY KEY (move_id)
    )
    WITH (OIDS=FALSE);

	CREATE TABLE logFolderCreations
    (	
        create_id serial NOT NULL,
	owner character varying(500) NOT NULL,
        path character varying(500) NOT NULL,
        date date,
        CONSTRAINT logFolderCreations_pkey PRIMARY KEY (create_id)
    )
    WITH (OIDS=FALSE);

	CREATE TABLE logDeletes
    (	
        del_id serial NOT NULL,
	owner character varying(500) NOT NULL,
        path character varying(500) NOT NULL,
        date date,
        CONSTRAINT logDeletes_pkey PRIMARY KEY (del_id)
    )
    WITH (OIDS=FALSE);
