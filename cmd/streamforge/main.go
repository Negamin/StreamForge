package main

import (
    "flag"
    "log"
    "github.com/Negamin/StreamForge/internal/config"
    "github.com/Negamin/StreamForge/internal/pipeline"
    "github.com/Negamin/StreamForge/internal/server"
)

func main() {
    configFile := flag.String("config", "examples/simple_pipeline.yaml", "Caminho do arquivo de configuração")
    flag.Parse()

    // Carrega a configuração
    cfg, err := config.LoadConfig(*configFile)
    if err != nil {
        log.Fatalf("Erro ao carregar configuração: %v", err)
    }
    if cfg == nil {
        log.Fatal("Configuração carregada é nula")
    }

    // Inicia o pipeline
    p, err := pipeline.NewPipeline(cfg)
    if err != nil {
        log.Fatalf("Erro ao criar pipeline: %v", err)
    }
    if p == nil || p.SourceChan == nil {
        log.Fatal("Pipeline ou SourceChan é nulo")
    }

    log.Println("Pipeline criado com sucesso, iniciando...")

    go p.Run()

    // Inicia o servidor de monitoramento
    srv := server.NewServer(p)
    if err := srv.Start(":8080"); err != nil {
        log.Fatalf("Erro ao iniciar servidor: %v", err)
    }
}
