package commands

import (
	"errors"
	"strings"
	"sync"

	"lynckx/das-boot-discord/pkg/dasbot/message"
)

type CommandHandler struct {
	mu       *sync.RWMutex
	Commands map[string]*Command
	// msgCh is used to listen for additional user input after a command was
	msgChs []chan *message.Message
}

type CommandFunction func(cmd *CommandHandler, msg *message.Message) error

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

func (cmd *CommandHandler) SendMessageToListeners(msg *message.Message) {
	for _, ch := range cmd.msgChs {
		ch <- msg
	}
}

func (cmd *CommandHandler) AddMessageListener() chan *message.Message {
	ch := make(chan *message.Message)
	cmd.mu.Lock()
	defer cmd.mu.Unlock()
	cmd.msgChs = append(cmd.msgChs, ch)
	return ch
}

func (cmd *CommandHandler) RemoveMessageListener(ch chan *message.Message) {
	cmd.mu.Lock()
	defer cmd.mu.Unlock()
	cmd.msgChs = append(cmd.msgChs, ch)
	for i, oldCh := range cmd.msgChs {
		if oldCh == ch {
			cmd.msgChs[i] = cmd.msgChs[len(cmd.msgChs)-1]
			cmd.msgChs[len(cmd.msgChs)-1] = ch
			cmd.msgChs = cmd.msgChs[:len(cmd.msgChs)-1]
		}
	}
	close(ch)
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

func (cmd *CommandHandler) Emit(msg *message.Message) error {
	cmd.mu.RLock()
	command, ok := cmd.Commands[strings.ToLower(msg.GetCommandListAfterPrefix()[0])]
	cmd.mu.RUnlock()
	if !ok {
		return errors.New("Command not found")
	}
	return command.f(cmd, msg)
}

func CreateCommandFunction(f func(cmd *CommandHandler, msg *message.Message) error) CommandFunction {
	return CommandFunction(f)
}

func (cmd CommandHandler) LoadStandartCommand() {
	cmd.On("logo", []string{"DB", "Boot", "DasBoot", "das_boot"}, DasBoot)
}
