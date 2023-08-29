#!/bin/bash

# 应用版本
APP_VERSION="0.0.1-beta"

# 应用名称
APP_NAME="app"

# 应用环境 dev/test/stage/prod
APP_ENV="dev"

# 目标系统 linux/windows/darwin
GO_OS="darwin"

# 目标指令集 amd64/arm64
GO_ARCH="amd64"

echo "Deleting build artifacts"

BUILD_DIR="deploy"

rm -rf $BUILD_DIR
mkdir -p $BUILD_DIR/conf
cp conf/$APP_ENV.yaml $BUILD_DIR/conf/$APP_ENV.yaml

echo "Installing dependencies"

go mod tidy

echo "Building ${APP_NAME} ${APP_VERSION} ${GO_OS}/${GO_ARCH}..."

BUILD_FLAGS="-X github.com/chubin518/kestrel-layout-basic/buildinfo.Version=$APP_VERSION \
-X github.com/chubin518/kestrel-layout-basic/buildinfo.Name=$APP_NAME \
-X github.com/chubin518/kestrel-layout-basic/buildinfo.Environment=$APP_ENV \
-X 'github.com/chubin518/kestrel-layout-basic/buildinfo.BuildTime=`date "+%Y-%m-%d %H:%M:%S"`' \
-X 'github.com/chubin518/kestrel-layout-basic/buildinfo.BuildVersion=`go version`'"

GOOS=$GO_OS GOARCH=$GO_ARCH \
go build -o $BUILD_DIR/$APP_NAME \
-ldflags "$BUILD_FLAGS" cmd/main.go cmd/wire_gen.go

echo "BUILD SUCCESS"

chmod a+x $BUILD_DIR/$APP_NAME

echo "Run App ./$BUILD_DIR/$APP_NAME"