rabbit:
	docker run -d --hostname my-rabbitMQ --name some-rabbit -p 15672:15672 -p5672:5672 rabbitmq:latest