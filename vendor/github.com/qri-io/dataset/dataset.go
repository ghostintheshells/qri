package dataset

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ipfs/go-datastore"
)

// Dataset is stored separately from prescriptive metadata stored in Resource structs
// to maximize overlap of the formal query & resource definitions.
// A Dataset must resolve to one and only one entity, specified by a `data` property.
// It's structure must be specified by a structure definition.
// This also creates space for subjective claims about datasets, and allows metadata
// to take on a higher frequency of change in contrast to the underlying definition.
// In addition, descriptive metadata can and should be author attributed
// associating descriptive claims about a resource with a cyptographic keypair which
// may represent a person, group of people, or software.
// This metadata format is also subject to massive amounts of change.
// Design goals should include making this compatible with the DCAT spec,
// with the one major exception that hashes are acceptable in place of urls.
type Dataset struct {
	// private storage for reference to this object
	path datastore.Key

	// Time this dataset was created. Required. Datasets are immutable, so no "updated"
	Timestamp time.Time `json:"timestamp"`
	// Structure of this dataset, required
	Structure *Structure `json:"structure"`
	// AbstractStructure is the abstract form of the structure field
	AbstractStructure *Structure `json:"abstractStructure,omitempty"`

	// Data is the path to the hash of raw data as it resolves on the network.
	Data datastore.Key `json:"data"`
	// Length is the length of the data object in bytes.
	// must always match & be present
	Length int `json:"length"`
	// Previous connects datasets to form a historical DAG
	Previous datastore.Key `json:"previous,omitempty"`

	// Title of this dataset
	Title string `json:"title,omitempty"`
	// Url to access the dataset
	AccessUrl string `json:"accessUrl,omitempty"`
	// Url that should / must lead directly to the data itself
	DownloadUrl string `json:"downloadUrl,omitempty"`
	// path to readme
	Readme datastore.Key `json:"readme,omitempty"`
	// Author
	Author    *User       `json:"author,omitempty"`
	Citations []*Citation `json:"citations"`
	Image     string      `json:"image,omitempty"`
	// Description follows the DCAT sense of the word, it should be around a paragraph of human-readable
	// text that outlines the
	Description string `json:"description,omitempty"`
	Homepage    string `json:"homepage,omitempty"`
	IconImage   string `json:"iconImage,omitempty"`
	// Identifier is for *other* data catalog specifications. Identifier should not be used
	// or relied on to be unique, because this package does not enforce any of these rules.
	Identifier string `json:"identifier,omitempty"`
	// License will automatically parse to & from a string value if provided as a raw string
	License *License `json:"license,omitempty"`
	// SemVersion this dataset?
	Version VersionNumber `json:"version,omitempty"`
	// String of Keywords
	Keywords []string `json:"keywords,omitempty"`
	// Contribute
	Contributors []*User `json:"contributors,omitempty"`
	// Languages this dataset is written in
	Language []string `json:"language,omitempty"`
	// Theme
	Theme []string `json:"theme,omitempty"`

	// QueryString is the user-inputted string of this query
	QueryString string `json:"queryString,omitempty"`
	// Query is a path to a query that generated this resource
	Query *Query `json:"query,omitempty"`
	// Syntax this query was written in
	QuerySyntax string `json:"querySyntax,omitempty"`
	// queryPlatform is an identifier for the operating system that performed the query
	QueryPlatform string `json:"queryPlatform,omitempty"`
	// QueryEngine is an identifier for the application that produced the result
	QueryEngine string `json:"queryEngine,omitempty"`
	// QueryEngineConfig outlines any configuration that would affect the resulting hash
	QueryEngineConfig map[string]interface{} `json:"queryEngineConfig,omitempty`
	// Resources is a map of dataset names to dataset references this query is derived from
	// all tables referred to in the query should be present here
	Resources map[string]*Dataset `json:"resources,omitempty"`
	// meta holds additional arbitrarty metadata not covered by the spec
	// when encoding & decoding json values here will be hoisted into the
	// Dataset object
	meta map[string]interface{}
}

// Meta gives access to additional metadata not covered by dataset metadata
func (d *Dataset) Meta() map[string]interface{} {
	if d.meta == nil {
		d.meta = map[string]interface{}{}
	}
	return d.meta
}

// Assign collapses all properties of a group of datasets onto one.
// this is directly inspired by Javascript's Object.assign
func (d *Dataset) Assign(datasets ...*Dataset) {
	for _, ds := range datasets {
		if ds == nil {
			continue
		}

		if ds.path.String() != "" {
			d.path = ds.path
		}
		if !ds.Timestamp.IsZero() {
			d.Timestamp = ds.Timestamp
		}

		if d.Structure == nil && ds.Structure != nil {
			d.Structure = ds.Structure
		} else if ds.Structure != nil {
			d.Structure.Assign(ds.Structure)
		}

		if d.AbstractStructure == nil && ds.AbstractStructure != nil {
			d.AbstractStructure = ds.AbstractStructure
		} else if ds.AbstractStructure != nil {
			d.AbstractStructure.Assign(ds.AbstractStructure)
		}

		if ds.Data.String() != "" {
			d.Data = ds.Data
		}
		if ds.Length != 0 {
			d.Length = ds.Length
		}
		if ds.Previous.String() != "" {
			d.Previous = ds.Previous
		}
		if ds.Title != "" {
			d.Title = ds.Title
		}
		if ds.AccessUrl != "" {
			d.AccessUrl = ds.AccessUrl
		}
		if ds.DownloadUrl != "" {
			d.DownloadUrl = ds.DownloadUrl
		}
		if ds.Readme.String() != "" {
			d.Readme = ds.Readme
		}
		if ds.Author != nil {
			d.Author = ds.Author
		}
		if ds.Citations != nil {
			d.Citations = ds.Citations
		}
		if ds.Image != "" {
			d.Image = ds.Image
		}
		if ds.Description != "" {
			d.Description = ds.Description
		}
		if ds.Homepage != "" {
			d.Homepage = ds.Homepage
		}
		if ds.IconImage != "" {
			d.IconImage = ds.IconImage
		}
		if ds.Identifier != "" {
			d.Identifier = ds.Identifier
		}
		if ds.License != nil {
			d.License = ds.License
		}
		if ds.Version != "" {
			d.Version = ds.Version
		}
		if ds.Keywords != nil {
			d.Keywords = ds.Keywords
		}
		if ds.Contributors != nil {
			d.Contributors = ds.Contributors
		}
		if ds.Language != nil {
			d.Language = ds.Language
		}
		if ds.Theme != nil {
			d.Theme = ds.Theme
		}
		if ds.QueryString != "" {
			d.QueryString = ds.QueryString
		}
		if ds.Query != nil {
			d.Query = ds.Query
		}
		if ds.QuerySyntax != "" {
			d.QuerySyntax = ds.QuerySyntax
		}
		if ds.QueryPlatform != "" {
			d.QueryPlatform = ds.QueryPlatform
		}
		if ds.QueryEngine != "" {
			d.QueryEngine = ds.QueryEngine
		}
		if ds.QueryEngineConfig != nil {
			d.QueryEngineConfig = ds.QueryEngineConfig
		}
		if ds.Resources != nil {
			d.Resources = ds.Resources
		}
		if ds.meta != nil {
			d.meta = ds.meta
		}
	}
}

// MarshalJSON uses a map to combine meta & standard fields.
// Marshalling a map[string]interface{} automatically alpha-sorts the keys.
func (d *Dataset) MarshalJSON() ([]byte, error) {
	// if we're dealing with an empty object that has a path specified, marshal to a string instead
	// TODO - check all fields
	if d.path.String() != "" && d.IsEmpty() {
		return d.path.MarshalJSON()
	}

	data := d.Meta()
	if d.AbstractStructure != nil {
		data["abstractStructure"] = d.AbstractStructure
	}
	if d.AccessUrl != "" {
		data["accessUrl"] = d.AccessUrl
	}
	if d.Author != nil {
		data["author"] = d.Author
	}
	if d.Citations != nil {
		data["citations"] = d.Citations
	}
	if d.Contributors != nil {
		data["contributors"] = d.Contributors
	}
	data["data"] = d.Data
	if d.Description != "" {
		data["description"] = d.Description
	}
	if d.DownloadUrl != "" {
		data["downloadUrl"] = d.DownloadUrl
	}
	if d.Homepage != "" {
		data["homepage"] = d.Homepage
	}
	if d.IconImage != "" {
		data["iconImage"] = d.IconImage
	}
	if d.Identifier != "" {
		data["identifier"] = d.Identifier
	}
	if d.Image != "" {
		data["image"] = d.Image
	}
	if d.Keywords != nil {
		data["keywords"] = d.Keywords
	}
	if d.Language != nil {
		data["language"] = d.Language
	}
	data["length"] = d.Length
	if d.License != nil {
		data["license"] = d.License
	}
	if d.Previous.String() != "" {
		data["previous"] = d.Previous
	}
	if d.Query != nil {
		data["query"] = d.Query
	}
	if d.QueryEngine != "" {
		data["queryEngine"] = d.QueryEngine
	}
	if d.QueryEngineConfig != nil {
		data["queryEngineConfig"] = d.QueryEngineConfig
	}
	if d.QueryPlatform != "" {
		data["queryPlatform"] = d.QueryPlatform
	}
	if d.QueryString != "" {
		data["queryString"] = d.QueryString
	}
	if d.QueryPlatform != "" {
		data["querySyntax"] = d.QuerySyntax
	}
	if d.Readme.String() != "" {
		data["readme"] = d.Readme
	}
	if d.Resources != nil {
		data["resources"] = d.Resources
	}
	data["structure"] = d.Structure
	if d.Theme != nil {
		data["theme"] = d.Theme
	}
	data["timestamp"] = d.Timestamp
	data["title"] = d.Title
	if d.Version != VersionNumber("") {
		data["version"] = d.Version
	}

	return json.Marshal(data)
}

// internal struct for json unmarshaling
type _dataset Dataset

// UnmarshalJSON implements json.Unmarshaller
func (d *Dataset) UnmarshalJSON(data []byte) error {
	// first check to see if this is a valid path ref
	var path string
	if err := json.Unmarshal(data, &path); err == nil {
		*d = Dataset{path: datastore.NewKey(path)}
		return nil
	}

	// TODO - I'm guessing what follows could be better
	ds := _dataset{}
	if err := json.Unmarshal(data, &ds); err != nil {
		return fmt.Errorf("error unmarshling dataset: %s", err.Error())
	}

	meta := map[string]interface{}{}
	if err := json.Unmarshal(data, &meta); err != nil {
		return fmt.Errorf("error unmarshaling dataset metadata: %s", err, err)
	}

	for _, f := range []string{
		"abstractStructure",
		"accessUrl",
		"author",
		"citations",
		"contributors",
		"data",
		"description",
		"downloadUrl",
		"homepage",
		"iconImage",
		"identifier",
		"image",
		"keywords",
		"language",
		"length",
		"license",
		"previous",
		"query",
		"queryEngine",
		"queryEngineConfig",
		"queryPlatform",
		"queryString",
		"querySyntax",
		"readme",
		"resources",
		"structure",
		"theme",
		"timestamp",
		"title",
		"version",
	} {
		delete(meta, f)
	}

	ds.meta = meta
	*d = Dataset(ds)
	return nil
}

func (ds *Dataset) IsEmpty() bool {
	return ds.Title == "" && ds.Description == "" && ds.Structure == nil && ds.Timestamp.IsZero() && ds.Previous.String() == ""
}

func (ds *Dataset) Path() datastore.Key {
	return ds.path
}

// UnmarshalDataset tries to extract a dataset type from an empty
// interface. Pairs nicely with datastore.Get() from github.com/ipfs/go-datastore
func UnmarshalDataset(v interface{}) (*Dataset, error) {
	switch r := v.(type) {
	case *Dataset:
		return r, nil
	case Dataset:
		return &r, nil
	case []byte:
		dataset := &Dataset{}
		err := json.Unmarshal(r, dataset)
		return dataset, err
	default:
		return nil, fmt.Errorf("couldn't parse dataset, value is invalid type")
	}
}