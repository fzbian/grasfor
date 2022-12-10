package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func Encrypt(plaintext string) (string, error) {
	// Create a new AES cipher
	block, err := aes.NewCipher([]byte("0123456789abcdef"))
	if err != nil {
		return "", err
	}

	// Encrypt the plaintext using the cipher
	ciphertext := make([]byte, len(plaintext))
	stream := cipher.NewCTR(block, []byte("abcdefghijklmnop"))
	stream.XORKeyStream(ciphertext, []byte(plaintext))

	// Return the encrypted ciphertext as a string
	return fmt.Sprintf("%x", ciphertext), nil
}

func Decrypt(ciphertext string) (string, error) {
	// Create a new AES cipher
	block, err := aes.NewCipher([]byte("0123456789abcdef"))
	if err != nil {
		return "", err
	}

	// Decrypt the ciphertext using the cipher
	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCTR(block, []byte("abcdefghijklmnop"))
	stream.XORKeyStream(plaintext, []byte(ciphertext))

	// Return the decrypted plaintext as a string
	return string(plaintext), nil
}
