package ruslog

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/Sirupsen/logrus"
)

const (
	// formatter types
	SIMPLE = "Simple"
	JSON   = "Json"
	TEXT   = "Text"
)

type (
	Formatter struct {
		Name      string
		Formatter logrus.Formatter
	}

	SimpleFormatter struct{}
)

func (f *SimpleFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	b := &bytes.Buffer{}

	fmt.Fprintf(b, "[%s]", entry.Time.String())
	fmt.Fprintf(b, " [%s]", strings.ToUpper(entry.Level.String()))
	fmt.Fprintf(b, " %s", entry.Message)

	for key, value := range entry.Data {
		if key != "time" && key != "level" && key != "msg" {
			fmt.Fprintf(b, " %s=%s", key, value)
		}
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}
