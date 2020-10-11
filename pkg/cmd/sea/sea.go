package sea

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {

	c := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Hugo",
		Long:  `All software has versions. This is Hugo's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
		},
	}

	c.AddCommand()
	cobra.OnInitialize(func() {})

	return c
}
