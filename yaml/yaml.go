package yaml

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/bearyinnovative/elrond"
)

type content struct {
	Language string `yaml:"language"`
	Text     string `yaml:"text"`
}

type translation struct {
	ID       string    `yaml:"id"`
	Contents []content `yaml:"contents"`
}

func From(data []byte) (*elrond.Bundle, error) {
	var ts []translation

	err := yaml.Unmarshal(data, &ts)

	if err != nil {
		return nil, err
	}

	bundle := elrond.NewBundle()

	for _, t := range ts {
		cs := make([]*elrond.Content, len(t.Contents))

		for i, c := range t.Contents {
			cs[i] = elrond.C(elrond.Language(c.Language), c.Text)
		}

		bundle.Add(t.ID, elrond.T(cs...))
	}

	return bundle, nil
}

func FromFile(filename string) (*elrond.Bundle, error) {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return From(data)
}
