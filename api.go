package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/urfave/negroni/v3"
)

// OpenLibraryResponse matches the top-level structure of the API response.
type OpenLibraryResponse struct {
	NumFound int       `json:"numFound"`
	Docs     []BookDoc `json:"docs"`
}

// BookDoc matches the structure of a single book document in the 'docs' array.
type BookDoc struct {
	Title            string   `json:"title"`
	AuthorName       []string `json:"author_name"`
	FirstPublishYear int      `json:"first_publish_year"`
	ISBN             []string `json:"isbn"`
}

// SimplifiedBook is the clean structure we want to return for each book.
type SimplifiedBook struct {
	Title            string `json:"Titulo"`
	AuthorName       string `json:"Autor"`
	FirstPublishYear int    `json:"Publicado em"`
}

// OurAPIResponse is the final response structure for our API.
type OurAPIResponse struct {
	BooksFound int              `json:"Livros encontrados"`
	Results    []SimplifiedBook `json:"Resultados"`
	FetchedAt  time.Time        `json:"Buscado em"`
}

// create a shared HTTP client for performance
var httpClient = &http.Client{
	Timeout: time.Second * 10,
}

// searchHandler queries the Open Library API
func searchHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Get the search query 'q' from the URL
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Por favor, forneça um termo de busca.", http.StatusBadRequest)
		return
	}

	log.Printf("Recebida solicitação de busca: '%s'", query)

	// 2. Make a request to the external API
	// We use url.QueryEscape to safely encode the query string
	externalURL := fmt.Sprintf("https://openlibrary.org/search.json?q=%s&limit=10", url.QueryEscape(query))

	log.Printf("Buscando dados de: %s", externalURL)

	resp, err := httpClient.Get(externalURL)
	if err != nil {
		http.Error(w, "Falha ao conectar com Open Library API", http.StatusInternalServerError)
		log.Printf("Erro de busca na API: %v", err)
		return
	}
	defer resp.Body.Close()

	// 3. Read and decode the external API's JSON response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Falha ao ler o corpo da resposta", http.StatusInternalServerError)
		return
	}

	var openLibraryResponse OpenLibraryResponse
	if err := json.Unmarshal(body, &openLibraryResponse); err != nil {
		http.Error(w, "Falha ao decodificar a resposta da API externa", http.StatusInternalServerError)
		return
	}

	// 4. Create our custom response by processing the results
	ourResponse := OurAPIResponse{
		BooksFound: openLibraryResponse.NumFound,
		FetchedAt:  time.Now(),
		Results:    []SimplifiedBook{}, // Initialize empty slice
	}

	// Loop through the books from Open Library and create our simplified version
	for _, book := range openLibraryResponse.Docs {
		simplified := SimplifiedBook{
			Title:            book.Title,
			FirstPublishYear: book.FirstPublishYear,
			AuthorName:       strings.Join(book.AuthorName, ", "), // Join author names into a single string
		}
		ourResponse.Results = append(ourResponse.Results, simplified)
	}

	log.Printf("Enviando resposta com %d resultados para o cliente.", len(ourResponse.Results))

	// 5. Send our JSON response back to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(ourResponse); err != nil {
		log.Printf("Erro ao codificar a resposta: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome! Try the /api/search?q=query endpoint. Example: /api/search?q=dune")
	})

	// Register our new search handler
	mux.HandleFunc("/api/search", searchHandler)

	n := negroni.Classic()
	n.UseHandler(mux)

	log.Println("Servidor iniciando em :3000...")
	if err := http.ListenAndServe(":3000", n); err != nil {
		log.Fatalf("Não foi possível iniciar o servidor: %s\n", err)
	}
}
