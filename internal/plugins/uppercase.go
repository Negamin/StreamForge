package main

import (
    "strings"
    "github.com/Negamin/StreamForge/internal/pipeline"
)

// Definição do Transformer
type UppercaseTransformer struct{}

// Implementação da interface Transformer
func (u *UppercaseTransformer) Transform(input chan string, output chan string) {
    for msg := range input {
        output <- strings.ToUpper(msg)
    }
    // NÃO feche `output` aqui! A pipeline é responsável por isso.
}


// Exporta a função NewTransformer
func NewTransformer() pipeline.Transformer {
    return &UppercaseTransformer{}
}
