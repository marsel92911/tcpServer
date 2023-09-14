package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"math"
	"math/big"
	"net"
	"os"
	"strconv"
)

func main() {
	conn, err := net.Dial("tcp", fmt.Sprintf("0.0.0.0:%s", getPort()))
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	quote, err := handleConnection(conn, err)
	if err != nil {
		fmt.Println("Error handling server connection:", err)
		os.Exit(1)
	}

	fmt.Printf("Received quote is: %s", string(quote))
}

func handleConnection(conn net.Conn, err error) ([]byte, error) {
	targetBits, challenge, err := getProofOfWorkData(conn)
	if err != nil {
		return nil, fmt.Errorf("error reading form tcp server: %w", err)
	}

	// Solution of PoW
	solution, err := solveProofOfWork(string(challenge), targetBits)

	// Sending solution to server
	_, err = conn.Write([]byte(solution + "\n"))
	if err != nil {
		return nil, fmt.Errorf("error sending solution to tcp server: %w", err)
	}

	quote, err := io.ReadAll(conn)
	if err != nil {
		return nil, fmt.Errorf("error reading quote to tcp server: %w", err)
	}

	return quote, nil
}

func solveProofOfWork(challenge string, targetBits int) (string, error) {
	var hashInt big.Int

	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	fmt.Printf("Challenge is: %s\n", challenge)

	for nonce := uint64(0); nonce <= math.MaxInt64; nonce++ {

		input := challenge + strconv.FormatUint(nonce, 10)
		hash := sha256.Sum256([]byte(input))
		fmt.Printf("\r%x", hash)

		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(target) == -1 {
			fmt.Printf("\nFound nonce is: %d\nresult hash is:\n%s\n", nonce, fmt.Sprintf("%x", hash))

			return strconv.FormatInt(int64(nonce), 10), nil
		} else {
			nonce++
		}
	}

	return "", fmt.Errorf("solution for challenge not found")
}

func getProofOfWorkData(conn net.Conn) (int, []byte, error) {
	reader := bufio.NewReader(conn)

	targetBitsStr, _, err := reader.ReadLine()
	if err != nil {
		return 0, nil, err
	}
	targetBits, err := strconv.Atoi(string(targetBitsStr))
	if err != nil {
		return 0, nil, err
	}

	challenge, _, err := reader.ReadLine()
	if err != nil {
		return 0, nil, err
	}
	return targetBits, challenge, nil
}
