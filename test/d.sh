#!/bin/bash

docker kill prometheus
docker kill alertmanager
docker rm prometheus
docker rm alertmanager

docker run --name=prometheus -v $(pwd)/stuff:/etc/config/  -d --net="host" prom/prometheus:v1.2.3 "-alertmanager.url=http://localhost:9093" "-config.file=/etc/config/prometheus.yml" "-storage.local.path=/prometheus" "-web.console.libraries=/usr/share/prometheus/console_libraries" "-web.console.templates=/usr/share/prometheus/consoles"
docker run --name=alertmanager -v $(pwd)/stuff:/etc/config/ -d --net="host" prom/alertmanager:v0.4.2 "-config.file=/etc/config/alertmanager.yml" "-storage.path=/alertmanager"

docker logs prometheus
docker logs alertmanager
