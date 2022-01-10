-- user.roles
CREATE TABLE public."user.roles" (
	id serial NOT NULL,
	role_name varchar(20) NOT NULL,
	description text NULL,
	CONSTRAINT user_roles_pk PRIMARY KEY (id),
	CONSTRAINT user_roles_un UNIQUE (role_name)
);
COMMENT ON TABLE public."user.roles" IS 'User role table';


-- users 
CREATE TABLE public.users (
	id uuid NOT NULL,
	username varchar(30) NOT NULL,
	first_name varchar(30) NOT NULL,
	last_name varchar(30) NULL,
	email varchar(100) NOT NULL,
	active bool NOT NULL DEFAULT false,
	role_id smallint null,
	CONSTRAINT users_pk PRIMARY KEY (id),
	CONSTRAINT users_un_uname UNIQUE (username),
	CONSTRAINT users_un_email UNIQUE (email),
	CONSTRAINT users_role_fk FOREIGN KEY (role_id) REFERENCES public."user.roles"(id) ON DELETE SET NULL ON UPDATE CASCADE
	
);
COMMENT ON TABLE public.users IS 'User account';


-- insert initial data to 'user.roles'
INSERT INTO "user.roles" (name, description)
VALUES
    ('Superuser', 'Super admin'),
    ('Admin', 'Administrator');
