package fstree

import "github.com/khulnasoft-lab/gobaseline/database/storage"

// Compile time interface checks.
var _ storage.Interface = &FSTree{}
