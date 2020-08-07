package main

import (
    "flag"
    "fmt"
    "math"
)

const (
    width, height = 600, 320
    cells         = 100
    xyrange       = 30.0
    xyscale       = width / 2 / xyrange
    zscale        = height * 0.4
    angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var curType = flag.String("t", "", "type")

func main() {
    flag.Parse()

    fmt.Printf(
        "<svg xmlns='http://www.w3.org/2000/svg' " +
        "style='stroke: grey; fill: white; stroke-width: 0.7' " +
        "width='%d' height='%d'>", width, height,
    )
    for i := 0; i < cells; i += 1 {
        for j := 0; j < cells; j += 1 {
            ax, ay := corner(i+1, j, *curType)
            bx, by := corner(i, j, *curType)
            cx, cy := corner(i, j+1, *curType)
            dx, dy := corner(i+1, j+1, *curType)

            if isFinite(ax) || isFinite(ay) || isFinite(bx) || isFinite(by) ||
                isFinite(cx) || isFinite(cy) || isFinite(dx) || isFinite(dy) {
                continue
            }

            fmt.Printf(
                "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
                ax, ay, bx, by, cx, cy, dx, dy,
            )
        }
    }
    fmt.Println("</svg>")
}

func corner(i, j int, curType string) (float64, float64) {
    x := xyrange * (float64(i) / cells - 0.5)
    y := xyrange * (float64(j) / cells - 0.5)

    var zFunc func(x, y float64) float64
    switch curType {
    case "eggbox":
        zFunc = eggbox
    case "saddle":
        zFunc = saddle
    default:
        zFunc = f
    }
    z := zFunc(x, y)

    sx := width / 2 + (x - y) * cos30 * xyscale
    sy := height / 2 + (x + y) * sin30 * xyscale - z * zscale
    return sx, sy
}

func eggbox(x, y float64) float64 {
    return 0.2 * (math.Cos(x) + math.Cos(y))
}

func saddle(x, y float64) float64 {
    a := 25.0
    b := 17.0
    a2 := a * a
    b2 := b * b
    return (y * y / a2 - x * x / b2)
}

func f(x, y float64) float64 {
    r := math.Hypot(x, y)
    return math.Sin(r) / r
}


func isFinite(f float64) bool {
    return math.IsInf(f, 0) || math.IsNaN(f)
}
