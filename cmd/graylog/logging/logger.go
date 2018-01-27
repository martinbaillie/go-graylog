package logging

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/crypto/ssh/terminal"
)

// Logger defines methods to implement for being a logger.
type Logger interface {
	Println(str string)
	String() string
}

// ColouredLogger implements Logger interface with colour support.
type ColouredLogger struct {
	name         string
	colourPrefix string
	factory      *ColouredLoggerFactory
}

// ColouredLoggerFactory is a factory for the (round robined) creation of
// coloured loggers. It is inspired by Docker/Moby's libcompose.
type ColouredLoggerFactory struct {
	maxLength int
	tty       bool
}

var (
	colourPrefix = make(chan string)
)

func generateColours() {
	i := 0
	colourOrder := []string{
		"36",   // cyan
		"33",   // yellow
		"32",   // green
		"35",   // magenta
		"31",   // red
		"34",   // blue
		"36;1", // intense cyan
		"33;1", // intense yellow
		"32;1", // intense green
		"35;1", // intense magenta
		"31;1", // intense red
		"34;1", // intense blue
	}

	for {
		colourPrefix <- fmt.Sprintf("\033[%sm%%s |\033[0m", colourOrder[i])
		i = (i + 1) % len(colourOrder)
	}
}

func init() {
	go generateColours()
}

// NewColouredLoggerFactory creates a new ColouredLoggerFactory.
func NewColouredLoggerFactory() *ColouredLoggerFactory {
	return &ColouredLoggerFactory{
		// Only colour if there's a TTY.
		tty: terminal.IsTerminal(int(os.Stdout.Fd())),
	}
}

// Create returns a new ColouredLogger where the colour is selected in a
// round robin fashion.
func (c *ColouredLoggerFactory) Create(name string) Logger {
	if c.maxLength < len(name) {
		c.maxLength = len(name)
	}

	return &ColouredLogger{
		name:         name,
		factory:      c,
		colourPrefix: <-colourPrefix,
	}
}

// Println implements the Println Logger method for ColouredLogger.
func (c *ColouredLogger) Println(str string) {
	if len(str) == 0 {
		return
	}
	logFmt, name := c.getLogFmt()
	message := fmt.Sprintf(logFmt, name, str)
	fmt.Println(message)
}

func (c *ColouredLogger) String() string {
	return c.name
}

// getLogFmt returns the log format and the name of the logger.
func (c *ColouredLogger) getLogFmt() (string, string) {
	pad := c.factory.maxLength

	logFmt := "%s | %s"
	if c.factory.tty {
		logFmt = c.colourPrefix + " %s"
	}

	name := fmt.Sprintf("%-"+strconv.Itoa(pad)+"s", c.name)

	return logFmt, name
}
