package volume

import (
	"github.com/concourse/baggageclaim/uidjunk"
	"code.cloudfoundry.org/lager"
)

type NamespacedStrategy struct {
	PreStrategy Strategy

	Namespacer uidjunk.Namespacer
}

func (strategy NamespacedStrategy) Materialize(logger lager.Logger, handle string, fs Filesystem) (FilesystemInitVolume, error) {
	volume, err := strategy.PreStrategy.Materialize(logger, handle, fs)
	if err != nil {
		return nil, err
	}

	err = strategy.Namespacer.Namespace(volume.DataPath())
	if err != nil {
		volume.Destroy()
		return nil, err
	}

	return volume, nil
}
