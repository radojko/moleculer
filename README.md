# Moleculer Go

<img src="https://moleculer-go-site.herokuapp.com/images/moleculer-gopher.jpg" alt="Moleculer Go" height="50"/>
🚀 Progressive microservices framework for Go <img src="https://golang.org/doc/gopher/frontpage.png" alt="Gopher" height="50"/>

Inspired and compatible with [Moleculer JS](https://github.com/moleculerjs/moleculer)

Simple, fast, light and fun to develop with. Also easy, very easy to test ;)

[![Gitter](https://badges.gitter.im/moleculer-go/community.svg)](https://gitter.im/moleculer-go/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)
[![Drone.io Build Status](https://cloud.drone.io/api/badges/moleculer-go/moleculer/status.svg)](https://cloud.drone.io/moleculer-go/moleculer)
[![Go Report Card](https://goreportcard.com/badge/github.com/moleculer-go/moleculer)](https://goreportcard.com/report/github.com/moleculer-go/moleculer)
[![Coverage -> Coveralls](https://coveralls.io/repos/github/moleculer-go/moleculer/badge.svg?branch=master)](https://coveralls.io/github/moleculer-go/moleculer?branch=master)
[![Coverage -> Codecov](https://codecov.io/gh/moleculer-go/moleculer/branch/develop/graph/badge.svg)](https://codecov.io/gh/moleculer-go/moleculer)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fmoleculer-go%2Fmoleculer.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fmoleculer-go%2Fmoleculer?ref=badge_shield)

Travis:

[![Travis Build Status](https://travis-ci.org/moleculer-go/moleculer.svg?branch=develop)](https://travis-ci.org/moleculer-go/moleculer)

**Website**: [gomicro.services](http://gomicro.services)

**Documentation**: [Docs](http://gomicro.services/docs/)

## Timeline

# Get Started

- [Database examples](https://moleculer-go-site.herokuapp.com/docs/0.1/moleculer-db.html)
- [WhatsApp App](https://github.com/moleculer-go/example-whatsapp)
- [Documentation](http://gomicro.services/docs/)
- [Benchmark](https://github.com/moleculer-go/benchmark)

## Example

![Example](https://moleculer-go-site.herokuapp.com/images/main-example.png)

# Roadmap

## v0.1.0 (MVP)

- Service Broker
- Transit and Transport
- Actions (request-reply)
- Events
- Mixins
- Load balancing for actions and events (random round-robin)
- Service registry & dynamic service discovery
- Versioned services
- Middlewares
- NATS Streaming Transporter
- JSON Serializer
- Examples :)

## v0.2.0 (Beta RC1)

- Action validators
- Support for streams
- More Load balancing implementations (cpu-usage, latency)
- Fault tolerance features (Circuit Breaker, Bulkhead, Retry, Timeout, Fallback)
- Built-in caching solution (memory, Redis)
- More transporters (gRPC, TCP, Redis, Kafka)
- More serializers (Avro, MsgPack, Protocol Buffer, Thrift)

## v0.3.0 (Beta)

- Performance and Optimization
- More DB Adaptors (SQLLite, Firebase, MySQL)
- CLI for Project Seed Generation

## v0.4.0 (Alpha)

- Event Sourcing Mixins

## v0.5.0 (Release)

# Installation

```bash
$ go get github.com/moleculer-go/moleculer
```

# Running examples

```bash

# simple moleculer db example with memory adaptor
$ go run github.com/moleculer-go/moleculer-db/examples/users

# simple moleculer db example with Mongo adaptor
$ go run github.com/moleculer-go/moleculer-db/examples/usersMongo

# complex moleculer db example with population of fields by other services
$ go run github.com/moleculer-go/moleculer-db/examples/populates


```
