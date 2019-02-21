initialize variable at runtime
```
$ go build -ldflags "-X main.str=123 -X main.str2=abc"  main.go
```

## Reference

* Golang : Setting variable value with ldflags : https://www.socketloop.com/tutorials/golang-setting-variable-value-with-ldflags/?utm_source=socketloop&utm_medium=search
