package formatter

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var titleStyle = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#FAFAFA")).
    Background(lipgloss.Color("#7D56F4")).
	Padding(0,1,0,1)

var definitionStyle = lipgloss.NewStyle().
    Foreground(lipgloss.Color("#FFFFFF")).
	PaddingLeft(1)

var lineSpacing = "\n\n"


type StyledDefinitionBuilder struct {
	title string
	wordDefintions []string
}

func (sdb *StyledDefinitionBuilder) AddTitle(title string) *StyledDefinitionBuilder {
	sdb.title = titleStyle.Render(title) + "\n" + "\n";
	return sdb
}

func (sdb *StyledDefinitionBuilder) AddDefinition(wordDefintion string) *StyledDefinitionBuilder {
	sdb.wordDefintions = append(sdb.wordDefintions, definitionStyle.Render(wordDefintion) + lineSpacing);
	return sdb
}

func (sdb *StyledDefinitionBuilder) AddDefinitions(wordDefintion []string) *StyledDefinitionBuilder {
	for _, wordDefintion := range wordDefintion {
		sdb.wordDefintions = append(sdb.wordDefintions, definitionStyle.Render(wordDefintion) + lineSpacing);
	}
	return sdb
}

func (sdb *StyledDefinitionBuilder) Build() string {
	var builder strings.Builder

	builder.WriteString(sdb.title)

	// Add a divider between these 
	for _, wordDefintion := range sdb.wordDefintions {
		builder.WriteString(wordDefintion)
	}

	return builder.String()
}

func NewDictionaryFormatter() *StyledDefinitionBuilder {
	return &StyledDefinitionBuilder{}
}
