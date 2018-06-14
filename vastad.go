package mongersvast

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//FromString load and unmarshal from string
func (v *VAST) FromString(body string) {
	body = strings.Replace(body, "\n", "", -1)
	body = strings.Replace(body, "\t", "", -1)
	body = strings.Replace(body, "\r", "", -1)
	xml.Unmarshal([]byte(body), v)
}

//ToString convert to string
func (v *VAST) ToString() (string, error) {
	//sanity check
	if v == nil {
		return "", ErrFailedToStringNilValue
	}
	w := new(bytes.Buffer)
	enc := xml.NewEncoder(w)
	enc.Indent("  ", "    ")
	if err := enc.Encode(v); err != nil {
		return "", fmt.Errorf("%s , %s", ErrFailedToString.Error(), err.Error())
	}
	return strings.TrimSpace(VastXMLHeader + "\n" + w.String()), nil
}

//FromXML an alias to fromString
func (v *VAST) FromXML(body string) {
	v.FromString(body)
}

//ToXML an alias to toString
func (v *VAST) ToXML() (string, error) {
	return v.ToString()
}

//Stringify an alias to toString
func (v *VAST) Stringify() (string, error) {
	return v.ToString()
}

//FromFile load from file
func (v *VAST) FromFile(filename string) {
	content, _ := ioutil.ReadFile(filename) //make sure the xml is readable and exists
	v.FromString(strings.TrimSpace(string(content)))
}

//ToFile save the xml into a file
func (v *VAST) ToFile(filename, body string) (bool, error) {
	var f *os.File
	var err error
	f, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return false, fmt.Errorf("%s , %s", ErrFailedFileSave.Error(), err.Error())
	}
	defer f.Close()
	_, err = f.Write([]byte(body))
	if err != nil {
		return false, fmt.Errorf("%s , %s", ErrFailedFileSave.Error(), err.Error())
	}
	return true, nil
}

//InLineAd inline ad template
func InLineAd(attrs AdAttributes, adSystem *AdSystem, title *AdTitle, desc *Description, verr *VASTError, imps []*Impression, creatives *Creatives) (req *VAST) {
	//minimal config
	req = &VAST{
		Version: VastXMLVer2,
		Ad: []*Ad{
			{InLine: &InLine{
				ID: "1",
				InLineWrapperData: InLineWrapperData{
					AdSystem:    adSystem,
					AdTitle:     title,
					Description: desc,
					Error:       verr,
					Impression:  imps,
					Creatives:   creatives,
				},
			},
			},
		},
	}
	//options
	req.FormatAdAttrs(attrs)
	return
}

//WrapperAd wrapper ad template
func WrapperAd(attrs AdAttributes, adSystem *AdSystem, title *AdTitle, desc *Description, verr *VASTError, imps []*Impression, creatives *Creatives, adURI *VASTAdTagURI) (req *VAST) {
	//minimal config
	req = &VAST{
		Version: VastXMLVer2,
		Ad: []*Ad{
			{Wrapper: &Wrapper{
				ID: "1",
				InLineWrapperData: InLineWrapperData{
					AdSystem:     adSystem,
					AdTitle:      title,
					Description:  desc,
					Error:        verr,
					Impression:   imps,
					Creatives:    creatives,
					VASTAdTagURI: adURI,
				},
			},
			},
		},
	}
	//options
	req.FormatAdAttrs(attrs)
	return
}

//FormatAdAttrs sync all possible options/attrs
func (v *VAST) FormatAdAttrs(attrs AdAttributes) {
	//just in case ;-)
	if v == nil {
		return
	}
	if len(v.Ad) <= 0 {
		return
	}
	//Ad attrs
	if kk, _ := attrs["ID"]; kk != "" {
		v.Ad[0].ID = kk
	}
	//Ad attrs
	if kk, _ := attrs["Sequence"]; kk != "" {
		v.Ad[0].Sequence = kk
	}
	//Ad attrs
	if kk, _ := attrs["ConditionalAd"]; kk != "" {
		v.Ad[0].ConditionalAd = kk
	}
	//VAST version
	if kk, _ := attrs["Version"]; kk != "" {
		switch kk {
		case VastXMLVer3:
			v.Version = VastXMLVer3
		case VastXMLVer4:
			v.Version = VastXMLVer4
			v.XMLNsXs = VastXMLNsXs
			v.XMLNs = VastXMLNs
		default:
			//Ver2.0
			v.Version = VastXMLVer2
		}
	}
}

//SetXMLHeaders set the xml headers simply
func (v *VAST) SetXMLHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/xml")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
	w.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
	w.Header().Set("Expires", "0")
	w.Header().Set("Access-Control-Allow-Origin", "*") //Google HTML5 SDK CORS Header
	//add CORS as per rubiconproject
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Max-Age", "10080")
}

//PushXML push content with proper xml hdrs
func (v *VAST) PushXML(w http.ResponseWriter) {
	//just in case ;-)
	if v == nil {
		return
	}
	xml, _ := v.Stringify()
	v.SetXMLHeaders(w)
	io.WriteString(w, xml)
}
