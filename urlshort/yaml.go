package urlshort

import (
	"gopkg.in/yaml.v3"
	"log"
)

type RedirectLink struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func parseYaml(b []byte) map[string]string {
	var r []RedirectLink
	err := yaml.Unmarshal(b, &r)
	if err != nil {
		log.Fatal("failed to unmarshal YAML\n")
	}

	m := make(map[string]string)
	for _, rl := range r {
		m[rl.Path] = rl.URL
	}
	return m
}
