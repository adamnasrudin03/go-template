package payload

type Translate struct {
	TargetLanguage string `json:"target_language"`
	OriginalText   string `json:"original_text" form:"original_text"`
	TranslatedText string `json:"translated_text"`
}
