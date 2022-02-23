package items

import (
	"oper/one-password/commands"

	easyjson "github.com/mailru/easyjson"
)

var op commands.OnePasswordCommands

func getItems(uuid string) (*Items, error) {
	com, err := op.GetItemsByContainer(uuid)
	if err != nil {
		return nil, err
	}
	buf := &Items{}
	err = easyjson.Unmarshal(com, buf)
	if err != nil {
		return nil, err
	}
	return buf, err
}
