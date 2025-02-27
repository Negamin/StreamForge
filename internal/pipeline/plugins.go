package pipeline

import (
    "errors"
    "plugin"
)

// LoadPlugin carrega um plugin Go a partir de um arquivo .so
func LoadPlugin(path string) (Transformer, error) {
    p, err := plugin.Open(path)
    if err != nil {
        return nil, err
    }

    // Busca a função exportada "NewTransformer"
    sym, err := p.Lookup("NewTransformer")
    if err != nil {
        return nil, err
    }

    // Converte para a assinatura correta
    newTransformerFunc, ok := sym.(func() Transformer)
    if !ok {
        return nil, errors.New("NewTransformer não tem a assinatura esperada")
    }

    instance := newTransformerFunc()
    if instance == nil {
        return nil, errors.New("NewTransformer retornou um valor nulo")
    }

    return instance, nil
}
