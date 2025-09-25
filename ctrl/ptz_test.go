package ctrl

import "testing"

func Test_Ptz(t *testing.T) {
	ptz := &Ptz{
		Zoom:      nil,
		Tilt:      nil,
		Pan:       nil,
		PlanSpeed: 0,
		TiltSpeed: 0,
		ZoomSpeed: 1,
	}
	t.Log(ptz.Value())
}
