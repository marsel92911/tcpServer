package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"math/big"
	"math/rand"
	"net"
	"strconv"

	"github.com/go-faker/faker/v4"
)

func handleClient(conn net.Conn, cfg *config) {
	defer conn.Close()

	challenge := generateChallenge()
	fmt.Println("Challenge is: ", challenge)

	nonce, err := sendChallengeAndReceiveNonce(conn, challenge, cfg.targetBits)
	if err != nil {
		fmt.Println("Receiving data failed: ", challenge)
		return
	}

	// Verify the PoW
	if !verifyProofOfWork(challenge, string(nonce), cfg.target) {
		fmt.Println("PoW verification failed.")
		return
	}

	_, err = conn.Write([]byte(faker.Sentence()))
	if err != nil {
		fmt.Println("Error sending quote to client:", err)
	}
}

func sendChallengeAndReceiveNonce(conn net.Conn, challenge string, targetBits int) ([]byte, error) {
	_, err := conn.Write([]byte(strconv.Itoa(targetBits) + "\n"))
	if err != nil {
		return nil, fmt.Errorf("error sending sending targetBits: %w", err)
	}

	_, err = conn.Write([]byte(challenge + "\n"))
	if err != nil {
		return nil, fmt.Errorf("error sending sending challenge: %w", err)
	}

	// Read the response from the client
	reader := bufio.NewReader(conn)
	nonce, _, err := reader.ReadLine()
	if err != nil {
		return nil, fmt.Errorf("error getting nonce: %w", err)

	}

	return nonce, nil
}

func verifyProofOfWork(challenge string, nonce string, target *big.Int) bool {
	var hashInt big.Int

	data := challenge + nonce
	fmt.Println("Data is: ", data)

	hash := sha256.Sum256([]byte(data))

	hashInt.SetBytes(hash[:])

	if hashInt.Cmp(target) == -1 {
		fmt.Printf("Found nonce is: %d\nresult hash is:\n%s\n", nonce, fmt.Sprintf("%x", hash))

		return true
	}

	return false
}

func generateChallenge() string {
	// Generate a random challenge
	return strconv.FormatInt(rand.Int63(), 10)
}
