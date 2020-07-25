package console

import (
	utils "github.com/BabichMikhail/golang-dev-utils"
)

type ICommandSlice []ICommand

func (slice ICommandSlice) GetIElementSlice() utils.IElementSlice {
	items := utils.IElementSlice{}
	for _, item := range slice {
		items = append(items, item)
	}

	return items
}

type IArgumentContainer interface {
	GetBool(name string) bool
	GetString(name string) string
	GetInt(name string) int
}

type ICommand interface {
	utils.IElement
	GetName() string
	RegisterArguments() interface{}
	Execute(container IArgumentContainer)
}

type BaseCommand struct {
	ICommand
}

func (cmd *BaseCommand) GetName() string {
	panic("Not implemented")
}

func (cmd *BaseCommand) RegisterArguments() interface{} {
	return struct{}{}
}

func (cmd *BaseCommand) Execute(container IArgumentContainer) {
	panic("Not implemented")
}
