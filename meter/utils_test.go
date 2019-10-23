package meter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type meterTest struct {
	NoTag    *Counter
	MaybeTag *Counter `meter:"YesTag"`
}

func TestTag(t *testing.T) {
	Poll("Poll", time.Second)
	var test meterTest
	Load(&test, "Load")
	require.NotNil(t, Get("Load.NoTag"))
	require.Nil(t, Get("Load.MaybeTag"))
	require.NotNil(t, Get("Load.YesTag"))
}
