cd fixtures && docker-compose down -v
sudo docker rm -f $(sudo docker ps -aq)
sudo docker network prune
sudo docker volume prune
docker-compose up -d
cd ~/go/src/application
rm application
go build
./application

# clipier

#npx caliper launch manager --caliper-workspace ./ --caliper-networkconfig networks/networkConfig.yaml --caliper-benchconfig benchmarks/myIdBenchmark.yaml --caliper-flow-only-test --caliper-fabric-gateway-enabled