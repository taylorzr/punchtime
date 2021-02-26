# TODO
- [ ] if someone is active now, count up until now as worked time
- [ ] basic web ui to show time worked by day / week
- [ ] haproxy? if frontend gets created
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

/punches
/hours?day=2021-02-22
/hours?start=2021-02-22&end=2021-02-23
/punches?user='Zach Taylor'

/hours
/hours?day=2021-02-22
/hours?start=2021-02-22&end=2021-02-23
/hours?user='Zach Taylor'
