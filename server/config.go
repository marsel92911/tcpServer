package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

const (
	defaultPort       = "12345"
	defaultTargetBits = 18
)

type config struct {
	port       string
	targetBits int
	target     *big.Int
}

func getConfig() *config {
	cfg := config{
		target: big.NewInt(1),
	}

	//flag.StringVar(&cfg.port, "port", "", "server port")
	//flag.IntVar(&cfg.targetBits, "target_bits", 0, "Adjust the number of leading zeros")
	//flag.Parse()
	cfg.port = os.Getenv("TCP_SERVER_PORT")
	targetBits := os.Getenv("TCP_SERVER_TARGET_BITS")

	if cfg.port == "" {
		fmt.Printf("Server port is not set: using default port %s\n", defaultPort)
		cfg.port = defaultPort
	}

	if targetBits == "" {
		fmt.Printf("target_bits is not set: using default value: %d\n", defaultTargetBits)
		cfg.targetBits = defaultTargetBits
	} else {
		var err error

		cfg.targetBits, err = strconv.Atoi(targetBits)
		if err != nil {
			fmt.Printf("target_bits parse error: %s using default value: %d\n", err.Error(), defaultTargetBits)
			cfg.targetBits = defaultTargetBits
		}
	}

	target := big.NewInt(1)
	cfg.target.Lsh(target, uint(256-cfg.targetBits))

	return &cfg
}
