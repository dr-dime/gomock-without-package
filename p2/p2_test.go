package p2

import "testing"

func TestP1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	NewMockP2(ctrl)
	// ...
}
