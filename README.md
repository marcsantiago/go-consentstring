# go-consentstring
Go implementation of the IAB consent string


This implements the IAB ConsentString for the IAB. This package has no affiliation with the IAB and is not officially supported. For more information see below.

https://github.com/InteractiveAdvertisingBureau/GDPR-Transparency-and-Consent-Framework/tree/master/Consent%20String%20SDK

Example useage:

```golang
// pass in a timeout for the http client
vl, err := LoadVendorList(5)
if err != nil {
    log.Fatal(err)
}

// take the vendorlist, cmp information, location of the consent string, and user's language
// base64 consent string that can be passed to SSPs and DSPs
consentString, err := BuildConsentStringFromVendorConsent(vl, 0, 0, 1, "EN")
if err != nil {
    log.Fatal(err)
}

fmt.Println(consentString)
```