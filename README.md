# Go Fundamentals

My journey learning Go, one folder per topic/exercise.

## Progress

- `helloworld/` — first Go program, basic syntax and `fmt.Println`
- `Cards/` — custom `deck` type ([]string) with methods for creating, printing, dealing, shuffling, and saving/loading a deck to/from file
- `StructsAndPointers/` — `person`/`account` structs, value vs pointer receivers, and struct relationships (a person owns accounts, each account points back to its owner)
- `Maps/` — map fundamentals (nil maps, comma-ok, delete/len, nested maps, sets, reference semantics, comparison) plus a small inventory app built on top
- `Interfaces/` — interface basics (bots example), plus edge cases: shared method signatures, embedding, `any`, type assertions/switches, the nil-interface gotcha, pointer-receiver satisfaction, comparison
- `Http/` — `net/http` client basics: a free, no-API-key weather lookup (geocode + forecast via Open-Meteo), plus `io.Reader`/`io.Writer` experiments against a live HTTP response
- `Assignment1-interfaces/` — `shape` interface satisfied by `square` and `triangle`
- `Assignment2-interfaces/` — CLI program that reads a file (via `os.Args`) and streams it to the terminal with `io.Copy`
- `Goroutines/` — goroutine fundamentals: fire-and-forget gotcha, `sync.WaitGroup`, buffered/unbuffered channels, channel directions, closing/`select`, race conditions vs `sync.Mutex`, a worker pool, and a continuous link-monitor as a real example
- `Context/` — `context.Context`: `WithCancel`/`WithTimeout`/`WithDeadline`, propagating cancellation into a goroutine doing real work, and the production graceful-shutdown pattern (`signal.NotifyContext`)
- `HttpServer/` — a real `net/http` REST API server (`cmd/`+`internal/` layout): routing, middleware chaining (structured logging via `log/slog`, panic recovery, API-key auth), request validation, env-var config, table-driven tests, graceful shutdown, and a tested Dockerfile
