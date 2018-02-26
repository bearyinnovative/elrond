Elrond
====

i18n translation library

QuickStart
----

```go
import "github.com/bearyinnovative/elrond"

func Translation() {
	t := elrond.T(
		erlond.C(elrond.EnUS, "Hello, world!"),
		elrond.C(elrond.ZhCN, "你好，世界！"))

	{
		c, ok := t.Content(elrond.EnUS)
		println(c)
		// output: Hello, world

	}
	{
		c, ok := t.Content(elrond.ZhCN)
		println(c)
		// output: 你好，世界！
	}
}

func Bundle() {
	b := elrond.NewBundle()
	b.Add("foo", elrond.T(
		erlond.C(elrond.EnUS, "Hello, world!"),
		elrond.C(elrond.ZhCN, "你好，世界！")))

	c, ok := b.MustGet("foo").Content(elrond.EnUS)
	println(c)
	// output: Hello, world
}

func FromJSON() {
	datas := []byte(`[{"id": "bar", "contents": [{"language": "Zh-CN", "text": "bar"}]}, {"id": "bar", "contents": [{"language": "En-US", "text": "foo"}, {"language": "Zh-CN", "text": "bar"}]}]`)

	bundle, err := json.From(d)

	if err != nil {
		panic(err)
	}

	c, ok := bundle.MustGet("foo").Content(elrond.ZhCN)
	println(c)
	// output: bar
}
```

License
----
MIT
