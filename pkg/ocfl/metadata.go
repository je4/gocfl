package ocfl

import (
	"github.com/je4/gocfl/v2/pkg/checksum"
	"time"
)

type FileMetadata struct {
	Checksums    map[checksum.DigestAlgorithm]string
	InternalName []string
	VersionName  map[string][]string
	Extension    map[string]any
}

type VersionMetadata struct {
	Created time.Time
	Message string
	Name    string
	Address string
}

type ObjectMetadata struct {
	ID              string
	DigestAlgorithm checksum.DigestAlgorithm
	Versions        map[string]*VersionMetadata
	Files           map[string]*FileMetadata
}

type StorageRootMetadata struct {
	Objects map[string]*ObjectMetadata
}
