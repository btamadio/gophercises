package urlshort

import (
	"gopkg.in/yaml.v3"
	"log"
)

type RedirectLink struct{
	Path string
	URL string
}

func parseYaml(b []byte) map[string]string{
	r := []RedirectLink{}
	err := yaml.Unmarshal(b, &r)
	if err != nil{
		log.Fatal("failed to unmarshal YAML\n")
	}

	m := make(map[string]string)
	for _, rl := range r{
		m[rl.Path] = rl.URL
	}
	return m
}
