package main

import (
	"os"

	contract "github.com/informalsystems/themis-contract/pkg/themis-contract"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	flagOutput string
)

func compileCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "compile [contract]",
		Short: "Compile a contract's sources to produce a PDF",
		Run: func(cmd *cobra.Command, args []string) {
			contractPath := defaultContractPath
			if len(args) > 0 {
				contractPath = args[0]
			}
			c, err := contract.Load(contractPath, ctx)
			if err != nil {
				log.Error().Msgf("Failed to load contract: %s", err)
				os.Exit(1)
			}
			err = c.Compile(flagOutput, ctx)
			if err != nil {
				log.Error().Msgf("Failed to compile contract: %s", err)
				os.Exit(1)
			}
			log.Info().Msg("Successfully compiled contract")
		},
	}
	cmd.PersistentFlags().StringVarP(&flagOutput, "output", "o", "contract.pdf", "where to write the output contract")
	return cmd
}
