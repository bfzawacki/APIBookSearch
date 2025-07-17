# Pesquisa de livros utilizando Open Library Search
Uma API em Go que recebe requisições em uma URL, faz requisições para uma API externa, e responde com um conjunto de informações obtidas.
Um script em Python para fazer requisições a este serviço em Go, formatar as respostas e exibi-las.

# Processo de desenvolvimento:

# Dificuldades:

# Utilização de Ferramentas de IA:

# Aprendizados

# Compilação e execução:
Clonar o repositório
    git clone https://github.com/bfzawacki/APIBookSearch
Entrar no folder do projeto
    cd APIBookSearch
Criar rede Docker
    docker network create rede-net
Buildar e rodar o servidor do API
    docker build -t my-api:1.0 -f Dockerfile.api 
    docker run -d --rm --name api-server --network my-app-network -p 3000:3000 my-api:1.0
Buildar e rodar Runner
    docker build -t my-runner:1.0 -f Dockerfile.runner .
Fazer pesquisas com qualquer query
    docker run --rm --network my-app-network -e API_HOST=api-server -e API_PORT=3000 my-runner:1.0 "Lord of the Rings"                
