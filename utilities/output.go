package utilities

import (
	"fmt"

	"github.com/fatih/color"
)

const ASCIILOGO = `
                                                                    
                                                      88 88         
                                    ,d                "" 88         
                                    88                   88         
,adPPYba, 8b,dPPYba,  88       88 MM88MMM 8b,dPPYba,  88 88   ,d8   
I8[    "" 88P'    "8a 88       88   88    88P'   '"8a 88 88 ,a8"    
 '"Y8ba,  88       d8 88       88   88    88       88 88 8888[      
aa    ]8I 88b,   ,a8" "8a,   ,a88   88,   88       88 88 88'"Yba,   
 '"YbbdP"' 88'YbbdP"'   'YbbdP'Y8   "Y888 88       88 88 88   'Y8a  
          88                                                        
          88
`

func PrintError(output string) {
	color.Red(output)
}

func PrintWarning(output string) {
	color.Yellow(output)
}

func Print(output string) {
	//We shouldn't assume the colour of the terminal they're using.
	fmt.Println(output)
}

func PrintPrompt(output string) {
	color.Magenta(output)
}

func PrintLogo() {
	color.Magenta(ASCIILOGO)
}
