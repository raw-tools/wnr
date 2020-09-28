package cmd

import (
	"github.com/noirbizarre/wnr/internal/tasks"
	"github.com/spf13/cobra"
)

// runCmd represents the set command
var runCmd = &cobra.Command{
	Use:   "run <task>",
	Short: "Execute a predefined task",
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

		done := make(chan struct{})

		// watchers.Watch(name)
		task := tasks.NewCommandTask("")
		task.Run()

		<-done

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
	rootCmd.AddCommand(runCmd)
}
