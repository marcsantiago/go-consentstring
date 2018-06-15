package consentstring

import (
	"encoding/base64"
	"time"
)

// VendorConsentFields represents the vendor consent fields required accoutned to IAB
// https://github.com/InteractiveAdvertisingBureau/GDPR-Transparency-and-Consent-Framework/blob/master/Consent%20string%20and%20vendor%20list%20formats%20v1.1%20Final.md#example-vendor-consent-string
// this is an incomplete object, only what was needed for my use case was implemented
type VendorConsentFields struct {
	Version           uint
	Created           time.Time
	LastUpdated       time.Time
	CmpID             uint
	CmpVersion        uint
	ConsentScreen     uint
	ConsentLanguage   string
	VendorListVersion uint
	PurposesAllowed   []uint
	MaxVendorID       uint
	EncodingType      uint
	DefaultConsent    uint
}

// CreateConsentString takes the values set in VendorConsentFields and generates a IAB consent string
func (v *VendorConsentFields) CreateConsentString() (string, error) {
	// 240 bits seems safe given the docs
	b := make(bit, 240)

	b.setNumber(versionBitOffset, versionBitSize, int(v.Version))
	b.setDateToDeciseconds(createdBitOffset, createdBitSize, v.Created)
	b.setDateToDeciseconds(updatedBitOffset, updatedBitSize, v.LastUpdated)
	b.setNumber(cmpIDOffset, cmpIDSize, int(v.CmpID))
	b.setNumber(cmpVersionOffset, cmpVersionSize, int(v.CmpVersion))
	b.setNumber(consentScreenSizeOffset, consentScreenSize, int(v.ConsentScreen))
	err := b.setSixBitString(consentLanguageOffset, consentLanguageSize, v.ConsentLanguage)
	if err != nil {
		return "", err
	}

	b.setNumber(vendorListVersionOffset, vendorListVersionSize, int(v.VendorListVersion))
	err = b.setPurposes(v.PurposesAllowed)
	if err != nil {
		return "", err
	}

	b.setNumber(maxVendorIDOffset, maxVendorIDSize, int(v.MaxVendorID))
	b.setNumber(encodingTypeOffset, encodingTypeSize, int(v.EncodingType))
	if v.EncodingType == vendorEncodingRange {
		if v.DefaultConsent == 1 {
			b.setBit(defaultConsentOffset)
		} else {
			b.unsetBit(defaultConsentOffset)
		}
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}
