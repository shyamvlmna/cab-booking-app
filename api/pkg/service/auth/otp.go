package auth

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strconv"

	"github.com/shayamvlmna/cab-booking-app/pkg/database/redis"
)

//generate and return a 4 digit random otp using crypto/rand
//return error if any
func GenerateOTP() (string, error) {
	nBig, e := rand.Int(rand.Reader, big.NewInt(8999))
	if e != nil {
		return "", e
	}
	return strconv.FormatInt(nBig.Int64()+1000, 10), nil
}

//create and assign a secret otp for the given number
func SetOtp(phone string) error {
	otp, err := GenerateOTP()
	if err != nil {
		return err
	}
	redis.Set(phone, otp)
	fmt.Printf("user signup otp for %s :%s", phone, otp)
	return nil
}

//ValidateOTP returns an error if the otp doesn't belong to the given number
func ValidateOTP(phone, otp string) error {
	value, err := redis.Get(phone)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if value != otp {
		return errors.New("invalid otp")
	}
	return nil
}

func StoreUser(phone string) {
	redis.Set("user", phone)
}
func GetUser() string {
	user, _ := redis.Get("user")
	return user
}

func StoreDriver(phone string) {
	redis.Set("driver", phone)
}

func GetDriver() string {
	driver, _ := redis.Get("driver")
	return driver
}
