
start:
	./.start_project
	# cd docker/dev && docker-compose up -d && docker-compose exec -it api /bin/ash

attach:
	cd docker/dev && docker-compose exec -it api /bin/ash

stop:
	cd docker/dev && docker-compose stop

down:
	cd docker/dev && docker-compose down

build:
	cd docker/dev && docker-compose up --build