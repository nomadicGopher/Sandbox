* `go test .`
    * To run specific tests, use `go test -run TestName`
    * To include additional details, add `-v` (for verbose)
    * To skip long running tests, add `-short`
    * To incude benchmarks, add `-bench`
    * To set a timeout duration, add `-timeout`
    * To include code coverage, add `-cover`
        * To output to an html file (in addition to std.Out), instead use `-coverprofile="CoverageReport.html"`
            * To open the html report in a browser, run `go tool cover -html="CoverageReport.html"`
* **Error functions**
    * **t.Error(args ...interface{})** - Fails with a message
    * **t.Errorf("", args ...interface{})** - Fails with a formatted message
    * **t.Fail()** - Fails with no message
    * **t.FailNow()** - Fails with no message & stops the test immediately
    * **t.Log(args ...interface{})** - Logs a message to the test output
    * **t.Fatalf("", args ...interface{})**  - Fails with a formatted message & stops the test execution
* **t.Parallel()** allows tests to run in parallel, and therefor must be nested within a sub-test (insise the test.Run(string, **func()**))
    * When doing this, it is good to instanciate a copy of the testCase since in ranges, the single variable is technically resues when looping over the range.  
* When working with a struct where you will expect specific values, you can create a **Pointer Reciever** like **.Equal()** to directly compare the structs & their values
* **reflect.DeepEqual()**: Used to compare two values deeply, including struct hierarchies, slices, maps, arrays, and interfaces. It compares all fields
recursively, including nested structs.
    * It may return false if the types are the same but not exactly identical in structure (e.g., nil slice vs. empty slice).
    * It’s not always the most efficient for performance-critical code.
* **Mocking**
    * **httptest** (standard package for API requests)
        * [**Unit testing HTTP handlers(Advanced Demo)**](https://www.linkedin.com/learning/unit-testing-in-go/unit-testing-http-handlers?autoplay=true&resume=false&u=2148298)
    