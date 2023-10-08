# GraphQL SANDBOX

GraphQL + クリーンアーキテクチャを試すための砂場

# requirements

- Docker 24.0.5
- Docker Compose 2.20.2
- go-task 3.30.1
- GoLang

# Usage

## Getting Started
```
# コンテナの起動
$ task docker-init

# cmd/main.go を実行
$ task run  # serve at localhost:8080
```

## linter の実行
```
# .golanci.yml を元に golangci-lint が実行される
$ task lint
```

## GraphQL
```
# gqlgen generate
$ task gqlgen
```

## 開発に必要なツールのインストール

```
# go-task のインストール
$ go install github.com/go-task/task/v3/cmd/task@v3.30.1

# 必要なツールのインストール
$ task install-tools
```
