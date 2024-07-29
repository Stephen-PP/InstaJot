package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

const (
	keySize    = 32    // AES-256
	nonceSize  = 12    // GCM standard nonce size
	saltSize   = 16    // Size of salt for PBKDF2
	iterations = 10000 // Number of iterations for PBKDF2
)

// deriveKey generates a 32-byte key from a password using PBKDF2
func deriveKey(password string, salt []byte) []byte {
	return pbkdf2.Key([]byte(password), salt, iterations, keySize, sha256.New)
}

// Encrypt takes a plaintext string and a password, and returns a base64-encoded encrypted string
func Encrypt(plaintext, password string) (string, error) {
	// Generate a random salt
	salt := make([]byte, saltSize)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return "", err
	}

	// Derive the key from the password
	key := deriveKey(password, salt)

	// Create cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Create nonce
	nonce := make([]byte, nonceSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Encrypt and seal
	ciphertext := gcm.Seal(nil, nonce, []byte(plaintext), nil)

	// Combine salt, nonce, and ciphertext
	result := make([]byte, saltSize+nonceSize+len(ciphertext))
	copy(result, salt)
	copy(result[saltSize:], nonce)
	copy(result[saltSize+nonceSize:], ciphertext)

	// Encode to base64
	encoded := base64.StdEncoding.EncodeToString(result)

	return encoded, nil
}

// Decrypt takes a base64-encoded encrypted string and a password, and returns the decrypted plaintext
func Decrypt(encodedCiphertext, password string) (string, error) {
	// Decode from base64
	decoded, err := base64.StdEncoding.DecodeString(encodedCiphertext)
	if err != nil {
		return "", err
	}

	// Ensure the decoded data is long enough
	if len(decoded) < saltSize+nonceSize {
		return "", errors.New("ciphertext too short")
	}

	// Extract salt, nonce, and ciphertext
	salt := decoded[:saltSize]
	nonce := decoded[saltSize : saltSize+nonceSize]
	ciphertext := decoded[saltSize+nonceSize:]

	// Derive the key from the password
	key := deriveKey(password, salt)

	// Create cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Decrypt
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
