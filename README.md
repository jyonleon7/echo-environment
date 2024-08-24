# echo-environment

## 環境変数
| 環境変数名   | default value | description                             |
|---------|---------------|-----------------------------------------|
| APP_ENV | `dev`         | `local`、`dev` など環境にあったものを想定しています。       |
| PORT    | `8080`        | listenするポートを指定してください。defaultは `8080`です。 |


## 初期設定
```
go mod init github.com/jyonleon7/echo-environment
```

```
go mod tidy
```


## OpenAPIのYAMLからサーバーinterfaceの自動生成
oapi-codegenを用いてopenapiのyamlからサーバー側のインターフェイスの自動生成を行っています。
今までに一度もoapi-codegeを利用したことがない場合、oapi-codegenをいれる必要があります。`go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest`

`make oapi`を実行すると`specification/openapi.yaml`で定義されたOpenAPIのserver interfaceが生成されます。
openapi.yamlに制約をかけばvalidationもされます。

1. `specification/openapi.yaml`を編集します。
2. `make oapi` を打ってserverのinterfaceを更新します。
3. `internal/handler/main.go` でhandlerの中身を実装します。
