package module

type Arg struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Declare struct {
	Root        string `json:"root"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Args        []Arg  `json:"args"`
}

type PluginConfig struct {
	Declare []Declare `json:"declare"`
}

type Plugin struct {
	ModuleInfo
	PluginConfig
}
