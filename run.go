package console

import (
	"flag"
	"fmt"
	utils "github.com/BabichMikhail/golang-dev-utils"
	"github.com/azer/snakecase"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type CommandHandler struct {
	DefaultCommand ICommand

	commands     ICommandSlice
	commandToRun ICommand
}

func (h *CommandHandler) SetCommands(commands ICommandSlice) {
	h.commands = commands
}

func (h *CommandHandler) NeedRunCommand() bool {
	if len(os.Args) > 1 {
		commandName := os.Args[1]
		// hack for flag.Parse
		os.Args = append(os.Args[:1], os.Args[2:]...)

		commandOrNil := utils.FindFirst(h.commands.GetIElementSlice(), func(command interface{}) bool {
			return command != nil && command.(ICommand).GetName() == commandName
		})

		if commandOrNil != nil {
			h.commandToRun = commandOrNil.(ICommand)
		}
	}

	if h.commandToRun == nil {
		h.commandToRun = h.DefaultCommand
	}

	return h.commandToRun != nil
}

// TODO validators
func (h *CommandHandler) Run() {
	if h.commandToRun == nil {
		panic("No command")
	}

	ac := newArgumentContainer()
	arguments := reflect.TypeOf(h.commandToRun.RegisterArguments())
	for i := 0; i < arguments.NumField(); i++ {
		field := arguments.Field(i)
		value := field.Tag.Get("default")
		usage := field.Tag.Get("usage")
		name := field.Tag.Get("name")
		if name == "" {
			name = snakecase.SnakeCase(field.Name)
		}

		if field.Type.Kind() == reflect.Bool {
			ac.addBool(name, flag.Bool(name, strings.ToLower(value) == "true", usage))
		} else if field.Type.Kind() == reflect.String {
			ac.addString(name, flag.String(name, value, usage))
		} else if field.Type.Kind() == reflect.Int {
			ac.addInt(name, flag.Int(name, utils.CheckNoError(strconv.Atoi(value)).(int), usage))
		} else {
			panic("Not implemented")
		}
	}

	flag.Parse()

	startedAt := time.Now()
	defer func() {
		fmt.Printf("Execution time: %s\n", time.Since(startedAt))
		if err := recover(); err != nil {
			panic(err)
		}
	}()

	h.commandToRun.Execute(ac)
}
