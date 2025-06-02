create table tasks_table(
	id serial primary key,
	description text not null,
	owner text,
	status int
);


insert into tasks_table(description, owner, status) values ('Criar conexão com bd', 'Joy', 0);
insert into tasks_table(description, owner, status) values ('Dar churu pros gatos', 'Joy', 0 );