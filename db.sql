create table persons
(
	id bigserial not null
		constraint persons_pk
			primary key,
	first_name varchar(255),
	last_name varchar(255),
	birth_date date
);

alter table persons owner to postgres;

INSERT INTO public.persons (id, first_name, last_name, birth_date) VALUES (1, 'Arif', 'Rakhman', '2010-01-01');
INSERT INTO public.persons (id, first_name, last_name, birth_date) VALUES (2, 'Dyta', 'Vina', '2010-01-02');


create table phone_number
(
	id bigserial not null
		constraint phone_number_pk
			primary key,
	person_id integer,
	phone_number varchar(30),
	type varchar(255)
);

alter table phone_number owner to postgres;

INSERT INTO public.phone_number (id, person_id, phone_number, type) VALUES (1, 1, '0821345678', 'smart phone');
INSERT INTO public.phone_number (id, person_id, phone_number, type) VALUES (2, 1, '021666666', 'work');
INSERT INTO public.phone_number (id, person_id, phone_number, type) VALUES (3, 2, '087654321', 'smart phone');
INSERT INTO public.phone_number (id, person_id, phone_number, type) VALUES (4, 2, '021777777', 'work');