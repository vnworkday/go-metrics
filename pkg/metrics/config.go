package metrics

import "github.com/vnworkday/go-metrics/pkg/tags"

type ConfigTag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Config struct {
	Host string      `json:"host"`
	Port int         `json:"port"`
	Tags []ConfigTag `json:"tags"`
}

var DefaultConfig = Config{
	Host: "localhost",
	Port: 4317,
}

func (c Config) isValid() bool {
	return c.Host != "" && c.Port > 0
}

// tags returns a list of Tag from the config that are valid (non-empty key and value).
func (c Config) tags() []tags.Tag {
	tagList := make([]tags.Tag, 0, len(c.Tags))

	for _, tag := range c.Tags {
		if tag.Key == "" || tag.Value == "" {
			continue
		}

		tagList = append(tagList, tags.NewTag(tag.Key, tag.Value))
	}

	return tagList
}
