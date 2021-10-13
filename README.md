# Dependency injection in Go for unit tests

Here's a quick demo/tutorial for proper dependency injection in order to enable quick unit tests. This ia based off the incredible ebook ["Everyday Golang"](https://openfaas.gumroad.com/l/everyday-golang) by [Alex Ellis](https://blog.alexellis.io/i-wrote-a-book-about-golang/).

Note how the _encoding/json_ dependency is abstracted in order to test the _utils.GetAstronauts_ method, which in the main source code calls a public API.