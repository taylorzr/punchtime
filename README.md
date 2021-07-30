# Punchtime

Very alpha project. Monitor slack presence to record "working" hours.

## FIXME
- [ ] Configure users to track
      Track by team?
      $ curl -sH "Authorization: Bearer $SLACK_TOKEN" 'https://slack.com/api/usergroups.list' | jq '.usergroups[] | select(.handle == "devops")'
      $ curl -sH "Authorization: Bearer $SLACK_TOKEN" 'https://slack.com/api/usergroups.users.list?usergroup=S013DRRKA4Q' | jq .
 .


## Configuration

Intended to be run on a raspberry pi. Create a file .envrc.local with the envs
- SERVER: pi's address
- SLACK_TOKEN: your slack token
- GETPING_TOKEN: your unique key from https://my.getping.info/

```
export SERVER='192.168.1.2'
export SLACK_TOKEN='<your-slack-token>'
export GETPING_TOKEN='<your-getping-token>'
```

## Raspberry Pi Installation

```
ansible-playbook punchtime.yml -i $SERVER,
```

Test by checking `make logs` and or `open http:$SERVER:8081/hours`

## Local Dev

Run Webserver: `make serve`
TODO: Run iteration on a loop
Test: `go test`

## TODO
- [ ] error service!
- [ ] move port to env and put it in overrides
- [ ] stats on user page, or page of stats?
      like avg first / last punch
      avg hours?
---
- [x] if someone is active now, count up until now as worked time
- [x] basic web ui to show time worked by day / week
- [x] systemd config including schedule > https://trstringer.com/systemd-timer-vs-cronjob/
- [x] go code to hit api and update db
- [x] ansible
- [x] makefile
- [x] service doesn't start after reboot?

## Frontend ideas
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


## Learned
- [x] alpinejs
- [x] deploying via ansible
- [x] creating systemd services & timers
- [x] basic sqlite usage, and lovely details about timezones
