package elog

import (
	"fmt"
	"log"
)

type ColorCode string

const (
	red   ColorCode = "1;31"
	green           = "1;32"
	blue            = "1;34"
	white           = "1;37"
)

func colorize(raw string, color ColorCode) string {
	return fmt.Sprintf("%s%s%s", "\033["+color+"m", raw, "\033[0m")
}

func Debug(format string, v ...interface{}) {
	format = fmt.Sprintf("[D] %q\n", format)
	format = colorize(format, blue)
	if v != nil {
		log.Printf(format, v...)
	} else {
		log.Print(format)
	}
}

func Info(format string, v ...interface{}) {
	format = fmt.Sprintf("[I] %q\n", format)
	format = colorize(format, green)
	if v != nil {
		log.Printf(format, v...)
	} else {
		log.Print(format)
	}
}

func Error(format string, v ...interface{}) {
	format = fmt.Sprintf("[E] %q\n", format)
	format = colorize(format, red)
	if v != nil {
		log.Printf(format, v...)
	} else {
		log.Print(format)
	}

}

func Fatal(format string, v ...interface{}) {
	format = fmt.Sprintf("[F] %q\n", format)
	format = colorize(format, red)
	if v != nil {
		log.Fatalf(format, v...)
	} else {
		log.Fatal(format)
	}
}
