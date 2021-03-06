package main

import (
	"context"

	"github.com/filecoin-project/specs-storage/storage"

	"github.com/EpiK-Protocol/go-epik/build"
	"github.com/filecoin-project/sector-storage"
)

type worker struct {
	*sectorstorage.LocalWorker
}

func (w *worker) Version(context.Context) (build.Version, error) {
	return build.APIVersion, nil
}

var _ storage.Sealer = &worker{}
