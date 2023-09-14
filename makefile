SHELL=/bin/bash

export DB_USER_WRITER=rani  
export DB_PASSWORD_WRITER=rani  
export DB_HOST_WRITER=127.0.0.1  
export DB_PROTOCOL_WRITER=tcp   
export DB_PORT_WRITER=3306 
export DB_NAME_WRITER=simpleinvitation  

export DB_USER_READER=rani  
export DB_PASSWORD_READER=rani  
export DB_HOST_READER=127.0.0.1  
export DB_PROTOCOL_READER=tcp 
export DB_PORT_READER=3306 
export DB_NAME_READER=simpleinvitation 


run-local-config: 
	go run cmd/main.go

run-local-db-up: 
	docker compose -f tools/local/database.yml up

run-local-db-down: 
	docker compose -f tools/local/database.yml down


### please kill your local 8080 port first and go main. 
run-smoke-test-local: 
	k6 run './tools/qa/smoke/ping.js'
	k6 run './tools/qa/smoke/gathering.js'

run-load-local: 
	k6 run './tools/qa/load/ping.js'
	k6 run './tools/qa/load/ping_controller.js'
	k6 run './tools/qa/load/ping_usecase.js'
	k6 run './tools/qa/load/ping_repository.js'

	
