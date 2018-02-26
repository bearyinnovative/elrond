package elrond

import (
	"testing"
	"text/template"

	b "github.com/smartystreets/goconvey/convey"
)

func TestContent_Language(t *testing.T) {
	b.Convey("get content language", t, func() {
		langs := []Language{
			ZhCN,
			EnUS,
		}

		for _, lang := range langs {
			c := Content{
				language: lang,
			}

			b.So(c.Language(), b.ShouldEqual, lang)
		}
	})
}

func TestContent_Parse(t *testing.T) {
	b.Convey("parse template", t, func() {
		type Data struct {
			Text   string
			Value  interface{}
			Output string
		}

		datas := []Data{
			{"foobar", nil, "foobar"},
			{"{{ .foo }} bar", nil, "<no value> bar"},
			{"{{ .foo }} bar", map[string]string{"foo": "baz"}, "baz bar"},
			{"foobar", map[string]string{"foo": "bar"}, "foobar"},
		}

		for _, data := range datas {
			c := C(ZhCN, data.Text)
			output, err := c.Parse(data.Value)

			b.So(err, b.ShouldBeNil)
			b.So(output, b.ShouldEqual, data.Output)
		}
	})
}

func TestContent_Template(t *testing.T) {
	b.Convey("get content raw template", t, func() {

		tmpl := template.New("testing")

		c := Content{template: tmpl}

		tmplGot := c.Template()

		b.So(tmplGot, b.ShouldEqual, tmpl)
	})
}

func TestContent_Text(t *testing.T) {
	b.Convey("get content text", t, func() {
		type Data struct {
			Text   string
			Output string
			Lang   Language
		}

		datas := []Data{
			{"{{.Foo}}", "<no value>", ZhCN},
			{"{{ .Bar }} testing", "<no value> testing", EnUS},
			{"{{ `{{ .Foo }}` }}", "{{ .Foo }}", ZhCN},
			{"foo", "foo", ZhCN},
			{"bar", "bar", ZhCN},
			{"测试", "测试", ZhCN},
			{"你好", "你好", ZhCN},
			{"", "", ZhCN},
		}

		for _, data := range datas {
			c := C(data.Lang, data.Text)

			output, err := c.Text()

			b.So(err, b.ShouldBeNil)
			b.So(output, b.ShouldEqual, data.Output)
		}
	})
}

func TestC(t *testing.T) {
	b.Convey("create a new content", t, func() {
		b.Convey("with valid template text", func() {
			type Data struct {
				Text string
				Lang Language
			}

			datas := []Data{
				{"{{.Foo}}", ZhCN},
				{"{{ .Bar }} testing", EnUS},
				{"foo", ZhCN},
				{"bar", ZhCN},
				{"测试", ZhCN},
				{"你好", ZhCN},
				{"", ZhCN},
			}

			for _, data := range datas {
				b.So(func() {
					c := C(data.Lang, data.Text)

					b.So(c.language, b.ShouldEqual, data.Lang)
				}, b.ShouldNotPanic)
			}
		})

		b.Convey("with invalid template text", func() {
			type Data struct {
				Text string
				Lang Language
			}

			datas := []Data{
				{"{{ if }}", ZhCN},
				{"{{ range 123 }} testing", EnUS},
			}

			for _, data := range datas {
				b.So(func() {
					C(data.Lang, data.Text)

				}, b.ShouldPanic)
			}
		})
	})
}
