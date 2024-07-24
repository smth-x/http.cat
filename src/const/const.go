package _const

const (
	FileName        = "res.jpg"
	UnknownMessage  = "Unknown message"
	StartMessage    = "Hello! If this is your first time using this bot, " + "first of all familiarize yourself with the capabilities of this bot by entering " + "the /commandList command. There you will find a list of all commands that are currently available to you"
	CaptionToImage  = "your image by code http"
	EmptyGetCommand = "Cannot use command get without argument. Write /commandList to find out the way to use commands"
)

type Command struct {
	Command        string
	Description    string
	ExampleOfUsing string
}

var (
	Commands = []Command{
		{Command: "/get", Description: "get an image by an http code", ExampleOfUsing: "/get 404"},
		{Command: "/commandList", Description: "shows all commands, that you can use"},
	}
)
