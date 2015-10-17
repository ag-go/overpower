package attack

/*

create table games(
	gid SERIAL PRIMARY KEY,
	owner varchar(20) NOT NULL UNIQUE,
	name varchar(20) NOT NULL,
	size int,
	turn int
);

create table factions (
	gid integer NOT NULL REFERENCES games ON DELETE CASCADE,
	fid SERIAL NOT NULL,
	owner varchar(20) NOT NULL,
	name varchar(20) NOT NULL,
	done bool NOT NULL,
	reports text[][],
	UNIQUE(gid, name),
	PRIMARY KEY(gid, fid)
);

create table planets (
	gid integer NOT NULL REFERENCES games ON DELETE CASCADE,
	pid integer NOT NULL,
	name varchar(20) NOT NULL,
	loc point NOT NULL,
	controller int,
	inhabitants int NOT NULL,
	resources int NOT NULL,
	parts int NOT NULL,
	UNIQUE(gid, name),
	PRIMARY KEY(gid, pid),
	FOREIGN KEY(gid, controller) REFERENCES factions ON DELETE CASCADE
);

create table ships (
	gid integer NOT NULL REFERENCES games ON DELETE CASCADE,
	fid int NOT NULL,
	sid SERIAL NOT NULL,
	size int NOT NULL,
	loc point NOT NULL,
	path point[],
	FOREIGN KEY(gid, fid) REFERENCES factions ON DELETE CASCADE,
	PRIMARY KEY(gid, fid, sid)
);

create table planetviews (
	gid integer NOT NULL REFERENCES games ON DELETE CASCADE,
	fid integer NOT NULL,
	pid integer NOT NULL,
	name varchar(20) NOT NULL,
	loc point NOT NULL,
	turn int,
	controller int,
	inhabitants int,
	resources int,
	parts int,
	FOREIGN KEY(gid, fid) REFERENCES factions ON DELETE CASCADE,
	FOREIGN KEY(gid, controller) REFERENCES factions ON DELETE CASCADE,
	FOREIGN KEY(gid, pid) REFERENCES planets ON DELETE CASCADE,
	PRIMARY KEY(gid, fid, pid)
);

create table shipviews (
	gid integer NOT NULL REFERENCES games ON DELETE CASCADE,
	controller integer NOT NULL,
	viewer integer NOT NULL,
	sid SERIAL NOT NULL,
	name varchar(20) NOT NULL,
	loc point,
	trail point[],
	size int NOT NULL,
	FOREIGN KEY(gid, controller) REFERENCES factions ON DELETE CASCADE,
	FOREIGN KEY(gid, viewer) REFERENCES factions ON DELETE CASCADE,
	FOREIGN KEY(gid, controller, sid) REFERENCES ships ON DELETE CASCADE,
	PRIMARY KEY(gid, viewer, sid)
);

create table orders (
	gid integer NOT NULL REFERENCES games ON DELETE CASCADE,
	fid integer NOT NULL,
	source integer NOT NULL,
	target integer NOT NULL,
	size integer NOT NULL,
	FOREIGN KEY(gid, fid) REFERENCES factions ON DELETE CASCADE,
	FOREIGN KEY(gid, source) REFERENCES planets ON DELETE CASCADE,
	FOREIGN KEY(gid, target) REFERENCES planets ON DELETE CASCADE,
	PRIMARY KEY(gid, fid, source, target)
);





drop table games, planets, factions, ships, shipviews, planetviews, orders CASCADE;



*/