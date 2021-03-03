# TODO
- [ ] if someone is active now, count up until now as worked time
- [ ] basic web ui to show time worked by day / week
- [ ] error service!
- [ ] move port to env and put it in overrides
---
- [x] systemd config including schedule > https://trstringer.com/systemd-timer-vs-cronjob/
- [x] go code to hit api and update db
- [x] ansible
- [x] makefile


# Frontend ideas
- [ ] dropdown for day selection
- [ ] dropdown for user selection
- [ ] some checkbox or filter for only working hours like 9-5 or 8-6
- [ ] per day metrics (or abstract to per time period?)
	- [ ] first/last punch for the day per user
	- [ ] availability during core hours 9-5 for each user
	- [ ] hours over entire day for each user

/punches
/hours?day=2021-02-22
/hours?start=2021-02-22&end=2021-02-23
/punches?user='Zach Taylor'

/hours
/hours?day=2021-02-22
/hours?start=2021-02-22&end=2021-02-23
/hours?user='Zach Taylor'


# Learned
- [x] deploying via ansible
- [x] creating systemd services & timers
- [x] basic sqlite usage, and lovely details about timezones
