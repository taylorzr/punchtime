select sum(julianday("out") - julianday("in")) from punches group by user_id;


-- Previous data structure, lol the query complexity is ridiculous because the data model was wrong
with first_active as (
  select id from punches
  where presence = 'active'
  and time > datetime(current_timestamp, 'start of day')
  and time < datetime(current_timestamp, '+1 day', 'start of day')
  order by id asc
  limit 1
),
timeblocks as (
  select
    row_number() over w as row
    , time
    , lag(time, 1, null) over w as previous_time
    , (julianday(time) - julianday(lag(time, 1, null) over (order by punches.id))) * 24 * 60 as minutes
  from punches, first_active
  where user_id = 1
  and punches.id >= first_active.id
  and time > datetime(current_timestamp, 'start of day')
  and time < datetime(current_timestamp, '+1 day', 'start of day')
  window w as (order by punches.id)
  order by punches.id asc
)
select *
from timeblocks
where row % 2 = 0
;
