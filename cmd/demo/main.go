package main

import (
	"os"

	gologger "github.com/nrf24l01/go-logger"
)

func main() {
	// дополнительные типы для теста (в "избранное")
	customTypes := []gologger.LogType{
		gologger.LogType("PAYMENT"),
		gologger.LogType("BILLING"),
		gologger.LogType("NOTIFICATION"),
		gologger.LogType("WEBHOOK"),
		gologger.LogType("INVENTORY"),
	}

	logger := gologger.NewLogger(
		os.Stdout,
		"my-service",
		gologger.WithCustomTypes(customTypes...),
		gologger.WithTypeColors(map[gologger.LogType]string{
			gologger.LogType("PAYMENT"): gologger.BgMagenta,
			gologger.LogType("BILLING"): gologger.BgGreen,
			gologger.LogType("WEBHOOK"): gologger.BgBlue,
		}),
	)

	logger.Log(gologger.LevelInfo, gologger.LogType("PAYMENT"), "charge completed", "req-124")

	logger.Log(gologger.LevelSuccess, gologger.TypeAuth, "New user registered: user.ds1@prodcontest.ru (9f444a9a) as USER", "8c62d745")
	logger.Log(gologger.LevelInfo, gologger.TypeHTTP, "→ 201 /api/v1/fraud/register 56ms 533B", "8c62d745")
	logger.Log(gologger.LevelInfo, gologger.TypeDB, "SELECT by name on fraud_rules → 1 row in 1ms", "67e2b6sa")
	logger.Log(gologger.LevelInfo, gologger.TypeCache, "INSERT rules fraud_rules updated 1 row in 6ms", "67e2b6sa")
	logger.Log(gologger.LevelSuccess, gologger.TypeInfo, "Created rule \"High amount\" (38e59185) priority=10 ENABLED", "67e2b6sa")
	logger.Log(gologger.LevelInfo, gologger.TypeHTTP, "→ 201 /api/v1/fraud-rules 8ms 116B", "bc7e6f77")
	logger.Log(gologger.LevelInfo, gologger.TypeDB, "SELECT by name on fraud_rules → 116 rows in 0ms", "bc7ebf77")
	logger.Log(gologger.LevelSuccess, gologger.TypeInfo, "Created rule \"RUB only\" (d5d36decd) priority=20 ENABLED", "bc7ebf77")

	logger.Log(gologger.LevelDebug, gologger.TypeDebug, "Debug: evaluated rule set 3 -> matched 0", "aa11bb22")
	logger.Log(gologger.LevelTrace, gologger.TypeHTTP, "Trace: headers={Authorization: Bearer ...}", "aa11bb22")
	logger.Log(gologger.LevelWarn, gologger.TypeAuth, "Warning: login attempts threshold approaching for user X", "dd33ee44")
	logger.Log(gologger.LevelOk, gologger.TypeInfo, "OK: background job finished successfully", "ff55gg66")

	logger.Log(gologger.LevelError, gologger.TypeError, "Failed to validate token: expired", "a1b2c3d4")
	logger.Log(gologger.LevelFatal, gologger.TypeDB, "Fatal: cannot connect to primary DB, shutting down", "z9y8x7w6")

	logger.Log(gologger.LevelInfo, gologger.TypeOther, "Processed event with unknown type", "zz11yy22")
}
