# GraphQL SANDBOX

GraphQL の砂場環境のリポジトリ

# requirements

- go-task 3.30.1 
- GoLang

# Usage

## ソースの実行

```
task run
```

## linter の実行
```
# .golanci.yml を元に golangci-lint が実行される
task lint
```

## GraphQL

```
task gqlgen
```

## 開発に必要なツールのインストール

```
# go-task のインストール
go install github.com/go-task/task/v3/cmd/task@v3.30.1

# 必要なツールのインストール
task install-tools
```
