package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type Item struct {
	Name  string      `yaml:"name"`
	Value interface{} `yaml:"value"`
}

// Index Handler
func (s *Server) handleIndex(c *gin.Context) {
	c.String(http.StatusOK, "YAML Converter Service")
}

// Metrics Handler
func (s *Server) handleMetrics(c *gin.Context) {
	filename, err := filepath.Abs(s.config.YAMLFilePath)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}

	var yamlStruct map[string][]Item
	err = yaml.Unmarshal(yamlFile, &yamlStruct)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}

	metrics := make([]string, 0)
	for metricName, labels := range yamlStruct {
		items := make([]string, len(labels))
		for i, label := range labels {
			items[i] = fmt.Sprintf("%v=%v", label.Name, label.Value)
		}
		metrics = append(metrics, fmt.Sprintf("%v{%s}", metricName, strings.Join(items, ", ")))
	}

	c.String(http.StatusOK, strings.Join(metrics, "\n"))
}
