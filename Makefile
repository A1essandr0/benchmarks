PROTOC := protoc
PROTO_IN_PATH := common/proto
GO_OUT_PATH := ./
PYTHON_OUT_PATH := ./


list:
	ls -la

proto:
	${PROTOC} \
	--go_out=${GO_OUT_PATH} \
    --go_opt=paths=source_relative \
    --go-grpc_out=${GO_OUT_PATH} \
    --go-grpc_opt=paths=source_relative \
    ${PROTO_IN_PATH}/*.proto

	python -m grpc_tools.protoc \
    --python_out=${PYTHON_OUT_PATH} \
    --pyi_out=${PYTHON_OUT_PATH} \
    --grpc_python_out=${PYTHON_OUT_PATH} \
    --proto_path . \
    ./${PROTO_IN_PATH}/*.proto




# build-fastapi-collector:
# 	bash ./fastapi-collector/build_image.sh

# build-faust-sender:
# 	bash ./faust_sender/build_image.sh