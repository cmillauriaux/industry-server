package model

import "fmt"

// Company
const COMPANY_REQUEST_ERROR = 1000
const COMPANY_NOT_FOUND = 1001
const COMPANY_NOT_VALID = 1002
const COMPANY_ALREADY_EXISTS = 1003

// Factory
const FACTORY_REQUEST_ERROR = 2000
const FACTORY_NOT_FOUND = 2001
const FACTORY_NOT_VALID = 2002
const FACTORY_DELETE_ERROR = 2003
const HIRE_EMPLOYEE_NUMBER_ERROR = 2004
const FIRE_EMPLOYEE_NUMBER_ERROR = 2005
const UPDATE_SALARY_ERROR = 2006

// Product Type
const PRODUCT_TYPE_NOT_FOUND = 3000

// Product
const PRODUCT_NOT_FOUND = 4000

// Offer
const OFFER_NOT_FOUND = 5000
const OFFER_DELETE_ERROR = 50001

type BusinessError struct {
	Code  int
	Msg   string
	Cause error
}

func NewBusinessError(code int, msg string, cause error) error {
	return &BusinessError{Code: code, Msg: msg, Cause: cause}
}

func (b *BusinessError) Error() string {
	if b.Cause != nil {
		return fmt.Sprintf("%d : %s, cause : %s", b.Code, b.Msg, b.Cause.Error())
	}
	return fmt.Sprintf("%d : %s", b.Code, b.Msg)
}
