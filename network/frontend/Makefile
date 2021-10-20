#set +x

# default path to be used in connection profile files
ABSPATH?=$(shell dirname $(shell dirname $(shell dirname ${PWD})))

WHICHOS := $(shell uname)

start:
	./scripts/start-frontend.sh

stop:
	./scripts/stop-frontend.sh

destroy:
	./scripts/destroy-frontend.sh
	sudo rm -rf ../../frontend/src/data/LER.json
	sudo echo "{\"LER\": []}" > LER.json
	sudo mv LER.json ../../frontend/src/data/
	sudo chmod 777 ../../frontend/src/data/LER.json

ifeq ($(WHICHOS),Darwin)
	rm *-e
endif
	export ABSPATH=${ABSPATH}
