package main

import (
	"fmt"
	"math/big"
	"math/rand"
)

const (
	q = 257        // Modulus
	n = 8          // Number of variables
	σ = 8.0        // Noise parameter
	m = "Hello, World!" // Message to encrypt
)

func main() {
	// Key Generation
	s := generateSecret()
	A := generatePublicKey()
	pk, sk := A, s // In practice, public key (pk) would be shared, and secret key (sk) kept secret

	// Encryption
	ciphertext, ciphertext2 := encrypt(pk, m)

	// Decryption
	decryptedMessage := decrypt(sk, ciphertext, ciphertext2)
	fmt.Println("Decrypted Message:", decryptedMessage)
}

// Generate a random secret vector
func generateSecret() []int {
	secret := make([]int, n)
	for i := range secret {
		secret[i] = rand.Intn(q)
	}
	return secret
}

// Generate a random public key (matrix)
func generatePublicKey() [][]int {
	A := make([][]int, n)
	for i := range A {
		A[i] = make([]int, n)
		for j := range A[i] {
			A[i][j] = rand.Intn(q)
		}
	}
	return A
}

// Encrypt a message using the public key
func encrypt(publicKey [][]int, message string) ([]*big.Int, []*big.Int) {
	r := generateRandomVector()
	c1 := make([]*big.Int, n)
	c2 := make([]*big.Int, n)

	for i := 0; i < n; i++ {
		c1[i] = big.NewInt(0)
		c2[i] = big.NewInt(0)

		for j := 0; j < n; j++ {
			aj := big.NewInt(int64(publicKey[i][j]))
			rj := big.NewInt(int64(r[j]))
			c1[i].Add(c1[i], new(big.Int).Mul(aj, rj))
		}

		c2[i] = big.NewInt(int64(message[i]) + int64(q / 2) + int64(r[i]))
	}

	return c1, c2
}

// Decrypt a ciphertext using the secret key
func decrypt(secretKey []int, ciphertext1 []*big.Int, ciphertext2 []*big.Int) string {
	message := make([]byte, n)

	for i := 0; i < n; i++ {
		c1i := ciphertext1[i]
		c2i := ciphertext2[i]

		si := secretKey[i]
		tmp := new(big.Int).Mul(c1i, big.NewInt(int64(si)))
		tmp.Mod(tmp, big.NewInt(q))
		tmp.Sub(c2i, tmp)
		tmp.Mod(tmp, big.NewInt(q))
		tmp.Add(tmp, big.NewInt(q/2))
		message[i] = byte(tmp.Int64())
	}

	return string(message)
}

// Generate a random vector of length n
func generateRandomVector() []int {
	randomVec := make([]int, n)
	for i := range randomVec {
		randomVec[i] = rand.Intn(σ)
	}
	return randomVec
}
