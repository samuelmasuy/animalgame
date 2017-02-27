package repo

type JsonAnimalQuestion struct {
	Id      int32  `json:"id,omitempty"`
	Content string `json:"content,omitempty"`
}

type JsonAnimalAnswer struct {
	Id      int32 `json:"id,omitempty"`
	Content bool  `json:"content,omitempty"`
}

type JsonError struct {
	Content string `json:"content"`
}

type JsonAnimalList struct {
	Content []string `json:"content"`
}
