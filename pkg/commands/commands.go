package commands

import (
	"errors"
	"fmt"
	"lynckx/das-boot-discord/pkg/dasbot/message"
	"strings"
)

type CommandHandler struct {
	Commands map[string]*Command
}

type commandFunction func(msg message.Message) error
type Command struct {
	f commandFunction
}

func CreateCommandHandler() CommandHandler {
	return CommandHandler{
		Commands: make(map[string]*Command),
	}
}

func (cmd *CommandHandler) On(id string, aliases []string, cmdFunc commandFunction) {
	newCommand := &Command{
		f: cmdFunc,
	}
	cmd.Commands[strings.ToLower(id)] = newCommand
	for _, alias := range aliases {
		cmd.Commands[strings.ToLower(alias)] = newCommand
	}
}

func (cmd *CommandHandler) Emit(msg message.Message) error {
	fmt.Println("Emit called")
	command, ok := cmd.Commands[strings.ToLower(msg.GetFirstCommand())]
	if !ok {
		return errors.New("Command not found")
	}
	return command.f(msg)
}

func CreateCommandFunction(f func(msg message.Message) error) commandFunction {
	return commandFunction(f)
}

func (cmd CommandHandler) LoadStandartCommand() {
	cmd.On("logo", []string{"DB", "Boot", "DasBoot", "das_boot"}, DasBoot)
}

// standart commands
func DasBoot(msg message.Message) error {
	var responseString []string
	responseString = append(responseString, "```\n"+`
	                                                                                
	                            &@@@@@@@@@\                                         
	                            &@@@@  @@@&\  &/@@@@@@\   /&@@@@%\                       
	                            /@@@@  %@@@@ @@/    @@@\ @@@@   @@@                 
	                            %@@@@  %@@@@   .//@@@@@  \@@@@@@\,                  
	      @@@@@@@@@@@@@@@@@@@@  #@@@@  &@@@/ @@@* /@@@@  .,  .@@@@      .;;//@@     
	      @@@@@@@@@@@@@@@@@@@@@ &@@@@@@@@@/  \@@@@@ @@@\  \@@@@@@*     @@@@@@@*     
	      @@@@@@@@    @@@@@@@@@   ..;//@@@\\;..       ..;//@@@\\;..    @@@@@@@#     
	      @@@@@@@@    @@@@@@@@@  /@@@@@@@@@@@@@@\   /@@@@@@@@@@@@@@\ @@@@@@@@@@@   
	      @@@@@@@@@@@@@@@@@@@@  @@@@@@@@@@@@@@@@@@ @@@@@@@@@@@@@@@@@@ @@@@@@@@@@@   
	      @@@@@@@@@@@@@@@@@@@   @@@@@@@(  .@@@@@@@ @@@@@@@.  @@@@@@@@  @@@@@@@*     
	      @@@@@@@@@/%@@@@@@@@@@ @@@@@@@&  *@@@@@@@ @@@@@@@.  &@@@@@@@  @@@@@@@%     
	      @@@@@@@@    &@@@@@@@@ @@@@@@@&  *@@@@@@@ @@@@@&@.  &@@@@*@@  @@@@@@@@     
	      @@@@@@@(    @@@@@@@@@ @@@@@@@&  ,@@@@@@@ @@@@@@@*  @@@@@@@@  @@@@@@@@     
	     @@@@@@@@@@@@@@@@@@@@@@ @@@@@@@@  @@@@@@@@ @@@@@@@@  @@@@@@@@  #@@@@@@@&&   
	     @&@@@@@@@@@@@@@@@@@@@   @@@@@@@@@@%@@@@@(  @@@@@@@@@@@@@@@@   &@@@@@@@@@.  
	     %%%%%%@&@%&&%%#(,        '';\\@@@//;''      '';\\@@@//;''      \@@@&@@@,   
	`+"\n```")
	responseString = append(responseString, "```\n"+`                                                                        
	                        @&*                                                     
	                         @@@@(                                                  
	                          @@@@@,                                                
	                          @@@@@&                                                
	     @@,               ,@@@@@@@@#                                               
	   .@@@@          #@@@@@@@@@@@@@@@@@@%                                          
	   '@@@@'      .@@@@@@@@@@@@@@@@@@@@@@@@,                                       
	   .@@@@@,    @@@@@@@@@@@@@@@@@@@@@@@@@@@@                                      
	    %@@@@@@%.@@@@@@@@@@@@@@@@@@@@@@@/@@@@@@                                     
	   @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@ @@@@@@@@@                                    
	 *@@@@@@@@@%@@@@@@@@@@@@@@@@@@@@@@@& @'@@@@@@@@@%,                              
	 #@@@@@/    @@@@@@@@@@@@@@@@@@@@@@@@ ./@@@@@@@@@@@@@@@@%,                       
	  @@@       ,@@@@@@@@@@@@@@@@.,@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@&,                
	   ,         '@@@@@@@@@@@@@@@ % @@@@@@@@@@/.&@@@@@@@@@@@@@@@@@@@@@@@@&,         
	               &@@@@@@@@@@@@@/'&& &@@@@@#          /@%  .%@@% %@@@@@@@@@@@@@@.  
	                 (@@@@@@@@@@@@/,&,.  ,*                                 ,@@@@@( 
	               *@@@@@@@@@@@@@@@@..                                              
	             #@@@@@@@@@@.                                                       
	                '#%%(.                                                          
	`+"\n```")
	return msg.Respond(responseString)
}
