package cmd

import (
    "strings"

    "github.com/spf13/cobra"
    "github.com/truecharts/public/clustertool/pkg/initfiles"
    "github.com/truecharts/public/clustertool/pkg/talassist"
)

var advTestCmdlongHelp = strings.TrimSpace(`
This command is mostly just for development usage and should NEVER be used by end-users.
`)

var testcmd = &cobra.Command{
    Use:   "test",
    Short: "tests specific code for developer usages",
    Long:  advTestCmdlongHelp,
    Run: func(cmd *cobra.Command, args []string) {
        initfiles.LoadTalEnv(false)
        talassist.LoadTalConfig()
        // err := fluxhandler.ProcessJSONFiles("./testdata/truenas_exports")
        // if err != nil {
        //  log.Info().Msg("Error:", err)
        // }
        RunApply(false, "", []string{})
    },
}

func init() {
    adv.AddCommand(testcmd)
}
