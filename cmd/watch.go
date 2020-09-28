package cmd

import (
	"github.com/noirbizarre/wnr/internal/config"
	"github.com/noirbizarre/wnr/internal/controler"
	"github.com/noirbizarre/wnr/internal/tasks"
	"github.com/noirbizarre/wnr/internal/watchers"
	"github.com/spf13/cobra"
)

// runCmd represents the set command
var watchCmd = &cobra.Command{
	Use:   "watch <watcher>",
	Short: "Execute a predefined watcher",
	Args:  cobra.MinimumNArgs(2),
	// Args:  cobra.ExactArgs(1),
	Annotations: map[string]string{
		"source": "true",
	},
	Run: func(cmd *cobra.Command, args []string) {
		patterns := args[:len(args)-1]
		cmdLine := args[len(args)-1]

		exit := make(chan struct{})

		task := tasks.NewCommandTask(cmdLine)
		cfg := config.New()
		ctrl, _ := controler.New(cfg, task)
		watcher := watchers.NewFileWatcher(patterns...)
		ctrl.Run(exit, watcher)

		<-exit

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
	rootCmd.AddCommand(watchCmd)
}
