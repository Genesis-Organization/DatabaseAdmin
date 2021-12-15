package controllers

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"

	tmpl "GenesisDAT/src/templates"
)

func ServeIndex(c *gin.Context) {
	buf := new(bytes.Buffer)
	tmpl.Index("Test Index Page", true, buf)
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(buf.String()))
}
