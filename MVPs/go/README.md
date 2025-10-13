* **go.dev**
  * [**Learn**](https://go.dev/learn)
  * [**Docs**](https://go.dev/doc)
* [**GoByExample**](https://gobyexample.com/)
* [**maxence-charriere/go-app**](https://github.com/maxence-charriere/go-app)
  * [**goapp-github-pages**](https://github.com/maxence-charriere/goapp-github-pages)
  * [**go-app-demo**](https://github.com/maxence-charriere/go-app-demo)

## Notes
* bufio.NewScanner(): More practical for simple transactions like single line. Returns variable after running if scanner.Scan() followed by scanner.Text() or scanner.Bytes().
* bufio.NewReader(): More practical for complex transactions like reading files, reponses. Returns variable directly from function.
* Use `go run .` to run all files in current dir
    * Add `go run -race .` to leverage a race detector
* **Variadic functions (parameters)** like sample(nums ...int) allow functions to accept any number of arguments of a specified type. You can call such functions with individual values (e.g., sample(item1, item2)) or by expanding a slice using ... (e.g., sample(slice...)). The ... is necessary to satisfy the variadic parameter when passing a slice. This is different from when the function is expecting a slice, like sample(nums []int), where you simply pass the slice (e.g., sample(slice)).