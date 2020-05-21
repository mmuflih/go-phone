package gophone

import (
	"strings"
)

/*
 * Created by M. Muflih Kholidin
 * Wed Aug 29 2018 9:51:11 AM
 * mmuflic@gmail.com
 * https://github.com/mmuflih
 **/

type PhoneNumber interface {
	Sanitize(phoneNumber string) string
	SanitizePlus(phoneNumber string) string
	Sanitize0(phoneNumber string) string
}

type lPhoneNumber struct {
	dialingCode string
}

func NewPhoneNumber(dialingCode string) PhoneNumber {
	return &lPhoneNumber{dialingCode}
}

func (pn lPhoneNumber) Sanitize(phoneNumber string) string {
	if phoneNumber == "" {
		return ""
	}
	phone := pn.sanitize(phoneNumber)
	return pn.dialingCode + phone[1:]
}

func (pn lPhoneNumber) SanitizePlus(phoneNumber string) string {
	if phoneNumber == "" {
		return ""
	}
	phone := pn.sanitize(phoneNumber)
	return "+" + pn.dialingCode + phone[1:]
}

func (pn lPhoneNumber) Sanitize0(phoneNumber string) string {
	phone := pn.sanitize(phoneNumber)
	return phone
}

func (pn lPhoneNumber) sanitize(phoneNumber string) string {
	if phoneNumber == "" {
		return ""
	}
	phone := strings.Replace(phoneNumber, " ", "", 0)
	phone = strings.Replace(phone, "-", "", 0)
	if phone[0:2] == pn.dialingCode {
		return "0" + phone[2:]
	} else if phone[0:3] == "+"+pn.dialingCode {
		return "0" + phone[3:]
	}
	return phone
}
