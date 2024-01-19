CURRENT_DIR=$(shell pwd)

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

proto-gen:
	./scripts/gen-proto.sh	${CURRENT_DIR}
	ls genproto/*.pb.go | xargs -n1 -IX bash -c "sed -e '/bool/ s/,omitempty//' X > X.tmp && mv X{.tmp,}"

# migrate create -ext sql -dir migrations -seq create_student_table
migrate-up:
	migrate -source file://./migrations -database 'postgres://komron:esdmk@localhost:5432/students?sslmode=disable' up
migrate-down:
	migrate -source file://./migrations -database 'postgres://komron:esdmk@localhost:5432/students?sslmode=disable' down	