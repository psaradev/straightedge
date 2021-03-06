package cmd

import (
	"bufio"
	"io"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
)

const (
	flagDryRun = "dry-run"
)

// keyCommands registers a sub-tree of commands to interact with
// local private key storage.
func keyCommands(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keys",
		Short: "Add or view local private keys",
		Long: `Keys allows you to manage your local keystore for tendermint.
    These keys may be in any format supported by go-crypto and can be
    used by light-clients, full nodes, or any other application that
    needs to sign with a private key.`,
	}
	addCmd := keys.AddKeyCommand()
	addCmd.RunE = runAddCmd
	cmd.AddCommand(
		keys.MnemonicKeyCommand(),
		addCmd,
		keys.ExportKeyCommand(),
		keys.ImportKeyCommand(),
		keys.ListKeysCmd(),
		keys.ShowKeysCmd(),
		flags.LineBreak,
		keys.DeleteKeyCommand(),
		keys.ParseKeyStringCommand(),
		keys.MigrateCommand(),
		flags.LineBreak,
	)

	cmd.PersistentFlags().String(flags.FlagHome, defaultNodeHome, "The application home directory")
	cmd.PersistentFlags().String(flags.FlagKeyringDir, "", "The client Keyring directory; if omitted, the default 'home' directory will be used")
	cmd.PersistentFlags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|test)")
	cmd.PersistentFlags().String(cli.OutputFlag, "text", "Output format (text|json)")
	return cmd
}

func getKeybase(cmd *cobra.Command, dryrun bool, buf io.Reader) (keyring.Keyring, error) {
	if dryrun {
		return keyring.NewInMemory(
			keyring.Option(func(options *keyring.Options) {
				options.SupportedAlgos = append(options.SupportedAlgos, hd.Sr25519)
				options.SupportedAlgosLedger = append(options.SupportedAlgosLedger, hd.Sr25519)
			}),
		), nil
	}

	return keyring.New(
		sdk.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), buf,
		keyring.Option(func(options *keyring.Options) {
			options.SupportedAlgos = append(options.SupportedAlgos, hd.Sr25519)
			options.SupportedAlgosLedger = append(options.SupportedAlgosLedger, hd.Sr25519)
		}),
	)
}

func runAddCmd(cmd *cobra.Command, args []string) error {
	inBuf := bufio.NewReader(cmd.InOrStdin())
	kb, err := getKeybase(cmd, viper.GetBool(flagDryRun), inBuf)
	if err != nil {
		return err
	}

	return keys.RunAddCmd(cmd, args, kb, inBuf)
}
