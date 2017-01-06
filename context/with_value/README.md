This example demonstrates how to use `context.WithValue` to pass data between http handlers through request context's for middleware functionality.

`context.WithValue` should only be used when actually necessary and shouldn't just be used to pass anything through http handlers, especially things such as database connections or session values, etc.

To test the example.
```bash
go run main.go
```

Then try these different URL's for different results:

- http://localhost:9000 <- `username is not set`
- http://localhost:9000/username/radovskyb <- `Hello, RADOVSKYB`
- http://localhost:9000/username/admin <- `Welcome Administrator`
