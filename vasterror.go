package mongersvast

var VASTErrorCodes = map[string]string{
	"100": "XML parsing error",
	"101": "VAST schema validation error",
	"102": "VAST version of response not supported",
	"200": "Trafficking error Video player received an Ad type that it was not expecting and/or cannot display",
	"201": "Video player expecting different linearity",
	"202": "Video player expecting different duration",
	"203": "Video player expecting different size",
	"300": "General Wrapper error",
	"301": "Timeout of VAST URI provided in Wrapper element, or of VAST URI provided in a subsequent Wrapper element (URI was either unavailable or reached a timeout as defined by the video player)",
	"302": "Wrapper limit reached, as defined by the video player Too many Wrapper responses have  been received with no InLine response",
	"303": "No Ads VAST response after one or more Wrappers",
	"400": "General Linear error Video player is unable to display the Linear Ad",
	"401": "File not found Unable to find Linear/MediaFile from URI",
	"402": "Timeout of MediaFile URI",
	"403": "Couldn’t find MediaFile that is supported by this video player, based on the attributes of the MediaFile element",
	"405": "Problem displaying MediaFile Video player found a MediaFile with supported type but couldn’t display it MediaFile may include: unsupported codecs, different MIME type than MediaFile@type, unsupported delivery method, etc",
	"500": "General NonLinearAds error",
	"501": "Unable to display NonLinear Ad because creative dimensions do not align with creative display area (ie creative dimension too large)",
	"502": "Unable to fetch NonLinearAds/NonLinear resource",
	"503": "Couldn’t find NonLinear resource with supported type",
	"600": "General CompanionAds error",
	"601": "Unable to display Companion because creative dimensions do not fit within Companiondisplay area (ie, no available space)",
	"602": "Unable to display Required Companion",
	"603": "Unable to fetch CompanionAds/Companion resource",
	"604": "Couldn’t find Companion resource with supported type",
	"900": "Undefined Error",
	"901": "General VPAID error",
}
