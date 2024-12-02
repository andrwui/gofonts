package font

import (
	"errors"
	"fmt"
)

type NameID uint16

const (
	NameIDCopyrightNotice     NameID = iota // Copyright notice
	NameIDFontFamily                        // Family name
	NameIDFontSubfamily                     // Subfamily name
	NameIDUniqueSubfamilyID                 // Unique subfamily identification
	NameIDFullName                          // Full name of the font
	NameIDVersionTable                      // Version of the name table
	NameIDPostScriptName                    // PostScript name of the font
	NameIDTrademark                         // Trademark notice
	NameIDManufacturer                      // Manufacturer name
	NameIDDesigner                          // Designer name
	NameIDDescription                       // Description of the typeface
	NameIDVendorURL                         // URL of the font vendor
	NameIDDesignerURL                       // URL of the font designer
	NameIDLicenseDesc                       // License description
	NameIDLicenseInfoURL                    // License information URL
	_                                       // Reserved
	NameIDPreferredFamily                   // Preferred Family
	NameIDPreferredSubfamily                // Preferred Subfamily
	NameIDCompatibleFull                    // Compatible Full (macOS only)
	NameIDSampleText                        // Sample text
	NameIDOpenTypeID20                      // Defined by OpenType
	NameIDOpenTypeID21                      // Defined by OpenType
	NameIDOpenTypeID22                      // Defined by OpenType
	NameIDOpenTypeID23                      // Defined by OpenType
	NameIDOpenTypeID24                      // Defined by OpenType
	NameIDVarPostScriptPrefix               // Variations PostScript Name Prefix
)

func (f *Font) GetFontName(id NameID) (string, error) {

	n, err := f.Name()
	if err != nil {
		return "", err
	}

	for _, record := range n.NameRecords {
		if int64(record.NameId) == int64(id) {
			return string(n.Name[record.Offset : record.Offset+record.Length]), nil
		}
	}
	return "", errors.New(fmt.Sprintf("Name with id %d not found", id))
}
