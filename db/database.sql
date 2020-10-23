CREATE DATABASE IF NOT EXISTS geomaps;

CREATE EXTENSION postgis;

create table city (
	id SERIAL primary key not null,
	name varchar(60) not null,
	state varchar(2) not null,
	state_name varchar(60) not null,
	geojson json not null,
	geom geometry not null,
)

create table area (
	id serial primary key not null,
	name varchar(120) not null,
	radius numeric not null,
	center geography not null
)

-- https://stackoverflow.com/questions/60039007/how-to-insert-geojson-data-to-geometry-field-in-postgresql
INSERT INTO mytable (json_column) VALUES ('{
    "type": "Point",
    "coordinates": [7.0069, 51.1623]
}'); 

insert into area (name, radius, center) values ($1, $2, ST_MakePoint($3, $4)) RETURNING id;

update municipio set geom = ST_GeomFromGeoJSON(geojson);
