# Versioning policy

We are following [Semantic versioning](https://semver.org/) in this library

We will mark these with Git Tags

# Changes in releases

## release 2.0.0

Breaking:
- Module path changed to `github.com/keytiles/lib-utils-golang/v2`
- Upgraded sets dependency to `github.com/keytiles/lib-sets-golang/v2` (`pkg/kt_sets`)
  - `MapValuesSet` and `StringExtractVariableNames` now return `*kt_sets.Set[...]` (pointers), matching lib-sets v2 constructors

New features:
- Applying new Keytiles lib standards
  - Introducing constant `LIB_NAME`
  - Based on the above introducing constants `PACKAGE_NAME` in all packages - used as a prefix for kt_errors.Fault sources and Logging

Upgrades:
- Golang 1.26.0 is used from now
- `github.com/keytiles/lib-sets-golang/v2` v2.0.1


## release 1.2.0

New features:

- Adding some simple but often useful map helpers: `MapKeys()` and `MapValuesSet()`

## release 1.1.0

New features:

- A new function `ValueFromPtr()` is added to simplify boilerplates. Pair of `Ptr()` function which - by the way - is deprecated from now. See method docs
  for reason.

## release 1.0.0

Initial release
