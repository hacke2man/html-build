package htmlbuild

import (
    "strconv"
    "fmt"
    "strings"
)

const (
    Auto string = "auto"
    Center = "center"
)

func Rule(selector string, declarations ...string) string {
    content := selector + " {"

    for _, declaration := range declarations {
        content += declaration  + ";"
    }

    return content + "}"
}

func Style(rules ...string) string {
    open := "<style>"
    content := ""

    for _, rule := range rules {
        content += rule
    }

    return open + content + "</style>"
}

func BackgroundColor(color string) string {
    return "background:" + color
}

func Margin(margins ...interface{}) string {
    var marginValues []string

    for _, margin := range margins {
        switch v := margin.(type) {
        case float64:
            marginValues = append(marginValues, fmt.Sprintf("%fpx", v))
        case string:
            marginValues = append(marginValues, v)
        default:
            // Handle other types or ignore
        }
    }

    return "margin: " + strings.Join(marginValues, " ")
}

func MarginInline(margins ...interface{}) string {
    var marginValues []string

    for _, margin := range margins {
        switch v := margin.(type) {
        case float64:
            marginValues = append(marginValues, fmt.Sprintf("%fpx", v))
        case string:
            marginValues = append(marginValues, v)
        default:
            // Handle other types or ignore
        }
    }

    return "margin-inline: " + strings.Join(marginValues, " ")
}

func MarginBlock(margins ...interface{}) string {
    var marginValues []string

    for _, margin := range margins {
        switch v := margin.(type) {
        case float64:
            marginValues = append(marginValues, fmt.Sprintf("%fpx", v))
        case string:
            marginValues = append(marginValues, v)
        default:
            // Handle other types or ignore
        }
    }

    return "margin-block: " + strings.Join(marginValues, " ")
}

func Width(number string) string {
    return "width:" + number
}

func MarginLeft(number string) string {
    return "margin-left:" + number
}

func MarginRight(number string) string {
    return "margin-right:" + number
}

func Color(color string) string {
    return "color:" + color
}

func FontSize(size string) string {
    return "font-size:" + size
}

func Rem(number float64) string {
    return fmt.Sprintf("%f", number) + "rem"
}

func Em(number float64) string {
    return fmt.Sprintf("%f", number) + "em"
}

func Px(number float64) string {
    return fmt.Sprintf("%f", number) + "px"
}

func Percent(number float64) string {
    return fmt.Sprintf("%f", number) + "%"
}

func Vw(number float64) string {
    return fmt.Sprintf("%f", number) + "vw"
}

func Vh(number float64) string {
    return fmt.Sprintf("%f", number) + "vh"
}

func Rgb(red int, green int, blue int) string {
    return "rgb(" + strconv.Itoa(red) + ", " + strconv.Itoa(green) + ", " + strconv.Itoa(blue) + ")"
}

func Rgba(red int, green int, blue int, alpha float64) string {
    return fmt.Sprintf("rgba(%d, %d, %d, %f)", red, green, blue, alpha)
}

func Hsl(deg int, sat int, value int) string {
    return "hsl(" + strconv.Itoa(deg) + " " + strconv.Itoa(sat) +"%% " + strconv.Itoa(value) + "%%)"
}
