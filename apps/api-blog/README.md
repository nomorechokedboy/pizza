## Getting Started

### Setup

Before running, you need to create a `.env` file in [root](./) with the
following content:

```bash
#!/bin/bash
export AUDIO_LINK=string
export AUDIO_KEY=string
```

The content can be taken from [notion](https://www.notion.so/env-8072d9b954434345a54c281f95c4ff63)

### Run

```bash
go mod download # install deps if run on local

docker compose up postgres minio-svc -d # run database and other service for local dev

go run main.go # run the main application
```

Open [http://localhost:8080](http://localhost:8080) with your browser to see the result.
