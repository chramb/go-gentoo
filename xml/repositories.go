package xml

import "encoding/xml"

type Repository struct {
	XMLName xml.Name `xml:"repositories"`
	Version float32  `xml:"version,attr"`
	Repo    []struct {
		XMLName     xml.Name `xml:"repo"`
		Priority    int      `xml:"priority,attr"`
		Quality     string   `xml:"quality,attr"` /// (core|stable|testing|experimental|graveyard)
		Status      string   `xml:"status,attr"`  // (official|unofficial)
		Name        string   `xml:"name"`
		Description struct {
			XMLName xml.Name `xml:"description"`
			Value   string   `xml:",chardata"`
		} `xml:"description"`
		Homepage string `xml:"homepage"`
		Owner    struct {
			XMLName xml.Name `xml:"owner"`
			Type    string   `xml:"type"`
			Email   string   `xml:"email"`
			Name    string   `xml:"name"`
		} `xml:"owner"`
		Source struct {
			XMLName xml.Name `xml:"source"`
			Type    string   `xml:"type,attr"`
			Value   string   `xml:",chardata"` // (bitbucket|cpan|cpan-module|cpe|cran|ctan|freedesktop-gitlab|gentoo|github|gitlab|gnome-gitlab|google-code|hackage|heptapod|launchpad|osdn|pear|pecl|pypi|rubygems|savannah|savannah-nongnu|sourceforge|sourcehut|vim)
		} `xml:"source"`
		Feed string `xml:"feed"`
	} `xml:"repo"`
}
