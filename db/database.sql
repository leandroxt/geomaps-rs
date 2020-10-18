CREATE DATABASE IF NOT EXISTS geomaps;

CREATE EXTENSION postgis;

create table city (
	id SERIAL primary key not null,
	name varchar(60) not null,
	state varchar(2) not null,
	state_name varchar(60) not null,
	geojson json not null,
	geom geometry
)

-- https://stackoverflow.com/questions/60039007/how-to-insert-geojson-data-to-geometry-field-in-postgresql
INSERT INTO mytable (json_column) VALUES ('{
    "type": "Point",
    "coordinates": [7.0069, 51.1623]
}'); 

update municipio set geom = ST_GeomFromGeoJSON(geojson);
