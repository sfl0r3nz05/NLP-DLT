#set +x

# default path to be used in connection profile files
ABSPATH?=$(shell dirname $(shell dirname $(shell dirname ${PWD})))

WHICHOS := $(shell uname)

start:
	./scripts/start-backend.sh
	sudo chmod 777 -R ../../backend/src/data/LER.json
	sudo chmod 777 -R ../../backend/src/data/listOfMNOs.json
	sudo chmod 777 -R ../../backend/src/data/COUNTRY.json

stop:
	./scripts/stop-backend.sh

destroy:
	./scripts/destroy-backend.sh
	sudo rm -rf ../../backend/src/wallet
	sudo rm -rf ../../backend/network
	sudo rm -rf ../../backend/src/data/listOfMNOs.json
	echo '[]' | sudo tee listOfMNOs.json >/dev/null
	sudo mv listOfMNOs.json ../../backend/src/data/
	sudo chmod 777 -R ../../backend/src/data/listOfMNOs.json

ifeq ($(WHICHOS),Darwin)
	rm *-e
endif
	export ABSPATH=${ABSPATH}
