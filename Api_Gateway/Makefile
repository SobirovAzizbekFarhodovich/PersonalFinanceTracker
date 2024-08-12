CURRENT_DIR=$(shell pwd)

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

exp:
	export DBURL='postgres://postgres:root@localhost:5432/staffer?sslmode=disable'

mig-run:
	migrate create -ext sql -dir migrations -seq create_table

mig-up:
	migrate -database 'postgres://postgres:root@localhost:5432/staffer?sslmode=disable' -path migrations up

mig-down:
	migrate -database 'postgres://postgres:root@localhost:5432/staffer?sslmode=disable' -path migrations down

gen-proto:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	export PATH="$PATH:$(go env GOPATH)/bin"
	protoc --go_out=. --go-grpc_out=. city_submodule/*.proto

gen-protoAll:
#proto fileni hammasini bittada generatsiya qilish 
	protoc --go_out=./ \
	--go-grpc_out=./ \
	city_submodule/*.proto

swag-gen:
	~/go/bin/swag init -g ./api/api.go -o docs

# PROTO_DIR=proto
# OUT_DIR=genproto

# proto-genn:
#     protoc -I=$(PROTO_DIR) --go_out=$(OUT_DIR) --go-grpc_out=$(OUT_DIR) $(PROTO_DIR)/*.proto
