package main

import (
	"log"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	logger := log.New(os.Stdout, "", 0)
	for {
		sleepDuration := time.Duration(randInt(1, 20)) * time.Second
		logInfof(logger, "==> Sleeping for %s", sleepDuration.String())
		time.Sleep(sleepDuration)
	}
}

func logInfof(l *log.Logger, msg string, args ...interface{}) {
	l.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + " ")
	l.Printf(msg, args...)
}

func randInt(min, max int) int {
	return rand.Intn(max-min) + min
}
