package PointerFactory

type WrapperObj struct {
	checksumSymbols []string
	TypeMAP         map[rune]TypeTag `json:"types"`

	GlobalVersion    string `json:"ver"`
	GlobalDateUpdate string `json:"update"`
	GlobalName       string `json:"name"`

	NumBase   int `json:"num_base"`
	YearPoint int `json:"year_point"`

	IncrementMAX uint64 `json:"increment_max"`
}

func CreateWrapper() WrapperObj {
	return wp
}

var wp WrapperObj

func init() {
	wp.GlobalVersion = GlobalVersion
	wp.GlobalDateUpdate = GlobalDateUpdate
	wp.GlobalName = GlobalName

	wp.NumBase = NumBase
	wp.YearPoint = YearPoint
}
