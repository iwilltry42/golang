# An idiomatic Golang client for the Akamai API

[![Build Status](https://travis-ci.org/akamai-api/golang.svg?branch=master)](https://travis-ci.org/akamai-api/golang)


### Project which sets up a server that listens to information from Akamai servers. It sends the data to Influxdb and shows the data in Grafana.

To see how to compile and run the application see our [documentation](./docs):
- [Akami setup](./docs/akamai.md)
- [Payload documentation](./docs/payload.md)
- [InfluxDB configuration](./docs/influxdb.md)
- [Docker configuration](./docs/docker.md)
- [Docker-Compose setup](./docs/docker-compose.md)


## Dependency Management

This project uses [dep](https://github.com/golang/dep "github.com") for dependency management.
See the README of dep for installation instructions: [https://github.com/golang/dep#setup](https://github.com/golang/dep#setup "github.com").
Execute `dep ensure [-update]` for installing/updating dependencies.