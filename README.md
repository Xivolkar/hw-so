# Hello World with Go Plugins

This repo has a sample on how Go plugins can be used to accomplish a hotswap of functionality.

# Use

1. Clone the repository
2. Open a terminal, switch to the cloned directory and type in -  `go run tester.go` to begin.
3. You can drop in a .so file (which has a `Machine` symbol that implements `DoSomething`) to the `libFolder` directory 
4. The test utility which watches the `libFolder` loads the plugin and executes the `DoSomething` function
