# Ginrest

This is a wrapper librari for gin-gonic rest response

## Use

```go
u := c.Request.RequestURI
r := rest.New(u, "").SetGin(c)
v1 := Record{
        Title:  "Hola",
        Object: "object.overrided",
}
v2 := Extra{
        Type: "an Extra",
}
components := rest.Payload{
        "object": "override this",
        "hola":   v1,
        "mundo":  v2,
}
// println(components)
r.Res(429, components, "")
```