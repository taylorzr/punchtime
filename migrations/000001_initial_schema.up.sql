create table if not exists users (
  id integer primary key
  , name text not null unique
  , slack_id text not null
);

create table if not exists punches (
  id integer primary key
  , user_id integer not null
  , "in" text not null -- format: ISO8601 e.g. "YYYY-MM-DD HH:MM:SS.SSS"
  , "out" text -- format: ISO8601 e.g. "YYYY-MM-DD HH:MM:SS.SSS"
);

insert or ignore into users (name, slack_id) values
('Zach Taylor', 'W0162LHTW2C')
, ('Scott Rakes', 'W014Y3GK6RM')
, ('Nat Williams', 'W014Y3GR60P')
, ('Dan Chrul', 'W015D22662F')
, ('Zach Bosteel', 'W015CRBN1PC')
, ('Matthew Dalton', 'W015BE718DT')
, ('Matt Hoelting', 'W01563TBQES')
, ('Mike Son', 'W015K187CGL')
, ('Molly Gallagher', 'W0162LJS0TA')
, ('Umang Kapadia', 'W014Y3J166B')
, ('Mike Ross', 'W015BE71YVB')
;
