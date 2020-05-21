# coderr
[![GoDoc](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/foxmeder/coderr?tab=doc)

golang error with error code

# Example

```go
err := coderr.Error(30,"error with code 30")

code := coderr.Code(err) // 30
```

