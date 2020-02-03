# Goava: Guava ported to Go

![Go build](https://github.com/abc-inc/goava/workflows/Go%20build/badge.svg)

Goava is heavily inspired by [Guava](https://github.com/google/guava),
which is a set of core libraries from Google that includes new collection types
(such as multimap and multiset), immutable collections, a graph library, and
utilities for concurrency, I/O, hashing, primitives, strings, and more!

## Features
- [x] [base/CaseFormat](https://github.com/google/guava/wiki/StringsExplained#caseformat) => [github.com/abc-inc/goava/base/casefmt](https://github.com/abc-inc/goava/tree/master/base/casefmt)
- [x] [base/CharMatcher](https://github.com/google/guava/wiki/StringsExplained#charmatcher) => [github.com/abc-inc/goava/base/runematcher](https://github.com/abc-inc/goava/tree/master/base/runematcher)
- [x] [base/Optional](https://github.com/google/guava/wiki/UsingAndAvoidingNullExplained#optional) => [github.com/abc-inc/goava/base/opt](https://github.com/abc-inc/goava/tree/master/base/opt)
- [x] [base/Preconditions](https://github.com/google/guava/wiki/PreconditionsExplained) => [github.com/abc-inc/goava/base/precond](https://github.com/abc-inc/goava/tree/master/base/precond)
- [x] [base/Stopwatch](https://guava.dev/releases/28.2-jre/api/docs/com/google/common/base/Stopwatch.html) => [github.com/abc-inc/goava/base/stopwatch](https://github.com/abc-inc/goava/tree/master/base/stopwatch)
- [ ] [base/Strings](https://github.com/google/guava/wiki/StringsExplained)
- [ ] [cache/Cache](https://github.com/google/guava/wiki/CachesExplained)
- [x] [collect/ComparisonChain](https://guava.dev/releases/28.2-jre/api/docs/com/google/common/collect/ComparisonChain.html) => [github.com/abc-inc/goava/collect/compchain](https://github.com/abc-inc/goava/tree/master/collect/compchain)
- [x] [collect/DiscreteDomain](https://github.com/google/guava/wiki/RangesExplained#discrete-domains) => [github.com/abc-inc/goava/collect/domain](https://github.com/abc-inc/goava/tree/master/collect/domain)
- [ ] [collect/Ordering](https://github.com/google/guava/wiki/OrderingExplained)
- [x] [collect/Sets](https://github.com/google/guava/wiki/CollectionUtilitiesExplained#sets) => [github.com/abc-inc/goava/collect/set](https://github.com/abc-inc/goava/tree/master/collect/set)
- [ ] [eventbus/EventBus](https://github.com/google/guava/wiki/EventBusExplained)
- [ ] [io/Files](https://github.com/google/guava/wiki/IOExplained#files)
- [ ] [math](https://github.com/google/guava/wiki/MathExplained)
- [x] [net/HostAndPort](https://guava.dev/releases/28.2-jre/api/docs/com/google/common/net/HostAndPort.html) => [github.com/abc-inc/goava/net/hostandport](https://github.com/abc-inc/goava/tree/master/net/hostandport)
- [ ] ...

## Adding Goava to your project
To add a dependency on Goava, install the latest version of the library:

```shell script
go get -u github.com/abc-inc/goava
```

Next, include Goava in your application (see links above):
```go
import "github.com/abc-inc/goava/<package>"
```

## Learn about Goava and Guava

- Guava users' guide, [Guava Explained](https://github.com/google/guava/wiki/Home)

## Links

- [Goava GitHub project](https://github.com/abc-inc/goava)
- [Issue tracker: Report a defect or feature request](https://github.com/abc-inc/goava/issues/new)

## Disclaimer
This is not an official Google product (experimental or otherwise).
It is not affiliated with the [Guava](https://github.com/google/guava) project.
