package logging

import (
	"fmt"
	"sort"
	"time"
)

const (
	// Format used when printing timestamps.
	// This is not configurable at the moment.
	defaultTsFmt = "2006-01-02 15:04:05.000"
)

// LoggableMessages is a slice of LoggableMessage.
type LoggableMessages []*LoggableMessage

// Merge flattens multiple LoggableMessages into one.
func (lms LoggableMessages) Merge(lmses ...LoggableMessages) LoggableMessages {
	var (
		current = len(lms)
		total   int
	)
	for _, m := range lmses {
		total += len(m)
	}

	// Nothing to merge.
	if total == 0 {
		return lms
	}

	// Grow lms if we have to.
	if cap(lms) < current+total {
		newLms := make(LoggableMessages, current+total)
		copy(newLms, lms)
		lms = newLms
	}

	// Merge.
	for _, m := range lmses {
		current += copy(lms[current:], m)
	}

	return lms
}

// Sort arranges the LoggableMessages slice by ascending date.
func (lms LoggableMessages) Sort() LoggableMessages {
	sort.Slice(lms, func(i, j int) bool {
		return msgTs(lms[i]).Before(msgTs(lms[j]))
	})

	return lms
}

// LoggableMessage models a Graylog message that can be logged to the terminal.
type LoggableMessage struct {
	Logger  Logger
	Message map[string]interface{}
}

// Log will log the message to the terminal. The `timestamps` parameter prints
// timestamps as well as the message. The `utc` parameter can be used to switch
// to UTC (User local timezone is used otherwise).
func (lm *LoggableMessage) Log(timestamps, utc bool) {
	if lm == nil {
		return
	}

	var m string
	if timestamps {
		var ts string
		if utc {
			ts = msgTs(lm).UTC().Format(defaultTsFmt)
		} else {
			ts = msgTs(lm).Local().Format(defaultTsFmt)
		}

		m = fmt.Sprintf("%s %s", ts, lm.Message["message"])
	} else {
		m = lm.Message["message"].(string)
	}

	lm.Logger.Println(m)
}

// msgTs returns the timestamp found within the Graylog message.
func msgTs(lm *LoggableMessage) (ts time.Time) {
	if lm != nil {
		if tsStr, ok := lm.Message["timestamp"].(string); ok {
			ts, _ = time.Parse(time.RFC3339, tsStr)
		}
	}
	return ts
}
