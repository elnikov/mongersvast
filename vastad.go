package mongersvast

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
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

//FromFile load and unmarshal from file
func (v *VAST) FromFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("%s , %s", ErrFailedFileOpen.Error(), err.Error())
	}
	return strings.TrimSpace(string(content)), nil
}

//ToFile save the xml into a file
func (v *VAST) ToFile(filename, body string) (bool, error) {
	var f *File
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
	inline := &InLine{}
	inline.AdSystem = adSystem
	inline.AdTitle = title
	inline.Description = desc
	inline.Error = verr
	inline.Impression = imps
	inline.Creatives = creatives
	//give 1 instance
	req = &VAST{
		Version: VastXMLVer2,
		Ad: []*Ad{
			{InLine: inline},
		},
	}
	//options
	if kk, _ := attrs["ID"]; kk != "" {
		req.Ad[0].ID = kk
	} else if kk, _ := attrs["Sequence"]; kk != "" {
		req.Ad[0].Sequence = kk
	} else if kk, _ := attrs["ConditionalAd"]; kk != "" {
		req.Ad[0].ConditionalAd = kk
	}
	return req
}

//WrapperAd wrapper ad template
func WrapperAd(attrs AdAttributes, adSystem *AdSystem, title *AdTitle, desc *Description, verr *VASTError, imps []*Impression, creatives *Creatives, adURI *VASTAdTagURI) (req *VAST) {
	//minimal config
	wrapper := &Wrapper{}
	wrapper.AdSystem = adSystem
	wrapper.AdTitle = title
	wrapper.Description = desc
	wrapper.Error = verr
	wrapper.Impression = imps
	wrapper.Creatives = creatives
	wrapper.VASTAdTagURI = adURI
	//give 1 instance
	req = &VAST{
		Version: VastXMLVer2,
		Ad: []*Ad{
			{Wrapper: wrapper},
		},
	}
	//options
	if kk, _ := attrs["ID"]; kk != "" {
		req.Ad[0].ID = kk
	} else if kk, _ := attrs["Sequence"]; kk != "" {
		req.Ad[0].Sequence = kk
	} else if kk, _ := attrs["ConditionalAd"]; kk != "" {
		req.Ad[0].ConditionalAd = kk
	}
	return req
}
