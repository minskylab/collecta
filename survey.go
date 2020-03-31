package collecta

type Span struct {
	After  *string `json:"after"`
	Before *string `json:"before"`
	Text   string  `json:"text"`
}

type Survey struct {
	Title     string                 `json:"title"`
	Metadata  map[string]interface{} `json:"metadata"`
	Questions []Question             `json:"questions"`
	Spans     []Span                 `json:"spans"`
}
