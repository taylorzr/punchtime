.DEFAULT_GOAL := help
.PHONY := help deploy serve sql ssh logs

help:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'

clean:
	rm -rf ./bin ./vendor Gopkg.lock

build: clean
	env GOOS=linux GOARCH=arm GOARM=5 go build

deploy: clean
	ansible-playbook punchtime.yml -i hosts.yml

restart:
	ssh pi@${SERVER} -t 'sudo systemctl restart punchtime_web.service'

serve:
	exec reflex -s -d none -r '\.go$$' -- go run . serve

sql:
	ssh pi@${SERVER} -t 'sudo sqlite3 -header -column /usr/local/share/punchtime/punchtime.db'

ssh:
	ssh pi@${SERVER}

logs:
	ssh pi@${SERVER} -t 'journalctl -u punchtime_recorder -u punchtime_web -n 100 -f'
