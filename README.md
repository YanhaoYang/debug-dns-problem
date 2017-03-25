# Debug a DNS problem with Go 1.7.5 in Docker container

In docker-compose.yml, three containers are defined to run the code in src/main.go.
In container `dns-test-175x` and `dns-test-18`, the code works fine, but in
container `dns-test-175`, it complains that it cannot connect to Elasticsearch:

    panic: no Elasticsearch node available

    goroutine 1 [running]:
    panic(0x653d80, 0xc42000d280)
      /usr/local/go/src/runtime/panic.go:500 +0x1a1
    main.main()
      /go/src/github.com/YanhaoYang/debug-dns-problem/main.go:17 +0x206
    exit status 2
