package google

import (
	"fmt"
	"time"

	dataproc "google.golang.org/api/dataproc/v1beta2"
)

type DataprocClusterOperationWaiter struct {
	Service *dataproc.Service
	CommonOperationWaiter
}

func (w *DataprocClusterOperationWaiter) QueryOp() (interface{}, error) {
	if w == nil {
		return nil, fmt.Errorf("Cannot query operation, it's unset or nil.")
	}
	return w.Service.Projects.Regions.Operations.Get(w.Op.Name).Do()
}

func dataprocClusterOperationWait(config *Config, op *dataproc.Operation, activity, userAgent string, timeout time.Duration) error {
	w := &DataprocClusterOperationWaiter{
		Service: config.NewDataprocBetaClient(userAgent),
	}
	if err := w.SetOp(op); err != nil {
		return err
	}
	return OperationWait(w, activity, timeout, config.PollInterval)
}
