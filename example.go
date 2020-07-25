package console

import (
	"fmt"
)

type ExampleCommand struct {
	BaseCommand
}

func (cmd *ExampleCommand) GetName() string {
	return "app:example"
}

func (cmd *ExampleCommand) RegisterArguments() interface{} {
	return struct {
		BoolWithDefault           bool   `default:"false"`
		StringWithDefaultAndUsage string `default:"Hello world" usage:"Hello, usage!"`
		BoolWithoutName           bool   `name:"hello"`
	}{}
}

func (cmd *ExampleCommand) Execute(container IArgumentContainer) {
	fmt.Println("Hello. I'm example command")
	fmt.Println(container.GetBool("bool_with_default"))
	fmt.Println(container.GetBool("hello"))
	fmt.Println(container.GetString("string_with_default_and_usage"))
}
