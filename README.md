# Html Build
Go library ment for creating html as a string, with better type checking/completion, then just making strings.

This has definitely been done before in a few different languages, and probably in go as well, but I haven't come across a go implementation yet.

The main use case I see for this is with htmx, but I'm sure there's many others

The current state of this library is no where near usable, as there is barely any coverage, and any breaking change is on the table

``` go

package main

import (
    . "mite.lol/htmlbuild"
)

func main() {
    content := Div (
        Props {"class='foo'", "cup='bar'"},
        "hello",
    )

    content += Style(
            Rule(".foo",
                BackgroundColor(Hsl(0, 50, 50)),
                MarginInline(Auto),
                Width(Rem(5)),
            ),
        )

    println(content)
}

```

output:
``` text
<div class='foo' cup='bar' >hello</div><style>.foo {background:hsl(0 50%% 50%%);margin-inline: auto;width:5.000000rem;}</style>
```
