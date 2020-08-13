-- we don't know how to generate schema main (class Schema) :(
create table funds
(
	system_id integer
		constraint funds_pk
			primary key autoincrement,
	name text not null,
	amount real not null,
	create_time integer,
	order_id integer,
	fund_type integer
);

create table orders
(
	system_id integer
		constraint orders_pk
			primary key autoincrement,
	customer_name text,
	file text,
	department text,
	progress text,
	create_time integer,
	deadline_time integer,
	order_status integer,
	price real,
	area real,
	note text,
	amount real,
	sum real,
	after text,
	maker_id integer
);

create table users
(
	system_id integer
		constraint users_pk
			primary key autoincrement,
	username text not null,
	password text not null,
	type integer not null,
	create_time integer,
	update_time integer
);

