### Handler  

Os handlers em Go atuam de maneira semelhante aos controladores em outros frameworks. Eles são responsáveis por executar a lógica da aplicação e gerenciar os cabeçalhos e corpos das respostas HTTP. Cada handler processa as requisições recebidas e determina como responder.  

### Router  

O roteador (ou **servemux** em Go) faz a associação entre padrões de URL e seus respectivos handlers. Ele direciona as requisições recebidas para o handler apropriado com base nas rotas definidas. Normalmente, uma aplicação possui um único **servemux** que gerencia todas as rotas.  

### Servidor Web  

Uma das vantagens do Go é seu servidor web embutido. Diferente de outras linguagens que exigem servidores externos como Nginx ou Apache, o Go permite que as aplicações lidem com requisições HTTP de forma nativa. Isso simplifica a implantação e facilita o gerenciamento de requisições diretamente dentro do código da aplicação.

