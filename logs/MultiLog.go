package logs

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	gelfClient = gelf.New(gelf.Config{
		GraylogPort:     514,
		GraylogHostname: "localhost",
	})
)

func LogsConfig(port int, hostName string) {
	gelfClient.GraylogPort = port
	gelfClient.GraylogHostname = hostName
}

func Println(a ...interface{}) {
	sprintf := fmt.Sprintf("%v", a)

	if strings.Contains(sprintf, "short_message") {
		if sprintf[0] == '[' {
			sprintf = sprintf[1 : len(sprintf)-1]
		}
	} else {
		sprintf = Create("gelf message").SetFullMessage(fmt.Sprintf("%#v", a)).SetHost("hooostteeee").SetLevel(1).ToJSON()

	}

	fmt.Println(a) // test
	gelfClient.Log(sprintf)
}

type Log struct {
	Version      string `json:"version"`
	Host         string `json:"host"`
	ShortMessage string `json:"short_message"`
	FullMessage  string `json:"full_message"`
	Timestamp    int64  `json:"timestamp"`
	Level        int    `json:"level"`
}

func Create(message string) *Log {
	return &Log{
		Version:      "1.1",
		Host:         "default",
		ShortMessage: message,
		Timestamp:    time.Now().Unix(),
		Level:        1,
	}
}

func (self *Log) SetTimestamp(timestamp int64) *Log {
	self.Timestamp = timestamp
	return self
}

func (self *Log) SetHost(host string) *Log {
	self.Host = host
	return self
}

func (self *Log) SetFullMessage(fullMessage string) *Log {
	self.FullMessage = fullMessage
	return self
}

func (self *Log) SetLevel(level int) *Log {
	self.Level = level
	return self
}

func (self *Log) ToJSON() string {
	message, err := json.Marshal(self)
	if err != nil {
		os.Exit(1)
	}
	return string(message)
}
