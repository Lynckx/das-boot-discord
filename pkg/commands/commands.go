package commands

func DasBoot() []string {
	var responseString []string
	asciiArt := "```\n" + `
	                                                                                
	                            &@@@@@@@@@\                                         
	                            &@@@@  @@@&\  &/@@@@@@\   /&@@@@%\                       
	                            /@@@@  %@@@@ @@/    @@@\ @@@@   @@@                 
	                            %@@@@  %@@@@   .//@@@@@  \@@@@@@\,                  
	      @@@@@@@@@@@@@@@@@@@@  #@@@@  &@@@/ @@@* /@@@@     ..@@@@      .;;//@@     
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
	` + "\n```"
	responseString = append(responseString, asciiArt)
	return responseString
}
