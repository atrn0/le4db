// Package Openapi provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package Openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9xXW28cNRT+K5HhAaRNZ5K0UM0bFAkiJKii8kIUVe7M2cS7M+OJ7Y2IopW6MwFShZuK",
	"1L4g0adQUqn0JsQl8GecTcO/QLazk5kd76XNbqu2L7vb2MfnfN93PvtsIZ9GCY0hFhx5W4j7axBh/fVD",
	"EEvAgW1gQWjMl0D/b8JoAkwQ84sVFqjfRECkv7zJoI489IZzGt45ie0UoqJ2DYnNBJCHMGN4E7XbNcRg",
	"vUUYBMhbLh+wki+m1xrgC7VbZUlpNCA99Zfx86I0Gp2QDmnNpAVc8FcEtTzXCWFnAk4EwcuUi884ML4E",
	"69W0YhyB+oQvcJSEar/s/Cw7+7Kzi/KQXDASr6qQLQ7sKgnKO3DIm0GjjnnYDOrVXX359kLUzNmDEy8y",
	"VMnbXwO/eZXE5UzmXfedWXdu1p2/MnfBc8977oXP3Xc910U1VKcswgJ5KMACZgWJwFagiUtbYtKB+0EL",
	"OVagUWJdrUit4MxDXA8ahA7c0CNzOAEa+xy/Ysmn5xYDDmFI6bNCzRrlopJ7GDZxgwd1Etqy11tsUvxb",
	"pj/Kzp2xAMVN3gjqmNs0WLMJPduXWSazbZnek9mBzO7JbEdmj227E0b88vZ5V//L15JYwCowO9r68F6U",
	"Wo5QsXAbyqp1y+63/hp0glXbmFMCjXo40j5GCHcgjMaYx3PAMwljYTxdlCRhy/r1sr9VZRhVV8C8OdAD",
	"zRabK3wr0+syvfGybfZ5FHNGLy4AWQLIKiCrO7+ytlktUe0hcZ3qK48I4yOEBfG1mfcuL6Ia2gDGdfug",
	"OZUNTSDGCUEeWjjnnltQkbFY06g4Gkvu9D8lV0F3QgDcZyQxvYgO/9p5+mT78I/rx3u/yPRm9/tb3X9v",
	"Ix2f6a2LAfLs71dFJvCExtywMe+6up1pLCDWR+EkCYmvVzsNbnrfvArHezNaHswaqHIJn36sEefgtxgR",
	"m8hbVsq4QpsQI295pa3Axqu8+vhVpFJuAeXo1gPZuS072wadChz9dxkyfAMX79Ngc2Ig2K5MS/192SpN",
	"p7/J9K5M/5TZDipqUbAWtCu0zU2YttI4Us33EgMsIHhe0tq1U4X35hCrtP/L7nYf7o4jah3mhai5N1Kd",
	"UcZmTKpA4Wypj8WgPQqTzv3jXx8/ffLgKPuye+fhmBBpj2E4AgGM6/TKsdWamcUPkDIy5Gk/6vmeh0xi",
	"FTHWCgj23y4rU2fEzKSTpkINhWZoHm4uMttTd076u8wO3pLpI9Ovb8v05uE/Px3t/DCAjXwGnpLrWCbt",
	"YaZTrEJ27pvUn92DzleB+oTOXDopaFw2DPSGDTWOTPQS/Ii+4DtwyrdfAaShPnr89X535yvZ+ebou0fH",
	"Bzc07XuK+XR3lMXmkE3dYSfurSPfBnntBhb7IyGve0qvg94kOKxDc88/4W8yvfmsrZkbpV7PNnp3SIuF",
	"yENrQiSe44TUx6HSpHfRveg6G3Oovyr1GG6vtP8PAAD//8KDYy6cFgAA",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
