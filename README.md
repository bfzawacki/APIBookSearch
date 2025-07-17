# Pesquisa de livros utilizando Open Library Search
Uma API em Go que recebe requisições em uma URL, faz requisições para uma API externa, e responde com um conjunto de informações obtidas.
Um script em Python para fazer requisições a este serviço em Go, formatar as respostas e exibi-las.

# Processo de desenvolvimento:
Após utilizar o código base para a utilização da framework negroni, ver esse vídeo: (youtube.com/watch?v=reN_okp2Gq4) que fala sobre como funciona o OpenLibrary e consultar o código do criador do vídeo, comecei meu projeto. Fiz o desenvolvimento do arquivo API.go e runner.py. Consegui acessar a API e rodar runner.py para fazer algumas pesquisas de teste, porém, ainda não deixando o usuário digitar suas próprias querys. Eu também ainda não estava utilizando o Docker, que foi adicionado após o código estar funcionando adequadamente. Para fazer o docker funcionar, precisei fazer alterações e criar files novas. Alterei o código para permitir que o usuário digite querys, ajustei os comentários e apaguei algumas coisas não utilizadas mais.

# Utilização de Ferramentas de IA:
Para auxiliar na implementação desse projeto, foi utilizado a ferramenta GeminiAI. 
Foi utilizado IA para a criação dos arquivos Docker (Dockerfile.api, Dockerfile.runner, entrypoint.sh) e para rodar o código também através dos comandos docker no terminal. Utilizar IA também facilitou muito para a detecção de problemas e erros no meu código.
Além disso, o Gemini me ajudou a construir e corrigir várias partes do código da aplicação em Go da API e do runner em python, principalmente para a criação de structs e da lógica do arquivo da API (fica evidenciado pelos comentários em inglês).

# Aprendizados
Para entender um pouco melhor da linguagem Go, utilizei o projeto Golings (github.com/mauricioabreu/golings) de estudo. Fiz todos os exercícios propostos e procurei entender a resposta de cada um.
Além disso, mesmo já tendo um pouco de experiência mexendo com arquivos json para fazer pesquisas, nunca havia feito isso sem ter os arquivos localmente em meu computador. Não possuía experiência em APIs ou com o framework negroni.
Utilizei o Docker pela primeira vez e aprendi a sua funcionalidade e utilidade para projetos como esse. Nunca me preocupei com a portabilidade de aplicações já que não é algo necessário na maioria dos projetos da faculdade.

# Dificuldades:
As principais dificuldades estavam no uso do Docker. Tive que debuggar bastante as files Docker e também fiquei confuso algumas vezes em como buildar e rodar meu código no Docker.
Para a implementação dos arquivos do API e runner também não foi trivial fazer a lógica e conexão entre eles. Por isso, foi essencial o uso de IA nesse momento para eu conseguir concluir.
A questão do tempo, o prazo não era curto, então, peço desculpas pela demora da entrega. Apenas consegui trabalhar nisso (incluindo os exercícios do Golings) na segunda-feira, dia 14, e quarta-feira, dia 16. 

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
