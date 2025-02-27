package pipeline

import (
    "fmt"
    "log"
    "sync"
    "github.com/Negamin/StreamForge/internal/config"
)

type Pipeline struct {
    SourceChan  chan string
    sinkChan    chan string
    transforms  []Transformer
    wg          sync.WaitGroup
}

type Transformer interface {
    Transform(input chan string, output chan string)
}

// NewPipeline cria um novo pipeline e carrega as transformações
func NewPipeline(cfg *config.Config) (*Pipeline, error) {
    if cfg == nil {
        return nil, logError("Configuração nula fornecida")
    }

    p := &Pipeline{
        SourceChan: make(chan string, 100),
        sinkChan:   make(chan string, 100),
    }

    log.Println("Pipeline inicializado com sucesso.")

    if p.SourceChan == nil || p.sinkChan == nil {
        return nil, logError("Erro ao inicializar canais do pipeline")
    }

    // Carrega transformações (plugins)
    if len(cfg.Pipeline.Transformations) == 0 {
        log.Println("Aviso: Nenhuma transformação definida")
    }

    for i, pluginPath := range cfg.Pipeline.Transformations {
        t, err := LoadPlugin(pluginPath)
        if err != nil {
            log.Printf("Erro ao carregar plugin %s: %v", pluginPath, err)
            return nil, err
        }
        if t == nil {
            return nil, fmt.Errorf("Erro: plugin %s retornou um valor nulo", pluginPath)
        }

        log.Printf("Plugin %d (%s) carregado com sucesso", i+1, pluginPath)
        p.transforms = append(p.transforms, t)
    }

    return p, nil
}

// Run inicia o pipeline, processando os dados da fonte até o destino
func (p *Pipeline) Run() {
    if p == nil {
        log.Fatal("Pipeline é nulo")
    }
    if p.SourceChan == nil {
        log.Fatal("Erro crítico: SourceChan não foi inicializado")
    }
    if p.sinkChan == nil {
        log.Fatal("Erro crítico: sinkChan não foi inicializado")
    }

    log.Println("Iniciando pipeline...")

    p.wg.Add(1)
    go p.source()

    // Encadeia transformações
    currentChan := p.SourceChan
    for i, t := range p.transforms {
        nextChan := make(chan string, 100)
        p.wg.Add(1)

        go func(t Transformer, in, out chan string, isLast bool) {
            defer p.wg.Done()
            t.Transform(in, out)
            if isLast { 
                close(out) // Só fecha o último canal da pipeline
            }
        }(t, currentChan, nextChan, i == len(p.transforms)-1)

        currentChan = nextChan
    }

    p.wg.Add(1)
    go p.sink(currentChan)
    p.wg.Wait()
}


// Gera os dados iniciais
func (p *Pipeline) source() {
    defer p.wg.Done()
    for i := 0; i < 10; i++ {
        p.SourceChan <- "hello world"
    }
    close(p.SourceChan)
}

// Consome os dados processados
func (p *Pipeline) sink(input chan string) {
    defer p.wg.Done()
    for msg := range input {
        log.Printf("Sink recebeu: %s", msg)
    }
}

// logError simplifica a criação de logs de erro
func logError(msg string) error {
    log.Println("Erro: ", msg)
    return fmt.Errorf(msg)
}
