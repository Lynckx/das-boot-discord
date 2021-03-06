package commands

import "lynckx/das-boot-discord/pkg/dasbot/message"

// standart commands
func DasBoot(_ *CommandHandler, msg *message.Message, _ Arguments) error {
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
