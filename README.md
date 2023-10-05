# go-bun-newrelic-example

Go の ORM である [bun](https://bun.uptrace.dev/guide/) を New Relic で計測するサンプルです。

## 使い方

### 1. New Relic のアカウントを作成する

[New Relic](https://newrelic.com/) でアカウントを作成します。

### 2. New Relic のアカウントにアプリケーションを登録する

[New Relic Go Agent](https://docs.newrelic.com/docs/apm/agents/go-agent/get-started/introduction-new-relic-go/) でアプリケーションを登録します。

### 3. docker compose でアプリと MySQL を起動する

```bash
$ docker-compose up --build
``` 

### 4. アプリケーションにリクエストして New Relic のダッシュボードで計測を確認する

```bash
$ curl http://localhost:8080
```

