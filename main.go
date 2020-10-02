package main

import (
	"crypto"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/sec51/twofactor"
)

func main() {
	CreateTokenEndpoint()
	otp, err := CreateOtp("luongdai246@gmail.com", "secret")
	if err != nil {
		fmt.Println(err.Error())
	}
	s, err := otp.OTP()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("\n" + s)
	if err := ValidateOtp(s, otp); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("match")
}

//create token
func CreateTokenEndpoint() {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "Dai",
		"password": "159753",
	})
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(tokenString)
}

//create otp
func CreateOtp(name string, secret string) (*twofactor.Totp, error) {
	otp, err := twofactor.NewTOTP(name, secret, crypto.SHA1, 6)
	if err != nil {
		return nil, err
	}
	return otp, nil
}

//validate otp
func ValidateOtp(otp string, totp *twofactor.Totp) error {
	err := totp.Validate(otp)
	if err != nil {
		return err
	}
	return nil
}
