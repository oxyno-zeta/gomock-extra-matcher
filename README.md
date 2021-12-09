<h1 align="center">gomock-extra-matcher</h1>

<p align="center">
  <a href="http://godoc.org/github.com/oxyno-zeta/gomock-extra-matcher" rel="noopener noreferer" target="_blank"><img src="https://img.shields.io/badge/godoc-reference-blue.svg" alt="Go Doc" /></a>
  <a href="https://circleci.com/gh/oxyno-zeta/gomock-extra-matcher" rel="noopener noreferer" target="_blank"><img src="https://circleci.com/gh/oxyno-zeta/gomock-extra-matcher.svg?style=svg" alt="CircleCI" /></a>
  <a href="https://goreportcard.com/report/github.com/oxyno-zeta/gomock-extra-matcher" rel="noopener noreferer" target="_blank"><img src="https://goreportcard.com/badge/github.com/oxyno-zeta/gomock-extra-matcher" alt="Go Report Card" /></a>
</p>
<p align="center">
  <a href="https://coveralls.io/github/oxyno-zeta/gomock-extra-matcher?branch=master" rel="noopener noreferer" target="_blank"><img src="https://coveralls.io/repos/github/oxyno-zeta/gomock-extra-matcher/badge.svg?branch=master" alt="Coverage Status" /></a>
  <a href="https://github.com/oxyno-zeta/gomock-extra-matcher/blob/master/LICENSE" rel="noopener noreferer" target="_blank"><img src="https://img.shields.io/github/license/oxyno-zeta/gomock-extra-matcher" alt="GitHub license" /></a>
  <a href="https://github.com/oxyno-zeta/gomock-extra-matcher/releases" rel="noopener noreferer" target="_blank"><img src="https://img.shields.io/github/v/release/oxyno-zeta/gomock-extra-matcher" alt="GitHub release (latest by date)" /></a>
</p>

---

## Menu

- [Menu](#menu)
- [Why ?](#why-)
- [How to use ?](#how-to-use-)
- [Matchers](#matchers)
  - [IntRangeMatcher](#intrangematcher)
    - [Explanation](#explanation)
    - [Example](#example)
  - [StringRegexpMatcher](#stringregexpmatcher)
    - [Explanation](#explanation-1)
    - [Example](#example-1)
  - [MapMatcher](#mapmatcher)
    - [Explanation](#explanation-2)
    - [Example](#example-2)
  - [StructMatcher](#structmatcher)
    - [Explanation](#explanation-3)
    - [Example](#example-3)
- [Thanks](#thanks)
- [Author](#author)
- [License](#license)

## Why ?

I've created this library because I cannot find such a library like this one and I don't want to create specific matcher for each structure, each map, ... in all tests that I have and can write.

Moreover, creating specific matcher need maintenance code and this is time consumption.

I'm sure that I'm not the only one that what to avoid all of this :) .

## How to use ?

Import the library in your tests files:

```go
import "github.com/oxyno-zeta/gomock-extra-matcher"
```

Use it in your gomock instance:

```go
mock.EXPECT().DoSomething(extra.StringRegexpMatcher(`^[a-z]+\[[0-9]+\]$`))
```

## Matchers

### IntRangeMatcher

#### Explanation

This matcher will allow to test that an int is inside a range. Here it is considered that input can be equal to the lower bound and the same for the upper bound.

#### Example

Here is an example of usage:

```go
mock.EXPECT().DoSomething(extra.IntRangeMatcher(lowerBound, upperBound))
```

### StringRegexpMatcher

#### Explanation

This matcher will allow to test that a string is validating a Regexp.

#### Example

Here is an example of usage:

```go
mock.EXPECT().DoSomething(extra.StringRegexpMatcher(`^[a-z]+\[[0-9]+\]$`))
```

### MapMatcher

#### Explanation

This matcher will allow to test map key and map value.

To this one, it is possible to give a gomock matcher to a key and also to the value.

The `Key` function is chainable. It is also possible to more than one test per key.

#### Example

Here are some example of usage:

```go
// Here we consider a map[string]string as input
mock.EXPECT().DoSomething(extra.MapMatcher().Key("key1", "value1").Key(gomock.Any(), "value1").Key("key1", gomock.Not("value2")))
```

### StructMatcher

#### Explanation

This matcher will allow to test public fields of a structure (Only public ones. Reflect can't manage private fields...).

To this matcher, it is possible to give either a gomock matcher or either a real value for validation.

The `Field` function is chainable. It is also possible to more than one test per field.

#### Example

```go
// Here we consider a struct as this one
/*
type Fake struct {
    Name string
    Data map[string]string
}
*/
mock.EXPECT().DoSomething(extra.StructMatcher().Field("Name", "value1").Field("Data", gomock.Eq(map[string]string{"fake":"value"})))
```

## Thanks

- My wife BH to support me doing this

## Author

- Oxyno-zeta (Havrileck Alexandre)

## License

Apache 2.0 (See in LICENSE)
