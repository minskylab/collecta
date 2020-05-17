package collecta

type PairMap struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

type Map struct {
	Content []PairMap `json:"content"`
}
