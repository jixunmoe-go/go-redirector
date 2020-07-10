package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type TemplateParams struct {
	Config
	PackageName    string
	Namespace      string
	Name           string
	NamespaceUpper string
	DirectoryURL   string
	FileURL        string

	ImportLine string
	SourceLine string
}

func main() {
	config = readConfig()
	initTemplates()

	fmt.Println("go-redirector: private go repo redirector")
	fmt.Printf("ver: %s-%s (build time: %s)\n\n", Version, BuildMode, BuildTime)

	if config.ListenAddr == "" {
		config.ListenAddr = "127.0.0.1:4001"
	}

	gin.SetMode(BuildMode)

	r := gin.New()
	r.GET("/:namespace/:name", func(c *gin.Context) {
		name := c.Param("name")
		namespace := c.Param("namespace")
		namespaceUpper := strings.ToUpper(namespace)
		packageName := fmt.Sprintf("%s/%s/%s", config.DeployHost, namespace, name)

		params := TemplateParams{
			Config:         config,
			PackageName:    packageName,
			Namespace:      namespace,
			Name:           name,
			NamespaceUpper: namespaceUpper,
			DirectoryURL:   "",
			FileURL:        "",
			ImportLine:     "",
			SourceLine:     "",
		}
		params.BaseWebURL = textTemplateToString(tplBaseWebURL, params)
		params.DirectoryURL = textTemplateToString(tplDirectoryURL, params)
		params.FileURL = textTemplateToString(tplFileURL, params)
		params.ImportLine = textTemplateToString(tplImportURL, params)
		params.SourceLine = textTemplateToString(tplSourceURL, params)

		c.Writer.WriteHeader(200)
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.Header("X-Hello", "github.com/jixunmoe/go-redirector")

		if isGo, ok := c.GetQuery("go-get"); ok && isGo == "1" {
			_ = tplGoGetPage.Execute(c.Writer, params)
		} else {
			_ = tplBrowserPage.Execute(c.Writer, params)
		}
	})

	r.GET("/:namespace", func(c *gin.Context) {
		c.Writer.WriteHeader(403)
	})

	r.GET("/", func(c *gin.Context) {
		c.Redirect(301, config.HomepageURL)
	})

	fmt.Printf("go-redirector is listening at %s\n", config.ListenAddr)
	_ = r.Run(config.ListenAddr)
}
