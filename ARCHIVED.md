## Archived Projects
  * 🔴 [RabbitMQ > Hello World](https://www.rabbitmq.com/tutorials/tutorial-one-go)
```yml
# docker-compose.yml

# Alternative:  docker run -d --name rabbitmq-container -p 5672:5672 -p 15672:15672 -v rabbitmq-data:/var/lib/rabbitmq rabbitmq:4-management
# Run this file in detached mode: docker-compose up -d
# To view logs later: docker compose logs

version: '3.8' # version of Docker Compose syntax
services:
  rabbitmq:
    image: rabbitmq:4-management
    container_name: rabbitmq-container
    ports:
      - "5672:5672"  # messaging
      - "15672:15672" # management dashboard
    volumes:
      - rabbitmq-data:/var/lib/rabbitmq  # uses an existing volume (rabbitmq-data) within the service.

volumes:
  rabbitmq-data:  # declares the volume (rabbitmq-data) to ensure it is created and managed properly by Docker.
```
  * 🔴 Redis
```yml
# docker-compose.yml

# Alternative:  docker run -d --name redis-container -p 6379:6379 -v redis-data:/data redis:latest redis-server --save 60 1 --loglevel warning
# Run this file in detached mode: docker-compose up -d
# To view logs later: docker compose logs

version: '3.8' # version of Docker Compose syntax
services:
  redis:
    image: redis:latest
    container_name: redis-container
    ports:
      - "6379:6379"
    volumes:
      - redis-data:/data  # uses an existing volume (redis-data) within the service.
    command: ["redis-server", "--save", "60", "1", "--loglevel", "warning"]

volumes:
  redis-data:  # declares the volume (redis-data) to ensure it is created and managed properly by Docker.
```
  * 🔴 SQLite ( _pkg.go.dev/modernc.org/sqlite_ )