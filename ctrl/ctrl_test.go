package ctrl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CombinedCode1(t *testing.T) {
	require.Equal(t, byte(0x0f), CombinedCode1(DeviceControl_Magic, 0x00))
}
