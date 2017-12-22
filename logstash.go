package logspout_amqp_rancher

import (
	"github.com/fsouza/go-dockerclient"
	"os"
	"strings"
)

// Parse the logstash fields env variables
func GetLogstashFields(container *docker.Container, parser *LogstashParser) map[string]string {
	if fields, ok := parser.logstashFields[container.ID]; ok {
		return fields
	}

	fieldsStr := os.Getenv("LOGSTASH_FIELDS")
	fields := map[string]string{}

	for _, e := range container.Config.Env {
		if strings.HasPrefix(e, "LOGSTASH_FIELDS=") {
			fieldsStr = strings.TrimPrefix(e, "LOGSTASH_FIELDS=")
		}
	}

	if len(fieldsStr) > 0 {
		for _, f := range strings.Split(fieldsStr, ",") {
			sp := strings.Split(f, "=")
			k, v := sp[0], sp[1]
			fields[k] = v
		}
	}

	parser.logstashFields[container.ID] = fields

	return fields
}

// NewHTTPAdapter creates an HTTPAdapter
func NewLogstashParser() *LogstashParser {
	return &LogstashParser{
		logstashFields: make(map[string]map[string]string),
	}
}

type LogstashParser struct {
	logstashFields map[string]map[string]string
}
