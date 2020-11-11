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

	"H4sIAAAAAAAC/9xYW28bRRT+K9bAA0huvUlbqPatF6lEqBAFeCGKqsnusb32emczMxsRRZZqb4CEcFMf",
	"WiGQ4KEKJZVKLxECFPgzk6Tpv0AzY693vesbWSdK85Bs7Jkz33znO9+ZnXVkkYZPPPA4Q+Y6YlYVGlg9",
	"3gqAcXYL+G1YAPWJT4kPlDv6P8eWv+Ez3PBdQCbCLqvbtTJmbt0uoyLia778mHHqeBXULCIPNyA5RbR+",
	"Fq1d0dpOD28WEYWVwKFgI3NRLtYJsBQNJcs1sLiMHEFdAAZ0FXOHeCwTNI0NUJvg0FAPb1IoIxO9Uerx",
	"UeqQUdLhY7Hlmh0QmFK8loKbWGY4YkIaA6DKbybFSEhjNDgVeDCqecL4JwwoW4CVNKyJslhEAQN6Z0Kp",
	"9OHthhgpgHiGUritKlj1O46XRDJrGO9cMGYuGLMfz1wxjcumceVT413TMFARlQltYI5MZGMOF7jTgKwN",
	"6rgk4HkH7ifNZViSRpzM0TKpKZ6Zi8t2zSEDJ3STOUbxRfzFt9xbNx5wSIakPlOpqRLGU9hdt45rzC47",
	"7ljk4Dqr2WXMxraecFeEoQg3RPuxCPdF+FiEmyJ8kTXbp46VnD5rqJ9orONxqAAdYlvdKMVot1ksvUfO",
	"i+V2kU7HcVX0XA03wpuT32qEJ7NbFUO6bQfUeG57IuFeGk+3CckOhh6cV+SvV6eoSG9NGyhm9YHtQk/J",
	"sodvRPuuaG+ddUf6P7I5YduKEZkgaLCKMrvZuW1NWftU7pSw+JXXoGAypYkZcaBWdkfqaoSk0jQ2i4iB",
	"FVCHr30k24dm7TpgClQ+2cAs6vjajdANCjZ43MEuKxBawJYFjBU4qYNXKBNauDY/h4r6RU0usqzDRItW",
	"OfdRU67peGWiDncO1xt0qO0td+avAmV6vRnJCPHBw76DTHTponHxktQE5lWFs6RKgZX6G3kFeBr8wd+b",
	"L/c2Dv68e7zzq2jfO/zu/uG/D5CKT9XUORuZ2e9rkkJgPvGY5mfWMJS4iMfBU0th33cdS40u1Zi2bt2P",
	"x3s7yjiuKKKSW/jw/UTGkLnYy9XiUlOmF1eYFEKl/62DoSVZk4RlMHN0/6loPRCtDU1RipP+SkNadMD4",
	"dWKv5cZEVkFnkNCHVlpS+3fRfiTaf4lwE8ULgtMAmqnczeScu8RhMI33BgXMwT5R5prFnta7Z8FMkb8K",
	"Hx0+2x5H3irMqei6e6zNQ9D6wJrio7Qu/8zZzVHEtJ4c//bi5d7To/Dzw1+ejcmTshyKG8CBMoUxGVuO",
	"KczdRNLXkKnsqdvATKSBpWRZjNHY7+lLU0+LfjuYSj4CpkhaH+U1ItyRJ4j2HyLcf0u0n+vyfVu07x38",
	"89PR5vcDUhLdAE3JhDLumYZ5UHwXovVEQ5/cki6nifqAFG50NjRRSjQ56ZSU9FkuszgSuxhdE7fhVIxD",
	"3zTkoNI4JVWS84Eh67ZhmvQMut04EU9Vkt1zOmwNbTnHX+4ebn4hWl8fffv8eH9L1cWOLI329qhulLj5",
	"OBXSculFVRKzvpGHqogEzU82CdFVy5RsLX2fM8zVombZSWk+fjahnSVY7hNjot/b4AKHYWehra9e/fAw",
	"m/ibanKX+jNu87nTVkR+MOww1DsGPdw6+nFvgDYDfnb8TKsUgvNfCTm19d7F/mk4cA5NXe0+6ulqDl3t",
	"SjKgbuemwSyVXGJhVw43rxpXjdLqDOpf9dr8HGouNf8LAAD//2NRNn5sHgAA",
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
