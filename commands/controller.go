package commands

import (
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

var GenerateCmd = &cobra.Command{
	Use:   "generation",
	Short: "Generate files for the project",
}

var controllerCmd = &cobra.Command{
	Use:   "controller [name]",
	Short: "Generate a new controller",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		controllerName := args[0]
		generateController(controllerName)
	},
}

func init() {
	GenerateCmd.AddCommand(controllerCmd)
}

func generateController(name string) {
	tmpl := `package controllers

import "net/http"

type {{.Name}}Controller struct{}

func (c *{{.Name}}Controller) Index(w http.ResponseWriter, r *http.Request) {
	// Handle index
}

func (c *{{.Name}}Controller) Show(w http.ResponseWriter, r *http.Request) {
	// Handle show
}
`
	data := struct {
		Name string
	}{
		Name: name,
	}

	t := template.Must(template.New("controller").Parse(tmpl))

	fileName := fmt.Sprintf("controllers/%s_controller.go", name)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	err = t.Execute(file, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Printf("Controller '%s' created successfully.\n", name)
}
