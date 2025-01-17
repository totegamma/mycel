package types

import (
	"errors"
	fmt "fmt"
	"regexp"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	NamePattern = `-a-z0-9\p{So}\p{Sk}`
)

func (secondLevelDomain SecondLevelDomain) ValidateName() (err error) {
	regex := regexp.MustCompile(fmt.Sprintf(`(^[%s]+$)`, NamePattern))
	if !regex.MatchString(secondLevelDomain.Name) {
		err = sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s", secondLevelDomain.Name)), ErrInvalidDomainName.Error())
	}
	return err
}

func (secondLevelDomain SecondLevelDomain) ValidateParent() (err error) {
	regex := regexp.MustCompile(fmt.Sprintf(`(^[%s]+[%[1]s\.]*[%[1]s]$)|^$`, NamePattern))
	if !regex.MatchString(secondLevelDomain.Parent) {
		err = sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s", secondLevelDomain.Parent)), ErrInvalidDomainParent.Error())
	}
	return err
}

func (secondLevelDomain SecondLevelDomain) Validate() (err error) {
	err = secondLevelDomain.ValidateName()
	if err != nil {
		return err
	}
	err = secondLevelDomain.ValidateParent()
	if err != nil {
		return err
	}
	return err
}

func ValidateWalletRecordType(walletRecordType string) (err error) {
	_, isFound := NetworkName_value[walletRecordType]
	if !isFound {
		err = sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s", walletRecordType)), ErrInvalidWalletRecordType.Error())
	}
	return err
}

func ValidateDnsRecordValue(dnsRecordFormat string, address string) (err error) {
	dnsRecordRegex, isFound := DnsRecordValueRegex()[dnsRecordFormat]
	if !isFound {
		panic(fmt.Sprintf("Dns record value format %s is not found in DnsRecordValueRegex", dnsRecordFormat))
	}
	regex := regexp.MustCompile(dnsRecordRegex)
	if !regex.MatchString(address) {
		err = sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s %s", dnsRecordFormat, address)), ErrInvalidDnsRecordValue.Error())
	}
	return err
}

func ValidateDnsRecordType(dnsRecordType string) (err error) {
	_, isFound := DnsRecordType_value[dnsRecordType]
	if !isFound {
		err = sdkerrors.Wrapf(errors.New(fmt.Sprintf("%s", dnsRecordType)), ErrInvalidDnsRecordType.Error())
	}
	return err
}
