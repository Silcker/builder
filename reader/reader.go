package reader

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type conf struct {
	Package struct {
		Build string `json:"build"`
		Start string `json:"start"`
	}
	Type    string `json:"type"`
	Volumes struct {
		Path string `json:"path"`
	}
}

func ReadConfig(path string) (*conf, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := &conf{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %w", path, err)
	}

	return c, err
}
