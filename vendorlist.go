package consentstring

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	vendorListURL = "https://vendorlist.consensu.org/vendorlist.json"
)

// VendorList can be downloaded from
// https://vendorlist.consensu.org/vendorlist.json
type VendorList struct {
	VendorListVersion uint       `json:"vendorListVersion"`
	LastUpdated       time.Time  `json:"lastUpdated"`
	Purposes          []Purposes `json:"purposes"`
	Features          []Features `json:"features"`
	Vendors           []Vendors  `json:"vendors"`
}

// Purposes ...
type Purposes struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Features ...
type Features struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Vendors ...
type Vendors struct {
	ID               uint   `json:"id"`
	Name             string `json:"name"`
	PolicyURL        string `json:"policyUrl"`
	PurposeIDS       []uint `json:"purposeIds"`
	LegIntPurposeIDS []uint `json:"legIntPurposeIds"`
	FeatureIDS       []uint `json:"featureIds"`
}

// LoadVendorList fetches the lastests vendor list from consensu.org
func LoadVendorList(timeoutSeconds int) (VendorList, error) {
	var v VendorList
	client := http.Client{Timeout: time.Duration(timeoutSeconds) * time.Second}
	res, err := client.Get("https://vendorlist.consensu.org/vendorlist.json")
	if err != nil {
		return v, err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return v, err
	}

	err = json.Unmarshal(b, &v)
	if err != nil {
		return v, err
	}

	return v, nil
}

// BuildConsentStringFromVendorConsent is a convience method
// it assumes version 1 of the IAB consent string, that all vendors should be allowed for all purposes
func BuildConsentStringFromVendorConsent(v VendorList, cmpID, cmpVersion, consentScreen uint, language string) (string, error) {
	vcf := &VendorConsentFields{
		Version:           1,
		Created:           time.Now(),
		CmpID:             cmpID,
		CmpVersion:        cmpVersion,
		LastUpdated:       v.LastUpdated,
		ConsentScreen:     consentScreen,
		ConsentLanguage:   language,
		VendorListVersion: v.VendorListVersion,
		PurposesAllowed:   []uint{1, 1, 1, 1, 1},
		MaxVendorID:       uint(len(v.Vendors)),
		EncodingType:      1,
		DefaultConsent:    1,
	}
	return vcf.CreateConsentString()
}
