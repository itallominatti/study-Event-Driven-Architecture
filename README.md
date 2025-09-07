Eventos:
    Situações que ocorreram no passado
    Normalmente deixa efeitos colaterais. Ex: Porta do carro abriu. Ligar a luz interna.
    Trabalhar de forma internalizada no software ou externalizada
    Domain Events: Eventos de domínio: Mudança no estado interno da aplicação / regra de negócios -> ex: Agregados


3 tipos de eventos
    1) Event Notification. Forma curta de comunicação. {"pedido": 1, "status": "aprovado"}
    2) Event Carried State Transfer: Formato completo para trafegar as informações. Steam de dados { "pedido": 1, "produtos": [{}], valor: 10.00, tax: "1%", "comprador": "Itallo" }
    3) Event sourcing: Armazenamento dos eventos baseado em uma linha do tempo. Possibilidade de replay

4. Event Colaboration

    Itallo comprou um produto -> Estoque do produto comprado -> muda o catálogo -> emite nota -> Separação


5. CQRS ( Command Query Responsibility Segregation) + Event sourcing
    CQS vs CQRS
        Nivel de granularidade



