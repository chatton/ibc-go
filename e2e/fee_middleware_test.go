package e2e

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/ibc-go/v3/e2e/e2efee"
	"github.com/cosmos/ibc-go/v3/e2e/testsuite"
	feetypes "github.com/cosmos/ibc-go/v3/modules/apps/29-fee/types"
	"github.com/cosmos/ibc-go/v3/testing/simapp"
	simappparams "github.com/cosmos/ibc-go/v3/testing/simapp/params"
	"github.com/strangelove-ventures/ibctest/chain/cosmos"
	"github.com/strangelove-ventures/ibctest/ibc"
	"github.com/strangelove-ventures/ibctest/test"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

var (
	encCfg    simappparams.EncodingConfig
	txBuilder client.TxBuilder
)

func init() {
	encCfg = simapp.MakeTestEncodingConfig()
	txBuilder = encCfg.TxConfig.NewTxBuilder()

}
func TestFeeMiddlewareTestSuite(t *testing.T) {
	suite.Run(t, new(FeeMiddlewareTestSuite))
}

type FeeMiddlewareTestSuite struct {
	testsuite.E2ETestSuite
}

// TODO: wip test, depends on https://github.com/strangelove-ventures/ibctest/issues/172
func (s *FeeMiddlewareTestSuite) _TestSyncSingleSender() {

	//t.Run("Relayer wallets can be recovered", s.AssertRelayerWalletsCanBeRecovered(ctx, relayer))
	//
	//srcRelayerWallet, dstRelayerWallet, err := s.GetRelayerWallets(relayer)
	//t.Run("Relayer wallets can be fetched", func(t *testing.T) {
	//	s.Req.NoError(err)
	//})
	//
	//
	//feetypes.MsgPayPacketFee{}
	//feetypes.NewMsgRegisterCounterpartyPayee()
	//transfertypes.NewMsgTransfer()
	//

	t := s.T()
	ctx := context.TODO()

	srcChain, _ := s.GetChains()

	s.CreateChainsRelayerAndChannel(ctx, e2efee.FeeMiddlewareChannelOptions())

	startingTokenAmount := int64(10_000_000)
	srcChainSenderOne := s.CreateUserOnSourceChain(ctx, startingTokenAmount)

	srChainDenom := srcChain.Config().Denom

	defaultRecvFee := sdk.Coins{sdk.Coin{Denom: srChainDenom, Amount: sdk.NewInt(100)}}
	defaultAckFee := sdk.Coins{sdk.Coin{Denom: srChainDenom, Amount: sdk.NewInt(200)}}
	defaultTimeoutFee := sdk.Coins{sdk.Coin{Denom: srChainDenom, Amount: sdk.NewInt(300)}}
	fee := feetypes.NewFee(defaultRecvFee, defaultAckFee, defaultTimeoutFee)

	payPacketFeeMsg := feetypes.NewMsgPayPacketFee(fee, "transfer", "channel-0", srcChainSenderOne.Bech32Address(srcChain.Config().Bech32Prefix), nil)

	err := txBuilder.SetMsgs(payPacketFeeMsg)
	s.Req.NoError(err)

	//txBuilder.SetGasLimit(...)
	//txBuilder.SetFeeAmount(...)
	txBuilder.SetMemo("ibc-test")

	//txBuilder.SetTimeoutHeight(...)
	//info, err := srcChain.ChainNodes[0].GetKey(srcChainSenderOne.KeyName)
	priv1, _, _ := testdata.KeyTestPubAddr()
	privs := []cryptotypes.PrivKey{priv1}
	accNums := []uint64{0} // The accounts' account numbers
	accSeqs := []uint64{0} // The accounts' sequence numbers

	var sigsV2 []signing.SignatureV2
	for i, priv := range privs {
		sigV2 := signing.SignatureV2{
			PubKey: priv.PubKey(),
			Data: &signing.SingleSignatureData{
				SignMode:  encCfg.TxConfig.SignModeHandler().DefaultMode(),
				Signature: nil,
			},
			Sequence: accSeqs[i],
		}

		sigsV2 = append(sigsV2, sigV2)
	}
	err = txBuilder.SetSignatures(sigsV2...)
	s.Req.NoError(err)

	// Second round: all signer infos are set, so each signer can sign.
	sigsV2 = []signing.SignatureV2{}
	for i, priv := range privs {
		signerData := xauthsigning.SignerData{
			ChainID:       srcChain.Config().ChainID,
			AccountNumber: accNums[i],
			Sequence:      accSeqs[i],
		}
		sigV2, err := tx.SignWithPrivKey(
			encCfg.TxConfig.SignModeHandler().DefaultMode(), signerData,
			txBuilder, priv, encCfg.TxConfig, accSeqs[i])
		s.Req.NoError(err)
		sigsV2 = append(sigsV2, sigV2)
	}
	err = txBuilder.SetSignatures(sigsV2...)
	s.Req.NoError(err)

	//txBuilder.GetTx(
	//s.Req.NoError(err)
	//
	txJSONBytes, err := encCfg.TxConfig.TxJSONEncoder()(txBuilder.GetTx())
	s.Req.NoError(err)

	txJSON := string(txJSONBytes)
	t.Logf("JSON: %s", txJSON)

}

func (s *FeeMiddlewareTestSuite) TestAsyncMultipleSenders() {
	t := s.T()
	ctx := context.TODO()

	srcChain, dstChain := s.GetChains()

	relayer, srcChainChannelInfo := s.CreateChainsRelayerAndChannel(ctx, e2efee.FeeMiddlewareChannelOptions())

	startingTokenAmount := int64(10_000_000)

	srcChainSenderOne := s.CreateUserOnSourceChain(ctx, startingTokenAmount)
	srcChainSenderTwo := s.CreateUserOnSourceChain(ctx, startingTokenAmount)
	dstChainWallet := s.CreateUserOnDestinationChain(ctx, startingTokenAmount)

	t.Run("Relayer wallets can be recovered", s.AssertRelayerWalletsCanBeRecovered(ctx, relayer))

	srcRelayerWallet, dstRelayerWallet, err := s.GetRelayerWallets(relayer)
	t.Run("Relayer wallets can be fetched", func(t *testing.T) {
		s.Req.NoError(err)
	})

	s.Req.NoError(test.WaitForBlocks(ctx, 10, srcChain, dstChain), "failed to wait for blocks")

	t.Run("Register Counter Party Payee", s.AssertCounterPartyPayeeCanBeRegistered(ctx, dstChain, dstRelayerWallet.Address, srcRelayerWallet.Address, srcChainChannelInfo.Counterparty.PortID, srcChainChannelInfo.Counterparty.ChannelID))
	t.Run("Verify Counter Party Payee", s.AssertCounterPartyPayeeCanBeVerified(ctx, dstChain, dstRelayerWallet.Address, srcChainChannelInfo.Counterparty.ChannelID, srcRelayerWallet.Address))
	t.Run("Test fee middleware with multiple senders", func(t *testing.T) {

		chain1WalletToChain2WalletAmount := ibc.WalletAmount{
			Address: dstChainWallet.Bech32Address(dstChain.Config().Bech32Prefix), // destination address
			Denom:   srcChain.Config().Denom,
			Amount:  10000,
		}

		var srcTx ibc.Tx

		t.Run("Test IBC transfer", func(t *testing.T) {
			t.Run("Send IBC transfer", func(t *testing.T) {
				var err error
				srcTx, err = srcChain.SendIBCTransfer(ctx, srcChainChannelInfo.ChannelID, srcChainSenderOne.KeyName, chain1WalletToChain2WalletAmount, nil)
				s.Req.NoError(err)
				s.Req.NoError(srcTx.Validate(), "source ibc transfer tx is invalid")
			})

			expected := startingTokenAmount - chain1WalletToChain2WalletAmount.Amount - srcChain.GetGasFeesInNativeDenom(srcTx.GasSpent)
			t.Run("Verify tokens have been escrowed", s.AssertChainNativeBalance(ctx, srcChain, srcChainSenderOne, expected))
		})

		t.Run("Test Packet Fees", func(t *testing.T) {

			recvFee := int64(50)
			ackFee := int64(25)
			timeoutFee := int64(10)

			t.Run("Pay packet fee", func(t *testing.T) {

				t.Run("Before paying packet fee there should be no incentivized packets", s.AssertEmptyPackets(ctx, srcChain, srcChainChannelInfo.PortID, srcChainChannelInfo.ChannelID))

				t.Run("Paying packet fee should succeed", func(t *testing.T) {
					s.Req.NoError(e2efee.PayPacketFee(ctx, srcChain, srcChainSenderOne.KeyName, srcChainChannelInfo.PortID, srcChainChannelInfo.ChannelID, 1, recvFee, ackFee, timeoutFee))
					// wait so that incentivised packets will show up
					time.Sleep(5 * time.Second)
				})

				t.Run("Paying packet fee with second sender should succeed", func(t *testing.T) {
					s.Req.NoError(e2efee.PayPacketFee(ctx, srcChain, srcChainSenderTwo.KeyName, srcChainChannelInfo.PortID, srcChainChannelInfo.ChannelID, 1, recvFee, ackFee, timeoutFee))
					// wait so that incentivised packets will show up
					time.Sleep(5 * time.Second)
				})

				t.Run("After paying packet fee there should be incentivized packets", func(t *testing.T) {
					packets, err := e2efee.QueryPackets(ctx, srcChain, srcChainChannelInfo.PortID, srcChainChannelInfo.ChannelID)
					s.Req.NoError(err)
					s.Req.Len(packets.IncentivizedPackets, 1)
					s.Req.Len(packets.IncentivizedPackets[0].PacketFees, 2)

					expectedRecv, expectedAck, exectedTimeout := convertFeeAmountsToCoins(srcChain.Config().Denom, recvFee, ackFee, timeoutFee)

					t.Run("test first packet fee", func(t *testing.T) {
						actualFee := packets.IncentivizedPackets[0].PacketFees[0].Fee
						s.Req.True(actualFee.RecvFee.IsEqual(expectedRecv))
						s.Req.True(actualFee.AckFee.IsEqual(expectedAck))
						s.Req.True(actualFee.TimeoutFee.IsEqual(exectedTimeout))
					})
					t.Run("test second packet fee", func(t *testing.T) {
						actualFee := packets.IncentivizedPackets[0].PacketFees[1].Fee
						s.Req.True(actualFee.RecvFee.IsEqual(expectedRecv))
						s.Req.True(actualFee.AckFee.IsEqual(expectedAck))
						s.Req.True(actualFee.TimeoutFee.IsEqual(exectedTimeout))
					})
				})

				expectedSenderOneBal := startingTokenAmount - chain1WalletToChain2WalletAmount.Amount - srcChain.GetGasFeesInNativeDenom(srcTx.GasSpent) - recvFee - ackFee - timeoutFee
				t.Run("Balance from first sender should be lowered by sum of recv ack and timeout and IBC transfer amount", s.AssertChainNativeBalance(ctx, srcChain, srcChainSenderOne, expectedSenderOneBal))

				expectedSenderTwoBal := startingTokenAmount - recvFee - ackFee - timeoutFee
				t.Run("Balance from second sender should be lowered by sum of recv ack and timeout (not IBC transfer amount)", s.AssertChainNativeBalance(ctx, srcChain, srcChainSenderTwo, expectedSenderTwoBal))
			})

			t.Run("Start relayer", func(t *testing.T) {
				s.StartRelayer(relayer)
			})

			s.Req.NoError(test.WaitForBlocks(ctx, 5, srcChain, dstChain), "failed to wait for blocks")

			t.Run("Packets should have been relayed", s.AssertEmptyPackets(ctx, srcChain, srcChainChannelInfo.PortID, srcChainChannelInfo.ChannelID))

			// once the relayer has relayed the packets, the timeout fee should be refunded.
			gasFee := srcChain.GetGasFeesInNativeDenom(srcTx.GasSpent)
			expectedSenderOneBal := startingTokenAmount - chain1WalletToChain2WalletAmount.Amount - gasFee - ackFee - recvFee
			t.Run("Verify timeout fee is refunded on successful relay of packets for first sender", s.AssertChainNativeBalance(ctx, srcChain, srcChainSenderOne, expectedSenderOneBal))

			expectedSenderTwoBal := startingTokenAmount - ackFee - recvFee
			t.Run("Verify timeout fee is refunded on successful relay of packets for second sender", s.AssertChainNativeBalance(ctx, srcChain, srcChainSenderTwo, expectedSenderTwoBal))
		})
	})
}

func (s *FeeMiddlewareTestSuite) TestAsyncSingleSender() {
	t := s.T()
	ctx := context.TODO()

	srcChain, dstChain := s.GetChains()

	relayer, srcChainChannelInfo := s.CreateChainsRelayerAndChannel(ctx, e2efee.FeeMiddlewareChannelOptions())

	startingTokenAmount := int64(10_000_000)

	srcChainWallet := s.CreateUserOnSourceChain(ctx, startingTokenAmount)
	dstChainWallet := s.CreateUserOnDestinationChain(ctx, startingTokenAmount)

	t.Run("Relayer wallets can be recovered", func(t *testing.T) {
		s.Req.NoError(s.RecoverRelayerWallets(ctx, relayer))
	})

	srcRelayerWallet, dstRelayerWallet, err := s.GetRelayerWallets(relayer)
	t.Run("Relayer wallets can be fetched", func(t *testing.T) {
		s.Req.NoError(err)
	})

	s.Req.NoError(test.WaitForBlocks(ctx, 10, srcChain, dstChain), "failed to wait for blocks")

	t.Run("Register Counter Party Payee", func(t *testing.T) {
		s.Req.NoError(e2efee.RegisterCounterPartyPayee(ctx, dstChain, dstRelayerWallet.Address, srcRelayerWallet.Address, srcChainChannelInfo.Counterparty.PortID, srcChainChannelInfo.Counterparty.ChannelID))
		// give some time for update
		time.Sleep(time.Second * 5)
	})

	t.Run("Verify Counter Party Payee", func(t *testing.T) {
		address, err := e2efee.QueryCounterPartyPayee(ctx, dstChain, dstRelayerWallet.Address, srcChainChannelInfo.Counterparty.ChannelID)
		s.Req.NoError(err)
		s.Req.Equal(srcRelayerWallet.Address, address)
	})

	chain1WalletToChain2WalletAmount := ibc.WalletAmount{
		Address: dstChainWallet.Bech32Address(dstChain.Config().Bech32Prefix), // destination address
		Denom:   srcChain.Config().Denom,
		Amount:  10000,
	}

	var srcTx ibc.Tx
	t.Run("Send IBC transfer", func(t *testing.T) {
		var err error
		srcTx, err = srcChain.SendIBCTransfer(ctx, srcChainChannelInfo.ChannelID, srcChainWallet.KeyName, chain1WalletToChain2WalletAmount, nil)
		s.Req.NoError(err)
		s.Req.NoError(srcTx.Validate(), "source ibc transfer tx is invalid")
	})

	t.Run("Verify tokens have been escrowed", func(t *testing.T) {
		actualBalance, err := s.GetSourceChainNativeBalance(ctx, srcChainWallet)
		s.Req.NoError(err)

		expected := startingTokenAmount - chain1WalletToChain2WalletAmount.Amount - srcChain.GetGasFeesInNativeDenom(srcTx.GasSpent)
		s.Req.Equal(expected, actualBalance)
	})

	recvFee := int64(50)
	ackFee := int64(25)
	timeoutFee := int64(10)

	t.Run("Pay packet fee", func(t *testing.T) {
		t.Run("Before paying packet fee there should be no incentivized packets", s.AssertEmptyPackets(ctx, srcChain, srcChainChannelInfo.PortID, srcChainChannelInfo.ChannelID))

		t.Run("Paying packet fee should succeed", func(t *testing.T) {
			s.Req.NoError(e2efee.PayPacketFee(ctx, srcChain, srcChainWallet.KeyName, srcChainChannelInfo.PortID, srcChainChannelInfo.ChannelID, 1, recvFee, ackFee, timeoutFee))

			// wait so that incentivised packets will show up
			time.Sleep(5 * time.Second)
		})

		t.Run("After paying packet fee there should be incentivized packets", func(t *testing.T) {
			packets, err := e2efee.QueryPackets(ctx, srcChain, srcChainChannelInfo.PortID, srcChainChannelInfo.ChannelID)
			s.Req.NoError(err)
			s.Req.Len(packets.IncentivizedPackets, 1)
			actualFee := packets.IncentivizedPackets[0].PacketFees[0].Fee

			expectedRecv, expectedAck, exectedTimeout := convertFeeAmountsToCoins(srcChain.Config().Denom, recvFee, ackFee, timeoutFee)
			s.Req.True(actualFee.RecvFee.IsEqual(expectedRecv))
			s.Req.True(actualFee.AckFee.IsEqual(expectedAck))
			s.Req.True(actualFee.TimeoutFee.IsEqual(exectedTimeout))
		})
	})

	t.Run("Balance should be lowered by sum of recv ack and timeout", func(t *testing.T) {
		// The balance should be lowered by the sum of the recv, ack and timeout fees.
		actualBalance, err := s.GetSourceChainNativeBalance(ctx, srcChainWallet)
		s.Req.NoError(err)

		expected := startingTokenAmount - chain1WalletToChain2WalletAmount.Amount - srcChain.GetGasFeesInNativeDenom(srcTx.GasSpent) - recvFee - ackFee - timeoutFee
		s.Req.Equal(expected, actualBalance)
	})

	t.Run("Start relayer", func(t *testing.T) {
		s.StartRelayer(relayer)
	})

	s.Req.NoError(test.WaitForBlocks(ctx, 5, srcChain, dstChain), "failed to wait for blocks")

	t.Run("Packets should have been relayed", s.AssertEmptyPackets(ctx, srcChain, srcChainChannelInfo.PortID, srcChainChannelInfo.ChannelID))

	t.Run("Verify timeout fee is refunded on successful relay of packets", func(t *testing.T) {

		actualBalance, err := s.GetSourceChainNativeBalance(ctx, srcChainWallet)
		s.Req.NoError(err)

		gasFee := srcChain.GetGasFeesInNativeDenom(srcTx.GasSpent)
		// once the relayer has relayed the packets, the timeout fee should be refunded.
		expected := startingTokenAmount - chain1WalletToChain2WalletAmount.Amount - gasFee - ackFee - recvFee
		s.Req.Equal(expected, actualBalance)
	})
}

func (s *FeeMiddlewareTestSuite) TestAsyncSingleSenderTimesOut() {
	t := s.T()
	ctx := context.TODO()

	srcChain, dstChain := s.GetChains()

	relayer, srcChainChannelInfo := s.CreateChainsRelayerAndChannel(ctx, e2efee.FeeMiddlewareChannelOptions())

	startingTokenAmount := int64(10_000_000)

	srcChainWallet := s.CreateUserOnSourceChain(ctx, startingTokenAmount)
	dstChainWallet := s.CreateUserOnDestinationChain(ctx, startingTokenAmount)

	t.Run("Relayer wallets can be recovered", func(t *testing.T) {
		s.Req.NoError(s.RecoverRelayerWallets(ctx, relayer))
	})

	srcRelayerWallet, dstRelayerWallet, err := s.GetRelayerWallets(relayer)
	t.Run("Relayer wallets can be fetched", func(t *testing.T) {
		s.Req.NoError(err)
	})

	s.Req.NoError(test.WaitForBlocks(ctx, 10, srcChain, dstChain), "failed to wait for blocks")

	t.Run("Register Counter Party Payee", func(t *testing.T) {
		s.Req.NoError(e2efee.RegisterCounterPartyPayee(ctx, dstChain, dstRelayerWallet.Address, srcRelayerWallet.Address, srcChainChannelInfo.Counterparty.PortID, srcChainChannelInfo.Counterparty.ChannelID))
		// give some time for update
		time.Sleep(time.Second * 5)
	})

	t.Run("Verify Counter Party Payee", func(t *testing.T) {
		address, err := e2efee.QueryCounterPartyPayee(ctx, dstChain, dstRelayerWallet.Address, srcChainChannelInfo.Counterparty.ChannelID)
		s.Req.NoError(err)
		s.Req.Equal(srcRelayerWallet.Address, address)
	})

	chain1WalletToChain2WalletAmount := ibc.WalletAmount{
		Address: dstChainWallet.Bech32Address(dstChain.Config().Bech32Prefix), // destination address
		Denom:   srcChain.Config().Denom,
		Amount:  10000,
	}

	var srcTx ibc.Tx
	t.Run("Send IBC transfer", func(t *testing.T) {
		var err error
		srcTx, err = srcChain.SendIBCTransfer(ctx, srcChainChannelInfo.ChannelID, srcChainWallet.KeyName, chain1WalletToChain2WalletAmount, &ibc.IBCTimeout{
			NanoSeconds: 100, // want it to timeout immediately
		})
		s.Req.NoError(err)
		s.Req.NoError(srcTx.Validate(), "source ibc transfer tx is invalid")
		time.Sleep(1 * time.Second) // cause timeout
	})

	t.Run("Verify tokens have been escrowed (relayer has not yet picked up the packet)", func(t *testing.T) {
		actualBalance, err := s.GetSourceChainNativeBalance(ctx, srcChainWallet)
		s.Req.NoError(err)

		expected := startingTokenAmount - chain1WalletToChain2WalletAmount.Amount - srcChain.GetGasFeesInNativeDenom(srcTx.GasSpent)
		s.Req.Equal(expected, actualBalance)
	})

	recvFee := int64(50)
	ackFee := int64(25)
	timeoutFee := int64(10)

	t.Run("Pay packet fee", func(t *testing.T) {
		t.Run("Before paying packet fee there should be no incentivized packets", s.AssertEmptyPackets(ctx, srcChain, srcChainChannelInfo.PortID, srcChainChannelInfo.ChannelID))

		t.Run("Paying packet fee should succeed", func(t *testing.T) {
			s.Req.NoError(e2efee.PayPacketFee(ctx, srcChain, srcChainWallet.KeyName, srcChainChannelInfo.PortID, srcChainChannelInfo.ChannelID, 1, recvFee, ackFee, timeoutFee))

			// wait so that incentivised packets will show up
			time.Sleep(5 * time.Second)
		})

		t.Run("After paying packet fee there should be incentivized packets", func(t *testing.T) {
			packets, err := e2efee.QueryPackets(ctx, srcChain, srcChainChannelInfo.PortID, srcChainChannelInfo.ChannelID)
			s.Req.NoError(err)
			s.Req.Len(packets.IncentivizedPackets, 1)
			actualFee := packets.IncentivizedPackets[0].PacketFees[0].Fee

			expectedRecv, expectedAck, exectedTimeout := convertFeeAmountsToCoins(srcChain.Config().Denom, recvFee, ackFee, timeoutFee)
			s.Req.True(actualFee.RecvFee.IsEqual(expectedRecv))
			s.Req.True(actualFee.AckFee.IsEqual(expectedAck))
			s.Req.True(actualFee.TimeoutFee.IsEqual(exectedTimeout))
		})
	})

	t.Run("Balance should be lowered by sum of recv ack and timeout", func(t *testing.T) {
		// The balance should be lowered by the sum of the recv, ack and timeout fees.
		actualBalance, err := s.GetSourceChainNativeBalance(ctx, srcChainWallet)
		s.Req.NoError(err)

		expected := startingTokenAmount - chain1WalletToChain2WalletAmount.Amount - srcChain.GetGasFeesInNativeDenom(srcTx.GasSpent) - recvFee - ackFee - timeoutFee
		s.Req.Equal(expected, actualBalance)
	})

	t.Run("Start relayer", func(t *testing.T) {
		s.StartRelayer(relayer)
	})

	s.Req.NoError(test.WaitForBlocks(ctx, 5, srcChain, dstChain), "failed to wait for blocks")

	t.Run("Funds recv and ack should be refunded as the packet timed out", func(t *testing.T) {
		actualBalance, err := s.GetSourceChainNativeBalance(ctx, srcChainWallet)
		s.Req.NoError(err)

		expected := startingTokenAmount - srcChain.GetGasFeesInNativeDenom(srcTx.GasSpent) - timeoutFee
		s.Req.Equal(expected, actualBalance)
	})
}

//  When packet is incentivized from single sender AND counterparty payee address is not set AND token transfer succeeds, then recv fees are refunded.
func (s *FeeMiddlewareTestSuite) TestAsyncSingleSenderNoCounterPartyAddress() {
	t := s.T()
	ctx := context.TODO()

	srcChain, dstChain := s.GetChains()

	relayer, srcChainChannelInfo := s.CreateChainsRelayerAndChannel(ctx, e2efee.FeeMiddlewareChannelOptions())

	startingTokenAmount := int64(10_000_000)

	srcChainWallet := s.CreateUserOnSourceChain(ctx, startingTokenAmount)
	dstChainWallet := s.CreateUserOnDestinationChain(ctx, startingTokenAmount)

	t.Run("Relayer wallets can be recovered", func(t *testing.T) {
		s.Req.NoError(s.RecoverRelayerWallets(ctx, relayer))
	})

	s.Req.NoError(test.WaitForBlocks(ctx, 10, srcChain, dstChain), "failed to wait for blocks")

	chain1WalletToChain2WalletAmount := ibc.WalletAmount{
		Address: dstChainWallet.Bech32Address(dstChain.Config().Bech32Prefix), // destination address
		Denom:   srcChain.Config().Denom,
		Amount:  10000,
	}

	var srcTx ibc.Tx
	t.Run("Send IBC transfer", func(t *testing.T) {
		var err error
		srcTx, err = srcChain.SendIBCTransfer(ctx, srcChainChannelInfo.ChannelID, srcChainWallet.KeyName, chain1WalletToChain2WalletAmount, nil)
		s.Req.NoError(err)
		s.Req.NoError(srcTx.Validate(), "source ibc transfer tx is invalid")
	})

	t.Run("Verify tokens have been escrowed", func(t *testing.T) {
		actualBalance, err := s.GetSourceChainNativeBalance(ctx, srcChainWallet)
		s.Req.NoError(err)

		expected := startingTokenAmount - chain1WalletToChain2WalletAmount.Amount - srcChain.GetGasFeesInNativeDenom(srcTx.GasSpent)
		s.Req.Equal(expected, actualBalance)
	})

	recvFee := int64(50)
	ackFee := int64(25)
	timeoutFee := int64(10)

	t.Run("Pay packet fee", func(t *testing.T) {
		t.Run("Before paying packet fee there should be no incentivized packets", s.AssertEmptyPackets(ctx, srcChain, srcChainChannelInfo.PortID, srcChainChannelInfo.ChannelID))

		t.Run("Paying packet fee should succeed", func(t *testing.T) {
			s.Req.NoError(e2efee.PayPacketFee(ctx, srcChain, srcChainWallet.KeyName, srcChainChannelInfo.PortID, srcChainChannelInfo.ChannelID, 1, recvFee, ackFee, timeoutFee))

			// wait so that incentivised packets will show up
			time.Sleep(5 * time.Second)
		})

		t.Run("After paying packet fee there should be incentivized packets", func(t *testing.T) {
			packets, err := e2efee.QueryPackets(ctx, srcChain, srcChainChannelInfo.PortID, srcChainChannelInfo.ChannelID)
			s.Req.NoError(err)
			s.Req.Len(packets.IncentivizedPackets, 1)
			actualFee := packets.IncentivizedPackets[0].PacketFees[0].Fee

			expectedRecv, expectedAck, exectedTimeout := convertFeeAmountsToCoins(srcChain.Config().Denom, recvFee, ackFee, timeoutFee)
			s.Req.True(actualFee.RecvFee.IsEqual(expectedRecv))
			s.Req.True(actualFee.AckFee.IsEqual(expectedAck))
			s.Req.True(actualFee.TimeoutFee.IsEqual(exectedTimeout))
		})
	})

	t.Run("Balance should be lowered by sum of recv ack and timeout", func(t *testing.T) {
		// The balance should be lowered by the sum of the recv, ack and timeout fees.
		actualBalance, err := s.GetSourceChainNativeBalance(ctx, srcChainWallet)
		s.Req.NoError(err)

		expected := startingTokenAmount - chain1WalletToChain2WalletAmount.Amount - srcChain.GetGasFeesInNativeDenom(srcTx.GasSpent) - recvFee - ackFee - timeoutFee
		s.Req.Equal(expected, actualBalance)
	})

	t.Run("Start relayer", func(t *testing.T) {
		s.StartRelayer(relayer)
	})

	s.Req.NoError(test.WaitForBlocks(ctx, 5, srcChain, dstChain), "failed to wait for blocks")

	t.Run("Packets should have been relayed", s.AssertEmptyPackets(ctx, srcChain, srcChainChannelInfo.PortID, srcChainChannelInfo.ChannelID))

	t.Run("Verify timeout fee and recv fee are refunded on successful relay of packets when there is no counter party address", func(t *testing.T) {

		actualBalance, err := s.GetSourceChainNativeBalance(ctx, srcChainWallet)
		s.Req.NoError(err)

		gasFee := srcChain.GetGasFeesInNativeDenom(srcTx.GasSpent)
		// once the relayer has relayed the packets, the timeout fee should be refunded.
		expected := startingTokenAmount - chain1WalletToChain2WalletAmount.Amount - gasFee - ackFee
		s.Req.Equal(expected, actualBalance)
	})
}

// Utility and assertion functions
// AssertCounterPartyPayeeCanBeRegistered attempts to register a counter party payee, and asserts there is no error.
func (s *FeeMiddlewareTestSuite) AssertCounterPartyPayeeCanBeRegistered(ctx context.Context, chain *cosmos.CosmosChain, relayerAddress, counterPartyPayee, portId, channelId string) func(t *testing.T) {
	return func(t *testing.T) {
		s.Req.NoError(e2efee.RegisterCounterPartyPayee(ctx, chain, relayerAddress, counterPartyPayee, portId, channelId))
		// give some time for update
		time.Sleep(time.Second * 5)
	}
}

// AssertCounterPartyPayeeCanBeVerified asserts that the given relayer address has the expected counter party address.
func (s *FeeMiddlewareTestSuite) AssertCounterPartyPayeeCanBeVerified(ctx context.Context, chain *cosmos.CosmosChain, relayerAddress, channelID, expectedAddress string) func(t *testing.T) {
	return func(t *testing.T) {
		actualAddress, err := e2efee.QueryCounterPartyPayee(ctx, chain, relayerAddress, channelID)
		s.Req.NoError(err)
		s.Req.Equal(expectedAddress, actualAddress)
	}
}

func convertFeeAmountsToCoins(denom string, recvFee, ackFee, timeoutFee int64) (sdk.Coins, sdk.Coins, sdk.Coins) {
	recvCoins := sdk.NewCoins(
		sdk.NewCoin(denom, sdk.NewInt(recvFee)),
	)
	ackCoins := sdk.NewCoins(
		sdk.NewCoin(denom, sdk.NewInt(ackFee)),
	)
	timeoutCoins := sdk.NewCoins(
		sdk.NewCoin(denom, sdk.NewInt(timeoutFee)),
	)

	return recvCoins, ackCoins, timeoutCoins
}