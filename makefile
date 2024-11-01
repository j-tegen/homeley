start:
	docker-compose -f docker/docker-compose.yaml up 

build:
	docker-compose -f docker/docker-compose.yaml up --build

stop:
	docker-compose -f docker/docker-compose.yaml down

clean:
	docker-compose -f docker/docker-compose.yaml down --volumes --remove-orphans