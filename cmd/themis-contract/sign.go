package main

import (
	"os"

	contract "github.com/informalsystems/themis-contract/pkg/themis-contract"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var flagSigId string

func signCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign [contract]",
		Short: "Sign a contract",
		Long:  "Sign a contract, optionally specifying the signatory as whom you want to sign",
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
			err = c.Sign(flagSigId, ctx)
			if err != nil {
				log.Error().Msgf("Failed to sign contract: %s", err)
				os.Exit(1)
			}
			log.Info().Msg("Successfully signed contract")
		},
	}
	cmd.PersistentFlags().StringVar(&flagSigId, "as", "", "the ID of the signatory on behalf of whom you want to sign")
	return cmd
}
