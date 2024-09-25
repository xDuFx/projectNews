create table news_table (
	id serial primary key,
	title text   default 'nothing',
	body text   default 'nothing',
	image text   default 'no path',
	mark text  default 'no mark',
	reliz text  default 'no push'
);
create table users_data (
	id serial primary key,
	login text unique not null ,
	passhash text  not null,
	status int not null default 0
);
create table personal_data (
	id serial primary key,
	email text  not null,
	username text  not null,
	born_data text not null
);