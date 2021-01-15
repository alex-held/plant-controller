--
-- Create Database
--
create table plant_dbs
(
    id bigserial not null
        constraint plant_dbs_pkey
            primary key,
    name text
);

alter table plant_dbs owner to postgres;

create table grow_config_dbs
(
    id bigserial not null
        constraint grow_config_dbs_pkey
            primary key,
    titel text,
    description text,
    germination_days bigint,
    growing_days bigint,
    plant_id bigint
        constraint grow_config_dbs_plant_id_fkey
            references plant_dbs
);

alter table grow_config_dbs owner to postgres;

create table tray_config_dbs
(
    id bigserial not null
        constraint tray_config_dbs_pkey
            primary key,
    plant_id bigint
        constraint tray_config_dbs_plant_id_fkey
            references plant_dbs,
    started_at timestamp with time zone
);

alter table tray_config_dbs owner to postgres;

create table tray_dbs
(
    id bigserial not null
        constraint tray_dbs_pkey
            primary key,
    slot bigint,
    is_active boolean,
    grow_config_id bigint
        constraint tray_dbs_grow_config_id_fkey
            references grow_config_dbs,
    tray_config_id bigint
        constraint tray_dbs_tray_config_id_fkey
            references tray_config_dbs
);

alter table tray_dbs owner to postgres;


--
-- Insert Data --
--
INSERT INTO public.plant_dbs (id, name) VALUES (1, 'Alfalfa');
INSERT INTO public.plant_dbs (id, name) VALUES (2, 'Mungo');
INSERT INTO public.plant_dbs (id, name) VALUES (3, 'Rucola');

INSERT INTO public.grow_config_dbs (id, titel, description, germination_days, growing_days, plant_id) VALUES (1, 'Alfalfa Config 123', 'Alfalfa Config 1', 3, 5, 1);

INSERT INTO public.tray_config_dbs (id, plant_id, started_at) VALUES (1, 1, '2021-01-14 11:45:19.721000');
INSERT INTO public.tray_config_dbs (id, plant_id, started_at) VALUES (2, 2, '2021-01-14 12:01:13.169474');
INSERT INTO public.tray_config_dbs (id, plant_id, started_at) VALUES (3, 1, '2021-01-01 23:00:00.000000');

INSERT INTO public.tray_dbs (id, slot, is_active, grow_config_id, tray_config_id) VALUES (1, 1, true, 1, 1);
INSERT INTO public.tray_dbs (id, slot, is_active, grow_config_id, tray_config_id) VALUES (2, 12, null, null, null);
