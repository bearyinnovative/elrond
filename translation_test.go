package elrond

import (
	"testing"

	b "github.com/smartystreets/goconvey/convey"
)

func TestTranslation_Content(t *testing.T) {
	b.Convey("get translation content by language", t, func() {
		b.Convey("no content", func() {
			tr := T()

			langs := []Language{
				ZhCN,
				EnUS,
			}

			for _, l := range langs {
				c, ok := tr.Content(l)

				b.So(c, b.ShouldBeNil)
				b.So(ok, b.ShouldBeFalse)
			}
		})

		b.Convey("one content", func() {
			contents := []*Content{
				C(ZhCN, "bar"),
			}

			tr := T(contents...)

			for _, c := range contents {
				c, ok := tr.Content(c.Language())

				b.So(c, b.ShouldEqual, c)
				b.So(ok, b.ShouldBeTrue)
			}

			noContentLangs := []Language{
				EnUS,
			}

			for _, l := range noContentLangs {
				c, ok := tr.Content(l)

				b.So(c, b.ShouldBeNil)
				b.So(ok, b.ShouldBeFalse)
			}
		})

		b.Convey("contents without duplicated language", func() {

			contents := []*Content{
				C(ZhCN, "foo"),
				C(EnUS, "bar"),
			}

			tr := T(contents...)

			for _, c := range contents {
				c, ok := tr.Content(c.Language())

				b.So(c, b.ShouldEqual, c)
				b.So(ok, b.ShouldBeTrue)
			}
		})
	})
}

func TestTranslation_Languages(t *testing.T) {
	b.Convey("get translation languages", t, func() {
		b.Convey("no content", func() {
			tr := T()
			b.So(tr.Languages(), b.ShouldBeEmpty)
		})

		b.Convey("one content", func() {
			c := C(ZhCN, "foobar")
			tr := T(c)

			b.So(tr.Languages(), b.ShouldHaveLength, 1)
			b.So(tr.Languages(), b.ShouldContain, c.Language())
		})

		b.Convey("contents without duplicated language", func() {
			c1 := C(ZhCN, "foo")
			c2 := C(EnUS, "bar")
			tr := T(c1, c2)

			b.So(tr.Languages(), b.ShouldHaveLength, 2)
			b.So(tr.Languages(), b.ShouldContain, c1.Language())
			b.So(tr.Languages(), b.ShouldContain, c2.Language())
		})

		b.Convey("contents with duplicated language", func() {
			c1 := C(ZhCN, "foo")
			c2 := C(ZhCN, "bar")
			tr := T(c1, c2)

			b.So(tr.Languages(), b.ShouldHaveLength, 1)
			b.So(tr.Languages(), b.ShouldContain, c2.Language())
		})
	})
}

func TestT(t *testing.T) {
	b.Convey("create a translation", t, func() {
		b.Convey("no content", func() {
			tr := T()

			b.So(tr.contents, b.ShouldBeEmpty)
		})

		b.Convey("one content", func() {
			c := C(ZhCN, "foobar")
			tr := T(c)

			b.So(tr.contents, b.ShouldHaveLength, 1)
			b.So(tr.contents[c.Language()], b.ShouldEqual, c)
		})

		b.Convey("contents without duplicated language", func() {
			c1 := C(ZhCN, "foo")
			c2 := C(EnUS, "bar")
			tr := T(c1, c2)

			b.So(tr.contents, b.ShouldHaveLength, 2)
			b.So(tr.contents[c1.Language()], b.ShouldEqual, c1)
			b.So(tr.contents[c2.Language()], b.ShouldEqual, c2)
		})

		b.Convey("contents with duplicated language", func() {
			c1 := C(ZhCN, "foo")
			c2 := C(ZhCN, "bar")
			tr := T(c1, c2)

			b.So(tr.contents, b.ShouldHaveLength, 1)
			b.So(tr.contents[c2.Language()], b.ShouldEqual, c2)
		})
	})
}
