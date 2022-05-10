// Package cdnUser GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package cdnUser

import "github.com/swaggo/swag"

const docTemplatecdnUser = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {}
}`

// SwaggerInfocdnUser holds exported Swagger Info so clients can modify it
var SwaggerInfocdnUser = &swag.Spec{
	Version:          "1.0",
	Host:             "petstore.swagger.io:8080",
	BasePath:         "/v2",
	Schemes:          []string{},
	Title:            "cdnAdmin的swagger文档",
	Description:      "This is a sample server Petstore server.",
	InfoInstanceName: "cdnUser",
	SwaggerTemplate:  docTemplatecdnUser,
}

func init() {
	swag.Register(SwaggerInfocdnUser.InstanceName(), SwaggerInfocdnUser)
}