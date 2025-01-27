# Documentação da Aplicação de Teste de Carga

## Resumo

Esta aplicação é um ferramental de teste de carga desenvolvido em Go. Ela permite realizar testes de carga HTTP utilizando múltiplas requisições simultâneas para um URL específico.

## Funcionalidades Principais

- Configuração de parâmetros via flags:
    - URL do serviço a ser testado
    - Número total de requisições
    - Número de chamadas simultâneas

- Execução concorrente das requisições
- Contagem e exibição dos códigos de status HTTP retornados
- Cálculo do tempo total gasto no teste

## Como Usar

1. Construa a imagem Docker executando:
   docker build -t load-test .


2. Execute o teste passando os parâmetros necessários:
   docker run -it --rm
   -e URL=http://exemplo.com
   -e REQUESTS=100
   -e CONCURRENCY=10
   load-test


## Configuração

### Parâmetros

- `URL`: URL do serviço a ser testado (obrigatório)
- `REQUESTS`: Número total de requisições (padrão: 100)
- `CONCURRENCY`: Número de chamadas simultâneas (padrão: 1)

### Construção

A aplicação utiliza uma imagem Docker base do Go Alpine. O processo de construção inclui:

1. Definição da imagem base
2. Criação do diretório de trabalho
3. Cópia dos arquivos necessários
4. Compilação da aplicação
5. Definição da entrada padrão

## Execução

Ao executar o contêiner, a aplicação realizará o seguinte:

1. Inicializar um canal de requisições
2. Criar goroutines para cada thread de execução
3. Enviar requisições ao URL especificado
4. Aguardar o término das requisições
5. Exibir o relatório final

## Relatório

O relatório incluirá informações como:

- URL testada
- Tempo total gasto
- Quantidade total de requisições realizadas
- Distribuição dos códigos de status HTTP retornados

## Considerações Finais

Esta ferramenta é útil para testar o desempenho de serviços web sob carga. Lembre-se de ajustar os parâmetros de acordo com suas necessidades específicas e do serviço sendo testado.
Esta documentação fornece uma visão geral clara da aplicação, incluindo seus principais recursos, como usar, construir e executar. Ela também aborda aspectos importantes como configuração, execução e interpretação dos resultados.

