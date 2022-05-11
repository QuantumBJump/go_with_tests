package interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	tt := []struct {
		name string
		want float64
	}{
		{
			name: "basic",
			want: 40.0,
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			rect := Rectangle{10.0, 10.0}
			got := Perimeter(rect)

			if got != test.want {
				t.Errorf("got %.2f, want %.2f", got, test.want)
			}
		})
	}

}

func TestArea(t *testing.T) {

	tt := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangles", shape: Rectangle{12.0, 7.0}, want: 84.0},
		{name: "circles", shape: Circle{10}, want: 314.1592653589793},
		{name: "triangles", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.want {
				t.Errorf("%#v got %g, want %g", test.shape, got, test.want)
			}
		})
	}
}
