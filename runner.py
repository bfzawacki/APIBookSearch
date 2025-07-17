import requests
import json
import os
import sys


# URL da API Go
# utilizar "os" para garantir funcionamento em sistemas operacionais diferentes e adicionar "/search" no fim
# API_URL = os.getenv("API_URL", "http://localhost:3000/api/search")
 
# Envia uma busca para a API e imprime a resposta JSON 
def search_books(query: str):

    api_host = os.getenv("API_HOST", "localhost")
    api_port = os.getenv("API_PORT", "3000")
    api_url = f"http://{api_host}:{api_port}/api/search"

    params = {"q": query}
    print(f"Solicitando URL: {api_url} com busca por: '{query}'")

    try:
        response = requests.get(api_url, params=params)
        response.raise_for_status() 

        data = response.json()
        print("Sucesso! Resposta da API:")
        print(json.dumps(data, indent=2))

    except requests.exceptions.RequestException as e:
        print(f"Ocorreu um erro: {e}")

 
if __name__ == "__main__":
    if len(sys.argv) > 1:
    
        search_query = " ".join(sys.argv[1:])
        search_books(search_query)
    else:
        print("Erro: Forneça um termo de busca como argumento.")
        print("Exemplo: docker run ... \"busca de livro\"")
    