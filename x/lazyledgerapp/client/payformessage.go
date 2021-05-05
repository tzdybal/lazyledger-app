package client

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client/tx"
	ctypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/lazyledger/lazyledger-app/x/lazyledgerapp/types"
)

func PayForMessage(ctx context.Context, namespaceID []byte, message []byte, from []byte) {
	sdkCtx := ctypes.UnwrapSDKContext(ctx)
	sdkCtx.
	// get info on the key
	keyInfo, err := clientCtx.Keyring.Key(accName)
	if err != nil {
		return err
	}
	// create the PayForMessage
	pfmMsg, err := types.NewMsgWirePayForMessage(
		namespace,
		message,
		keyInfo.GetPubKey().Bytes(),
		&types.TransactionFee{}, // transaction fee is not yet used
		types.SquareSize,
	)
	if err != nil {
		return err
	}

	// sign the PayForMessage's ShareCommitments
	err = pfmMsg.SignShareCommitments(accName, clientCtx.Keyring)
	if err != nil {
		return err
	}

	// run message checks
	if err = pfmMsg.ValidateBasic(); err != nil {
		return err
	}
	return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), pfmMsg)
}
