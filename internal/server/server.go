package server

import (
    "github.com/gin-gonic/gin"
    "github.com/Negamin/StreamForge/internal/pipeline"
)

type Server struct {
    pipeline *pipeline.Pipeline
}

func NewServer(p *pipeline.Pipeline) *Server {
    return &Server{pipeline: p}
}

func (s *Server) Start(addr string) error {
    r := gin.Default()

    r.GET("/metrics", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status": "running",
            "queue":  len(s.pipeline.SourceChan), // Corrigido para SourceChan
        })
    })

    return r.Run(addr)
}