# Mesher
jjjjjjjjjjjjj
[![Build Status](https://travis-ci.org/go-chassis/mesher.svg?branch=master)](https://travis-ci.org/go-chassis/mesher) [![Coverage Status](https://coveralls.io/repos/github/go-chassis/mesher/badge.svg?branch=master)](https://coveralls.io/github/go-chassis/mesher?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/go-chassis/mesher)](https://goreportcard.com/report/github.com/go-chassis/mesher) [![GoDoc](https://godoc.org/github.com/go-chassis/mesher?status.svg)](https://godoc.org/github.com/go-chassis/mesher) [![HitCount](http://hits.dwyl.io/go-chassis/mesher.svg)](http://hits.dwyl.io/go-chassis/mesher) [![Join Slack](https://img.shields.io/badge/Join-Slack-orange.svg)](https://join.slack.com/t/go-chassis/shared_invite/enQtMzk0MzAyMjEzNzEyLTRjOWE3NzNmN2IzOGZhMzZkZDFjODM1MDc5ZWI0YjcxYjM1ODNkY2RkNmIxZDdlOWI3NmQ0MTg3NzBkNGExZGU)

A service mesh implementation based on [go chassis](https://github.com/ServiceComb/go-chassis).

One big advantage of Mesher is it is able to 
work with go-chassis in same service mesh control plane like Istio, without control plane they can work 
together with ServiceComb Service center.
So if you choose go as your service language you can use go-chassis to gain better performance, and you can freely use 
other programing language which suit your service scenario the most

Mesher support both linux and windows OS, 
that makes possible that .Net service can work with java, go, python service in distributed system easily

# Features
- go-chassis: Mesher has all of features of [go chassis](https://github.com/ServiceComb/go-chassis)
a go micro service framework
- Admin API：Listen on isolated port, let user to query runtime information 


# Get started
Refer to [mesher-examples](https://github.com/go-chassis/mesher-examples)

### How to build and run

1. Install ServiceComb [service-center](https://github.com/ServiceComb/service-center/releases)

2. go build mesher.go

3. ./mesher

# Documentations

https://mesher.readthedocs.io/en/latest/
