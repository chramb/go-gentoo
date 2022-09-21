package xml

import "encoding/xml"

type Mirrors struct {
	XMLName     xml.Name `xml:"mirrors"`
	Mirrorgroup []struct {
		XMLName     xml.Name `xml:"mirrorgroup"`
		Region      string   `xml:"region,attr"`
		Country     string   `xml:"country,attr"`
		CountryName string   `xml:"countryname,attr"`
		Mirror      []struct {
			XMLName     xml.Name `xml:"mirror"`
			City        string   `xml:"city,attr"`
			Coordinates string   `xml:"coordinates,attr"`
			GentooBug   string   `xml:"gentoo-bug,attr"`
			Name        string   `xml:"name"`
			Uri         []struct {
				XMLName  xml.Name `xml:"uri"`
				IPv4     string   `xml:"ipv4,attr"`     // (Y|y|N|n) default "y"
				IPv6     string   `xml:"ipv6,attr"`     // (Y|y|N|n) default "n"
				Partial  string   `xml:"partial,attr"`  // (Y|y|N|n) default "n"
				Protocol string   `xml:"protocol,attr"` // (http|ftp|rsync) default "http"
			} `xml:"uri"`
		} `xml:"mirror"`
	} `xml:"mirrorgroup"`
}
