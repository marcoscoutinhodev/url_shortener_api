package entity

import (
	"errors"
	"net/url"
	"regexp"
)

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password should be of 8 characters long")
	}

	done, err := regexp.MatchString("([a-z])+", password)
	if err != nil {
		return err
	}

	if !done {
		return errors.New("password should contain atleast one lower case character")
	}

	done, err = regexp.MatchString("([A-Z])+", password)
	if err != nil {
		return err
	}

	if !done {
		return errors.New("password should contain atleast one upper case character")
	}

	done, err = regexp.MatchString("([0-9])+", password)
	if err != nil {
		return err
	}

	if !done {
		return errors.New("password should contain atleast one digit")
	}

	done, err = regexp.MatchString("([!@#$%^&*.?-])+", password)
	if err != nil {
		return err
	}

	if !done {
		return errors.New("password should contain atleast one special character")
	}

	return nil
}

func isUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
