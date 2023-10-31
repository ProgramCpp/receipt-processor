# API documentation

Open API documentation for receipt processor.

# run

docker pull swaggerapi/swagger-ui
docker run -p 1200:8080 -e SWAGGER_JSON=/api.yml -v /<absolute_path>/api.yml:/api.yml swaggerapi/swagger-ui

## usage

swagger UI is hosted [here](httlp://localhost:1200/).