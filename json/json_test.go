package json

import (
	"testing"

	b "github.com/smartystreets/goconvey/convey"
)

func TestFrom(t *testing.T) {
	b.Convey("create translation bundle from json", t, func() {
		b.Convey("empty translations", func() {
			datas := [][]byte{
				[]byte(`null`),
				[]byte(`[]`),
			}

			for _, d := range datas {
				bundle, err := From(d)

				b.So(err, b.ShouldBeNil)
				b.So(bundle.Size(), b.ShouldBeZeroValue)
			}
		})

		b.Convey("invalid format", func() {
			datas := [][]byte{
				[]byte(``),
				[]byte(`{}`),
				[]byte(`[{"id": 123}]`),
			}

			for _, d := range datas {
				bundle, err := From(d)

				b.So(err, b.ShouldNotBeNil)
				b.So(bundle, b.ShouldBeNil)
			}
		})

		b.Convey("create a bundle", func() {
			datas := [][]byte{
				[]byte(`[{"id": "foo", "contents": [{"language": "Zh-CN", "text": "bar"}]}]`),
				[]byte(`[{"id": "foo", "contents": [{"language": "Zh-CN", "text": "bar"}]}, {"id": "bar", "contents": [{"language": "En-US", "text": "foo"}, {"language": "Zh-CN", "text": "bar"}]}]`),
			}

			for _, d := range datas {
				bundle, err := From(d)

				b.So(err, b.ShouldBeNil)
				b.So(bundle, b.ShouldNotBeNil)
			}
		})
	})
}

func TestFromFile(t *testing.T) {
	b.Convey("create bundle from json file", t, func() {
		b.Convey("from no exist file", func() {
			bundle, err := FromFile("./foobar.json")
			b.So(bundle, b.ShouldBeNil)
			b.So(err, b.ShouldNotBeNil)
		})

		b.Convey("from testing json file", func() {
			bundle, err := FromFile("./json_test.json")
			b.So(err, b.ShouldBeNil)
			b.So(bundle, b.ShouldNotBeNil)
		})
	})
}
