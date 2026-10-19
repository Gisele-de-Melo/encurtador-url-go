package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	env string
}

func New(env string) *Logger {
	return &Logger{env: env}
}

func (l *Logger) Info(msg string, fields ...interface{}) {
	log.Println(l.format("INFO", msg, fields...))
}

func (l *Logger) Error(msg string, fields ...interface{}) {
	log.Println(l.format("ERROR", msg, fields...))
}

func (l *Logger) Fatal(msg string, fields ...interface{}) {
	log.Println(l.format("FATAL", msg, fields...))
	os.Exit(1)
}

func (l *Logger) format(level, msg string, fields ...interface{}) string {
	// Ambiente de desenvolvimento: logs mais legíveis
	if l.env == "development" {
		return formatDev(level, msg, fields...)
	}

	// Produção: logs mais estruturados
	return formatProd(level, msg, fields...)
}

func formatDev(level, msg string, fields ...interface{}) string {
	out := "[" + level + "] " + msg
	for i := 0; i < len(fields); i += 2 {
		if i+1 < len(fields) {
			out += " | " + fields[i].(string) + "=" + toString(fields[i+1])
		}
	}
	return out
}

func formatProd(level, msg string, fields ...interface{}) string {
	out := `{"level":"` + level + `","message":"` + msg + `"`
	for i := 0; i < len(fields); i += 2 {
		if i+1 < len(fields) {
			out += `,"` + fields[i].(string) + `":"` + toString(fields[i+1]) + `"`
		}
	}
	out += "}"
	return out
}

func toString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}
