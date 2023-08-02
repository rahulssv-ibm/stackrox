package usagecsv

import (
	"encoding/csv"
	"fmt"
	"io"
	"time"

	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/protoconv"
)

var (
	zeroTime = time.Time{}

	log = logging.LoggerForModule()
)

func writeCSV(metrics <-chan *storage.Usage, iow io.Writer) error {
	csvWriter := csv.NewWriter(iow)
	csvWriter.UseCRLF = true

	record := []string{"Timestamp", "Nodes", "CPU Units"}
	if err := csvWriter.Write(record); err != nil {
		return errors.Wrap(err, "failed to write CSV header")
	}
	for m := range metrics {
		record[0] = protoconv.ConvertTimestampToTimeOrDefault(m.GetTimestamp(), zeroTime).UTC().Format(time.RFC3339)
		record[1] = fmt.Sprint(m.GetNumNodes())
		record[2] = fmt.Sprint(m.GetNumCpuUnits())
		if err := csvWriter.Write(record); err != nil {
			return errors.Wrap(err, "failed to write CSV record")
		}
	}
	csvWriter.Flush()
	return errors.Wrap(csvWriter.Error(), "failed to flush CSV buffer")
}
