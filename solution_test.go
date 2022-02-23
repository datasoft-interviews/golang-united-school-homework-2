package square

import (
  "testing"
  "math"
)

func Test1(t *testing.T) {
    want := 4.0
    if got := CalcSquare(2, SidesSquare); got != want {
        t.Errorf("CalcSquare() = %f, want %f", got, want)
    }
}

func Test2(t *testing.T) {
    want := 0.0
    if got := CalcSquare(2, 1); got != want {
        t.Errorf("CalcSquare() = %f, want %f", got, want)
    }
}

func Test3(t *testing.T) {
    want := 28.274334
    if got := CalcSquare(3, SidesCircle); math.Abs(got-want) > 0.0001 {
        t.Errorf("CalcSquare() = %f, want %f", got, want)
    }
}

func Test4(t *testing.T) {
    want := 6.9282032
    if got := CalcSquare(4, SidesTriangle); math.Abs(got-want) > 0.0001 {
        t.Errorf("CalcSquare() = %f, want %f", got, want)
    }
}