package elrond

import (
	"bytes"
	"fmt"
	"text/template"
)

// Content for specified language
type Content struct {
	language Language

	template *template.Template
}

// Language of content
func (c *Content) Language() Language {
	return c.language
}

// Parse template
func (c *Content) Parse(v interface{}) (string, error) {
	var buf bytes.Buffer
	err := c.template.Execute(&buf, v)

	return buf.String(), err
}

// Template get content raw template
func (c *Content) Template() *template.Template {
	return c.template
}

// Text get plain text
func (c *Content) Text() (string, error) {
	return c.Parse(nil)
}

// C create a content, content is empty plain text by default
func C(language Language, text string) *Content {

	tmpl, err := template.New("translation").Parse(text)

	if err != nil {
		panic(fmt.Sprintf("failed parse template: %s", err.Error()))
	}

	c := Content{
		language: language,
		template: tmpl,
	}

	return &c
}
