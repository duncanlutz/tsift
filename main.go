package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type InterfaceDoc struct {
	Name        string
	Description string
	Properties  []PropertyDoc
}

type PropertyDoc struct {
	Name     string
	Type     string
	Required bool
}

func main() {
	var files string
	var directory string
	var outputFlag string
	outputFile := new(string)

	flag.StringVar(&files, "f", "", "Comma separated list of files to process")
	flag.StringVar(&files, "files", "", "Comma separated list of files to process")
	flag.StringVar(&directory, "d", "", "Directory to process")
	flag.StringVar(&directory, "directory", "", "Directory to process")
	flag.StringVar(&outputFlag, "o", "", "Output file")
	flag.StringVar(&outputFlag, "output", "", "Output file")

	flag.Parse()

	if outputFlag == "" {
		outputFile = nil
	} else {
		outputFile = &outputFlag
	}

	interfaces := []InterfaceDoc{}

	if files != "" {
		interfaces = append(interfaces, processFiles(files)...)
	} else if directory != "" {
		interfaces = append(interfaces, processDirectory(directory)...)
	} else {
		fmt.Println("You must provide either a list of files or a directory to process")
		os.Exit(1)
	}

	outputMarkdown(interfaces, outputFile)
}

func processFiles(files string) []InterfaceDoc {
	result := []InterfaceDoc{}

	fileList := strings.Split(files, ",")
	for _, file := range fileList {
		interfaces, err := processFile(file)
		if err != nil {
			fmt.Printf("Error processing file %s: %v\n", file, err)
			os.Exit(1)
		}

		result = append(result, interfaces...)
	}

	return result
}

func processDirectory(directory string) []InterfaceDoc {
	interfaces, err := findInterfaces(directory)
	if err != nil {
		fmt.Printf("Error processing directory: %v\n", err)
		os.Exit(1)
	}

	return interfaces
}

func findInterfaces(root string) ([]InterfaceDoc, error) {
	var interfaces []InterfaceDoc

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && (strings.HasSuffix(path, ".ts") || strings.HasSuffix(path, ".tsx")) {
			fileInterfaces, err := processFile(path)
			if err != nil {
				return fmt.Errorf("error processing file %s: %v", path, err)
			}
			interfaces = append(interfaces, fileInterfaces...)
		}

		return nil
	})

	return interfaces, err
}

func processFile(filePath string) ([]InterfaceDoc, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var interfaces []InterfaceDoc
	var currentInterface *InterfaceDoc
	var docComment strings.Builder
	var inInterface bool

	scanner := bufio.NewScanner(file)
	interfaceRegex := regexp.MustCompile(`^interface\s+(\w+)\s*{`)
	propertyRegex := regexp.MustCompile(`^\s*(\w+)(\?)?:\s*([^;]+)`)

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)

		if matches := interfaceRegex.FindStringSubmatch(trimmedLine); matches != nil {
			currentInterface = &InterfaceDoc{
				Name:        matches[1],
				Description: strings.TrimSpace(docComment.String()),
				Properties:  make([]PropertyDoc, 0),
			}
			inInterface = true
			docComment.Reset()
			continue
		}

		if inInterface && trimmedLine == "}" {
			interfaces = append(interfaces, *currentInterface)
			currentInterface = nil
			inInterface = false
			continue
		}

		if inInterface {
			if matches := propertyRegex.FindStringSubmatch(trimmedLine); matches != nil {
				property := PropertyDoc{
					Name:     matches[1],
					Type:     matches[3],
					Required: matches[2] == "", // If there's no "?", it's required
				}
				currentInterface.Properties = append(currentInterface.Properties, property)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return interfaces, nil
}

func outputMarkdown(interfaces []InterfaceDoc, file *string) {
	if file != nil {
		f, err := os.Create(*file)
		if err != nil {
			fmt.Printf("Error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()

		os.Stdout = f
	}

	for _, iface := range interfaces {
		fmt.Printf("# Interface: %s\n\n", iface.Name)

		if iface.Description != "" {
			fmt.Printf("%s\n\n", iface.Description)
		}

		fmt.Println("| Property | Type | Required |")
		fmt.Println("|----------|------|----------|")

		for _, prop := range iface.Properties {
			required := "Yes"
			if !prop.Required {
				required = "No"
			}
			fmt.Printf("| %s | %s | %s |\n", prop.Name, prop.Type, required)
		}

		fmt.Printf("\n\n")
	}
}
