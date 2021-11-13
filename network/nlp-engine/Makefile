#set +x

# default path to be used in connection profile files
ABSPATH?=$(shell dirname $(shell dirname $(shell dirname ${PWD})))

WHICHOS := $(shell uname)

start:
	./scripts/start-nlp-engine.sh

stop:
	./scripts/stop-nlp-engine.sh

destroy:
	./scripts/destroy-nlp-engine.sh

ifeq ($(WHICHOS),Darwin)
	rm *-e
endif
	export ABSPATH=${ABSPATH}
