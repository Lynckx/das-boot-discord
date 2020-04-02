package commands

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/renstrom/shortuuid"

	"lynckx/das-boot-discord/pkg/dasbot/message"
)

type CommandHandler struct {
	mu       *sync.RWMutex
	Commands map[string]*Command
	// msgCh is used to listen for additional user input after a command was
	msgChs map[string]chan *message.Message
}

type Listener struct {
	id string
	Ch chan *message.Message
}

type CommandFunction func(cmd *CommandHandler, msg *message.Message, args Arguments) error

type Arguments []interface{}

type CommandLoader struct {
	Id      string
	Alias   []string
	CmdFunc CommandFunction
}
type Command struct {
	f CommandFunction
}

func CreateCommandHandler() *CommandHandler {
	return &CommandHandler{
		mu:       &sync.RWMutex{},
		Commands: make(map[string]*Command),
	}
}

func (cmd *CommandHandler) Broadcast(msg *message.Message) {
	cmd.SendMessageToListeners(msg)
	go func() {
		if msg.HasBotPrefix() {
			if err := cmd.Emit(strings.ToLower(msg.GetCommandListAfterPrefix()[0]), msg, Arguments{}); err != nil {
				fmt.Printf("Got an error %v\n", err)
			}
			fmt.Printf("received message: %v From User: %+v\n", msg.GetMessageContent(), msg.GetAuthor())
		}
	}()
}

func (cmd *CommandHandler) SendMessageToListeners(msg *message.Message) {
	for _, ch := range cmd.msgChs {
		ch <- msg
	}
}

func (cmd *CommandHandler) AddMessageListener() Listener {
	ch := make(chan *message.Message)
	cmd.mu.Lock()
	defer cmd.mu.Unlock()
	id := shortuuid.New()
	cmd.msgChs[id] = ch
	return Listener{
		id: id,
		Ch: ch,
	}
}

func (cmd *CommandHandler) RemoveMessageListener(l Listener) {
	cmd.mu.Lock()
	defer cmd.mu.Unlock()
	for id, oldCh := range cmd.msgChs {
		if id == l.id {
			delete(cmd.msgChs, id)
			close(oldCh)
			break
		}
	}
}

func (cmd *CommandHandler) LoadCommandsFrom(newCommands []*CommandLoader) {
	for _, newCommand := range newCommands {
		cmd.On(newCommand.Id, newCommand.Alias, newCommand.CmdFunc)
	}
}

func (cmd *CommandHandler) LoadCommand(newCommand *CommandLoader) {
	cmd.On(newCommand.Id, newCommand.Alias, newCommand.CmdFunc)
}

func (cmd *CommandHandler) On(id string, aliases []string, cmdFunc CommandFunction) {
	newCommand := &Command{
		f: cmdFunc,
	}
	cmd.mu.Lock()
	cmd.Commands[strings.ToLower(id)] = newCommand
	cmd.mu.Unlock()
	for _, alias := range aliases {
		cmd.Commands[strings.ToLower(alias)] = newCommand
	}
}

func (cmd *CommandHandler) Emit(cmdId string, msg *message.Message, args Arguments) error {
	cmd.mu.RLock()
	command, ok := cmd.Commands[cmdId]
	cmd.mu.RUnlock()
	if !ok {
		return errors.New("Command not found")
	}
	return command.f(cmd, msg, args)
}

func CreateCommandFunction(f func(cmd *CommandHandler, msg *message.Message, args Arguments) error) CommandFunction {
	return CommandFunction(f)
}

func (cmd CommandHandler) LoadStandartCommand() {
	cmd.On("logo", []string{"DB", "Boot", "DasBoot", "das_boot"}, DasBoot)
}
