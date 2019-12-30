package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"time"

	"./ntp"
)

const (
	Sunday    = 0
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: lockedbox3 [encrypt|decrypt] [secret|cypher]")
		return
	}

	timeNow, err := getTime()
	fmt.Println("")

	if err != nil {
		timeNow = time.Now() // Fall back to system time only if ntp time fails
	}

	keyString := getSeededRandom()
	key := []byte(keyString[0:16])

	if os.Args[1] == "decrypt" {
		// Decryption only allowed at certain times
		if !((timeNow.Weekday() == Friday) &&
			(timeNow.Hour() >= 17)) {
			fmt.Println("Only available on Fri after 5pm")
			return
		}

		text := decrypt(key, os.Args[2])
		fmt.Println(text)

	} else if os.Args[1] == "encrypt" {
		cryptoText := encrypt(key, os.Args[2])
		fmt.Println(cryptoText)
	} else {
		fmt.Println("Usage: lockedbox3 [encrypt|decrypt] [secret|cypher]")
		return
	}
}

func getSeededRandom() (result string) {
	// NOTE: Change these values after compilation to make the key harder to find
	value := (math.Pow(47, 8) * math.Pow(32, 5) * math.Pow(3, 4) * math.Pow(156, 4)) / math.Pow(7, 3)
	result = strconv.FormatFloat(value, 'f', -1, 64)
	return
}

func getTime() (result time.Time, err error) {
	response, err := ntp.Query("time.nist.gov") //"0.beevik-ntp.pool.ntp.org")
	result = time.Now().Add(response.ClockOffset)

	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	return
}

// encrypt string to base64 crypto using AES
func encrypt(key []byte, text string) string {
	// key := []byte(keyText)
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

// decrypt from base64 to decrypted string
func decrypt(key []byte, cryptoText string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}
