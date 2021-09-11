
#set +x

start:
	docker-compose -f docker-compose.yml up -d
	
destroy:
	docker-compose -f docker-compose.yml down

erase:
	docker-compose -f docker-compose.yml down -v
	sudo rm -rf ../dashboards/7/dashboard/*