# prom-graf
A **Golang Application** Monitoring System With **Prometheus** And **Grafana**.

### Cautions:
- While implementing this project, the following software versios were used
  - **Docker** - version `20.10.22`, build 3a2c30b
  - **Docker Compose** - version `1.29.0`, build 07737305
- Make sure the directories `~/.docker-volumes/prometheus` and `~/.docker-volumes/grafana` exist in the host machine.
- Allow WRITE permission to the persistent data volumes `~/.docker-volumes/prometheus` and `~/.docker-volumes/grafana`
  for **prometheus** and **grafana** respectively in the host machine.

### Ports
- Prometheus is running in the port `9090` inside the container, and mapped to `9091` in the host machine.
- Grafana is running in the port `3000` inside the container, and mapped to `3001` in the host machine.
- App is running in the port `2020` inside the container, and mapped to `2021` in the host machine.

### How To Run
- Run the following docker-compose command to start the project:
``` bash
docker-compose up -d
```
