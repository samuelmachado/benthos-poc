#!/bin/sh
CREDENTIALS_DIR="/c/teste/benthoskey.json"

echo "======Montando imagens======"
docker-compose -f resources/docker/docker-compose.yml up -d
echo "======Imagens montadas======"
echo "======Aguardando 5s======"
sleep 5
echo "======Setando credencial======"
export GOOGLE_APPLICATION_CREDENTIALS=$CREDENTIALS_DIR
echo "======Produzindo mensagem no PUB/SUB======"
cd resources/producer && go run main.go
echo "======Aguardando 10s======"
sleep 10
echo "Obtendo dados do Elastic Search"
echo "============="
curl -H 'Content-Type: application/json' -X GET http://localhost:9200/benthos_index/_search?pretty
echo "Done!"