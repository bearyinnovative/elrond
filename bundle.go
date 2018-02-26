package elrond

import (
	"fmt"
)

// Bundle bundle translations
type Bundle struct {
	translations map[string]*Translation
}

// Add translation to bundle
func (b *Bundle) Add(id string, t *Translation) {
	b.translations[id] = t
}

// Get translation by id
func (b *Bundle) Get(id string) (translation *Translation, found bool) {
	translation, found = b.translations[id]
	return
}

func (b *Bundle) Size() int {
	return len(b.translations)
}

// MustGet translation, if not exists then panic
func (b *Bundle) MustGet(id string) *Translation {
	t, found := b.Get(id)

	if !found {
		panic(fmt.Sprintf("translation %s not found", id))
	}

	return t
}

func NewBundle() *Bundle {
	return &Bundle{
		translations: make(map[string]*Translation),
	}
}
