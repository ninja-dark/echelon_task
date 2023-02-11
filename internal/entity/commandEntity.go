package entity

type RemoteCommand struct {
	Cmd   string `json:"cmd"`
	Os    string `json:"os"`
	Stdin string `json:"stdin"`
}
