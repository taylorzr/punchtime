clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean
	ansible-playbook deploy.yml -i hosts.yml

serve:
	reflex -d none -s -R vendor. -r \.go$ -- go run . serve

sql:
	ssh pi@192.168.1.2 -t 'sqlite3 /usr/local/share/punchtime/punchtime.db'

ssh:
	ssh pi@192.168.1.2

logs:
	ssh pi@192.168.1.2 -t 'journalctl -u punchtime -n 100 -f'
