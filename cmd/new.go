package cmd

import (
	"github.com/spf13/cobra"
)

// newCmd represents the set command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Initialise a new WnR configuration",
	Args:  cobra.ExactArgs(1),
	Annotations: map[string]string{
		"source": "true",
	},
	Run: func(cmd *cobra.Command, args []string) {
		// user := user.Current()
		// project := user.CurrentProject()

		// if project == nil {
		// 	log.Fatalf("No active project")
		// }

		name := args[0]
		println(name)

		// // project.Load()
		// // key := fmt.Sprintf("scripts.%s", name)
		// // if !project.Config.IsSet(key) {
		// // 	log.Fatalf("Unknown script %s", name)
		// // }
		// // script := project.Config.GetString(key)

		// // session := shell.NewSession(false)
		// // session.AddCommand(script)
		// // code := user.Shell().Exec(session)
		// os.Exit(code)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
