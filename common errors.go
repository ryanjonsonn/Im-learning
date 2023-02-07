func MakeEncodingConfig() appparams.EncodingConfig {
	encodingConfig := appparams.MakeEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
  // increment sequence of all signers
		for _, addr := range sigTx.GetSigners() {
			AccountTrack[addr.String()]++
			if AccountTrack[addr.String()] > maxTxs {
				panic(fmt.Sprintf("maximum txs in block is %d ", maxTxs))
			}
