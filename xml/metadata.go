package xml

import "encoding/xml"

type longDescription struct {
	XMLName  xml.Name `xml:"longdesription"`
	Restrict string   `xml:"restrict,attr"`
	Lang     string   `xml:"lang"`
	Value    string   `xml:",chardata"`
}

type CatMetadata struct {
	XMLName         xml.Name          `xml:"catmetadata"`
	LongDescription []longDescription `xml:"longdescription"`
}

type PkgMetadata struct {
	XMLName         xml.Name          `xml:"pkgmetadata"`
	LongDescription []longDescription `xml:"longdescription"`
	Maintainer      []struct {
		XMLName     xml.Name `xml:"maintainer"`
		Restrict    string   `xml:"restrict,attr"`
		Type        string   `xml:"type,attr"`    // (person|project|unknown) "unknown"
		Proxied     string   `xml:"proxied,attr"` // (yes|no|proxy) "no"
		Email       string   `xml:"email"`
		Name        string   `xml:"name"`
		Description struct {
			XMLName xml.Name `xml:"description"`
			Lang    string   `xml:"lang,attr"`
			Content string   `xml:",chardata"`
		} `xml:"description"`
	} `xml:"maintainer"`
	Slots []struct {
		XMLName xml.Name `xml:"slots"`
		Lang    string   `xml:"lang,attr"`
		Slot    []struct {
			XMLName xml.Name `xml:"slot"`
			Name    string   `xml:"name,attr"`
		} `xml:"slot"`
		SubSlot string `xml:"subslot"`
	} `xml:"slots"`
	// StabilizeAllArches ??? {
	// 	XMLName  xml.Name `xml:"stabilize-allarches"`
	// 	Restrict string   `xml:"restrict,attr"`
	// } `xml:"stabilize-allarches"`
	Use []struct {
		XMLName xml.Name `xml:"use"`
		Lang    string   `xml:"lang,attr"`
		Flag    []struct {
			XMLName  xml.Name `xml:"flag"`
			Name     string   `xml:"name,attr"`
			Restrict string   `xml:"restrict,attr"`
			Content  string   `xml:",chardata"`
		} `xml:"flag"`
	} `xml:"use"`
	Upstream struct {
		Maintainer []struct {
			XMLName  xml.Name `xml:"maintainer"`
			Status   string   `xml:"status,attr"`
			Restrict string   `xml:"restrict,attr"`
			Name     string   `xml:"name"`
			Email    string   `xml:"email"`
		} `xml:"maintainer"`
		Changelog string `xml:"changelog"`
		Doc       []struct {
			XMLName xml.Name `xml:"doc"`
			Lang    string   `xml:"lang,attr"`
			Content string   `xml:",chardata"`
		} `xml:"doc"`
		BugsTo   string `xml:"bugs-to"`
		RemoteID []struct {
			XMLName xml.Name `xml:"remote-id"`
			Type    string   `xml:"type,attr"` // (bitbucket|cpan|cpan-module|cpe|cran|ctan|gentoo|github|gitlab|google-code|hackage|heptapod|launchpad|osdn|pear|pecl|pypi|rubygems|sourceforge|sourcehut|vim)
			Content string   `xml:",chardata"`
		} `xml:"remote-id"`
	} `xml:"upstream"`
}
