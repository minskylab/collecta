package collecta

type Span struct {
	After  *string
	Before *string
	Text   string
}

type Survey struct {
	Title     string
	Metadata  map[string]interface{}
	Questions []Question
	Spans     []Span
}
