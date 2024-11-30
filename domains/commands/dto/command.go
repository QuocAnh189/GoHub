package dto

type Command struct {
	ID   string `json:"id" `
	Name string `json:"name" `
}

type CommandInFunctionRes struct {
	Commands []Command `json:"commands"`
}

type CommandNotInFunctionRes struct {
	Commands []Command `json:"commands"`
}
