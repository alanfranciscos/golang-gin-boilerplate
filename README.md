# Go Gin Boilerplate

Este é um boilerplate padronizado para microsserviços em Go, utilizando o framework Gin Gonic e seguindo os princípios da Arquitetura Hexagonal (Ports and Adapters). O projeto foi desenhado para ser resiliente, testável e observável desde o primeiro dia.

## Funcionalidades

- **Arquitetura Hexagonal:** Isolamento total da lógica de negócio das dependências de infraestrutura.
- **Observabilidade:** Integração nativa com OpenTelemetry (OTel) para rastreamento (tracing) e métricas.
- **Monitoramento:** Endpoint de /health padronizado retornando status, uptime e versão.
- **Qualidade de Código:** 
    - Testes unitários obrigatórios com cobertura mínima de 90%.
    - Static analysis com golangci-lint.
    - Formatação padronizada com go fmt.
- **Automação:** Ciclo de vida completo gerenciado via Makefile.
- **Conteinerização:** Dockerfile multi-stage otimizado para produção.
- **CI/CD:** Workflow do GitHub Actions configurado para testes automatizados.

## Tecnologias Utilizadas

- [Go](https://go.dev/) (v1.21+)
- [Gin Gonic](https://gin-gonic.com/) - Framework Web
- [OpenTelemetry](https://opentelemetry.io/) - Observabilidade
- [godotenv](https://github.com/joho/godotenv) - Gerenciamento de variáveis de ambiente
- [Docker](https://www.docker.com/) - Conteinerização

## Estrutura de Pastas

```text
/cmd/api           # Ponto de entrada da aplicação
/internal/domain   # Entidades e modelos de domínio (sem dependências externas)
/internal/application # Casos de uso e interfaces de portas (Ports)
/internal/infrastructure # Adaptadores (Web, Telemetria, Config)
/docs              # Documentação técnica e diagramas de arquitetura
/api               # Definições de API (OpenAPI/Swagger)
/pkg               # Bibliotecas compartilhadas
```

## Como Começar

### Pré-requisitos
- Go 1.21 ou superior
- Docker (opcional, para execução em container)

### Instalação
1. Clone o repositório:
```bash
git clone https://github.com/alanfranciscos/golang-gin-boilerplate.git
cd golang-gin-boilerplate
```

2. Configure as variáveis de ambiente:
```bash
cp .env.example .env
# Edite o arquivo .env com suas configurações
```

3. Instale as dependências:
```bash
go mod tidy
```

## Makefile (Comandos Disponíveis)

| Comando | Descrição |
| :--- | :--- |
| `make run` | Executa a aplicação localmente |
| `make build` | Compila o binário na pasta `bin/` |
| `make test` | Executa todos os testes e exibe cobertura |
| `make format` | Formata o código fonte |
| `make lint` | Executa a análise estática (lint) |
| `make docker-build` | Cria a imagem Docker da aplicação |
| `make docker-run` | Executa a aplicação via Docker |

## Guardrails e Princípios

Este projeto segue rigorosos padrões de engenharia detalhados no arquivo [AGENT.md](./AGENT.md):

1. **Sem ORM:** Consultas ao banco de dados devem ser feitas em SQL Puro.
2. **Isolamento:** A camada de domínio nunca deve importar frameworks externos.
3. **Commits:** Padrão Conventional Commits em inglês.
4. **Configuração:** Todas as variáveis do .env são obrigatórias para o startup.

## Documentação

Para detalhes profundos sobre a arquitetura e fluxos de dados, consulte a pasta [/docs](./docs/ARCHITECTURE.md).
