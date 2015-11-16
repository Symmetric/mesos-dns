# Mesos-DNS [![Circle CI](https://circleci.com/gh/mesosphere/mesos-dns.svg?style=svg)](https://circleci.com/gh/mesosphere/mesos-dns) [![Coverage Status](https://coveralls.io/repos/mesosphere/mesos-dns/badge.svg?branch=master&service=github)](https://coveralls.io/github/mesosphere/mesos-dns?branch=master) [![GoDoc](https://godoc.org/github.com/mesosphere/mesos-dns?status.svg)](https://godoc.org/github.com/mesosphere/mesos-dns) [![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/mesosphere/mesos-dns?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)
Mesos-DNS enables [DNS](http://en.wikipedia.org/wiki/Domain_Name_System) based service discovery in [Apache Mesos](http://mesos.apache.org/) clusters.

![Architecture
Diagram](http://mesosphere.github.io/mesos-dns/img/architecture.png)

## Status
This project is undergoing heavy refactoring, hardening and testing and
is considered **alpha**. We have adopted a [semantic versioning](http://semver.org/) scheme which clearly communicates its status and progression over time.

## Compatibility
`mesos-N` tags mark the start of support for a specific Mesos version while
maintaining backwards compatibility with the previous major version. 
This release breaks compatibility with handling NetworkInfo on versions
before 0.26. It is compatible with Mesos 0.26 and newer.

## Installing
The official distribution and installation channel is pre-compiled binaries available in [Github releases](https://github.com/mesosphere/mesos-dns/releases).

## Building
Building the **master** branch from source should always succeed but doesn't provide
the same stability and compatibility guarantees as releases.

All branches and pull requests are tested by [Circle-CI](https://circleci.com/gh/mesosphere/mesos-dns), which also
outputs artifacts for Mac OS X, Windows, and Linux via cross-compilation.

You will need [Go](https://golang.org/) *1.5* or later to build the project.
All dependencies are vendored using `Godeps`. You must first install it in order to build from source.

```shell
$ go get github.com/tools/godep
$ godep go build ./...
```

## Testing
```shell
$ godep go test -race ./...
```

## Documentation
Detailed documentation on how to configure, operate and use Mesos-DNS
under different scenarios and environments is available in http://mesosphere.github.io/mesos-dns/.

## Contributing
Contributions are welcome. Please refer to [CONTRIBUTING.md](CONTRIBUTING.md) for
guidelines.

## Contact
For any discussion that isn't well suited for Github [issues](https://github.com/mesosphere/mesos-dns/issues),
please use our [mailing list](https://groups.google.com/forum/#!forum/mesos-dns) or our public [chat room](https://gitter.im/mesosphere/mesos-dns).

## License
This project is [Apache License 2.0](LICENSE).
