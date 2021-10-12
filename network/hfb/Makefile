
#set +x

HLNETWORK?=el-network
# default path to be used in connection profile files
ABSPATH?=$(shell dirname $(shell dirname $(shell dirname ${PWD})))

WHICHOS := $(shell uname)

start:
	make generate ABSPATH=${ABSPATH}
	./scripts/start-hfb.sh
	
destroy:
	./scripts/destroy-network.sh
	sudo rm -rf ../../backend/src/wallet
	sudo rm -rf connection-profile-[0-9].yaml
	sudo rm -rf connectionProfile.json
	sudo rm -rf docker-compose.yaml
	sudo rm -rf ../../dashboards

rmchaincode:
	docker images | grep test | awk '{print $1}' | xargs docker rmi

generate:
	cp ./templates/docker-compose-TEMPLATE.yaml docker-compose.yaml
	sed -i -e "s/HLNETWORK/${HLNETWORK}/g" docker-compose.yaml
	cp ./templates/connectionProfile-TEMPLATE.json connectionProfile.json
	sed -i -e "s/HLNETWORK/${HLNETWORK}/g" connectionProfile.json
	sed -i -e "s~ABSPATH~${ABSPATH}~g" connectionProfile.json
	for NUMBER in 1 2 3 4 ; \
	do \
		cp ./templates/connection-profile-TEMPLATE.yaml connection-profile-$${NUMBER}.yaml; \
		sed -i -e  "s/HLNETWORK/${HLNETWORK}/g" connection-profile-$${NUMBER}.yaml; \
		sed -i -e  "s/ORG/Org$${NUMBER}/g" connection-profile-$${NUMBER}.yaml; \
		sed -i -e  "s~ABSPATH~${ABSPATH}~g" connection-profile-$${NUMBER}.yaml; \
	done
ifeq ($(WHICHOS),Darwin)
	rm *-e
endif
	export ABSPATH=${ABSPATH}
