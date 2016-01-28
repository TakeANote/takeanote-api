CURRENT_DIRECTORY := $(shell pwd)
COMPOSE := $(shell which docker-compose)
FLAVOR = development.yml

start:
	$(COMPOSE) -f $(FLAVOR) up -d

clean:
	$(COMPOSE) -f $(FLAVOR) rm --force -v

stop:
	$(COMPOSE) -f $(FLAVOR) stop

status:
	$(COMPOSE) -f $(FLAVOR) ps

logs:
	$(COMPOSE) -f $(FLAVOR) logs

restart:
	$(COMPOSE) -f $(FLAVOR) stop api web
	$(COMPOSE) -f $(FLAVOR) start api web

.PHONY: clean start stop status logs restart
