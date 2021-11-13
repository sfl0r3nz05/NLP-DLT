#set +x

# default path to be used in connection profile files
ABSPATH?=$(shell dirname $(shell dirname $(shell dirname ${PWD})))

WHICHOS := $(shell uname)

start:
	./scripts/start-monitoring.sh

stop:
	./scripts/stop-monitoring.sh

destroy:
	./scripts/destroy-monitoring.sh

ifeq ($(WHICHOS),Darwin)
	rm *-e
endif
	export ABSPATH=${ABSPATH}
