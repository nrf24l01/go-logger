# go-logger

Цветной консольный логгер с уровнями и типами событий.

## Запуск демо

```bash
go run ./cmd/demo
```

## Использование

```go
logger := gologger.NewLogger(
	os.Stdout,
	"my-service",
	gologger.WithCustomTypes(gologger.LogType("PAYMENT")),
)
logger.Log(gologger.LevelInfo, gologger.TypeHTTP, "→ 200 /health 2ms", "req-123")
logger.Log(gologger.LevelInfo, gologger.LogType("PAYMENT"), "charge completed", "req-124")
```
