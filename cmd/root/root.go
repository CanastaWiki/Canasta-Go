package cmd

import (
	"fmt"
	"os"

	addCmd "github.com/CanastaWiki/Canasta-CLI/cmd/add"
	backupCmd "github.com/CanastaWiki/Canasta-CLI/cmd/backup"
	configCmd "github.com/CanastaWiki/Canasta-CLI/cmd/config"
	createCmd "github.com/CanastaWiki/Canasta-CLI/cmd/create"
	deleteCmd "github.com/CanastaWiki/Canasta-CLI/cmd/delete"
	devmodeCmd "github.com/CanastaWiki/Canasta-CLI/cmd/devmode"
	exportCmd "github.com/CanastaWiki/Canasta-CLI/cmd/export"
	extensionCmd "github.com/CanastaWiki/Canasta-CLI/cmd/extension"
	gitopsCmd "github.com/CanastaWiki/Canasta-CLI/cmd/gitops"
	importCmd "github.com/CanastaWiki/Canasta-CLI/cmd/import"
	listCmd "github.com/CanastaWiki/Canasta-CLI/cmd/list"
	maintenanceCmd "github.com/CanastaWiki/Canasta-CLI/cmd/maintenance"
	removeCmd "github.com/CanastaWiki/Canasta-CLI/cmd/remove"
	restartCmd "github.com/CanastaWiki/Canasta-CLI/cmd/restart"
	sitemapCmd "github.com/CanastaWiki/Canasta-CLI/cmd/sitemap"
	skinCmd "github.com/CanastaWiki/Canasta-CLI/cmd/skin"
	startCmd "github.com/CanastaWiki/Canasta-CLI/cmd/start"
	stopCmd "github.com/CanastaWiki/Canasta-CLI/cmd/stop"
	upgradeCmd "github.com/CanastaWiki/Canasta-CLI/cmd/upgrade"
	versionCmd "github.com/CanastaWiki/Canasta-CLI/cmd/version"

	"github.com/CanastaWiki/Canasta-CLI/internal/config"
	"github.com/CanastaWiki/Canasta-CLI/internal/logging"

	"github.com/spf13/cobra"
)

const deprecationWarning = `WARNING: This is the final Canasta CLI 3.x (Go) release. Future versions
ship as Canasta CLI 4.x (Python/Ansible). To migrate:
    curl -fsSL https://get.canasta.wiki | bash
See https://canasta.wiki/wiki/Help:Installation.
(Set CANASTA_SUPPRESS_DEPRECATION_WARNING=1 to silence.)
`

func printDeprecationWarning() {
	if os.Getenv("CANASTA_SUPPRESS_DEPRECATION_WARNING") != "" {
		return
	}
	fmt.Fprint(os.Stderr, deprecationWarning)
}

var verbose bool

var rootCmd = &cobra.Command{
	Use:   "canasta",
	Short: "A CLI tool for managing Canasta instances.",
	Long: `Canasta CLI manages Canasta MediaWiki instances using Docker Compose
or Kubernetes. It supports creating, importing, starting, stopping, upgrading,
and backing up multiple Canasta instances, including wiki farms with multiple
wikis per instance.`,
	SilenceErrors: true,
	SilenceUsage:  true,
	PersistentPreRun: func(_ *cobra.Command, _ []string) {
		printDeprecationWarning()
		logging.SetVerbose(verbose)
		logging.Print("Verbose logging enabled")
	},
	Run: func(cmd *cobra.Command, _ []string) {
		_ = cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		logging.Fatal(err)
	}
}

// GetRootCmd returns the root command for documentation generation
func GetRootCmd() *cobra.Command {
	return rootCmd
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")

	rootCmd.AddCommand(configCmd.NewCmd())
	rootCmd.AddCommand(createCmd.NewCmd())
	rootCmd.AddCommand(deleteCmd.NewCmd())
	rootCmd.AddCommand(devmodeCmd.NewCmd())
	rootCmd.AddCommand(exportCmd.NewCmd())
	rootCmd.AddCommand(extensionCmd.NewCmd())
	rootCmd.AddCommand(gitopsCmd.NewCmd())
	rootCmd.AddCommand(importCmd.NewCmd())
	rootCmd.AddCommand(listCmd.NewCmd())
	rootCmd.AddCommand(maintenanceCmd.NewCmd())
	rootCmd.AddCommand(restartCmd.NewCmd())
	rootCmd.AddCommand(backupCmd.NewCmd())
	rootCmd.AddCommand(sitemapCmd.NewCmd())
	rootCmd.AddCommand(skinCmd.NewCmd())
	rootCmd.AddCommand(startCmd.NewCmd())
	rootCmd.AddCommand(stopCmd.NewCmd())
	rootCmd.AddCommand(upgradeCmd.NewCmd())
	rootCmd.AddCommand(versionCmd.NewCmd())
	rootCmd.AddCommand(addCmd.NewCmd())
	rootCmd.AddCommand(removeCmd.NewCmd())
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// Add config file location to help output
	defaultHelp := rootCmd.HelpFunc()
	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		defaultHelp(cmd, args)
		if cmd == rootCmd {
			configDir, err := config.GetConfigDir()
			if err != nil {
				cmd.Printf("\nConfig file: (error: %s)\n", err)
			} else {
				cmd.Printf("\nConfig file: %s/conf.json\n", configDir)
			}
		}
	})
}
