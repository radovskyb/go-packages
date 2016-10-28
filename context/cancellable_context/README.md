This example demonstrates how to use cancellable contexts to prevent goroutine leaks.

To see that all of the goroutines are cancelled correctly, run with:
```bash
go run -race main.go
```
