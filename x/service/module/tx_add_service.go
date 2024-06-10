package service

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/pokt-network/poktroll/x/service/types"
	sharedtypes "github.com/pokt-network/poktroll/x/shared/types"
)

var _ = strconv.Itoa(0)

func CmdAddService() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-service <service_id> <service_name> <compute_units_per_relay>",
		Short: "Add a new service to the network",
		Long: `Add a new service to the network that will be available for applications,
gateways and suppliers to use. The service id MUST be unique but the service name doesn't have to be.

Example:
$ poktrolld tx service add-service "svc1" "service_one" --keyring-backend test --from $(SUPPLIER) --node $(POCKET_NODE) --home $(POKTROLLD_HOME)`,
		Args: cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			serviceIdStr := args[0]
			serviceNameStr := args[1]

			computeUnitsPerRelay := sharedtypes.DefaultComputeUnitsPerRelay
			// if compute units per relay argument is provided
			if len(args) > 2 {
				computeUnitsPerRelay, err = strconv.ParseUint(args[2], 10, 64)
				if err != nil {
					return types.ErrServiceInvalidComputUnitsPerRelay.Wrapf("unable to parse as uint64: %s", args[2])
				}
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddService(
				clientCtx.GetFromAddress().String(),
				serviceIdStr,
				serviceNameStr,
				computeUnitsPerRelay,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
