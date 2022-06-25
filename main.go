package hardcandy

type Collection struct {
}

type Attribute struct {
	Trait    string
	Value    string // or int??
	Scarcity uint16 // 0 .. 100_00 percent
}
