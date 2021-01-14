create table plant
(
    plant_id integer unique primary key,
    name     text not null
);

create table grow_config
(
    grow_config_id integer unique primary key,
    plant_id integer references plant,
    germination_days integer not null,
    grow_days integer not null
);


create table tray_config
(
    tray_config_id integer unique primary key,
    plant_id       integer references plant,
    started_at     date not null
);

CREATE table tray
(
    tray_id        integer unique primary key,
    slot           integer                        not null,
    isActive       bool                           not null,
    tray_config_id integer references tray_config null,
    grow_config_id integer references grow_config null
);
