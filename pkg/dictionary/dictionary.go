package dictionary

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Dictionary interface {
	Define(word string) (Definition, error)
}

type dictionary struct {

}

func (d *dictionary) Define(word string) (Definition, error) {
	url := fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", word)
	resp, err := http.Get(url)
    if err != nil {
        return Definition{}, err
    }
	defer resp.Body.Close()

	// Read the response body into a byte array
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return Definition{}, err
    }

	// Unmarshal the JSON byte array into a Definition struct
	var definition Definition
	err = json.Unmarshal(body, &definition)
	if err != nil {
        return Definition{}, err
	}

	return definition, nil
}

func GetWordDefintions(definition Definition) ([]string) {
	var wordDefintions []string

	// hellish lol
	for _, def := range definition {
		for _, definition := range def.Meanings {
			for i, def := range definition.Definitions {
				if i == 2 {
					break
				}
				wordDefintions = append(wordDefintions, def.Definition)
			}
		}
	}

	return wordDefintions
}

func NewDictionary() Dictionary {
	return &dictionary{}
} 
