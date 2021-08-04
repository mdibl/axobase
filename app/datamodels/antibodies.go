package datamodels

type AntibodyRecord struct {
	AntigenTarget       string `json:"antigen_target"`
	CellTypes           string `json:"cell_type"`
	ExampleImages       string `json:"example_images"`
	Vendor              string `json:"vendor"`
	SecondaryValidation string `json:"secondary_validation"`
	Dilution            string `json:"dilution"`
	Fixation            string `json:"fixation"`
	Application         string `json:"application"`
	Isotype             string `json:"isotype"`
	CloneName           string `json:"clone_name"`
	SpeciesReactivity   string `json:"species_reactivity"`
	Description         string `json:"description"`
	CatalogNum          string `json:"catalog_num"`
	VendorPageLink      string `json:"vendor_page_link"`
	PublicationLinks    string `json:"publication_links"`
	Notes               string `json:"notes"`
	SubmittedBy         string `json:"submitted_by"`
	Rating              int64  `json:"rating"`
}
