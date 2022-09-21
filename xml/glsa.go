package xml

// TODO: p,b and other tags should keep order
import "encoding/xml"

type Glsa struct {
	XMLName  xml.Name `xml:"glsa"`
	Id       string   `xml:"id,attr"`
	Title    string   `xml:"title"`
	Synopsis string   `xml:"synopsis"`
	Product  struct {
		XMLName xml.Name `xml:"product"`
		Type    string   `xml:"type,attr"` // (ebuild,informational,infrastructure)
		Content string   `xml:",chardata"`
	} `xml:"product"` //
	Announced string `xml:"announced"`
	Revised   struct {
		XMLName xml.Name `xml:"revised"`
		Count   int      `xml:"count,attr"`
		Content string   `xml:"content"`
	} `xml:"revised"`
	Bug      []uint `xml:"bug"`
	Access   string `xml:"access"` /*?*/
	Affected struct {
		XMLName xml.Name `xml:"affected"`
		Package []struct {
			XMLName   xml.Name `xml:"package"`
			Name      string   `xml:"name,attr"`
			Auto      string   `xml:"auto,attr"`
			Arch      string   `xml:"arch,attr"`
			Vunerable struct {
				XMLName xml.Name `xml:"vunerable"`
				Range   string   `xml:"range,attr"` // (le|lt|eq|gt|ge|rlt|rle|rgt|rge)
				Slot    string   `xml:"slot,attr"`  // default "*"
				Content string   `xml:",chardata"`
			} `xml:"vunerable"`
			Unaffected struct {
				XMLName xml.Name `xml:"vunerable"`
				Range   string   `xml:"range,attr"` // (le|lt|eq|gt|ge|rlt|rle|rgt|rge)
				Slot    string   `xml:"slot,attr"`  // default "*"
				Name    string   `xml:"name,attr"`
				Content string   `xml:",chardata"`
			} `xml:"unaffected"`
		} `xml:"package"` // when product@type = ( 'ebuild' || 'portage' )
		Service []struct {
			XMLName xml.Name `xml:"service"`
			Type    string   `xml:"type,attr"`  // (rsync|web|mirror)
			Fixed   string   `xml:"fixed,attr"` // (yes|no)
			Content string   `xml:",chardata"`
		} `xml:"service"` // when product@type = 'infrastructure'
	} `xml:"affected"`
	Background  xhtml_thingy `xml:"background"` /*?*/
	Description xhtml_thingy `xml:"description"`
	Impact      xhtml_thingy `xml:"impact"`
	Workaround  xhtml_thingy `xml:"workaround"`
	Resolution  xhtml_thingy `xml:"resolution"`
	References  []uri        `xml:"references"`
	// License /*? EMPTY */ `xml:"license"`
	Metadata struct {
		Tag       string `xml:"tag,attr"`
		Revision  string `xml:"revision,attr"`
		Author    string `xml:"author,attr"`
		Timestamp string `xml:"timestamp,attr"`
	} `xml:"metadata"`
}

// TODO: add some universal struct for (p|code|ul|ol)
type xhtml_thingy struct {
}

type uri struct {
	XMLName xml.Name `xml:"uri"`
	Link    string   `xml:"link,attr"`
	Content string   `xml:",chardata"`
}
type mail struct {
	XMLName xml.Name `xml:"mail"`
	Link    string   `xml:"link"`
	Content string   `xml:",chardata"`
}
type p struct {
	XMLName xml.Name `xml:"p"`
	Mail    []mail   `xml:"mail"`

	Content string `xml:",chardata"`
}
