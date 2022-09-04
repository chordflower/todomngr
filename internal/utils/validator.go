package utils

import (
	"net"
	"net/url"
	"regexp"
	"time"

	"emperror.dev/errors"
	"github.com/ShiraazMoollatjie/goluhn"
	date "github.com/bykof/gostradamus"
	"github.com/emirpasic/gods/lists/arraylist"
	isd "github.com/jbenet/go-is-domain"
)

type number interface {
	~float64 | ~float32 | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

var (
	alphanumericRegex = regexp.MustCompile(`^(?:\pL|\pN|\s)+$`)
	base64Regex       = regexp.MustCompile(`^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{4})$`)
	lowercaseRegex    = regexp.MustCompile(`^(?:\p{Ll}|\s)+$`)
	uppercaseRegex    = regexp.MustCompile(`^(?:\p{Lu}|\s)+$`)
	emailRegex        = regexp.MustCompile("^[a-z0-9!#$%&'*+/=?^_`{|}~-]+(\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@([a-z0-9!#$%&'*+/=?^_`{|}~-]+(\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\\[[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\])$")
	guidRegex         = regexp.MustCompile("^(?:[0-9A-Fa-f]{32})|(?:(?:\\{|\\()?[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{12}(?:\\}|\\))?)|(?:\\{0x[0-9A-Fa-f]{8},0x[0-9A-Fa-f]{4},0x[0-9A-Fa-f]{4},\\{(?:0x[0-9A-Fa-f]{2},){7}0x[0-9A-Fa-f]{2}\\}\\})$")
)

// Validate represents a validator
type Validate struct {
	errors arraylist.List
}

// NewValidator creates a new validator
func NewValidator() *Validate {
	return &Validate{
		errors: *arraylist.New(),
	}
}

// Check checks if the given condition is true, and it it is adds the given error
func (v *Validate) Check(c bool, errMsg string) {
	if !c {
		v.errors.Add(errMsg)
	}
}

// IsPresent checks if the given field is not nil
func (v *Validate) IsPresent(field any, errMsg string) {
	v.Check(field != nil, errMsg)
}

// IsNotPresent checks if the given field is nil
func (v *Validate) IsNotPresent(field any, errMsg string) {
	v.Check(field == nil, errMsg)
}

// IsAlphaNumeric checks if the given string is a unicode alphanumeric string
func (v *Validate) IsAlphaNumeric(field string, errMsg string) {
	v.Check(alphanumericRegex.MatchString(field), errMsg)
}

// IsBase64 checks if the given string is a base64 string
func (v *Validate) IsBase64(field string, errMsg string) {
	v.Check(base64Regex.MatchString(field), errMsg)
}

// IsLowercase checks if the given string is a unicode lowercase string
func (v *Validate) IsLowercase(field string, errMsg string) {
	v.Check(lowercaseRegex.MatchString(field), errMsg)
}

// IsUppercase checks if the given string is a unicode uppercase string
func (v *Validate) IsUppercase(field string, errMsg string) {
	v.Check(uppercaseRegex.MatchString(field), errMsg)
}

// IsCreditCard checks if the given string is a valid credit card number using Luhn algorithm
func (v *Validate) IsCreditCard(field string, errMsg string) {
	v.Check(goluhn.Validate(field) == nil, errMsg)
}

// IsDomain checks if the given string is a valid domain
func (v *Validate) IsDomain(field string, errMsg string) {
	v.Check(isd.IsDomain(field), errMsg)
}

// IsEmail checks if the given string is a valid email
func (v *Validate) IsEmail(field string, errMsg string) {
	v.Check(emailRegex.MatchString(field), errMsg)
}

// IsGUID checks if the given string is a valid guid aka uuid
func (v *Validate) IsGUID(field string, errMsg string) {
	v.Check(guidRegex.MatchString(field), errMsg)
}

// IsHostname checks if the given string is a valid hostname
func (v *Validate) IsHostname(field string, errMsg string) {
	v.Check(isd.IsDomain(field) || net.ParseIP(field) != nil, errMsg)
}

// IsIP checks if the given string is a valid ip
func (v *Validate) IsIP(field string, errMsg string) {
	v.Check(net.ParseIP(field) != nil, errMsg)
}

// IsStdDate checks if the given string is a valid iso8601 date
func (v *Validate) IsStdDate(field string, errMsg string) {
	_, err := date.Parse(field, date.Iso8601TZ)
	v.Check(err == nil, errMsg)
}

// IsDuration checks if the given string is a valid duration
func (v *Validate) IsDuration(field string, errMsg string) {
	_, err := time.ParseDuration(field)
	v.Check(err == nil, errMsg)
}

// IsSize checks if the given string has the given size
func (v *Validate) IsSize(field string, size uint32, errMsg string) {
	v.Check(uint32(len(field)) == size, errMsg)
}

// IsNotEmpty checks if the given string is not empty
func (v *Validate) IsNotEmpty(field string, errMsg string) {
	v.Check(uint32(len(field)) != 0, errMsg)
}

// IsEmpty checks if the given string is empty
func (v *Validate) IsEmpty(field string, errMsg string) {
	v.Check(uint32(len(field)) == 0, errMsg)
}

// IsBetween checks if the given string is between the given size
func (v *Validate) IsBetween(field string, min, max uint32, errMsg string) {
	var size uint32 = uint32(len(field))
	v.Check(size >= min && size <= max, errMsg)
}

// IsURL checks if the given string is a valid url
func (v *Validate) IsURL(field string, errMsg string) {
	_, err := url.Parse(field)
	v.IsNotPresent(err, errMsg)
}

// IsGreaterThan checks if the given number is greater than the other given number
func (v *Validate) IsGreaterThan(field, min any, errMsg string) {
	switch field.(type) {
	case float64:
		v.Check(field.(float64) > min.(float64), errMsg)
	case float32:
		v.Check(field.(float32) > min.(float32), errMsg)
	case int:
		v.Check(field.(int) > min.(int), errMsg)
	case int8:
		v.Check(field.(int8) > min.(int8), errMsg)
	case int16:
		v.Check(field.(int16) > min.(int16), errMsg)
	case int32:
		v.Check(field.(int32) > min.(int32), errMsg)
	case int64:
		v.Check(field.(int64) > min.(int64), errMsg)
	case uint:
		v.Check(field.(uint) > min.(uint), errMsg)
	case uint8:
		v.Check(field.(uint8) > min.(uint8), errMsg)
	case uint16:
		v.Check(field.(uint16) > min.(uint16), errMsg)
	case uint32:
		v.Check(field.(uint32) > min.(uint32), errMsg)
	case uint64:
		v.Check(field.(uint64) > min.(uint64), errMsg)
	}
}

// IsLessThan checks if the given number is less than the other given number
func (v *Validate) IsLessThan(field, max any, errMsg string) {
	switch field.(type) {
	case float64:
		v.Check(field.(float64) < max.(float64), errMsg)
	case float32:
		v.Check(field.(float32) < max.(float32), errMsg)
	case int:
		v.Check(field.(int) < max.(int), errMsg)
	case int8:
		v.Check(field.(int8) < max.(int8), errMsg)
	case int16:
		v.Check(field.(int16) < max.(int16), errMsg)
	case int32:
		v.Check(field.(int32) < max.(int32), errMsg)
	case int64:
		v.Check(field.(int64) < max.(int64), errMsg)
	case uint:
		v.Check(field.(uint) < max.(uint), errMsg)
	case uint8:
		v.Check(field.(uint8) < max.(uint8), errMsg)
	case uint16:
		v.Check(field.(uint16) < max.(uint16), errMsg)
	case uint32:
		v.Check(field.(uint32) < max.(uint32), errMsg)
	case uint64:
		v.Check(field.(uint64) < max.(uint64), errMsg)
	}
}

// IsBetweenNumbers checks if the given number is between the given numbers
func (v *Validate) IsBetweenNumbers(field, min, max any, errMsg string) {
	switch field.(type) {
	case float64:
		v.Check(field.(float64) <= max.(float64) && field.(float64) >= min.(float64), errMsg)
	case float32:
		v.Check(field.(float32) <= max.(float32) && field.(float32) >= min.(float32), errMsg)
	case int:
		v.Check(field.(int) <= max.(int) && field.(int) >= min.(int), errMsg)
	case int8:
		v.Check(field.(int8) <= max.(int8) && field.(int8) >= min.(int8), errMsg)
	case int16:
		v.Check(field.(int16) <= max.(int16) && field.(int16) >= min.(int16), errMsg)
	case int32:
		v.Check(field.(int32) <= max.(int32) && field.(int32) >= min.(int32), errMsg)
	case int64:
		v.Check(field.(int64) <= max.(int64) && field.(int64) >= min.(int64), errMsg)
	case uint:
		v.Check(field.(uint) <= max.(uint) && field.(uint) >= min.(uint), errMsg)
	case uint8:
		v.Check(field.(uint8) <= max.(uint8) && field.(uint8) >= min.(uint8), errMsg)
	case uint16:
		v.Check(field.(uint16) <= max.(uint16) && field.(uint16) >= min.(uint16), errMsg)
	case uint32:
		v.Check(field.(uint32) <= max.(uint32) && field.(uint32) >= min.(uint32), errMsg)
	case uint64:
		v.Check(field.(uint64) <= max.(uint64) && field.(uint64) >= min.(uint64), errMsg)
	}
}

// IsNegative checks if the given number is negative
func (v *Validate) IsNegative(field any, errMsg string) {
	switch field.(type) {
	case float64:
		v.Check(field.(float64) < 0, errMsg)
	case float32:
		v.Check(field.(float32) < 0, errMsg)
	case int:
		v.Check(field.(int) < 0, errMsg)
	case int8:
		v.Check(field.(int8) < 0, errMsg)
	case int16:
		v.Check(field.(int16) < 0, errMsg)
	case int32:
		v.Check(field.(int32) < 0, errMsg)
	case int64:
		v.Check(field.(int64) < 0, errMsg)
	}
}

// IsPositive checks if the given number is positive
func (v *Validate) IsPositive(field any, errMsg string) {
	switch field.(type) {
	case float64:
		v.Check(field.(float64) > 0, errMsg)
	case float32:
		v.Check(field.(float32) > 0, errMsg)
	case int:
		v.Check(field.(int) > 0, errMsg)
	case int8:
		v.Check(field.(int8) > 0, errMsg)
	case int16:
		v.Check(field.(int16) > 0, errMsg)
	case int32:
		v.Check(field.(int32) > 0, errMsg)
	case int64:
		v.Check(field.(int64) > 0, errMsg)

	}
}

// IsPort checks if the given number can be a TCP/UDP port
func (v *Validate) IsPort(field any, errMsg string) {
	switch field.(type) {
	case int:
		v.IsBetweenNumbers(field, int(0), int(65535), errMsg)
	case int32:
		v.IsBetweenNumbers(field, int32(0), int32(65535), errMsg)
	case int64:
		v.IsBetweenNumbers(field, int64(0), int64(65535), errMsg)
	case uint:
		v.IsBetweenNumbers(field, uint(0), uint(65535), errMsg)
	case uint16:
		v.IsBetweenNumbers(field, uint16(0), uint16(65535), errMsg)
	case uint32:
		v.IsBetweenNumbers(field, uint32(0), uint32(65535), errMsg)
	case uint64:
		v.IsBetweenNumbers(field, uint64(0), uint64(65535), errMsg)

	}
}

// IsDateDefined checks if the given date field is defined aka not zero
func (v *Validate) IsDateDefined(field any, errMsg string) {
	switch date := field.(type) {
	case time.Time:
		v.Check(!date.IsZero(), errMsg)
	case date.DateTime:
		v.Check(!date.Time().IsZero(), errMsg)
	}
}

// IsDateBefore checks if the given date field is before the min date
func (v *Validate) IsDateBefore(field, min any, errMsg string) {
	switch field.(type) {
	case time.Time:
		v.Check(field.(time.Time).Before(min.(time.Time)), errMsg)
	case date.DateTime:
		v.Check(field.(date.DateTime).Time().Before(min.(date.DateTime).Time()), errMsg)
	}
}

// IsDateAfter checks if the given date field is after the max date
func (v *Validate) IsDateAfter(field, max any, errMsg string) {
	switch field.(type) {
	case time.Time:
		v.Check(field.(time.Time).After(max.(time.Time)), errMsg)
	case date.DateTime:
		v.Check(field.(date.DateTime).Time().After(max.(date.DateTime).Time()), errMsg)
	}
}

// HasErrors returns if there are any errors so far
func (v *Validate) HasErrors() bool {
	return !v.errors.Empty()
}

// ErrorNumber returns the number of errors detected so far
func (v *Validate) ErrorNumber() uint {
	return uint(v.errors.Size())
}

// AllValid aglomerates all validation errors into one and returns an error
func (v *Validate) AllValid() error {
	if !v.errors.Empty() {
		errs := make([]error, v.errors.Size())
		v.errors.Each(func(index int, value interface{}) {
			errs = append(errs, errors.New(value.(string)))
		})
		return errors.Combine(errs...)
	}
	return nil
}
