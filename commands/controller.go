package commands

import (
	"fmt"
	"os"
	"strings"
	"text/template"
	"unicode"

	"github.com/brandaoplaster/amethyst/templates"
	"github.com/spf13/cobra"
)

// Define o comando 'GenerateCmd' como uma variável exportada
var GenerateCmd = &cobra.Command{
	Use:   "generation",
	Short: "Generate files for the project",
}

// Define o subcomando 'controllerCmd'
var controllerCmd = &cobra.Command{
	Use:   "controller [name] [actions...]",
	Short: "Generate a new controller with specified actions",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		controllerName := args[0]
		actions := args[1:]
		generateController(controllerName, actions)
	},
}

func init() {
	// Adiciona o subcomando 'controllerCmd' ao comando 'GenerateCmd'
	GenerateCmd.AddCommand(controllerCmd)
}

// Função para gerar um novo arquivo de controller
func generateController(name string, actions []string) {
	tmpl := templates.GenerationController()
	data := struct {
		Name    string
		Actions []string
	}{
		Name:    name,
		Actions: actions,
	}

	data.Actions = CapitalizeSlice(data.Actions)

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

	fmt.Printf("Controller '%s' with actions %s created successfully.\n", name, strings.Join(actions, ", "))
}

func Capitalize(s string) string {
	var result strings.Builder
	capitalize := true

	for _, r := range s {
		if unicode.IsSpace(r) {
			capitalize = true
			result.WriteRune(r)
		} else if capitalize && unicode.IsLetter(r) {
			result.WriteRune(unicode.ToUpper(r))
			capitalize = false
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()
}

func CapitalizeSlice(slice []string) []string {
	for i, str := range slice {
		slice[i] = Capitalize(str)
	}
	return slice
}
