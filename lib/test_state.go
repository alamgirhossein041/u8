package lib

import (
	"io"

	"github.com/sirupsen/logrus"
	"github.com/uvite/u8/metrics"
)

type InitState struct {
	RuntimeOptions RuntimeOptions
	Registry       *metrics.Registry
	BuiltinMetrics *metrics.BuiltinMetrics
	KeyLogger      io.Writer

	// TODO: replace with logrus.FieldLogger when all of the tests can be fixed
	Logger *logrus.Logger
}

// TestRunState contains the pre-init state as well as all of the state and
// options that are necessary for actually running the test.
type TestRunState struct {
	*InitState

	Options Options
	Runner  Runner // TODO: rename to something better, see type comment
	RunTags *metrics.TagSet

	// TODO: add other properties that are computed or derived after init, e.g.
	// thresholds?
}
