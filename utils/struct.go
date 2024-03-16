package utils

type Register struct {
	Identifier  string    `json:"identifier"`
	Role        string    `json:"role"`
	EntryPoint  string    `json:"entry_point"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	Preview     string    `json:"preview"`
	Declare     []Declare `json:"declare"`
}
type Declare struct {
	Root        string        `json:"root"`
	Description string        `json:"description"`
	URL         string        `json:"url"`
	Args        []interface{} `json:"args"`
}
