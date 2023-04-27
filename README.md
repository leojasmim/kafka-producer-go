# kafka-producer-go
Exemplo de API para envio de evento ao Kafka

- Pré Requisitos
  - Go 1.19.8
  - Docker
  
Para iniciar a aplicação execute o seguinte comando no diretório do projeto

~~~ bash
docker-compose up --build
~~~

No terminal poderá ser realizado o envio da mensagem com o comando cUrl

~~~ bash
curl --location 'http://localhost:3000/message' \
--header 'Content-Type: application/json' \
--data '{
    "message":"Exemplo de mensagem"
}'
~~~

A mensagem poderá ser verificada através do _client_ **Kafdrop** disponível em ``http:\\loalhost:19000`

![image](https://user-images.githubusercontent.com/18493760/234979404-02f783af-be19-4b8b-8acf-84be9e4d522c.png)
