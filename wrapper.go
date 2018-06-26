package mongersvast

//SetWrapperAd set the minimum WrapperAd
func (v *VAST) SetWrapperAd(wrapperID, followAdditionalWrappers, allowMultipleAds, fallbackOnNoAd string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//add the wrapper
	v.Ad[v.GetAdPos()].Wrapper = &Wrapper{
		ID: wrapperID,
		FollowAdditionalWrappers: followAdditionalWrappers,
		AllowMultipleAds:         allowMultipleAds,
		FallbackOnNoAd:           fallbackOnNoAd,
		InLineWrapperData:        InLineWrapperData{},
	}
	//good ;-)
	return v
}

//SetVASTAdTagURI set the VASTAdTagURI
func (v *VAST) SetVASTAdTagURI(adID, adValue string) *VAST {
	//min config
	if v == nil {
		v = &VAST{
			Version: VastXMLVer2,
		}
		v.SetAd(VastXMLVer2, "", "", "")
	}
	v.FormatAd()
	//set 1
	data := &VASTAdTagURI{
		ID:    adID,
		Value: adValue,
	}
	//check which type
	if v.Ad[v.GetAdPos()].Wrapper != nil {
		v.Ad[v.GetAdPos()].Wrapper.InLineWrapperData.VASTAdTagURI = data
	} else if v.Ad[v.GetAdPos()].InLine != nil {
		v.Ad[v.GetAdPos()].InLine.InLineWrapperData.VASTAdTagURI = data
	}
	//good ;-)
	return v
}
