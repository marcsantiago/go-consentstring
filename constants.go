package consentstring

const (
	vendorEncodingRange = 1

	versionBitOffset = 0
	versionBitSize   = 6

	createdBitOffset = 6
	createdBitSize   = 36

	updatedBitOffset = 42
	updatedBitSize   = 36

	cmpIDOffset = 78
	cmpIDSize   = 12

	cmpVersionOffset = 90
	cmpVersionSize   = 12

	consentScreenSizeOffset = 102
	consentScreenSize       = 6

	consentLanguageOffset = 108
	consentLanguageSize   = 12

	vendorListVersionOffset = 120
	vendorListVersionSize   = 12

	purposesOffset = 132
	purposesSize   = 24

	maxVendorIDOffset = 156
	maxVendorIDSize   = 16

	encodingTypeOffset = 172
	encodingTypeSize   = 1

	vendorBitfieldOffset = 173

	defaultConsentOffset = 173

	numEntriesOffset = 174
	numEntriesSize   = 12

	rangeEntryOffset = 186

	vendorIDSize = 16
)
