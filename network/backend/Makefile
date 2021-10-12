#set +x

# default path to be used in connection profile files
ABSPATH?=$(shell dirname $(shell dirname $(shell dirname ${PWD})))

WHICHOS := $(shell uname)

start:
	./scripts/start-backend.sh

stop:
	./scripts/stop-backend.sh

destroy:
	./scripts/destroy-backend.sh
	sudo rm -rf ../../backend/src/wallet
	sudo rm -rf ../../backend/network

ifeq ($(WHICHOS),Darwin)
	rm *-e
endif
	export ABSPATH=${ABSPATH}
