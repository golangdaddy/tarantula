package validation

type Language struct {
	Code string 					`json:"code"`
	Native string 					`json:"native"`
	Names map[string]string 		`json:"names"`
	Reverse bool 					`json:"reversed"`
}

func Lang(code, native string, rev bool) *Language { return &Language{code, native, map[string]string{}, rev} }

var languageMap map[string]*Language

func Languages() map[string]*Language {

	return map[string]*Language{
		"EN":	Lang("EN", "English", false),
		"FR":	Lang("FR", "Français", false),
		"DE":	Lang("DE", "Deutsch", false),
		"ES":	Lang("ES", "Español", false),
		"DA":	Lang("DA", "Danske", false),
		"AR":	Lang("AR", "العربية", true),
		"SV":	Lang("SV", "Svenska", false),
		"HE":	Lang("HE", "עברית", true),
		"TR":	Lang("TR", "Turkish", false),
		"CA":	Lang("CA", "Catalan", false),
		"RO":	Lang("RO", "Romanian", false),
		"ID":	Lang("ID", "Bahasa Indonesia", false),
		"CS":	Lang("CS", "Czech", false),
		"MT":	Lang("MT", "Malti", false),
		"HI":	Lang("HI", "Hindi", false),
		"ET":	Lang("ET", "Estonian", false),
	}
}