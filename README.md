# Go Slices

## Introduction

The `slices` package provides a set of utility functions to work with slices and maps in Go. The functions in this
package can be used to perform common operations such as mapping, filtering, and finding elements in slices and maps.

## Installation

To install the package, use the following command:<br>
```go get github.com/SergeySlonimsky/slices```<br>

Then import the package in your Go file:<br>
```go import "github.com/SergeySlonimsky/slices"```

## Usage

### Array functions

```ArrayMap(arr []I, callback func(key int, value I) T) []T```<br>
ArrayMap creates a new array populated with the results of calling a provided function on every element in the calling
array.

```ArrayForEach(arr []I, callback func(key int, value I))```<br>
ArrayForEach executes a provided function once for each array element.

```ArrayFilter(arr []I, callback func(key int, value I) bool) []I```<br>
ArrayFilter creates a copy of a given array, filtered down to just the elements from the given array that pass the test
implemented by the provided function.

```ArrayConcat(arr ...[]I) []I```<br>
ArrayConcat creates a new array with all the elements of the provided arrays.

```ArrayEvery(arr []I, callback func(value I) bool) bool```<br>
ArrayEvery tests whether all elements in the array pass the test implemented by the provided function.

```ArrayUniq(arr []I) []I```<br>
ArrayUniq creates a new array with all unique values from provided array. Provided array should contain comparable
values.

```ArrayHashUniq(arr []I, hashFunc func(value I) string) []I```<br>
ArrayHashUniq creates a new array with all unique values comparable by provided hashFunc. HashFunc should return unique
string for every unique element.

```ArrayFind(arr []I, callback func(key int, value I) bool) (I, bool)```<br>
ArrayFind returns the first element in the provided array that satisfies the provided testing function. If no values
satisfy the testing function, empty value and false are returned.

```ArrayFindIndex(arr []I, callback func(key int, value I) bool) (int, bool)```<br>
ArrayFindIndex returns the index of the first element in the provided array that satisfies the provided testing
function. If no values satisfy the testing function, -1 and false are returned.

### Map functions

```MapWalk(arr map[I]K, callback func(key I, value K) T) []T```<br>
MapWalk creates a new array populated with the results of calling a provided function on every element in the calling
map.

```MapForEach(arr map[I]K, callback func(key I, value K))```<br>
MapForEach executes a provided function once for each map element.

```MapFilter[I comparable, K any](arr map[I]K, callback func(key I, value K) bool) map[I]K```<br>
MapFilter creates a new map with all the elements from the original map that pass the test implemented by the provided
function.

```MapKeys(arr map[I]K) []I```<br>
MapKeys creates a new slice with all the keys from the original map.

```MapValues(arr map[I]K) []K```<br>
MapValues creates a new slice with all the values from the original map.

`MapWalk, MapForEach, MapFilter, MapKeys, MapValues` functions<br>
These functions have the same behavior as Array functions, but works with maps.

## License
This package is released under the <a href="https://opensource.org/licenses/MIT">MIT License</a>.
