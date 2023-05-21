package tddexample

import "testing"

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("Got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkPerimeter(t, rectangle, 36.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkPerimeter(t, circle, 62.8318530718)
	})
}

//func TestArea(t *testing.T) {
//	t.Run("Rectangle area", func(t *testing.T) {
//		r := Rectangle{10, 20.0}
//
//		got := Area(r)
//		want := 200.0
//
//		if got != want {
//			t.Errorf("got %.3f want %.3f, rec: %v", got, want, r)
//		}
//	})
//}
