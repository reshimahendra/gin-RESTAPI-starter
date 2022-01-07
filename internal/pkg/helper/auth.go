package helper

import (
	crand "crypto/rand"
	"encoding/base64"
	"errors"
	mRand "math/rand"
	"net/mail"
	"time"

	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

var (
    minSecureKeyLength   = 8
    secureKeyTooShortErr = errors.New("Lenght of the given 'secure key' does not meet the requirement.")
)

// checkSecureKeyLength will checking length of the 'Secure Key'
func checkSecureKeyLength(length int) error {
    if length < minSecureKeyLength {
        return secureKeyTooShortErr
    }

    return nil
}

// fallbackInsecureKey will give fallback value for insecure key 
// It will generated once 'GenerateSecureKey' resulting error 
func fallbackInsecureKey(length int) (string, error) {
    const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789" +
		"~!@#$%^&*()_+{}|<>?,./:"

        if err := checkSecureKeyLength(length); err != nil {
            return "", err
        }

        var seededRand *mRand.Rand = mRand.New(mRand.NewSource(time.Now().UnixNano()))
        fbk := make([]byte, length)
        for i := range fbk {
            fbk[i] = charset[seededRand.Intn(len(charset))]
        }

        return string(fbk), nil
}

// GenerateSecureKey will create 'Secure Key' with given length
func GenerateSecureKey(length int) (string, error) {
    gsk := make([]byte, length)

    if err := checkSecureKeyLength(length); err != nil {
        return "", err
    }

    _, err := crand.Read(gsk)
    if err != nil {
        return fallbackInsecureKey(length)
    }

    encryptKey := base64.StdEncoding.EncodeToString(gsk)
    return encryptKey, nil
}

// HashPassword will generated hashed password so it wont easily be roken by unauthorized person
func HashPassword(password string) (hashed string, err error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    hashed = string(bytes)
    return
}

// CheckPasswordHash will compare 'hashed' password with the 'input' password
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

    return err == nil
}

// EmailIsValid will check whether given email was valid
func EmailIsValid(email string) (isValid bool) {
    _, err := mail.ParseAddress(email)

    isValid = err == nil

    return
}

// PasswordTooShort will check whether password length is not match the minimum 
// password lenght for the user account
func PasswordTooShort(password string) (isPasswordtooShort bool) {
    return len(password) < viper.GetInt("account.minimum_password_length")
}
