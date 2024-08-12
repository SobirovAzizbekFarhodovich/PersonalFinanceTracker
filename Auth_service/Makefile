CURRENT_DIR=$(shell pwd)
DBURL="postgres://azizbek:123@localhost:5432/farimsh?sslmode=disable"
exp:
	export DBURL="postgres://azizbek:123@localhost:5432/farimsh?sslmode=disable"

mig-up:
	migrate -path migrations -database ${DBURL} -verbose up

mig-down:
	migrate -path migrations -database ${DBURL} -verbose down


mig-create:
	migrate create -ext sql -dir migrations -seq create_table

mig-insert:
	migrate create -ext sql -dir migrations -seq insert_table

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}
swag-gen:
	~/go/bin/swag init -g ./api/api.go -o docs force 1	