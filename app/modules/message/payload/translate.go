package payload

type Translate struct {
	TargetLanguage string `json:"target_language"`
	Text           string `json:"text" form:"text"`
}
