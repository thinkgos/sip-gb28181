package ctrl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Ptz(t *testing.T) {
	ptz := &Ptz{
		Zoom:      nil,
		Tilt:      nil,
		Pan:       nil,
		PlanSpeed: 0,
		TiltSpeed: 0,
		ZoomSpeed: 1,
	}
	require.Equal(t, "A50F0100000010C5", ptz.Value())
}
