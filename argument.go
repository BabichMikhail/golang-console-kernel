package console

import (
	"fmt"
	utils "github.com/BabichMikhail/golang-dev-utils"
)

type ArgumentContainer struct {
	IArgumentContainer

	boolValues   map[string]*bool
	stringValues map[string]*string
	intValues    map[string]*int
}

func (ac *ArgumentContainer) addBool(name string, ptr *bool) {
	ac.boolValues[name] = ptr
}

func (ac *ArgumentContainer) GetBool(name string) bool {
	value, ok := ac.boolValues[name]
	utils.CheckTrue(ok, fmt.Sprintf("Argument '%s' not found", name))
	return *value
}

func (ac *ArgumentContainer) SetBool(name string, value bool) {
	ac.boolValues[name] = &value
}

func (ac *ArgumentContainer) addString(name string, ptr *string) {
	ac.stringValues[name] = ptr
}

func (ac *ArgumentContainer) GetString(name string) string {
	value, ok := ac.stringValues[name]
	utils.CheckTrue(ok, fmt.Sprintf("Argument '%s' not found", name))
	return *value
}

func (ac *ArgumentContainer) SetString(name string, value string) {
	ac.stringValues[name] = &value
}

func (ac *ArgumentContainer) addInt(name string, ptr *int) {
	ac.intValues[name] = ptr
}

func (ac *ArgumentContainer) GetInt(name string) int {
	value, ok := ac.intValues[name]
	utils.CheckTrue(ok, fmt.Sprintf("Argument '%s' not found", name))
	return *value
}

func (ac *ArgumentContainer) SetInt(name string, value int) {
	ac.intValues[name] = &value
}

func newArgumentContainer() *ArgumentContainer {
	ac := new(ArgumentContainer)
	ac.boolValues = map[string]*bool{}
	ac.stringValues = map[string]*string{}
	ac.intValues = map[string]*int{}
	return ac
}
