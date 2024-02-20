package htmlbuild

type Props = []string

func ElementBuilder(name string, components []interface{}) string {
    open := "<" + name
    props := ""
    content := ""

    for _, c := range components {
        switch c.(type) {
        case Props:
            for _, prop := range c.(Props) {
                props += " " + prop
            }
        case string:
            content += c.(string)
        }
    }

    open += props + " >"

    return open + content + "</" + name + ">"
}


func Div(components ...interface{}) string {
    return ElementBuilder("div", components)
}

func P(components ...interface{}) string {
    return ElementBuilder("p", components)
}

func Span(components ...interface{}) string {
    return ElementBuilder("span", components)
}

func A(components ...interface{}) string {
    return ElementBuilder("a", components)
}

func Img(components ...interface{}) string {
    return ElementBuilder("img", components)
}

func H1(components ...interface{}) string {
    return ElementBuilder("h1", components)
}

func Ul(components ...interface{}) string {
    return ElementBuilder("ul", components)
}

func Li(components ...interface{}) string {
    return ElementBuilder("li", components)
}

