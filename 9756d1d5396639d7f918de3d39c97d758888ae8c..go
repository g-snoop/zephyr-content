type ArtifactRecord struct {
	ID           int            `json:"id,omitempty"`            // Database Id
	UUID         string         `json:"uuid"`                    // UUID provide w/previous registration
	Name         string         `json:"name"`                    //  name
	Alias        string         `json:"short_id,omitempty"`      // 15 < len string starts with alphanumeric or '_' followed by any number of alphanumeric, '_' or "."
	Label        string         `json:"label,omitempty"`         // Display name
	Checksum     string         `json:"checksum"`                // <host_address:port> in  http://<host_address:port>
	OpenChain    string         `json:"openchain,omitempty"`     // True if aritfact was prepared under an OpenChain program
	ContentType  string         `json:"content_type,omitempty"`  // Source, notices, envelope, data, spdx, envelope, other
	Timestamp    string         `json:"timestamp,omitempty"`     // Timestamp in UTC format of content at that point in time
	ArtifactList []ArtifactItem `json:"artifact_list,omitempty"` // Aritfact list used only with envelopes to list its artifacts
	URIList      []URIRecord    `json:"uri_list, omitempty"`     // URI references
}

type ArtifactItem struct {
	UUID string `json:"uuid"` // Artifact Universal Unique IDentifier
	Path string `json:"path"` // Path of artifact within the envelope
}

type URIRecord struct {
	Version     string `json:"version"`
	Checksum    string `json:"checksum"`
	ContentType string `json:"content_type"`   // text, envelope, binary, archive
	Size        string `json:"size,omitempty"` // size in bytes
	URIType     string `json:"uri_type"`       // e.g., http, ipfs
	Location    string `json:"location"`       // actual link
}

Add artifact to ledger 