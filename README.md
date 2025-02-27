# StreamForge

StreamForge é um pipeline de processamento de dados baseado em plugins para transformar fluxos de texto. Ele permite a fácil integração de transformadores personalizados através de plugins compilados.

## 📌 Características
- 🔌 Suporte a plugins dinâmicos
- 🚀 Baseado em Go com suporte a Goroutines
- 📡 API REST com **Gin**
- 📊 Monitoramento de métricas

---

## 📦 Instalação

### 1️⃣ **Clone o repositório**
```sh
 git clone https://github.com/seu-usuario/StreamForge.git
 cd StreamForge
```

### 2️⃣ **Instale as dependências**
Certifique-se de ter **Go 1.18+** instalado:
```sh
go mod tidy
```

### 3️⃣ **Compile os plugins**
O StreamForge usa plugins para transformar dados. Compile o plugin `uppercase` como exemplo:
```sh
go build -buildmode=plugin -o internal/plugins/uppercase.so internal/plugins/uppercase.go
```

### 4️⃣ **Compile o projeto**
```sh
go build -o streamforge cmd/streamforge/main.go
```

---

## 🚀 Uso

### **Executar o StreamForge**
```sh
./streamforge
```

### **Exemplo de entrada e saída**
O pipeline pode processar strings da fonte e enviá-las para o destino:
```sh
2025/02/27 19:44:55 Pipeline inicializado com sucesso.
2025/02/27 19:44:55 Plugin 1 (internal/plugins/uppercase.so) carregado com sucesso
2025/02/27 19:44:55 Sink recebeu: HELLO WORLD
```

---

## ⚙️ Configuração (config.yaml)
O arquivo `config.yaml` define a estrutura do pipeline:
```yaml
pipeline:
  source: "memory://simple"
  transformations:
    - "internal/plugins/uppercase.so"
  sink: "memory://simple"
```

| Parâmetro         | Descrição                                  |
|------------------|------------------------------------------|
| `source`         | Fonte dos dados                          |
| `transformations`| Lista de plugins para transformar os dados |
| `sink`           | Destino dos dados processados            |

---

## 🔌 Criando um Plugin

Os plugins devem implementar a interface `Transformer`:

```go
package main

import "strings"

type UppercaseTransformer struct{}

func (u *UppercaseTransformer) Transform(input chan string, output chan string) {
    for msg := range input {
        output <- strings.ToUpper(msg)
    }
}

func NewTransformer() *UppercaseTransformer {
    return &UppercaseTransformer{}
}
```

Para compilar:
```sh
go build -buildmode=plugin -o internal/plugins/uppercase.so internal/plugins/uppercase.go
```

---

## 🛠 API Endpoints

O StreamForge expõe uma API REST para monitoramento:

| Método | Rota      | Descrição           |
|--------|----------|-------------------|
| GET    | `/metrics` | Retorna o status do pipeline |

Exemplo de resposta:
```json
{
    "status": "running",
    "queue": 10
}
```

---

## 🏗 Estrutura do Projeto
```
StreamForge/
│── cmd/
│   ├── streamforge/
│   │   └── main.go  # Ponto de entrada do programa
│
│── internal/
│   ├── config/
│   │   └── config.go  # Carregamento do config.yaml
│   ├── pipeline/
│   │   ├── pipeline.go  # Lógica do pipeline
│   │   ├── plugins.go   # Carregamento dos plugins
│   ├── plugins/
│   │   ├── uppercase.go # Exemplo de plugin
│
│── config.yaml  # Configuração do pipeline
│── go.mod  # Dependências do projeto
│── go.sum  # Hash das dependências
```


## 📜 Licença
Este projeto é licenciado sob a **MIT License**. Sinta-se livre para usá-lo e contribuir! 😃

