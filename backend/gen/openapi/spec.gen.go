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

	"H4sIAAAAAAAC/9xYXW8cNRT9KyvDA0ib7iRpoZo3GqQSIUEV4IUoqpyZu4l3Z8cT2xsRRSN1ZwKkCl/q",
	"Q/uCBE+hpFLplxCgwJ9xNg3/AtnenczueD/azIJKXzJR7et7zzk+9zq7yKOtiIYQCo7cXcS9TWhh/Xkd",
	"xAqlLb4C+teI0QiYIOY3pv5HfRAB5uN1BnXkotdq5wFrvWg1FQfFVSR2IkAuwozhHRTHVcRgq00Y+Mhd",
	"7YVcy1bR9QZ4Qm273gYuuMoHOLBtLAgNR6SVWzB1diZ8LvbkVPPHjM+4JAR7OZaB4w3KxSccGF+BrWJa",
	"IW6B+gmf4VYUqP2y84PsHMnOAcpCcsFIuKFCtjmwm8Qf3IED3vQbdcyDpl8v7hrKtx+ias4enXieoULe",
	"3iZ4zZskHMxkwXHemnPm55yFj+evuM5l17nyqfO26zioiuqUtbBALvKxgDlBWmAr0MSlbVF24GHQAo4V",
	"aJRYVytSCzjzANf9BqEjN/TJHE+Axj7DL1/y+bn5gGMYUvosULNJuSjkHgRN3OB+nQRTgYObvOHXMbfp",
	"qWoTbXok01SmezJ5INNjmT6Q6b5Mn9p2R4x4g9sXHP0vW0tCARvA7Mjpw/tRqlm1NpTeo7P0MR29VBsr",
	"RHz1r92GEmpRjZg3R949s8VmjF/L5JZMbv/X1/tlZH9BD8gBOQCQTUWq4Qwqfut/ICQrZZhTAo16MBHv",
	"CVCPhNGME9P17QtZ4OJ0Djhgfras7S3hlfX3YolxFXHw2oyInY+UEZsCrwFmwNSXD9xjJDL2iZYY+BAK",
	"ggNeoayCPQ84rwjahLBSp6zyzo1lVDVTuDpk3YTJDt0UIkKxOpOEdapbOxFGeYT54Xpv/zYwbs6bVwDQ",
	"CEIcEeSixUvOpUVVDBabOs+avru8NtxwNkAUkz/5Y//5s72T326dHf4kkzvdb+92/7qHdHymty77yLVP",
	"60rbwCMacoPPguPoW09DAaE+CkdRQDy9utbgpteYzjbdbGxpqxqowRI+fH+AMeSunnO1uhYrevEGLzZH",
	"JSPKLZic3n0kO/dkZ8+AU0Bj2PyQURhwcY36O6VhYPNYS/lD2apblPwik/sy+V2m+yivfsHaEBdYmy+Z",
	"tYFxpZjvEgMswH9JzuLqub77ry2rsP9O73cfH0wjaR3mX9Fy/+F4MRGbt2ABidqu+rHsx5Mg6Tw8+/np",
	"82ePTtPPuz8+nhIhbTAMt0AA4zq7wdhqTWX5XaRcDLnajPo+6yKTWEGK1RyAw611beaEmId3yUyoh6+Z",
	"/cc7i0wPVYtLfpXp8RsyeWIu65syuXPy5/en+9+NICN758/Icix/TRjnOPkqZOehSf3FDehyEagPaGWp",
	"V9CUZBjkDRnqvVZq/7M98mZpGaMelWV2vxxQY4307Muj7v4XsvPV6TdPzo5va+YPFfnJwSSPzcqYucWW",
	"ba4TR4OsdIOKfUbIyp7RcNB/OYy7o5np9+gr53a+4OXMnFIvZ9v9HtJmQW8Cdmu1gHo4UIp0rzpXndr2",
	"PBouSk3C8Vr8TwAAAP//1H5MieEWAAA=",
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
