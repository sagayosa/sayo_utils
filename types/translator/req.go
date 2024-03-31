package translatortypes

// POST /translate
type TranslateReq struct {
	SourceText     string `json:"sourceText"`
	SourceLanguage string `json:"sourceLanguage"`
	TargetLanguage string `json:"targetLanguage"`
}

// GET /desktop/selected
type GetSelectedReq struct{}

// PUT /desktop/translator_model
type ModelTranslatorReq struct {
	Text string `json:"text"`
}

// GET /desktop/blur_trigger
type BlurTriggerReq struct{}

// POST /open_translator
type OpenTranslatorReq struct{}

// GET /setting
type GetSettingReq struct{}

// PUT /setting
type UpdateSettingReq struct {
	Config *Config `json:"config"`
}

type Config struct {
	Tencent            *Tencent `json:"tencent"`
	Model              string   `json:"model"`
	TranslationAnytime bool     `json:"translationAnytime"`
}

type Tencent struct {
	SecretId  string `json:"secretId"`
	SecretKey string `json:"secretKey"`
}

// POST /polish
type PolishReq struct {
	SourceText string `json:"sourceText"`
}

// POST /open_polisher
type OpenPolisherReq struct{}
