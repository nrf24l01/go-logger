# go-logger

Цветной консольный логгер с уровнями и типами событий.
![example](/docs/example.png)

## Запуск демо

```bash
go run ./cmd/demo
```

## Использование

```go
logger := gologger.NewLogger(
	os.Stdout,
	"loggodemo",
	gologger.WithCustomTypes(gologger.LogType("PAYMENT")),
)
logger.Log(gologger.LevelInfo, gologger.LevelInfo, "→ 200 /health 2ms", "req-123")
logger.Log(gologger.LevelInfo, gologger.LogType("PAYMENT"), "charge completed", "req-124")
```
