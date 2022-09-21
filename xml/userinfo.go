package xml

import "encoding/xml"

type UserList struct {
	XMLName xml.Name `xml:"userlist"`
	User    []struct {
		XMLName  xml.Name `xml:"user"`
		Username string   `xml:"username,attr"`
		ID       uint     `xml:"ID,attr"`
		Realname struct {
			XMLName    xml.Name `xml:"realname"`
			Fullname   string   `xml:"fullname,attr"`
			Firstname  string   `xml:"firstname"`
			Familyname struct {
				XMLName xml.Name `xml:"familyname"`
				Sort    string   `xml:"sort,attr"`
				// check if used anywhere, if not remove to make <familyname/> just string
				Content string `xml:",chardata"`
			} `xml:"familyname"`
		} `xml:"realname"`
		PGPKey   []string `xml:"PGPKey"`
		Alias    []string `xml:"alias"`
		Email    []string `xml:"email"`
		Joined   []string `xml:"joined"`
		Retired  []string `xml:"retired"`
		Status   string   `xml:"status"`
		Roles    string   `xml:"roles"`
		Location struct {
			XMLName   xml.Name `xml:"location"`
			Latitude  string   `xml:"latitude,attr"`
			Longitude string   `xml:"longitude,attr"`
			Content   string   `xml:",chardata"`
		} `xml:"location"`
	} `xml:"user"`
}
