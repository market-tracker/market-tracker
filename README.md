# MARKET-TRACKER

## Description of the project

[websocket tool](https://github.com/nhooyr/websocket)

## Setup the project

### Init the project

#### Development

Ensure to install [pre-commit](https://pre-commit.com/#install) and follow the instructions of [conventional-pre-commit](https://github.com/compilerla/conventional-pre-commit)

First, it must be created a volume. This is necessary to avoid to install all the dependencies all the time when the project was initialized. If there are some problems with dependencies, it can be tested to remove the volume then create it again, and try again with the development.

```bash
docker volume create market-tracker-volume
```

With the next commands, it can be installed the project, please note, that the project will run in the port 8080 in your machine, and in the port 3000 in the docker image. It can be modified if it is what is wanted.

```bash
docker build -t market-tracker -f development.Dockerfile .
docker run --name market --mount source=market-tracker-volume,target=/go -v $PWD:/home/market-tracker -p 8080:3000 -d market-tracker
```

For install new packages is necessary to use de docker daemon like a bridge. All of this to avoid to install the packages in the machine and not in the docker container. The important thing to show here is the record of the packages installed will be in the go.mod file.

```bash
docker exec market go get -u github.com/gin-gonic/gin
```

#### Production

With this commands it can be executed the project for production

```bash
docker build -t market-tracker .
docker run --name market -p 8080:3000 -d market-tracker
```

