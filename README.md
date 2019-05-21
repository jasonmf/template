# tp - Template Processor

`tp` is a trivial command-line template processor for golang/pkg/text/template. It takes its data as a JSON file specified in the first argument and the template specified in the second argument. Output is to stdout.

## Building

```
go build -o /tmp/tp tp.go
```

## Usage

Given `vars.json`:

```
{
        "SomeVar": 5,
        "Other": [
                {
                        "foo": "a",
                        "bar": 1
                },
                {
                        "foo": "b",
                        "bar": 2
                }
        ]
}
```


and `test.tmpl`:

```
SomeVAR: {{.SomeVar}}

{{range .Other}}
---------------------------------------
FOO: {{.foo}}
BAR: {{.bar}}
{{end}}
```

You can:

```
/tmp/tp vars.json test.tmpl
```

and you should get

```
SomeVAR: 5


---------------------------------------
FOO: a
BAR: 1

---------------------------------------
FOO: b
BAR: 2

```

