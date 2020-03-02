package urlshort

import (
	"gopkg.in/yaml.v3"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url, ok := pathsToUrls[r.URL.Path]
		if ok {
			http.Redirect(w, r, url, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	yams, err := parseYaml(yml)
	if err != nil {
		return nil, err
	}
	m := makeMap(yams)
	return MapHandler(m, fallback), nil
}

type RedirectLink struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func parseYaml(b []byte) ([]RedirectLink, error) { //map[string]string {
	var r []RedirectLink
	err := yaml.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func makeMap(r []RedirectLink) map[string]string {

	m := make(map[string]string)

	for _, rl := range r {
		m[rl.Path] = rl.URL
	}
	return m
}
