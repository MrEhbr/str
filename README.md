# str

[![Go](https://github.com/MrEhbr/str/actions/workflows/go.yml/badge.svg)](https://github.com/MrEhbr/str/actions/workflows/go.yml)
[![License](https://img.shields.io/badge/license-Apache--2.0%20%2F%20MIT-%2397ca00.svg)](https://github.com/MrEhbr/str/blob/master/COPYRIGHT)
[![codecov](https://codecov.io/gh/MrEhbr/str/branch/master/graph/badge.svg)](https://codecov.io/gh/MrEhbr/str)
![Made by Alexey Burmistrov](https://img.shields.io/badge/made%20by-Alexey%20Burmistrov-blue.svg?style=flat)

## Install

```console
go get -u github.com/MrEhbr/str
```

### Usage

```go
import "github.com/MrEhbr/str"
```

## str.ToCamelCase

Transform string to camel case

```go

str.ToCamelCase("foo_bar", true) // FooBar
```

## str.ToSnackCase

Transform string to snack case

```go

str.ToSnackCase("Foo_bar", false) // foo_bar
```

## License

© 2022 [Alexey Burmistrov]

Licensed under the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0) ([`LICENSE`](LICENSE)). See the [`COPYRIGHT`](COPYRIGHT) file for more details.

`SPDX-License-Identifier: Apache-2.0`
