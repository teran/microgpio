# microgpio
Microservice to control GPIO pins via REST

[![Go Report](https://goreportcard.com/badge/github.com/teran/microgpio)](https://goreportcard.com/report/github.com/teran/microgpio)
[![Layers size](https://images.microbadger.com/badges/image/teran/microgpio.svg)](https://hub.docker.com/r/teran/microgpio/)
[![Recent build commit](https://images.microbadger.com/badges/commit/teran/microgpio.svg)](https://hub.docker.com/r/teran/microgpio/)
[![Docker Automated build](https://img.shields.io/docker/automated/teran/microgpio.svg)](https://hub.docker.com/r/teran/microgpio/)
[![GoDoc](https://godoc.org/github.com/teran/microgpio?status.svg)](https://godoc.org/github.com/teran/microgpio)

# Docker
```
docker run -d -p8080:8080 -v /sys:/sys teran/microgpio:armv7-latest
```

Please note microgpio is built for ARMv7 at the moment.
If you need any other platform(supported by Go, obviously;), feel free to report issue or create a pull request.
