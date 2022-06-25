package hardcandy

type MetadataFile struct {
	URI  string `json:"uri"`
	Type string `json:"type"`
}

type MetadataProperties struct {
	Category string         `json:"category"`
	Files    []MetadataFile `json:"files"`
}

type MetadataAttribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

type Metadata struct {
	Image       string              `json:"image"`
	Description string              `json:"description"`
	Properties  MetadataProperties  `json:"properties"`
	Attributes  []MetadataAttribute `json:"attributes"`
}
