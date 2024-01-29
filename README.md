# RabbitMQ Playground

This repository setup a base communication between service 1 ( a node.js App) and service 2 (a go app) via RabbitMQ.

## Run Locally

- Pull RabbitMQ docker image:

```bash
 docker pull rabbitmq:3-management

```

- Run the docker image

```bash
docker run --rm -it -p 15672:15672 -p 5672:5672 rabbitmq:3-management

```

- Run the node.js service

```bash
cd service1 && npm install && node index.js
```

- Run the go service

```bash
cd service2 && go download && go run main.go
```

Now you all set, You should see `Hello world from go server`

## Important Links

- [RabbitMQ go client](https://github.com/rabbitmq/amqp091-go)
- [RabbitMQ Node.js client](https://www.npmjs.com/package/amqplib)
- [Implement RabbitMQ with Docker in 20 min from Architect](https://www.architect.io/blog/2021-01-19/rabbitmq-docker-tutorial/)
- [The Complete Guide to Context in Golang: Efficient Concurrency Management](https://medium.com/@jamal.kaksouri/the-complete-guide-to-context-in-golang-efficient-concurrency-management-43d722f6eaea)
- [Kafka vs. RabbitMQ - who wins and why? | Systems Design Interview 0 to 1 with Ex-Google SWE](https://youtu.be/_5mu7lZz5X4?si=Ss3U7OFVfDgPweQH)

## LICENSE

MIT
