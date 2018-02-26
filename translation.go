package elrond

// Translation represents a content for some language.
type Translation struct {
	contents map[Language]*Content
}

// Content get content by language
func (t *Translation) Content(lang Language) (content *Content, ok bool) {
	content, ok = t.contents[lang]
	return
}

// Languages list all languages of the translation
func (t *Translation) Languages() []Language {
	langs := make([]Language, len(t.contents))

	i := 0
	for l := range t.contents {
		langs[i] = l
		i++
	}

	return langs
}

// T create a new Translation
func T(contents ...*Content) *Translation {

	t := Translation{
		contents: make(map[Language]*Content, len(contents)),
	}

	for _, c := range contents {
		t.contents[c.Language()] = c
	}

	return &t
}
