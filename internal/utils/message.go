package utils

import (
	"fmt"

	aurora "github.com/logrusorgru/aurora/v3"
)

type messager struct {
	debugEnabled bool
}

func compose(manyv ...func(arg interface{}) aurora.Value) func(arg interface{}) aurora.Value {
	return func(arg interface{}) aurora.Value {
		for _, v := range manyv {
			arg = v(arg)
		}
		return arg.(aurora.Value)
	}
}

var (
	info           = aurora.BrightBlue
	warn           = aurora.BrightYellow
	err            = aurora.BrightRed
	defaultMessage = &messager{
		debugEnabled: true,
	}
)

// Info prints an info message
func Info(msg string, args ...any) {
	fmt.Printf(info("[info] "+msg).String()+"\n", args...)
}

// Warning prints a warning message
func Warning(msg string, args ...any) {
	fmt.Printf(warn("[warn] "+msg).String()+"\n", args...)
}

// Error prints an error message
func Error(msg string, args ...any) {
	fmt.Printf(err("[error] "+msg).String()+"\n", args...)
}

// Debug prints an warning message
func Debug(msg string, args ...any) {
	if defaultMessage.debugEnabled {
		fmt.Printf("[debug] "+msg+"\n", args...)
	}
}

// DisableDebug disables the debug messages
func DisableDebug() {
	defaultMessage.debugEnabled = false
}
