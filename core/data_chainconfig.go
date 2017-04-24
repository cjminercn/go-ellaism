// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// This file contains configuration literals.

package core

import (
	"math/big"

	"github.com/ethereumproject/go-ethereum/common"
)

var DefaultChainName = "mainnet"
var DefaultTestChainName = "testnet"

// DefaultConfig is the Ethereum Classic standard setup.
var DefaultConfig = &ChainConfig{
	Forks: []*Fork{
		{
			Name:  "Homestead",
			Block: big.NewInt(1150000),
			Features: []*ForkFeature{
				&ForkFeature{
					ID: "homestead",
					Options: ChainFeatureConfigOptions{
						"gastable": `{
							"extcodesize":     20,
							"extcodecopy":     20,
							"balance":         20,
							"sload":           50,
							"calls":           40,
							"suicide":         0,
							"expbyte":         10
						}`,
						"difficulty": `{
							"name": "homestead",
							"options": {}
						}`,
					},
				},
			},
		}, {
			Name:         "ETF",
			Block:        big.NewInt(1920000),
			RequiredHash: common.HexToHash("94365e3a8c0b35089c1d1195081fe7489b528a84b22199c916180db8b28ade7f"),
			Features:     []*ForkFeature{},
		}, {
			Name:     "GasReprice",
			Block:    big.NewInt(2500000),
			Features: []*ForkFeature{DefaultGasRepriceFeature},
		}, {
			Name:  "Diehard",
			Block: big.NewInt(3000000),
			Features: []*ForkFeature{
				DefaultEIP155Feature,
				DefaultDiehardGasRepriceFeature,
				{
					ID: "ecip1010Default",
					Options: ChainFeatureConfigOptions{
						"difficulty": `{
							"name": "diehard",
							"options": {}
						}`,
					},
				},
				{ // ecip1010 bomb delay
					ID:    "explosionDefault",
					Block: big.NewInt(0).Add(big.NewInt(3000000), DefaultBombDelayLength),
					Options: ChainFeatureConfigOptions{
						"difficulty": `{
							"name": "explosion",
							"options": {
								"delay": ` + DefaultBombDelayLength.String() + `
							}
						}`,
					},
				},
			},
		},
	},
	BadHashes: []*BadHash{
		{
			// consensus issue that occurred on the Frontier network at block 116,522, mined on 2015-08-20 at 14:59:16+02:00
			// https://blog.ethereum.org/2015/08/20/security-alert-consensus-issue
			Block: big.NewInt(116522),
			Hash:  common.HexToHash("05bef30ef572270f654746da22639a7a0c97dd97a7050b9e252391996aaeb689"),
		},
	},
	ChainId: big.NewInt(61),
}

// TestConfig is the semi-official setup for testing purposes.
var TestConfig = &ChainConfig{
	Forks: []*Fork{
		{
			Name:  "Homestead",
			Block: big.NewInt(494000),
			Features: []*ForkFeature{
				{
					ID: "homestead",
					Options: ChainFeatureConfigOptions{
						"gastable": `{
							"extcodesize":     20,
							"extcodecopy":     20,
							"balance":         20,
							"sload":           50,
							"calls":           40,
							"suicide":         0,
							"expbyte":         10
						}`,
						"difficulty": `{
							"name": "homestead",
							"options": {}
						}`,
					},
				},
			},
		},
		{
			Name:     "GasReprice",
			Block:    big.NewInt(1783000),
			Features: []*ForkFeature{DefaultGasRepriceFeature},
		},
		{
			Name:     "ETF",
			Block:    big.NewInt(1885000),
			Features: []*ForkFeature{},
		},
		{
			Name:  "Diehard",
			Block: big.NewInt(1915000),
			Features: []*ForkFeature{
				DefaultEIP155Feature,
				DefaultDiehardGasRepriceFeature,
				{
					ID: "ecip1010Default",
					Options: ChainFeatureConfigOptions{
						"difficulty": `{
							"name": "diehard",
							"options": {}
						}`,
					},
				},
				{ // ecip1010 bomb delay
					ID:    "explosionDefault",
					Block: big.NewInt(0).Add(big.NewInt(1915000), DefaultBombDelayLength),
					Options: ChainFeatureConfigOptions{
						"difficulty": `{
							"name": "explosion",
							"options": {
								"delay": ` + DefaultBombDelayLength.String() + `
							}
						}`,
					},
				},
			},
		},
	},
	BadHashes: []*BadHash{
		{
			// consensus issue at Testnet #383792
			// http://ethereum.stackexchange.com/questions/10183/upgraded-to-geth-1-5-0-bad-block-383792
			Block: big.NewInt(383792),
			Hash:  common.HexToHash("9690db54968a760704d99b8118bf79d565711669cefad24b51b5b1013d827808"),
		},
		{
			// chain followed by non-diehard testnet
			Block: big.NewInt(1915277),
			Hash:  common.HexToHash("3bef9997340acebc85b84948d849ceeff74384ddf512a20676d424e972a3c3c4"),
		},
	},
	ChainId: big.NewInt(62),
}
