# Pesquisa de livros utilizando Open Library Search
Uma API em Go que recebe requisições em uma URL, faz requisições para uma API externa, e responde com um conjunto de informações obtidas.
Um script em Python para fazer requisições a este serviço em Go, formatar as respostas e exibi-las.

# Processo de desenvolvimento:

# Dificuldades:

# Utilização de Ferramentas de IA:

# Aprendizados
Para entender um pouco melhor da linguagem Go, utilizei o projeto Golings (github.com/mauricioabreu/golings) de estudo. Fiz todos os exercícios propostos e procurei entender a resposta de cada um.
Além disso, mesmo já tendo um pouco de experiência mexendo com arquivos json para fazer pesquisas, nunca havia feito isso sem ter os arquivos localmente em meu computador. Não possuía experiência em APIs ou com o framework negroni.
Utilizei o Docker pela primeira vez e aprendi a sua funcionalidade e utilidade para projetos como esse. Nunca me preocupei com a portabilidade de aplicações já que não é algo necessário na maioria dos projetos da faculdade.

# Compilação e execução:
1.**Clonar o repositório**: 
  * git clone github.com/bfzawacki/APIBookSearch

2.**Entrar no folder do projeto**: 
  * cd APIBookSearch

3.**Criar rede Docker**:
  * docker network create my-app-network

4.**Buildar e rodar o servidor do API**:
  * docker build -t my-api:1.0 -f Dockerfile.api 
  * docker run -d --rm --name api-server --network my-app-network -p 3000:3000 my-api:1.0

5.**Buildar e rodar Runner**:
  * docker build -t my-runner:1.0 -f Dockerfile.runner .

6.**Fazer pesquisas com qualquer query**:
   * docker run --rm --network my-app-network -e API_HOST=api-server -e API_PORT=3000 my-runner:1.0 "Lord of the Rings"
