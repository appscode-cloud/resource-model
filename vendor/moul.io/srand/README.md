# srand

:smile: Initialize random seed in Golang.

[![CircleCI](https://circleci.com/gh/moul/srand.svg?style=shield)](https://circleci.com/gh/moul/srand)
[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/moul.io/srand)
[![License](https://img.shields.io/badge/license-Apache--2.0%20%2F%20MIT-%2397ca00.svg)](https://github.com/moul/srand/blob/master/COPYRIGHT)
[![GitHub release](https://img.shields.io/github/release/moul/srand.svg)](https://github.com/moul/srand/releases)
[![Go Report Card](https://goreportcard.com/badge/moul.io/srand)](https://goreportcard.com/report/moul.io/srand)
[![CodeFactor](https://www.codefactor.io/repository/github/moul/srand/badge)](https://www.codefactor.io/repository/github/moul/srand)
[![codecov](https://codecov.io/gh/moul/srand/branch/master/graph/badge.svg)](https://codecov.io/gh/moul/srand)
[![GolangCI](https://golangci.com/badges/github.com/moul/srand.svg)](https://golangci.com/r/github.com/moul/srand)
[![Sourcegraph](https://sourcegraph.com/github.com/moul/srand/-/badge.svg)](https://sourcegraph.com/github.com/moul/srand?badge)
[![Made by Manfred Touron](https://img.shields.io/badge/made%20by-Manfred%20Touron-blue.svg?style=flat)](https://manfred.life/)


## Usage

```golang
import "math/rand"
import "moul.io/srand"

func init() {
    // cryptographically secure initializer
    rand.Seed(srand.Secure())

// simple seed initializer
    rand.Seed(srand.Fast())

    // simple seed initializer overridable by the $SRAND env var
    rand.Seed(srand.Overridable("SRAND"))
}
```

## Install

```console
$ go get -u moul.io/srand
```

## License

© 2019 [Manfred Touron](https://manfred.life)

Licensed under the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0) ([`LICENSE-APACHE`](LICENSE-APACHE)) or the [MIT license](https://opensource.org/licenses/MIT) ([`LICENSE-MIT`](LICENSE-MIT)), at your option. See the [`COPYRIGHT`](COPYRIGHT) file for more details.

`SPDX-License-Identifier: (Apache-2.0 OR MIT)`
