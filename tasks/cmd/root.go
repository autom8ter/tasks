package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "tasks",
	Long: `
------------------------------------------------------------------------------------------ 
TTTTTTTTTTTTTTTTTTTTTTT                                kkkkkkkk                          
T:::::::::::::::::::::T                                k::::::k                          
T:::::::::::::::::::::T                                k::::::k                          
T:::::TT:::::::TT:::::T                                k::::::k                          
TTTTTT  T:::::T  TTTTTTaaaaaaaaaaaaa      ssssssssss    k:::::k    kkkkkkk  ssssssssss   
        T:::::T        a::::::::::::a   ss::::::::::s   k:::::k   k:::::k ss::::::::::s  
        T:::::T        aaaaaaaaa:::::ass:::::::::::::s  k:::::k  k:::::kss:::::::::::::s 
        T:::::T                 a::::as::::::ssss:::::s k:::::k k:::::k s::::::ssss:::::s
        T:::::T          aaaaaaa:::::a s:::::s  ssssss  k::::::k:::::k   s:::::s  ssssss 
        T:::::T        aa::::::::::::a   s::::::s       k:::::::::::k      s::::::s      
        T:::::T       a::::aaaa::::::a      s::::::s    k:::::::::::k         s::::::s   
        T:::::T      a::::a    a:::::assssss   s:::::s  k::::::k:::::k  ssssss   s:::::s 
      TT:::::::TT    a::::a    a:::::as:::::ssss::::::sk::::::k k:::::k s:::::ssss::::::s
      T:::::::::T    a:::::aaaa::::::as::::::::::::::s k::::::k  k:::::ks::::::::::::::s 
      T:::::::::T     a::::::::::aa:::as:::::::::::ss  k::::::k   k:::::ks:::::::::::ss  
      TTTTTTTTTTT      aaaaaaaaaa  aaaa sssssssssss    kkkkkkkk    kkkkkkksssssssssss    
------------------------------------------------------------------------------------------

A gRPC powered task management server
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
