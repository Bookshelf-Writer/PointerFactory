package PointerFactory

type TypeTag byte

const TypeNone TypeTag = 0

func init() {
	wp.TypeMAP = map[rune]TypeTag{
		'0': TypeNone,
	}

}

//////////////////////////////////////////////////////////////////////////////

/* Парсинг типа по символу */
func (w *WrapperObj) ParseType(char rune) TypeTag {
	t, v := w.TypeMAP[char]
	if !v {
		return TypeNone
	}

	return t
}
