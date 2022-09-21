package xml

import "encoding/xml"

type Projects struct {
	XMLName xml.Name `xml:"projects"`
	Project []struct {
		XMLName     xml.Name `xml:"project"`
		Email       string   `xml:"email"`
		Name        string   `xml:"name"`
		Url         string   `xml:"url"`
		Description string   `xml:"description"`
		Subproject  []struct {
			XMLName        xml.Name `xml:"subproject"`
			InheritMembers bool     `xml:"inherit-members,attr"` // (0|1) default "0"
			Ref            string   `xml:"ref,attr"`
			Content        string   `xml:",chardata"`
		} `xml:"subproject"`
		Member []struct {
			XMLName xml.Name `xml:"member"`
			IsLead  bool     `xml:"is-lead,attr"` // (0|1) default "0"
			Email   string   `xml:"email"`
			Name    string   `xml:"name"`
			Role    string   `xml:"role"`
			Content string   `xml:",chardata"`
		} `xml:"member"`
	}
}
