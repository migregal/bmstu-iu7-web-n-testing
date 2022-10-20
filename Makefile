.PHONY: lint
lint:
	$(MAKE) -C ./backend lint

.PHONY: unit-test
unit-test:
	$(MAKE) -C ./backend unit-test

.PHONY: smoke-test
smoke-test:
	$(MAKE) -C ./backend smoke-test

.PHONY: load-test
load-test:
	$(MAKE) -C ./testing load-test

.PHONY: generate-crypto
generate-crypto:
	$(MAKE) -C ./backend generate-crypto
	$(MAKE) -C ./nginx generate-crypto

.PHONY: dev-start
dev-start:
	docker compose -f ./docker/docker-compose.dev.yml up --build

.PHONY: prod-start
prod-start:
	docker compose -f ./docker/docker-compose.prod.yml up --build -d

.PHONY: prod-stop
prod-stop:
	docker compose -f ./docker/docker-compose.prod.yml down
