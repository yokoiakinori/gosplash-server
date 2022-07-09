project = gosplash
up:
	docker-compose up -d
build:
	docker-compose build --no-cache
stop:
	docker-compose stop
down:
	docker-compose down
app:
	docker-compose exec app sh
nginx:
	docker exec -it $(project)_nginx /bin/bash
mysql:
	docker exec -it $(project)_mysql /bin/bash
