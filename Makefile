.DEFAULT_GOAL := build

clean:
	rm -rf ./bin ./vendor Gopkg.lock

build: clean
	env GOOS=linux GOARCH=arm GOARM=5 go build

deploy: clean
	ansible-playbook punchtime.yml -i hosts.yml

# FIXME: Doesn't seem to auto-reload inside this makefile :(
serve:
	reflex -d none -s -R vendor. -r \.go$ -- go run . serve

sql:
	ssh pi@192.168.1.2 -t 'sudo sqlite3 -header -column /usr/local/share/punchtime/punchtime.db'

ssh:
	ssh pi@192.168.1.2

logs:
	ssh pi@192.168.1.2 -t 'journalctl -u punchtime_recorder -u punchtime_web -n 100 -f'
