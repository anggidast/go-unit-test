 # Go Unit Testing

Go has it own built in package for unit test, called `testing`

https://pkg.go.dev/testing

## Package Structs

- ### testing.T

  this struct is main struct used for unit testing in Go

- ### testing.M

  Go provided this struct for lifecycle testing use

- ### testing.B

  struct which used for benchmarking (measure source code program performance/speed)

## Naming Rule

Go has some special rules for unit test

- to make unit test file, the file name should ended with `_test`, example: `hello_world_test.go`

- function name in unit test should started with Test, example: `TestHelloWorld()`, and the function should has testing parameter `(t *testing.T)` and has no return value

## Running Unit Test

To running unit test, use these commands in terminal:

- `go test`

- `go test -v` to run test with function detail

- `go test -v -run TestFunctionName` to test only specific function (using prefix)

## Fail Test Functions

Don't use `panic` when test is fail

- `Fail()` : when test is fail, the function will continue execute other unit test

- `FailNow()` : will stop the test when the test is fail

- `Error()` : same with `Fail()` but it has message

- `Fatal()` : same with `FailNow()` but it has message

## Assertion

Assertion is used to replace if else condition in unit test

Go has no built in package for assertion in unit test

One of most popular assertion library is Testify which has `assert` and `require` package

https://github.com/stretchr/testify

When using assert from testify, error handling in unit test doesn't need to define, because `assert` already include `Fail()` and `require` include `FailNow()`.

## Skip Test

Skip test used to ignore the test when some condition

```go
if some_condition {
  t.Skip("Unit test can't run in this condition")
}

// test logic start here
```

## Lifecycle in Test

Lifecycle allow developers to insert some logic before and after run the test.

Lifecylce in test only run in `TestMain(m *testing.M)` function

Go will execute `TestMain` automatically in every test run in a package. `TestMain` function only executed once in every Go packages, not in every unit test functions.

```go
func TestMain(m *testing.M) {
  // do something before the test run
  fmt.Println("before unit test")

  m.Run() // this line will execute all unit tests in the package

  // do something after the test run
  fmt.Println("after unit test")
}
```

## Sub Test

Unit testing in Go is support to make unit test function in unit test function

to run specific sub test, use this command in terminal:

`go test -run TestFunctionName/SubTestName`

## Table Test

Iterate the unit test to run some test with dynamic sub test, parameter and output from array





