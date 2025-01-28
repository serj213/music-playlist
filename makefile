.PHONY: dc, run

dc:
	docker-compose up --remove-orphans --build

run:
	Config=config/local.yml go run cmd/main.go