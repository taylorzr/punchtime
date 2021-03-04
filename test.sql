-- first / last punch of the day
select u.name, min("in"), max("out") from punches p
join users u on u.id = p.user_id
where "in" between '2021-03-03T00:00:00Z' and '2021-03-04T06:00:00Z'
group by u.name
order by min("in") asc
;


-- hours worked today including now
select
	u.name
	, coalesce(sum(julianday(min(coalesce("out", datetime('now')), '2021-03-04T23:00:00Z')) - julianday("in")) * 24 , 0) as hours
from punches p
join users u on p.user_id = u.id
where "in" between '2021-03-03T15:00:00Z' and '2021-03-03T23:00:00Z'
group by u.name
order by hours desc
;


-- hours worked today not including now
select
	u.name
	, coalesce(sum(julianday("out") - julianday("in")) * 24 , 0) as hours
from punches p
join users u on p.user_id = u.id
where "in" between '2021-03-01T06:00:00Z' and '2021-03-02T06:00:00Z'
group by u.name
order by hours desc
;


-- punches for a specific user
select
  u.name
  , coalesce("out", datetime('now'))
  , "in"
from punches p
join users u on p.user_id = u.id
where "in" between '2021-03-01T06:00:00Z' and '2021-03-02T06:00:00Z'
and u.name = 'Zach Taylor'
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
