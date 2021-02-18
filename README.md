# TODO
- [ ] basic web ui to show time worked by day / week
- [ ] haproxy? if frontend gets created
- [ ] error handling
---
- [x] systemd config including schedule > https://trstringer.com/systemd-timer-vs-cronjob/
- [x] go code to hit api and update db
- [x] ansible


# Setup Notes CLEAN THESE UP / ansible
$ sudo vim /usr/local/testservice.sh
echo heyoooooooooo
$ sudo vim /etc/systemd/system/punchtime.service
$ sudo systemctl start punchtime
$ journalctl -u punchtime
verify heyooo is there
$ sudo vim /etc/systemd/system/punchtime.timer
$ systemd-analyze calendar '*:0/1'
$ sudo systemctl start punchtime.timer
$ systemctl list-timers
verify my timer is in there
$ sudo systemctl enable punchtime.timer

env GOOS=linux GOARCH=arm GOARM=5 go build
