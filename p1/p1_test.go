package p1

import "testing"

func TestP1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	NewMockP1(ctrl)
	// ...
}
