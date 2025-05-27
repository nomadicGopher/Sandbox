* reflect.DeepEqual(): Used to compare two values deeply, including struct hierarchies, slices, maps, arrays, and interfaces. It compares all fields
recursively, including nested structs.
    * It may return false if the types are the same but not exactly identical in structure (e.g., nil slice vs. empty slice).
    * t’s not always the most efficient for performance-critical code.
* Use `go run .` to run all files in current dir
    * Add `go run -race .` to leverage a race detector
* **Testing**
    * `go test .`
        * To run specific tests, use `go test -run TestName`
        * To include additional details, add `-v` (for verbose)
        * To skip long running tests, add `-short`
        * To incude benchmarks, add `-bench`
        * To set a timeout duration, add `-timeout`
        * To include code coverage, add `-cover`
            * To output to an html file (in addition to std.Out), instead use `-coverprofile="CoverageReport.html"`
                * To open the html report in a browser, run `go tool cover -html="CoverageReport.html"`
    * run.Parallel() allows tests to run in parallel, and therefor must be nested within a sub-test (insise the test.Run(string, **func()**))
        * When doing this, it is good to instanciate a copy of the testCase since in ranges, the single variable is technically resues when looping over the range.  
        ```go
        // FileName == something_test.go
        // Multiple test cases
        func TestSomething(test *testing.T) {
            testCases := []struct {
                Description    string
                Input1         int
                Input2         int
                ExpectedOutput int
            }{
                //{"Fail first", 1, 1, 99},
                {"Sample pass", 2, 2, 4},
            }

            for _, testCase := range testCases {
                testCase := testCase // Re-initilaize to prevent race condition, due to parallel

                test.Run(testCase.Description, func(test *testing.T) {
                    test.Parallel()

                    // Sample function found in main.go. Using a function isn't necessary, nor is setting input variables since they are in scope. Just for demonstration.
                    actualOutput := func(input1, input2 int) int {
                        return input1 + input2
                    }(testCase.Input1, testCase.Input2)

                    if actualOutput != testCase.ExpectedOutput {
                        test.Errorf("%s: %d + %d = %d; want %d.", testCase.Description, testCase.Input1, testCase.Input2, actualOutput, testCase.ExpectedOutput)
                    }
                })
            }
        }

        // Singular test
        func TestSomethingAlternative(test *testing.T) {
            input := 2
            // expectedOutput := 99 // Fail first
            expectedOutput := 4
            actualOutput := func(input int) int {
                return input + input
            }(input)

            if actualOutput != expectedOutput {
                test.Errorf("Test something: %d + %d = %d; want %d.", input, input, actualOutput, expectedOutput)
            }
        }
        ```
    * **Error functions**
        * **t.Error(args ...interface{})** - Fails with a message
        * **t.Errorf("", args ...interface{})** - Fails with a formatted message
        * **t.Fail()** - Fails with no message
        * **t.FailNow()** - Fails with no message & stops the test immediately
        * **t.Log(args ...interface{})** - Logs a message to the test output
        * **t.Fatalf("", args ...interface{})**  - Fails with a formatted message & stops the test execution
    * When working with a struct where you will expect specific values, you can create a Pointer Reciever like Equal() to directly compare the structs & their values
        ```go
        type restaurant struct {
            Name    string
            Cuisine string
        }

        func (r restaurant) Equal(other restaurant) bool {
            return r.Name == other.Name && r.Cuisine == other.Cuisine
        }

        func TestRestaurantWithCustomEquals(t *testing.T) {
            expected := restaurant{Name: "Joes Pizza", Cuisine: "Italian"}
            actual := restaurant{Name: "Mama Putt", Cuisine: "Indian"}

            if !expected.Equal(actual) {
                t.Errorf("Expected %v, but got %v", expected, actual)
            }
        }
        ```


    