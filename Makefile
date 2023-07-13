
create:
	./.start_project

start: create 
	cd docker/dev && docker-compose up -d && docker-compose exec -it api /bin/ash

build: create
	cd docker/dev && docker-compose up --build

attach:
	cd docker/dev && docker-compose exec -it api /bin/ash

stop:
	cd docker/dev && docker-compose stop

down:
	cd docker/dev && docker-compose down

run_start:
	cd docker/dev && docker-compose up -d && docker-compose exec api /start.sh