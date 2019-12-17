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

type phoneNumber struct {
	dialingCode string
}

func NewPhoneNumber(dialingCode string) PhoneNumber {
	return &phoneNumber{dialingCode}
}

func (pn phoneNumber) Sanitize(phoneNumber string) string {
	phone := pn.sanitize(phoneNumber)
	return pn.dialingCode + phone[1:]
}

func (pn phoneNumber) SanitizePlus(phoneNumber string) string {
	phone := pn.sanitize(phoneNumber)
	return "+" + pn.dialingCode + phone[1:]
}

func (pn phoneNumber) Sanitize0(phoneNumber string) string {
	phone := pn.sanitize(phoneNumber)
	return phone
}

func (pn phoneNumber) sanitize(phoneNumber string) string {
	phone := strings.Replace(phoneNumber, " ", "", 0)
	phone = strings.Replace(phone, "-", "", 0)
	if phone[0:2] == pn.dialingCode {
		return "0" + phone[2:]
	} else if phone[0:3] == "+"+pn.dialingCode {
		return "0" + phone[3:]
	}
	return phone
}
