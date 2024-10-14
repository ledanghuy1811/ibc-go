package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper *Keeper
}

// NewMigrator returns Migrator instance for the state migration.
func NewMigrator(k *Keeper) Migrator {
	return Migrator{
		keeper: k,
	}
}

// AssertChannelCapabilityMigrations checks that all channel capabilities generated using the interchain accounts controller port prefix
// are owned by the controller submodule and ibc.
func (m Migrator) AssertChannelCapabilityMigrations(ctx sdk.Context) error {
	if m.keeper != nil {
		logger := m.keeper.Logger(ctx)
		// filteredChannels := m.keeper.channelKeeper.GetAllChannelsWithPortPrefix(ctx, icatypes.ControllerPortPrefix)
		// for _, ch := range filteredChannels {
		// name := host.ChannelCapabilityPath(ch.PortId, ch.ChannelId)
		capability, _ := m.keeper.scopedKeeper.GetCapability(ctx, "capabilities/ports/icacontroller-orai10lv5lcj22zfnfqrqyqq73w5nvkjvmqa9a277lhc0vyu2hpm9ngssenvaaq/channels/channel-55")
		logger.Error(fmt.Sprintf("capability hehre: %v", capability))
		// if !found {
		// 	logger.Error(fmt.Sprintf("failed to find capability: %s", name))
		// 	return errorsmod.Wrapf(capabilitytypes.ErrCapabilityNotFound, "failed to find capability: %s", name)
		// }

		// isAuthenticated := m.keeper.scopedKeeper.AuthenticateCapability(ctx, capability, name)
		// if !isAuthenticated {
		// 	logger.Error(fmt.Sprintf("expected capability owner: %s", controllertypes.SubModuleName))
		// 	return errorsmod.Wrapf(capabilitytypes.ErrCapabilityNotOwned, "expected capability owner: %s", controllertypes.SubModuleName)
		// }

		// m.keeper.SetMiddlewareEnabled(ctx, ch.PortId, ch.ConnectionHops[0])
		// logger.Info("successfully migrated channel capability", "name", name)
		// }
	}
	return nil
}

// MigrateParams migrates the controller submodule's parameters from the x/params to self store.
// func (m Migrator) MigrateParams(ctx sdk.Context) error {
// 	if m.keeper != nil {
// 		params := controllertypes.DefaultParams()
// 		if m.keeper.legacySubspace != nil {
// 			m.keeper.legacySubspace.GetParamSetIfExists(ctx, &params)
// 		}
// 		m.keeper.SetParams(ctx, params)
// 		m.keeper.Logger(ctx).Info("successfully migrated ica/controller submodule to self-manage params")
// 	}
// 	return nil
// }
