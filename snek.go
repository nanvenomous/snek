package snek

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	completion string
	version    bool
	shells     = []string{"bash", "zsh", "fish", "powershell"}
)

func RunShellCompletion(rtCmd *cobra.Command) {
	if completion != "" {
		switch completion {
		case shells[0]:
			rtCmd.Root().GenBashCompletion(os.Stdout)
		case shells[1]:
			rtCmd.Root().GenZshCompletion(os.Stdout)
		case shells[2]:
			rtCmd.Root().GenFishCompletion(os.Stdout, true)
		case shells[3]:
			rtCmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		default:
			fmt.Println("not a recognized shell")
			os.Exit(1)
		}
		os.Exit(0)
	}
}

func InitRoot(rtCmd *cobra.Command, cfgFile string, cfgFileName string) {
	rtCmd.SilenceUsage = true
	rtCmd.SilenceErrors = true
	rtCmd.CompletionOptions.DisableDefaultCmd = true

	rtCmd.Flags().BoolVarP(
		&version, "version", "v", false,
		"show the version of this binary",
	)
	completionFlag := "completion"
	rtCmd.Flags().StringVar(&completion, completionFlag, "", "generate shell completion")
	rtCmd.RegisterFlagCompletionFunc(completionFlag, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return shells, cobra.ShellCompDirectiveDefault
	})

	rtCmd.Flags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is $HOME/.config/%s.yaml)", cfgFileName))

}
