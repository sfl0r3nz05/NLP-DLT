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
	echo '{"LER": []}' | sudo tee LER.json >/dev/null
	sudo mv LER.json ../../frontend/src/data/
	sudo chmod 777 -R ../../frontend/src/data/LER.json

ifeq ($(WHICHOS),Darwin)
	rm *-e
endif
	export ABSPATH=${ABSPATH}
