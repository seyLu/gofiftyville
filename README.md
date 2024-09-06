<div align="center">
    <img height=150 src="https://raw.githubusercontent.com/seyLu/gofiftyville/main/detective-golang.svg" alt="gofiftyville icon">
    <h1>Gofiftyville</h1>
    <p>Fiftyville API written in Go.</p>
    <p>
        <a href="https://seylu.github.io/gofiftyville/docs"><img src="https://img.shields.io/badge/api-docs-68d6e1" alt="Docs"></a>
        <a href="https://goreportcard.com/report/github.com/seyLu/gofiftyville"><img src="https://goreportcard.com/badge/github.com/seyLu/gofiftyville" alt="Go Report Card"></a>
        <a href="https://github.com/seyLu/gofiftyville/actions/workflows/test.yml"><img src="https://github.com/seyLu/gofiftyville/actions/workflows/test.yml/badge.svg" alt="Test badge"></a>
        <a href="https://codecov.io/github/seyLu/gofiftyville"><img src="https://codecov.io/github/seyLu/gofiftyville/graph/badge.svg?token=ZBUH7E9IYH"/></a>
        <a href="https://github.com/seyLu/gofiftyville/blob/main/LICENSE"><img src="https://img.shields.io/github/license/seyLu/gofiftyville.svg" alt="MIT License"></a>
    </p>
    <p>
        <a href="https://github.com/seyLu/gofiftyville/issues/new">Report Bug</a>
        ·
        <a href="https://github.com/seyLu/gofiftyville/issues/new">Request Feature</a>
        ·
        <a href="https://github.com/seyLu/gofiftyville/discussions">Ask Question</a>
    </p>
</div>

<br>

### API Endpoints

The list of valid api endpoints are available at the [docs](https://seylu.github.io/gofiftyville/docs).

<br>

### Developing Locally

#### Running the server

```bash
go run ./cmd/server
```

You can access the server at `localhost:8080/api/v1/:endpoint`.

#### Using Docker

##### 1. Supply the environment variables

```bash
cp .env.example .env
```

##### 2. Run docker compose

```bash
docker-compose up -d
```

You can access the server at `localhost/api/v1/:endpoint` or `localhost:8080/api/v1/:endpoint`.
