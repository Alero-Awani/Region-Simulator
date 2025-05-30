package dto

// UserLogin and UserSignUp struct validation and parsing - Validate Request input
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignUp struct {
	UserLogin
	Phone string `json:"phone"`
}

type VerificationCodeInput struct {
	Code int `json:"code"`
}

type SellerInput struct {
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	PhoneNumber       string `json:"phone_number"`
	BankAccountNumber uint   `json:"bankAccountNumber"`
	SwiftCode         string `json:"swiftCode"`
	PaymentType       string `json:"paymentType"`
}

type AddressInput struct {
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	City         string `json:"city"`
	PostCode     string `json:"post_code"`
	Country      string `json:"country"`
}

type ProfileInput struct {
	FirstName    string       `json:"first_name"`
	LastName     string       `json:"last_name"`
	Email        string       `json:"email"`
	AddressInput AddressInput `json:"address"`
}
