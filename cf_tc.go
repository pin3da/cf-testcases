package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strings"
)

func getAtt(t html.Token, att string) string {
	for _, a := range t.Attr {
		if a.Key == att {
			return a.Val
		}
	}
	return ""
}

func getAttNode(t *html.Node, att string) string {
	for _, a := range t.Attr {
		if a.Key == att {
			return a.Val
		}
	}
	return ""
}

func getDataTestCase(t *html.Node) []string {
	var data []string
	n := t.FirstChild.NextSibling
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			data = append(data, c.Data)
		}
	}
	return data
}

func writeToFile(data []string, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	for _, d := range data {
		f.WriteString(d + "\n")
	}
	f.Close()
}

func downloadTestCases(c chan string, url string, problemId string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	b := resp.Body
	defer b.Close() // close Body when the function returns

	tree, _ := html.Parse(b)

	idIn := 1
	idOut := 1

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" {
			if getAttNode(n, "class") == "input" {
				data := getDataTestCase(n)
				writeToFile(data, fmt.Sprintf("%s%d.in", problemId, idIn))
				idIn++
				return
			} else if getAttNode(n, "class") == "output" {
				data := getDataTestCase(n)
				writeToFile(data, fmt.Sprintf("%s%d.ans", problemId, idOut))
				idOut++
				return
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(tree)
	c <- "problem " + problemId + " downloaded"
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage %s contest_id\n", os.Args[0])
		os.Exit(0)
	}
	contestNumber := os.Args[1]
	cfPrefix := "/contest/" + contestNumber
	resp, err := http.Get("http://codeforces.com" + cfPrefix)

	if err != nil {
		log.Fatal(err)
		return
	}

	seen := make(map[string]bool)
	c := make(chan string)
	total := 0
	b := resp.Body
	defer b.Close() // close Body when the function returns

	z := html.NewTokenizer(b)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}
		if tt == html.StartTagToken {
			t := z.Token()
			if t.Data != "a" {
				continue
			}
			problemUrl := getAtt(t, "href")
			if strings.HasPrefix(problemUrl, cfPrefix+"/problem") {
				nextUrl := "http://codeforces.com" + problemUrl
				a := strings.Split(nextUrl, "/")
				if len(a) != 7 {
					continue
				}
				letter := strings.ToLower(a[6])
				if !seen[letter] {
					seen[letter] = true
					go downloadTestCases(c, nextUrl, letter)
					total = total + 1
				}
			}
		}
	}
	for total > 0 {
		fmt.Println(<-c)
		total--
	}
}
