<h1 align="center">DontPanic Daily 42 São Paulo 💻</h1>

<h2 align="center">Descrição do projeto</h2>

O jogo consiste em adivinhar a equação escondida da qual sempre resulta em `42`. Cada tentativa deve resultar no numero `42`, e a cada tentativa valida, uma dica será exibida ao jogador.

Por baixo dos panos, ao enviar uma tentativa, consultamos uma API que irá verificar se o calculo está correto ou não, devolvendo uma resposta adequada ao jogador.


## 📁 Acesso ao projeto
A fim de subir a aplicação e jogar o jogo, clone o repositório em sua máquina:
```bash
git clone https://github.com/42sp/42labs-dontpanic_daily-59882565.git dont_panic_daily
```
Mova para o diretorio da aplicação:
```bash
cd dont_panic_daily
```

Para subir o servidor localmente, so é necessario o servico __Docker__ instalado. Uma vez que voce tenha o serviço em execução, basta executar o seguinte comando no diretorio raiz do projeto:
```bash
docker compose up
```
A partir desse instante, quatro serviços estaram em execução: Um servidor MUX para o front-end, servindo as paginas estaticas; Um servidor para o back-end, responsavel pela logica do jogo e suas conexões; Um container executando PostgreSQL, a fim de armazenar informações referentes a solução do dia. E um servidor de _Proxy reverso_, responsavel por fazer o encaminhamento das requests para o ambiente containerizado.

O servidor irá escutar na porta `5011` providenciada pela 42sp.
Uma vez que o servidor esteja em execução, para acessar o jogo, basta acessarmos a URL na respectiva porta:
```
http://labs-bexs-u8968b-y.42sp.org.br:5011/
```

A fim de verificar os end-points que a API providencia, basta acessar o seguinte end-point:
```
http://labs-bexs-u8968b-y.42sp.org.br:5011/swagger/index.html
```


## 🛠️ Tecnologias utilizadas
- **Linguagens**: Golang, JavaScript;
- **Bibliotecas**: gin-gonic, gin-gonic-cors;
- **Tecnologias**: HTML, CSS, Docker, PostgreSQL, Vagrant, Swagger;
- **Ferramentas**: Make, VS Code.


## 🧪 Testes unitários
A aplicação vem com um conjunto extenso de testes unitários referentes a lógica do jogo. Um `makefile` é disponibilizado para subir o servidor, e também para executar todos testes unitários.

Para isso, basta executar o seguinte comando no path `services/server/`
```bash
make test
```

---
## Como jogar
Existe uma equação que resultara no valor `42`, como no exemplo abaixo:

![](imgs/exemple0.png)

No entanto, essa equação, embora resultando no valor `42`, ela nao é a nossa solução. Por conta disso o jogo exibe algumas dicas para que o jogador possa pensar em outra potencial solução:

![](imgs/exemple1.png)

Onde:

__`X`__: Indica que o digito/operador não faz parte da equação escondida;

__`T`__: Indica que o digito/operador faz parte da equação escondida, no entanto, não está exatamente naquela posição em específico;

__`C`__: Indica que o digito/operador faz parte da equação escondida, e está na posição correta.


Com essas informações o jogador pode tentar outras soluções até que eventualmente chegue a solução correta:

![](imgs/exemple2.png)


## Regras
Tenha em mente que existe ordem de precedência nas operações. Por exemplo, multiplicação e divisão serão perfomadas antes de adição e subtração.

Qualquer entrada invalida será devidamente invalidada, não gerando quaisquer dicas, apenas um aviso sobre o respectivo erro.

Propriedade comutativa __não__ é aplicada! Ou seja, `2+3 != 3+2`. A aplicação irá considerar a ordem dos operandos com base na solução providenciada pela API.


---

## Desenvolvedores

<div align="left">
<a href="https://github.com/ichmi"  target="_blank"><img src="https://avatars.githubusercontent.com/u/56981089?v=4" width=115 ></a>

[<sub>frosa-ma | Fábio</sub>](https://github.com/ichmi)

</div>
