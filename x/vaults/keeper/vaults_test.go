package keeper_test

import (
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/vaults/types"
)

func (s *KeeperTestSuite) TestVaultsStore() {
	s.SetupTest()

	v := types.Vault{
		Owner:            s.TestAccs[0].String(),
		Debt:             sdk.NewCoin("atom", math.NewInt(1000)),
		CollateralLocked: sdk.NewCoin("atom", math.NewInt(1000)),
		Status:           types.LIQUIDATED,
	}
	err := s.k.SetVault(s.Ctx, v)
	s.Require().NoError(err)

	vault, err := s.k.GetVault(s.Ctx, 0)
	s.Require().NoError(err)

	s.Require().Equal(v, vault)
}

func (s *KeeperTestSuite) TestCreateNewVault() {
	s.SetupTest()
	var (
		denom         = "atom"
		coin          = sdk.NewCoin(denom, math.NewInt(1000))
		coinMintToAcc = sdk.NewCoin(denom, math.NewInt(1000000))
		maxDebt       = math.NewInt(10000)
	)

	tests := []struct {
		name       string
		setup      func()
		denom      string
		owner      sdk.AccAddress
		collateral sdk.Coin
		mint       sdk.Coin
	}{
		{
			name: "success",
			setup: func() {
				err := s.k.ActiveCollateralAsset(s.Ctx, denom, math.LegacyMustNewDecFromStr("0.1"), math.LegacyMustNewDecFromStr("0.1"), maxDebt)
				s.Require().NoError(err)

				err = s.App.BankKeeper.MintCoins(s.Ctx, types.ModuleName, sdk.NewCoins(coinMintToAcc))
				s.Require().NoError(err)
				err = s.App.BankKeeper.SendCoinsFromModuleToAccount(s.Ctx, types.ModuleName, s.TestAccs[0], sdk.NewCoins(coinMintToAcc))
				s.Require().NoError(err)
			},
			denom:      denom,
			owner:      s.TestAccs[0],
			collateral: coin,
			mint:       coin,
		},
	}
	for _, t := range tests {
		s.Run(t.name, func() {
			t.setup()
			err := s.k.CreateNewVault(s.Ctx, t.denom, t.owner, t.collateral, t.mint)
			s.Require().NoError(err)

			vm, err := s.k.GetVaultManager(s.Ctx, denom)
			s.Require().NoError(err)
			s.Require().NotEqual(maxDebt, vm.MintAvailable)
		})
	}
}

func (s *KeeperTestSuite) TestRepayDebt() {
	s.SetupTest()
	var (
		denom         = "atom"
		coin          = sdk.NewCoin(denom, math.NewInt(1000))
		coinMintToAcc = sdk.NewCoin(denom, math.NewInt(1000000))
		maxDebt       = math.NewInt(10000)
	)

	tests := []struct {
		name    string
		setup   func()
		vaultID uint64
		sender  sdk.AccAddress
		mint    sdk.Coin
	}{
		{
			name: "success",
			setup: func() {
				err := s.k.ActiveCollateralAsset(s.Ctx, denom, math.LegacyMustNewDecFromStr("0.1"), math.LegacyMustNewDecFromStr("0.1"), maxDebt)
				s.Require().NoError(err)

				vault := types.Vault{
					Owner:            s.TestAccs[0].String(),
					Debt:             sdk.NewCoin(denom, maxDebt),
					CollateralLocked: sdk.NewCoin(denom, maxDebt),
					Status:           types.ACTIVE,
				}
				err = s.k.SetVault(s.Ctx, vault)
				s.Require().NoError(err)

				err = s.App.BankKeeper.MintCoins(s.Ctx, types.ModuleName, sdk.NewCoins(coinMintToAcc))
				s.Require().NoError(err)
				err = s.App.BankKeeper.SendCoinsFromModuleToAccount(s.Ctx, types.ModuleName, s.TestAccs[0], sdk.NewCoins(coinMintToAcc))
				s.Require().NoError(err)
			},
			vaultID: 0,
			sender:  s.TestAccs[0],
			mint:    coin,
		},
	}
	for _, t := range tests {
		s.Run(t.name, func() {
			t.setup()
			err := s.k.RepayDebt(s.Ctx, t.vaultID, t.sender, t.mint)
			s.Require().NoError(err)

			vm, err := s.k.GetVaultManager(s.Ctx, denom)
			s.Require().NoError(err)
			s.Require().NotEqual(maxDebt, vm.MintAvailable)
		})
	}
}

func (s *KeeperTestSuite) TestDepositToVault() {
	s.SetupTest()
	var (
		denom         = "atom"
		coin          = sdk.NewCoin(denom, math.NewInt(1000))
		coinMintToAcc = sdk.NewCoin(denom, math.NewInt(1000000))
		maxDebt       = math.NewInt(10000)
	)

	tests := []struct {
		name       string
		setup      func()
		vaultId    uint64
		sender     sdk.AccAddress
		collateral sdk.Coin
	}{
		{
			name: "success",
			setup: func() {
				err := s.k.ActiveCollateralAsset(s.Ctx, denom, math.LegacyMustNewDecFromStr("0.1"), math.LegacyMustNewDecFromStr("0.1"), maxDebt)
				s.Require().NoError(err)

				vault := types.Vault{
					Owner:            s.TestAccs[0].String(),
					Debt:             sdk.NewCoin(denom, maxDebt),
					CollateralLocked: sdk.NewCoin(denom, maxDebt),
					Status:           types.ACTIVE,
				}
				err = s.k.SetVault(s.Ctx, vault)
				s.Require().NoError(err)

				err = s.App.BankKeeper.MintCoins(s.Ctx, types.ModuleName, sdk.NewCoins(coinMintToAcc))
				s.Require().NoError(err)
				err = s.App.BankKeeper.SendCoinsFromModuleToAccount(s.Ctx, types.ModuleName, s.TestAccs[0], sdk.NewCoins(coinMintToAcc))
				s.Require().NoError(err)
			},
			vaultId:    0,
			sender:     s.TestAccs[0],
			collateral: coin,
		},
	}
	for _, t := range tests {
		s.Run(t.name, func() {
			t.setup()
			err := s.k.DepositToVault(s.Ctx, t.vaultId, t.sender, t.collateral)
			s.Require().NoError(err)

			vault, err := s.k.GetVault(s.Ctx, t.vaultId)
			s.Require().NoError(err)
			s.Require().NotEqual(maxDebt, vault.CollateralLocked)
		})
	}
}

func (s *KeeperTestSuite) TestWithdrawFromVault() {
	s.SetupTest()
	var (
		denom         = "atom"
		coin          = sdk.NewCoin(denom, math.NewInt(1000))
		coinMintToAcc = sdk.NewCoin(denom, math.NewInt(1000000))
		maxDebt       = math.NewInt(10000)
	)

	tests := []struct {
		name       string
		setup      func()
		vaultId    uint64
		sender     sdk.AccAddress
		collateral sdk.Coin
	}{
		{
			name: "success",
			setup: func() {
				err := s.k.ActiveCollateralAsset(s.Ctx, denom, math.LegacyMustNewDecFromStr("0.1"), math.LegacyMustNewDecFromStr("0.1"), maxDebt)
				s.Require().NoError(err)

				vault := types.Vault{
					Owner:            s.TestAccs[0].String(),
					Debt:             sdk.NewCoin(denom, maxDebt),
					CollateralLocked: sdk.NewCoin(denom, maxDebt),
					Status:           types.ACTIVE,
				}
				err = s.k.SetVault(s.Ctx, vault)
				s.Require().NoError(err)

				err = s.App.BankKeeper.MintCoins(s.Ctx, types.ModuleName, sdk.NewCoins(coinMintToAcc))
				s.Require().NoError(err)
				err = s.App.BankKeeper.SendCoinsFromModuleToAccount(s.Ctx, types.ModuleName, s.TestAccs[0], sdk.NewCoins(coinMintToAcc))
				s.Require().NoError(err)
			},
			vaultId:    0,
			sender:     s.TestAccs[0],
			collateral: coin,
		},
	}
	for _, t := range tests {
		s.Run(t.name, func() {
			t.setup()
			err := s.k.WithdrawFromVault(s.Ctx, t.vaultId, t.sender, t.collateral)
			s.Require().NoError(err)

			vault, err := s.k.GetVault(s.Ctx, t.vaultId)
			s.Require().NoError(err)
			s.Require().NotEqual(maxDebt, vault.CollateralLocked)
		})
	}
}

func (s *KeeperTestSuite) TestUpdateVaultsDebt() {
	s.SetupTest()
	var (
		denom              = "atom"
		maxDebt            = math.NewInt(10000)
		feeStabilityUpdate = math.LegacyMustNewDecFromStr("0.5")
	)

	tests := []struct {
		name    string
		setup   func()
		vaultId uint64
	}{
		{
			name: "success",
			setup: func() {
				vault := types.Vault{
					Owner:            s.TestAccs[0].String(),
					Debt:             sdk.NewCoin(denom, maxDebt),
					CollateralLocked: sdk.NewCoin(denom, maxDebt),
					Status:           types.ACTIVE,
				}
				err := s.k.SetVault(s.Ctx, vault)
				s.Require().NoError(err)

				// update params
				uP := types.DefaultParams()
				uP.StabilityFee = feeStabilityUpdate
				err = s.k.SetParams(s.Ctx, uP)
				s.Require().NoError(err)
			},
			vaultId: 0,
		},
	}
	for _, t := range tests {
		s.Run(t.name, func() {
			t.setup()
			err := s.k.UpdateVaultsDebt(s.Ctx)
			s.Require().NoError(err)

			// expect
			expectDebtAmount := math.LegacyNewDecFromInt(maxDebt).Add(math.LegacyNewDecFromInt(maxDebt).Mul(feeStabilityUpdate)).TruncateInt()
			vault, err := s.k.GetVault(s.Ctx, t.vaultId)
			s.Require().NoError(err)
			s.Require().Equal(expectDebtAmount.String(), vault.Debt.Amount.String())
		})
	}
}

func (s *KeeperTestSuite) TestLiquidate() {
	// s.SetupTest()

	vaultOwnerAddr := sdk.AccAddress([]byte("addr1_______________"))

	tests := []struct {
		name               string
		liquidation        types.Liquidation
		expVaultStatus     []types.VaultStatus
		shortfallAmount    sdk.Coin
		moduleBalances     sdk.Coins
		reserveBalances    sdk.Coins
		vaultOwnerBalances sdk.Coins
	}{
		{
			name: "single vault - sold all, enough to cover debt",
			liquidation: types.Liquidation{
				Denom: "atom",
				LiquidatingVaults: []*types.Vault{
					{
						Owner:            vaultOwnerAddr.String(),
						Debt:             sdk.NewCoin(types.DefaultMintDenom, math.NewInt(25_000_000)),
						CollateralLocked: sdk.NewCoin("atom", math.NewInt(5_000_000)), // lock 5 ATOM at price 8, ratio = 160%
						Status:           types.LIQUIDATING,
						LiquidationPrice: math.LegacyNewDecWithPrec(7, 0), // liquidate at price 7, ratio = 140%
					},
				},
				// Sold all at price 7,
				// Sold = 35
				// Remain collateral = 0
				VaultLiquidationStatus: map[uint64]*types.VaultLiquidationStatus{
					0: {
						Sold:             sdk.NewCoin(types.DefaultMintDenom, math.NewInt(35_000_000)),
						RemainCollateral: sdk.NewCoin("atom", math.ZeroInt()),
					},
				},
			},
			expVaultStatus: []types.VaultStatus{types.LIQUIDATED},
			reserveBalances: sdk.NewCoins(sdk.NewCoin(types.DefaultMintDenom, math.NewInt(10_000_000))),
		},
		{
			name: "single vault - sold all, not enough to cover debt",
			liquidation: types.Liquidation{
				Denom: "atom",
				LiquidatingVaults: []*types.Vault{
					{
						Owner:            vaultOwnerAddr.String(),
						Debt:             sdk.NewCoin(types.DefaultMintDenom, math.NewInt(25_000_000)),
						CollateralLocked: sdk.NewCoin("atom", math.NewInt(5_000_000)), // lock 5 ATOM at price 8, ratio = 160%
						Status:           types.LIQUIDATING,
						LiquidationPrice: math.LegacyNewDecWithPrec(7, 0), // liquidate at price 7, ratio = 140%
					},
				},
				// Sold all at price 4,
				// Sold = 20
				// Remain collateral = 0
				VaultLiquidationStatus: map[uint64]*types.VaultLiquidationStatus{
					0: {
						Sold:             sdk.NewCoin(types.DefaultMintDenom, math.NewInt(20_000_000)),
						RemainCollateral: sdk.NewCoin("atom", math.ZeroInt()),
					},
				},
			},
			expVaultStatus:  []types.VaultStatus{types.LIQUIDATED},
			shortfallAmount: sdk.NewCoin(types.DefaultMintDenom, math.NewInt(5_000_000)),
		},
		{
			name: "single vault - remain collateral, enough to cover debt",
			liquidation: types.Liquidation{
				Denom: "atom",
				LiquidatingVaults: []*types.Vault{
					{
						Owner:            vaultOwnerAddr.String(),
						Debt:             sdk.NewCoin(types.DefaultMintDenom, math.NewInt(25_000_000)),
						CollateralLocked: sdk.NewCoin("atom", math.NewInt(5_000_000)), // lock 5 ATOM at price 8, ratio = 160%
						Status:           types.LIQUIDATING,
						LiquidationPrice: math.LegacyNewDecWithPrec(7, 0), // liquidate at price 7, ratio = 140%
					},
				},
				// Sold 1 at 7
				// Sold 2 at 6.5
				// Sold 1 at 6
				// Sold = 26
				// Remain collateral = 1
				VaultLiquidationStatus: map[uint64]*types.VaultLiquidationStatus{
					0: {
						Sold:             sdk.NewCoin(types.DefaultMintDenom, math.NewInt(26_000_000)),
						RemainCollateral: sdk.NewCoin("atom", math.NewInt(1_000_000)),
					},
				},
			},
			expVaultStatus: []types.VaultStatus{types.LIQUIDATED},
			reserveBalances: sdk.NewCoins(
				sdk.NewCoin(types.DefaultMintDenom, math.NewInt(1_000_000)),
				sdk.NewCoin("atom", math.LegacyNewDec(25_000_000).QuoInt(math.NewInt(7)).Mul(math.LegacyNewDecWithPrec(5, 2)).TruncateInt()), // (25/7)*0.05
			),
			vaultOwnerBalances: sdk.NewCoins(
				sdk.NewCoin("atom", math.NewInt(1_000_000).Sub(math.LegacyNewDec(25_000_000).QuoInt(math.NewInt(7)).Mul(math.LegacyNewDecWithPrec(5, 2)).TruncateInt())),
			), // remain - penalty
		},
		{
			name: "single vault - remain collateral, not enough to cover debt, can reconstitute vault",
			liquidation: types.Liquidation{
				Denom: "atom",
				LiquidatingVaults: []*types.Vault{
					{
						Owner:            vaultOwnerAddr.String(),
						Debt:             sdk.NewCoin(types.DefaultMintDenom, math.NewInt(25_000_000)),
						CollateralLocked: sdk.NewCoin("atom", math.NewInt(5_000_000)), // lock 5 ATOM at price 8, ratio = 160%
						Status:           types.LIQUIDATING,
						LiquidationPrice: math.LegacyNewDecWithPrec(7, 0), // liquidate at price 7, ratio = 140%
					},
				},
				// Sold = 0
				// Remain collateral = 5
				VaultLiquidationStatus: map[uint64]*types.VaultLiquidationStatus{
					0: {
						Sold:             sdk.NewCoin(types.DefaultMintDenom, math.ZeroInt()),
						RemainCollateral: sdk.NewCoin("atom", math.NewInt(5_000_000)),
					},
				},
			},
			expVaultStatus: []types.VaultStatus{types.ACTIVE},
			reserveBalances: sdk.NewCoins(
				// penalty
				sdk.NewCoin("atom", math.LegacyNewDec(25_000_000).QuoInt(math.NewInt(7)).Mul(math.LegacyNewDecWithPrec(5, 2)).TruncateInt()), // (25/7)*0.05
			),
		},
		{
			name: "single vault - remain collateral, not enough to cover debt, can reconstitute vault",
			liquidation: types.Liquidation{
				Denom: "atom",
				LiquidatingVaults: []*types.Vault{
					{
						Owner:            vaultOwnerAddr.String(),
						Debt:             sdk.NewCoin(types.DefaultMintDenom, math.NewInt(25_000_000)),
						CollateralLocked: sdk.NewCoin("atom", math.NewInt(5_000_000)), // lock 5 ATOM at price 8, ratio = 160%
						Status:           types.LIQUIDATING,
						LiquidationPrice: math.LegacyNewDecWithPrec(7, 0), // liquidate at price 7, ratio = 140%
					},
				},
				// Sold = 0
				// Remain collateral = 5
				VaultLiquidationStatus: map[uint64]*types.VaultLiquidationStatus{
					0: {
						Sold:             sdk.NewCoin(types.DefaultMintDenom, math.ZeroInt()),
						RemainCollateral: sdk.NewCoin("atom", math.NewInt(5_000_000)),
					},
				},
			},
			expVaultStatus: []types.VaultStatus{types.ACTIVE},
			reserveBalances: sdk.NewCoins(
				// penalty
				sdk.NewCoin("atom", math.LegacyNewDec(25_000_000).QuoInt(math.NewInt(7)).Mul(math.LegacyNewDecWithPrec(5, 2)).TruncateInt()), // (25/7)*0.05
			),
		},
		{
			name: "single vault - remain collateral, not enough to cover debt, can not reconstitute vault",
			liquidation: types.Liquidation{
				Denom: "atom",
				LiquidatingVaults: []*types.Vault{
					{
						Owner:            vaultOwnerAddr.String(),
						Debt:             sdk.NewCoin(types.DefaultMintDenom, math.NewInt(25_000_000)),
						CollateralLocked: sdk.NewCoin("atom", math.NewInt(5_000_000)), // lock 5 ATOM at price 8, ratio = 160%
						Status:           types.LIQUIDATING,
						LiquidationPrice: math.LegacyNewDecWithPrec(7, 0), // liquidate at price 7, ratio = 140%
					},
				},
				// Sold 1 at 7
				// Sold = 7
				// Remain collateral = 4
				VaultLiquidationStatus: map[uint64]*types.VaultLiquidationStatus{
					0: {
						Sold:             sdk.NewCoin(types.DefaultMintDenom, math.NewInt(7_000_000)),
						RemainCollateral: sdk.NewCoin("atom", math.NewInt(4_000_000)),
					},
				},
			},
			expVaultStatus: []types.VaultStatus{types.LIQUIDATED},
			reserveBalances: sdk.NewCoins(
				// penalty
				sdk.NewCoin("atom", math.NewInt(4_000_000)), // (25/7)*0.05
			),
			shortfallAmount: sdk.NewCoin(types.DefaultMintDenom, math.NewInt(18_000_000)),
		},
	}
	for _, t := range tests {
		s.Run(t.name, func() {
			s.SetupTest()

			for _, vault := range t.liquidation.LiquidatingVaults {
				vaultId, vaultAddr := s.App.VaultsKeeper.GetVaultIdAndAddress(s.Ctx)
				fmt.Println("vaultId", vaultId)
				vault.Id = vaultId
				vault.Address = vaultAddr.String()

				err := s.App.VaultsKeeper.SetVault(s.Ctx, *vault)
				s.Require().NoError(err)

				// Fund collateral locked for vault
				lockCoins := sdk.NewCoins(t.liquidation.VaultLiquidationStatus[vaultId].RemainCollateral)
				s.App.BankKeeper.MintCoins(s.Ctx, types.ModuleName, lockCoins)
				s.App.BankKeeper.SendCoinsFromModuleToAccount(s.Ctx, types.ModuleName, vaultAddr, lockCoins)

				// Fund sold coins to vault Module
				soldCoins := sdk.NewCoins(t.liquidation.VaultLiquidationStatus[vaultId].Sold)
				s.App.BankKeeper.MintCoins(s.Ctx, types.ModuleName, soldCoins)
			}

			err, isShortfall, shortfallAmount := s.App.VaultsKeeper.Liquidate(s.Ctx, t.liquidation)
			fmt.Println("errrrr", err, isShortfall, shortfallAmount)

			if t.reserveBalances != nil {
				reserveModuleAddr := s.App.AccountKeeper.GetModuleAddress(types.ReserveModuleName)
				reserveBalance := s.App.BankKeeper.GetAllBalances(s.Ctx, reserveModuleAddr)
				fmt.Println("reserve balances", reserveBalance)
				fmt.Println("test case reserve balances", t.reserveBalances)
				// t.reserveBalances.Sort()
				s.Require().Equal(reserveBalance, t.reserveBalances)
			}

			if t.vaultOwnerBalances != nil {
				ownerBalances := s.App.BankKeeper.GetAllBalances(s.Ctx, vaultOwnerAddr)
				s.Require().Equal(ownerBalances, t.vaultOwnerBalances)
			}

			if !t.shortfallAmount.IsNil() {
				s.Require().True(isShortfall)
				s.Require().Equal(t.shortfallAmount, shortfallAmount)
			}

			for i, vault := range t.liquidation.LiquidatingVaults {
				updatedVault, err := s.App.VaultsKeeper.GetVault(s.Ctx, vault.Id)
				fmt.Println("updated vault", updatedVault)
				s.Require().NoError(err)
				s.Require().Equal(updatedVault.Status, t.expVaultStatus[i])
			}


		})
	}
}

// func (s *KeeperTestSuite) TestGetLiquidateVaults() {
// 	s.SetupTest()
// 	var (
// 		denom1        = "atom"
// 		denom2        = "osmo"
// 		coin          = sdk.NewCoin(denom1, math.NewInt(1000))
// 		coinMintToAcc = sdk.NewCoin(denom1, math.NewInt(1000000))
// 		maxDebt       = math.NewInt(10000)
// 	)

// 	tests := []struct {
// 		name       string
// 		setup      func()
// 		vaultId    uint64
// 		sender     sdk.AccAddress
// 		collateral sdk.Coin
// 	}{
// 		{
// 			name: "success",
// 			setup: func() {
// 				err := s.k.ActiveCollateralAsset(s.Ctx, denom1, math.LegacyMustNewDecFromStr("0.1"), math.LegacyMustNewDecFromStr("0.1"), maxDebt)
// 				s.Require().NoError(err)
// 				err = s.k.ActiveCollateralAsset(s.Ctx, denom2, math.LegacyMustNewDecFromStr("0.1"), math.LegacyMustNewDecFromStr("0.1"), maxDebt)
// 				s.Require().NoError(err)

// 				vault := types.Vault{
// 					Owner:            s.TestAccs[0].String(),
// 					Debt:             sdk.NewCoin(denom1, maxDebt),
// 					CollateralLocked: sdk.NewCoin(denom1, maxDebt),
// 					Status:           types.ACTIVE,
// 				}
// 				err = s.k.SetVault(s.Ctx, vault)
// 				s.Require().NoError(err)

// 				err = s.App.BankKeeper.MintCoins(s.Ctx, types.ModuleName, sdk.NewCoins(coinMintToAcc))
// 				s.Require().NoError(err)
// 				err = s.App.BankKeeper.SendCoinsFromModuleToAccount(s.Ctx, types.ModuleName, s.TestAccs[0], sdk.NewCoins(coinMintToAcc))
// 				s.Require().NoError(err)
// 			},
// 			vaultId:    0,
// 			sender:     s.TestAccs[0],
// 			collateral: coin,
// 		},
// 	}
// 	for _, t := range tests {
// 		s.Run(t.name, func() {
// 			t.setup()
// 			vaults, prices, err := s.k.GetLiquidateVaults(s.Ctx)
// 			s.Require().NoError(err)

// 			// current price = 1, vaults is empty,
// 			s.Require().Equal(2, len(prices))
// 			s.Require().Equal(0, len(vaults))
// 		})
// 	}
// }
