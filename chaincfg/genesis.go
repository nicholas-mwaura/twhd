// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"time"

	"github.com/nicholas-mwaura/twhd/chaincfg/chainhash"
	"github.com/nicholas-mwaura/twhd/wire"
)

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network, regression test network, and test network (version 3).
var genesisCoinbaseTx = wire.MsgTx{
	Version: 1,
	TxIn: []*wire.TxIn{
		{
			PreviousOutPoint: wire.OutPoint{
				Hash:  chainhash.Hash{},
				Index: 0xffffffff,
			},
			SignatureScript: []byte{
				0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04, 0x4c, /* |.......L| */
				0x51, 0x54, 0x72, 0x75, 0x73, 0x74, 0x6c, 0x65, /* |QTrustle| */
				0x73, 0x73, 0x20, 0x52, 0x65, 0x77, 0x61, 0x72, /* |ss Rewar| */
				0x64, 0x73, 0x20, 0x46, 0x6f, 0x72, 0x20, 0x53, /* |ds For S| */
				0x75, 0x73, 0x74, 0x61, 0x69, 0x6e, 0x61, 0x62, /* |ustainab| */
				0x6c, 0x65, 0x20, 0x45, 0x6e, 0x65, 0x72, 0x67, /* |le Energ| */
				0x79, 0x20, 0x41, 0x64, 0x6f, 0x70, 0x74, 0x69, /* |y Adopti| */
				0x6f, 0x6e, 0x20, 0x2d, 0x20, 0x55, 0x4e, 0x20, /* |on - UN | */
				0x53, 0x44, 0x47, 0x73, 0x20, 0x37, 0x20, 0x39, /* |SDGs 7 9| */
				0x20, 0x31, 0x31, 0x20, 0x31, 0x33, 0x20, 0x2c, /* |11 13 , | */
				0x20, 0x32, 0x2f, 0x30, 0x32, 0x2f, 0x32, 0x30, /* |2/02/20 | */
				0x32, 0x30,                                     /* |20| */
			},
			Sequence: 0xffffffff,
		},
	},
	TxOut: []*wire.TxOut{
		{
			Value: 0x12a05f200,
			PkScript: []byte{
				0x41, 0x04, 0xf9, 0x22, 0x79, 0x3e, 0x1e, 0x9f, /* |A.."y>..| */
				0xd9, 0x53, 0x40, 0x3b, 0x47, 0x5b, 0xbf, 0xe8, /* |.S@;G[..| */
				0xea, 0x32, 0x16, 0xb8, 0xb7, 0x08, 0x16, 0x68, /* |.2.....h| */
				0x53, 0xc0, 0x2f, 0x3f, 0x6b, 0x95, 0x23, 0x2c, /* |S./?k.#,| */
				0xf9, 0xe3, 0x6b, 0xcc, 0x29, 0x91, 0xdb, 0xef, /* |..k.)...| */
				0x13, 0x88, 0xf4, 0x2b, 0x5a, 0x72, 0x70, 0x0e, /* |...+Zrp.| */
				0xc8, 0x37, 0x6a, 0x79, 0xee, 0x4f, 0xfc, 0x00, /* |.7jy.O..| */
				0x23, 0x55, 0xee, 0x39, 0x0d, 0x19, 0x2b, 0x27, /* |#U.9..+'| */
				0xc9, 0x71, 0xac,                               /* |.q.| */
			},
		},
	},
	LockTime: 0,
}

// genesisHash is the hash of the first block in the block chain for the main
// network (genesis block).
var genesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x8a, 0xd6, 0xe3, 0x24, 0xbe, 0xc1, 0x30, 0xd4,
	0x48, 0xe3, 0x08, 0x85, 0x6d, 0x11, 0xa9, 0xd4,
	0xb5, 0xfd, 0x19, 0xc2, 0x04, 0x80, 0x07, 0x9a,
	0x00, 0xe8, 0xca, 0xc9, 0x42, 0x61, 0x00, 0x00,
})

//genesisMerkleRoot is the hash of the first transaction in the genesis block
//for the main network.
var genesisMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xab, 0x4c, 0x83, 0xb8, 0xe4, 0x5c, 0x5b, 0xdb,
	0x44, 0x23, 0x2f, 0x17, 0xea, 0x6a, 0xf5, 0xc0,
	0x2e, 0xe0, 0x40, 0x97, 0x14, 0xc1, 0xe0, 0x9b,
	0xd8, 0xed, 0x12, 0x09, 0xb7, 0x4c, 0x04, 0xb8,
})

//genesisBlock defines the genesis block of the block chain which serves as the
//public transaction ledger for the main network.
var genesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: genesisMerkleRoot,        // b8044cb70912edd89be0c1149740e02ec0f56aea172f2344db5b5ce4b8834cab
		Timestamp:  time.Unix(1580658322, 0), // 2020-02-02 15:45:22 +0000 UTC
		Bits:       0x1f00ffff,               // 520159231 [0000ffff0000000000000000000000000000000000000000000000000000]
		Nonce:      0xf9f8    ,               // 63992
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// regTestGenesisHash is the hash of the first block in the block chain for the
// regression test network (genesis block).
var regTestGenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x8f, 0x82, 0x13, 0x61, 0x10, 0x80, 0xc7, 0xa9,
	0xf0, 0xd9, 0xd2, 0x85, 0x4e, 0xef, 0x83, 0x0a,
	0xde, 0xdc, 0x04, 0xdf, 0x2d, 0x1a, 0xfa, 0x5b,
	0x3e, 0xd0, 0x20, 0x18, 0x01, 0x30, 0xc3, 0xc4,
})


// regTestGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the regression test network.  It is the same as the merkle root for
// the main network.
//var regTestGenesisMerkleRoot = genesisMerkleRoot
var regTestGenesisMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{
        0x3b, 0xa3, 0xed, 0xfd, 0x7a, 0x7b, 0x12, 0xb2,
	0x7a, 0xc7, 0x2c, 0x3e, 0x67, 0x76, 0x8f, 0x61,
	0x7f, 0xc8, 0x1b, 0xc3, 0x88, 0x8a, 0x51, 0x32,
	0x3a, 0x9f, 0xb8, 0xaa, 0x4b, 0x1e, 0x5e, 0x4a,
})

// regTestGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the regression test network.
var regTestGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: regTestGenesisMerkleRoot, // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1580662895, 0), // 2020-02-02 17:01:35 +0000 UTC
		Bits:       0x207fffff,               // 545259519 [7fffff0000000000000000000000000000000000000000000000000000000000]
		Nonce:      1,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// testNet3GenesisHash is the hash of the first block in the block chain for the
// test network (version 3).
var testNet3GenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x07, 0xc5, 0xe2, 0x09, 0x09, 0x0e, 0xc6, 0x0d,
	0xb9, 0x4f, 0xca, 0xfc, 0xb9, 0xcc, 0xcc, 0xc5,
	0x12, 0x1e, 0x13, 0x3e, 0x4b, 0x9e, 0xe6, 0xaa,
	0x80, 0x88, 0x95, 0x78, 0x7d, 0x15, 0x00, 0x00,
})

// testNet3GenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the test network (version 3).  It is the same as the merkle root
// for the main network.
//var testNet3GenesisMerkleRoot = genesisMerkleRoot
var testNet3GenesisMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{
        0xab, 0x4c, 0x83, 0xb8, 0xe4, 0x5c, 0x5b, 0xdb,
	0x44, 0x23, 0x2f, 0x17, 0xea, 0x6a, 0xf5, 0xc0,
	0x2e, 0xe0, 0x40, 0x97, 0x14, 0xc1, 0xe0, 0x9b,
	0xd8, 0xed, 0x12, 0x09, 0xb7, 0x4c, 0x04, 0xb8,
})
// testNet3GenesisBlock defines the genesis block of the block chain which
// serves as the public transaction ledger for the test network (version 3).
var testNet3GenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},          // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: testNet3GenesisMerkleRoot, // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1580661362, 0),  // 2020-02-02 17:36:02 +0000 UTC
		Bits:       0x1f00ffff,                //520159231 [00000000ffff0000000000000000000000000000000000000000000000000000]
		Nonce:      0x26331,                   //156465
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// simNetGenesisHash is the hash of the first block in the block chain for the
// simulation test network.
var simNetGenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xf6, 0x7a, 0xd7, 0x69, 0x5d, 0x9b, 0x66, 0x2a,
	0x72, 0xff, 0x3d, 0x8e, 0xdb, 0xbb, 0x2d, 0xe0,
	0xbf, 0xa6, 0x7b, 0x13, 0x97, 0x4b, 0xb9, 0x91,
	0x0d, 0x11, 0x6d, 0x5c, 0xbd, 0x86, 0x3e, 0x68,
})

// simNetGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the simulation test network.  It is the same as the merkle root for
// the main network.
var simNetGenesisMerkleRoot = genesisMerkleRoot

// simNetGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the simulation test network.
var simNetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: simNetGenesisMerkleRoot,  // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1401292357, 0), // 2014-05-28 15:52:37 +0000 UTC
		Bits:       0x207fffff,               // 545259519 [7fffff0000000000000000000000000000000000000000000000000000000000]
		Nonce:      2,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// sigNetGenesisHash is the hash of the first block in the block chain for the
// signet test network.
var sigNetGenesisHash = chainhash.Hash{
	0xf6, 0x1e, 0xee, 0x3b, 0x63, 0xa3, 0x80, 0xa4,
	0x77, 0xa0, 0x63, 0xaf, 0x32, 0xb2, 0xbb, 0xc9,
	0x7c, 0x9f, 0xf9, 0xf0, 0x1f, 0x2c, 0x42, 0x25,
	0xe9, 0x73, 0x98, 0x81, 0x08, 0x00, 0x00, 0x00,
}

// sigNetGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the signet test network. It is the same as the merkle root for
// the main network.
var sigNetGenesisMerkleRoot = genesisMerkleRoot

// sigNetGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the signet test network.
var sigNetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: sigNetGenesisMerkleRoot,  // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1598918400, 0), // 2020-09-01 00:00:00 +0000 UTC
		Bits:       0x1e0377ae,               // 503543726 [00000377ae000000000000000000000000000000000000000000000000000000]
		Nonce:      52613770,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}
