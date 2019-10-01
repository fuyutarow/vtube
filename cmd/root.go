package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var rootCmd = &cobra.Command{
	Use:   "vtube",
	Short: "A brief description of your application",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(
		&cobra.Command{
			Use:     "play",
			Short:   "Plays a song with the given name or url",
			Aliases: []string{"p"},
			Run:     PlayCmd,
		})

	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "status",
			Short: "Show status",
			Run:   StatusCmd,
		})

	rootCmd.AddCommand(
		&cobra.Command{
			Use:     "resume",
			Aliases: []string{"re"},
			Short:   "Resume paused music",
			Run: func(cmd *cobra.Command, args []string) {
				pid := GetMPlayerPID()
				SendSig("CONT", pid)
			},
		})

	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "pause",
			Short: "Pauses the currently playing track",
			Run: func(cmd *cobra.Command, args []string) {
				pid := GetMPlayerPID()
				SendSig("STOP", pid)
			},
		})

	rootCmd.AddCommand(
		&cobra.Command{
			Use:     "skip",
			Aliases: []string{"s"},
			Short:   "Skips the currently playing song",
			Run: func(cmd *cobra.Command, args []string) {
				pid := GetMPlayerPID()
				SendSig("KILL", pid)
			},
		})

	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "ui",
			Short: "Open GUI",
			Run:   UiCmd,
		})

}

func SendSig(signal, pid string) {
	cmd := exec.Command("kill", fmt.Sprintf("-%s", signal), pid)

	out, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
