package tempconv

import (
    "flag"
    "fmt"
)

type Celsius float64
type Fahrenheit float64

type celsiusFlag struct { Celsius }

func (f *celsiusFlag) Set(s string) error {
    var unit string
    var value float64

    fmt.Sscanf(s, "%f%s", &value, &unit)

    switch unit {
    case "C", "ªC":
        f.Celsius = Celsius(value)
        return nil
    case "F", "ªF":
        f.Celsius = FToC(Fahrenheit(value))
        return nil
    }

    return fmt.Errorf("invalid temperature %q", s)
}

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (c Celsius) String() string    { return fmt.Sprintf("%gªC", c) }

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
    f := celsiusFlag{value}
    flag.CommandLine.Var(&f, name, usage)
    return &f.Celsius
}
