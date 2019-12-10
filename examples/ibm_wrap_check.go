package main

import (
	"fmt"

	ibm_crypto "github.com/IBM-Bluemix/golang-openssl-wrapper/crypto" // For encryption, e.g. EVP
	ibm_rand "github.com/IBM-Bluemix/golang-openssl-wrapper/rand"     // PRNG
)

func encrypt(plaintext, key string) string {
	var sLen, eLen int
	ivLen := 12

	// Print plaintext string
	fmt.Println("Plaintext:", plaintext)

	// Create new EVP_CIPHER_CTX instances
	ctxEncrypt := ibm_crypto.EVP_CIPHER_CTX_new()

	// Panic if either EVP_CIPHER_CTX fails to create
	if ctxEncrypt == nil {
		panic("ctxEncrypt is nil")
	}

	defer ibm_crypto.EVP_CIPHER_CTX_cleanup(ctxEncrypt)

	// Initialize the EVP_CIPHER_CTX instances
	ibm_crypto.EVP_CIPHER_CTX_init(ctxEncrypt)

	// Create random IV for nondeterministic encryption
	ivBuf := make([]byte, ivLen)
	_, e := ibm_rand.Read(ivBuf)
	if e != nil {
		panic(e)
	}

	// Pass the IV into the encrypted string to be used when decoding
	encrypted := string(ivBuf)

	fmt.Println("Starting encrypt using AES 256 CBC in FIPS mode")

	// Initialize the ctxEncrypt context for encryption
	resp := ibm_crypto.EVP_EncryptInit_ex(ctxEncrypt,
		ibm_crypto.EVP_aes_256_gcm(),
		ibm_crypto.SwigcptrStruct_SS_engine_st(0),
		&[]byte(key)[0],
		&ivBuf[0])

	fmt.Println("Success:", resp == 1)

	// Make a buffer with enough size for the plaintext plus one block
	bufEncrypt := make([]byte, len(plaintext)+ctxEncrypt.GetCipher().GetBlock_size())

	// Update the cipher with some content
	resp = ibm_crypto.EVP_EncryptUpdate(ctxEncrypt, bufEncrypt, &sLen, &[]byte(plaintext)[0], len(plaintext))

	fmt.Println("Success:", resp == 1)

	// Append encrypted data to encrypted string
	encrypted += string(bufEncrypt[:sLen])

	// Finalize the cipher to flush any remaining data
	resp = ibm_crypto.EVP_EncryptFinal_ex(ctxEncrypt, bufEncrypt, &eLen)
	fmt.Println("Success:", resp == 1)

	// Append any remaining data to the encrypted string
	encrypted += string(bufEncrypt[:eLen])

	return encrypted
}

func decrypt(encrypted, key string) string {
	var sLen, eLen int
	ivLen := 12

	ctxDecrypt := ibm_crypto.EVP_CIPHER_CTX_new()
	if ctxDecrypt == nil {
		panic("ctxDecrypt is nil")
	}
	defer ibm_crypto.EVP_CIPHER_CTX_cleanup(ctxDecrypt)
	ibm_crypto.EVP_CIPHER_CTX_init(ctxDecrypt)

	fmt.Println("Starting decrypt...")
	// Grab the IV from the encrypted string
	iv := string([]byte(encrypted)[:ivLen])

	// Slice the encrypted string to begin after the iv
	encrypted = encrypted[ivLen:]

	// Initialize the ctxDecrypt context for decryption
	resp := ibm_crypto.EVP_DecryptInit_ex(ctxDecrypt,
		 ibm_crypto.EVP_aes_256_gcm(),
		 ibm_crypto.SwigcptrStruct_SS_engine_st(0),
		 &[]byte(key)[0], &[]byte(iv)[0])

	fmt.Println("Success:", resp == 1)

	// Make a buffer the exact size of the encrypted text
	bufDecrypt := make([]byte, len(encrypted))

	// Update the cipher with the encrypted string
	resp = ibm_crypto.EVP_DecryptUpdate(ctxDecrypt, bufDecrypt, &sLen, &[]byte(encrypted)[0], len(encrypted))
	fmt.Println("Success:", resp == 1)

	// Append decrypted data to decrypted string
	decrypted := string(bufDecrypt[:sLen])

	// Finalize the cipher to flush any remaining data
	if ibm_crypto.EVP_DecryptFinal_ex(ctxDecrypt, bufDecrypt, &eLen) == 1 {		
		decrypted += string(bufDecrypt[:eLen])// Append any remaining data to decrypted string
	}

	return decrypted
}

func main() {
	fmt.Println("starting test...")

	ibm_crypto.FIPS_mode_set(1)
	ibm_crypto.ERR_load_crypto_strings()
	ibm_crypto.OpenSSL_add_all_algorithms()
	ibm_crypto.OPENSSL_config("")

	key := "1234567890ABCDEF1234567890ABCDEF"
	plaintext := "My super super super super duper long string to be encrypted"

	encrypted := encrypt(plaintext, key)
	fmt.Println("Encypted string:", encrypted)

	// Print decoded string
	decrypted := decrypt(encrypted, key)
	fmt.Println("Got decrypted string:", decrypted)
}
