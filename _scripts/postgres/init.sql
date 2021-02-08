GRANT ALL PRIVILEGES ON DATABASE levee TO docker;

CREATE TABLE jobs (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    partner_id VARCHAR UNIQUE NOT NULL ,
	title VARCHAR NOT NULL,
    category_id VARCHAR NOT NULL,
	status 	varchar NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL
);
