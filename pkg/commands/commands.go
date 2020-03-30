package commands

import (
	"errors"
	"fmt"
	"lynckx/das-boot-discord/pkg/dasbot/message"
	"strings"
)

type commandFunction func(msg message.Message) error
type Command struct {
	f commandFunction
}

var (
	Commands = make(map[string]*Command)
)

func On(id string, aliases []string, cmdFunc commandFunction) {
	cmd := &Command{
		f: cmdFunc,
	}
	Commands[strings.ToLower(id)] = cmd
	for _, alias := range aliases {
		Commands[strings.ToLower(alias)] = cmd
	}
}

func Emit(msg message.Message) error {
	fmt.Println("Emit called")
	cmd, ok := Commands[strings.ToLower(msg.GetFirstCommand())]
	fmt.Printf("Command exsists: %v, recieved command: %v\n", ok, cmd)
	if !ok {
		fmt.Println("returning")
		return errors.New("Command not found")
	}
	return cmd.f(msg)
}

func CreateCommandFunction(f func(msg message.Message) error) commandFunction {
	return commandFunction(f)
}

func LoadStandartCommand() {
	On("logo", []string{"DB", "Boot", "DasBoot", "das_boot"}, DasBoot)
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
