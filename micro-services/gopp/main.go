package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Item represents a single data record
type Item struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// getStopWords returns a set of stop words for the given language.
func getStopWords(language string) map[string]bool {
	englishStopWords := map[string]bool{
		"i": true, "me": true, "my": true, "myself": true,
		"we": true, "our": true, "ours": true, "ourselves": true,
		"you": true, "your": true, "yours": true, "yourself": true, "yourselves": true,
		"he": true, "him": true, "his": true, "himself": true,
		"she": true, "her": true, "hers": true, "herself": true,
		"it": true, "its": true, "itself": true,
		"they": true, "them": true, "their": true, "theirs": true, "themselves": true,
		"what": true, "which": true, "who": true, "whom": true,
		"this": true, "that": true, "these": true, "those": true,
		"am": true, "is": true, "are": true, "was": true, "were": true,
		"be": true, "been": true, "being": true,
		"have": true, "has": true, "had": true, "having": true,
		"do": true, "does": true, "did": true, "doing": true,
		"a": true, "an": true, "the": true,
		"and": true, "but": true, "if": true, "or": true, "because": true, "as": true,
		"until": true, "while": true,
		"of": true, "at": true, "by": true, "for": true, "with": true, "about": true,
		"against": true, "between": true, "into": true, "through": true,
		"during": true, "before": true, "after": true, "above": true, "below": true,
		"to": true, "from": true, "up": true, "down": true,
		"in": true, "out": true, "on": true, "off": true, "over": true, "under": true,
		"again": true, "further": true, "then": true, "once": true,
		"here": true, "there": true, "when": true, "where": true, "why": true, "how": true,
		"all": true, "any": true, "both": true, "each": true, "few": true, "more": true,
		"most": true, "other": true, "some": true, "such": true,
		"no": true, "nor": true, "not": true, "only": true, "own": true, "same": true,
		"so": true, "than": true, "too": true, "very": true,
		"s": true, "t": true, "can": true, "will": true, "just": true,
		"don": true, "should": true, "now": true,
	}

	swedishStopWords := map[string]bool{
		"och": true, "det": true, "att": true, "i": true, "en": true,
		"jag": true, "hon": true, "som": true, "han": true, "på": true,
		"den": true, "med": true, "var": true, "sig": true, "för": true,
		"så": true, "till": true, "är": true, "men": true, "ett": true,
		"om": true, "hade": true, "av": true, "där": true, "min": true,
		"mig": true, "sin": true, "sitt": true, "sina": true, "hans": true,
		"hennes": true, "detta": true, "denna": true, "efter": true,
		"innan": true,
	}

	if strings.ToLower(language) == "swedish" {
		return swedishStopWords
	}
	return englishStopWords
}

// removeStopWords tokenizes the text, filters out any stop words, and rejoins the text.
func removeStopWords(text string, spacer string, language string) string {
	stopWords := getStopWords(language)
	words := strings.Fields(text)
	filteredWords := []string{}
	for _, w := range words {
		lowerW := strings.ToLower(w)
		if !stopWords[lowerW] {
			filteredWords = append(filteredWords, w)
		}
	}
	return strings.Join(filteredWords, spacer)
}

// loadJSON reads JSON data from the specified file and unmarshals it into a slice of Item.
func loadJSON(filePath string) ([]Item, error) {
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var items []Item
	if err := json.Unmarshal(fileData, &items); err != nil {
		return nil, err
	}
	return items, nil
}

// writeJSON marshals the data and writes it to the specified file.
func writeJSON(data []Item, filePath string) error {
	outputData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, outputData, 0644)
}

// filterData processes each item by lowercasing the text and removing stop words.
func filterData(inFilePath, outFilePath, language string) error {
	items, err := loadJSON(inFilePath)
	if err != nil {
		return err
	}
	filtered := []Item{}
	for _, item := range items {
		title := strings.ToLower(item.Title)
		description := strings.ToLower(item.Description)

		title = removeStopWords(title, " ", language)
		description = removeStopWords(description, " ", language)

		filtered = append(filtered, Item{
			Title:       title,
			Description: description,
		})
	}
	// Create the output directory if it does not exist
	outputDir := filepath.Dir(outFilePath)
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}
	return writeJSON(filtered, outFilePath)
}

func main() {
	// Define command-line flags for input, output, and language.
	inputFile := flag.String("input", ".data/wo_data.json", "Path to the input JSON file")
	outputFile := flag.String("output", ".data/cleaned_wo_data.json", "Path to the output JSON file")
	language := flag.String("lang", "swedish", "Language for stop words removal (e.g., english, swedish)")
	flag.Parse()

	if err := filterData(*inputFile, *outputFile, *language); err != nil {
		log.Fatalf("Error filtering data: %v", err)
	}
	fmt.Printf("Data successfully written to %s\n", *outputFile)
}
