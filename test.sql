select u.name, sum(julianday("out") - julianday("in")) * 24 * 60 as minutes
from punches p
join users u on p.user_id = u.id
where "in" > datetime(current_timestamp, 'start of day')
and "out" < datetime(current_timestamp, '+1 day', 'start of day')
group by u.name
;

select "in", "out", (julianday("out") - julianday("in")) * 24 * 60 as minutes
from punches p
join users u on u.id = p.user_id
where name = 'Zach Bosteel'
;


select
  u.name
  , strftime('%Y-%m-%dT00:00:00', "in")
  , sum(julianday("out") - julianday("in")) * 24 as hours
from punches p
join users u on p.user_id = u.id
-- where u.name = 'Zach Taylor'
where strftime('%Y-%m-%dT00:00:00', "in") = strftime('%Y-%m-%dT00:00:00', 'now', '-6 hours', 'start of day')
group by u.name, datetime(strftime('%Y-%m-%dT00:00:00', "in"))
-- order by datetime(strftime('%Y-%m-%dT00:00:00', "in")) asc
order by hours desc
;

select
  u.name
  , strftime('%Y-%m-%dT00:00:00', "in")
  , "in"
  , "out"
from punches p
join users u on p.user_id = u.id
-- where u.name = 'Zach Taylor'
where strftime('%Y-%m-%dT00:00:00', "in") = strftime('%Y-%m-%dT00:00:00', 'now', '-6 hours', 'start of day')
order by u.name
;

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
