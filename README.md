My solutions to the exercises from the book [Learn Concurrent Programming With Go](https://www.manning.com/books/learn-concurrent-programming-with-go).

[![](https://github.com/asarkar/concurrent-programming-with-go/workflows/CI/badge.svg)](https://github.com/asarkar/concurrent-programming-with-go/actions)

Official GitHub repo: https://github.com/cutajarj/ConcurrentProgrammingWithGo

## Contents

2. [Dealing with threads](ch02)
3. [Thread communication using memory sharing](ch03)
4. [Synchronization with mutexes](ch04)

## Development

To avoid `go` directory created under `$HOME`, run:
```
go env -w GOPATH="$HOME/.local/share/go"
```

Test and lint:
```
./.github/run.sh <directory>
```

## License

Released under [Apache License v2.0](LICENSE).