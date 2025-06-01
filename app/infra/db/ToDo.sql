create table tasks_table(
	id serial primary key,
	description text not null
);

insert into tasks_table(description) values ('Criar conexão com bd');
insert into tasks_table(description) values ('Dar churu pros gatos');