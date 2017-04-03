#HTML Generator

Go & Angular Markup Language - Created by Alex Breadman & Ben Wall

No more writing HTML, allows building easy hybrid single-page-application/server-side-rendering.

#Supports

HTML W3C tags & attributes

AngularJS tags & attributes

Angular-Material tags & attributes

#Usage

```

  element := g.HTML().Add(
    g.HEAD().Add(
      g.TITLE("My Server-Side-rendered page),
      g.FAVICON("/favicon.ico"),
    )
    g.HEADER().Add(
      g.H(1, "Hello world!"),
    ),
    g.BODY().Add(
      g.P("My name is Alex."),
    ),
    g.FOOTER().Add(
      g.A("Click here to go to golang.org").Href("https://www.golang.org"),
    ),
  )

  templateBytes, err := element.Render()

```

#More info coming soon...
