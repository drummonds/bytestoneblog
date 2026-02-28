// afp2book concatenates AFP content into a single markdown document
// suitable for pandoc conversion to PDF via typst.
//
// Usage:
//
//	go run ./cmd/afp2book              # writes combined markdown to stdout
//	go run ./cmd/afp2book -pdf book.pdf # generates PDF via pandoc+typst
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// Chapter defines a section of the book with its content directory
// and ordered sub-pages.
type Chapter struct {
	Dir      string   // relative to content/AFP/
	SubPages []string // filenames in order (empty = just _index.md)
}

// bookOrder defines the chapter ordering for the AFP book.
// This matches the structure in content/AFP/_index.md.
var bookOrder = []Chapter{
	{Dir: ".", SubPages: nil}, // AFP intro
	{Dir: "Historical Accounting", SubPages: []string{
		"Kushim the accountant/_index.md",
		"26Sheep/_index.md",
		"NegativeNumbers/_index.md",
		"AlJabrandPacioli/_index.md",
		"Pacioli/_index.md",
	}},
	{Dir: "Counting", SubPages: nil},
	{Dir: "Movements", SubPages: []string{
		"Boundaries.md",
		"Conventions.md",
		"GraphTheory.md",
		"Oranges.md",
		"OrangesAsMoney.md",
		"OrangesAsMoney_ac.md",
		"OrangesAsMoneyCreditsDebits.md",
		"MeaningOfMoney/_index.md",
		"PaciolioGotItWrong/_index.md",
	}},
	{Dir: "Balances Credits and Debits", SubPages: nil},
	{Dir: "Accounting Records", SubPages: nil},
	{Dir: "Accounting Entries", SubPages: nil},
	{Dir: "Double Entry bookkeeping", SubPages: nil},
	{Dir: "Plain text accounting", SubPages: []string{
		"JournalAsPlainText.md",
		"Times/_index.md",
	}},
	{Dir: "Time series and accounts", SubPages: nil},
	{Dir: "Accounting Tricks", SubPages: nil},
	{Dir: "British American terminology", SubPages: nil},
	{Dir: "Free and commercial software packages", SubPages: nil},
	{Dir: "References", SubPages: []string{
		"Pacioloi translated by Geijbeck/_index.md",
	}},
}

var (
	// Shortcode patterns
	reMermaid      = regexp.MustCompile(`(?s)\{\{[<%]\s*mermaid\s*[>%]\}\}(.*?)\{\{[<%]\s*/mermaid\s*[>%]\}\}`)
	reMermaidFig   = regexp.MustCompile(`(?s)\{\{[<%]\s*mermaid_fig\s+[^>]*[>%]\}\}(.*?)\{\{[<%]\s*/mermaid_fig\s*[>%]\}\}`)
	reAccountClass = regexp.MustCompile(`\{\{[<%]\s*account_class\s*[>%]\}\}`)
	reRef          = regexp.MustCompile(`\{\{[<%]\s*ref\s+(\S+)\s*[>%]\}\}`)
	reFrontmatter  = regexp.MustCompile(`(?s)\A---\n.*?\n---\n*`)
	reRawHTML      = regexp.MustCompile(`(?s)\{\{[<%]\s*rawhtml\s*[>%]\}\}(.*?)\{\{[<%]\s*/rawhtml\s*[>%]\}\}`)
)

func main() {
	pdfOut := flag.String("pdf", "", "output PDF path (requires pandoc+typst)")
	contentDir := flag.String("content", "content/AFP", "AFP content directory")
	flag.Parse()

	var buf bytes.Buffer

	// Book metadata header
	fmt.Fprintln(&buf, "---")
	fmt.Fprintln(&buf, "title: Accounting for Programmers")
	fmt.Fprintln(&buf, "author: Humphrey Drummond")
	fmt.Fprintln(&buf, "---")
	fmt.Fprintln(&buf)

	for _, ch := range bookOrder {
		dir := filepath.Join(*contentDir, ch.Dir)

		// Read chapter _index.md
		indexPath := filepath.Join(dir, "_index.md")
		if ch.Dir == "." {
			indexPath = filepath.Join(*contentDir, "_index.md")
		}
		if err := appendFile(&buf, indexPath, 1); err != nil {
			log.Printf("warning: skipping %s: %v", indexPath, err)
			continue
		}

		// Read sub-pages
		for _, sub := range ch.SubPages {
			subPath := filepath.Join(dir, sub)
			if err := appendFile(&buf, subPath, 2); err != nil {
				log.Printf("warning: skipping %s: %v", subPath, err)
			}
		}
	}

	if *pdfOut == "" {
		io.Copy(os.Stdout, &buf)
		return
	}

	// Write temp markdown file and run pandoc
	tmpFile, err := os.CreateTemp("", "afp-book-*.md")
	if err != nil {
		log.Fatalf("creating temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(buf.Bytes()); err != nil {
		log.Fatalf("writing temp file: %v", err)
	}
	tmpFile.Close()

	args := []string{
		tmpFile.Name(),
		"--pdf-engine=typst",
		"--toc",
		"-V", "mainfont=DejaVu Sans",
		"-o", *pdfOut,
	}
	cmd := exec.Command("pandoc", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("pandoc: %v", err)
	}
	fmt.Fprintf(os.Stderr, "wrote %s\n", *pdfOut)
}

// appendFile reads a markdown file, strips frontmatter, processes shortcodes,
// and appends it to the buffer. headingLevel promotes the first heading.
func appendFile(buf *bytes.Buffer, path string, headingLevel int) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// Strip frontmatter, extract title
	title := extractTitle(data)
	content := reFrontmatter.ReplaceAll(data, nil)

	// Process shortcodes
	content = processShortcodes(content)

	// Write chapter/section heading from frontmatter title
	if title != "" {
		prefix := strings.Repeat("#", headingLevel)
		fmt.Fprintf(buf, "%s %s\n\n", prefix, title)
	}

	buf.Write(content)
	fmt.Fprintln(buf)
	fmt.Fprintln(buf)
	return nil
}

// extractTitle pulls the title from YAML frontmatter.
func extractTitle(data []byte) string {
	lines := strings.Split(string(data), "\n")
	inFront := false
	for _, line := range lines {
		if strings.TrimSpace(line) == "---" {
			if inFront {
				return ""
			}
			inFront = true
			continue
		}
		if inFront && strings.HasPrefix(line, "title:") {
			t := strings.TrimPrefix(line, "title:")
			t = strings.TrimSpace(t)
			t = strings.Trim(t, "\"'")
			return t
		}
	}
	return ""
}

// processShortcodes converts or strips Hugo shortcodes.
func processShortcodes(content []byte) []byte {
	// Convert mermaid_fig to labeled code blocks (must be before generic mermaid)
	content = reMermaidFig.ReplaceAllFunc(content, func(match []byte) []byte {
		inner := reMermaidFig.FindSubmatch(match)
		if len(inner) < 2 {
			return match
		}
		body := bytes.TrimSpace(inner[1])
		// Strip account_class from inside
		body = reAccountClass.ReplaceAll(body, nil)
		return fmt.Appendf(nil, "```mermaid\n%s\n```", body)
	})

	// Convert mermaid to code blocks
	content = reMermaid.ReplaceAllFunc(content, func(match []byte) []byte {
		inner := reMermaid.FindSubmatch(match)
		if len(inner) < 2 {
			return match
		}
		body := bytes.TrimSpace(inner[1])
		body = reAccountClass.ReplaceAll(body, nil)
		return fmt.Appendf(nil, "```mermaid\n%s\n```", body)
	})

	// Strip remaining account_class
	content = reAccountClass.ReplaceAll(content, nil)

	// Convert ref shortcodes to plain text references
	content = reRef.ReplaceAllFunc(content, func(match []byte) []byte {
		sub := reRef.FindSubmatch(match)
		if len(sub) < 2 {
			return match
		}
		ref := string(sub[1])
		return fmt.Appendf(nil, "[%s](/afp/references/#%s)", ref, strings.ToLower(ref))
	})

	// Strip rawhtml blocks
	content = reRawHTML.ReplaceAll(content, nil)

	return content
}
