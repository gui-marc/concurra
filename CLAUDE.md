# events-autoscalling-platform

Event-driven autoscaling platform with a full Grafana observability stack for local development.

## Observability Stack

Start with: `docker compose up -d`

| Service    | URL                    | Purpose              |
|------------|------------------------|----------------------|
| Grafana    | http://localhost:3000  | Dashboards (no auth) |
| Prometheus | http://localhost:9090  | Metrics              |
| Loki       | http://localhost:3100  | Logs                 |
| Tempo      | http://localhost:3200  | Traces               |

OTLP ingestion on Tempo: gRPC `4317`, HTTP `4318`.

## Config files

```
config/
  observability/
    grafana/datasources/datasources.yml   # auto-provisions all 3 datasources
    prometheus/prometheus.yml             # scrape config (extend with your services)
    loki/loki.yml                         # single-node filesystem storage
    tempo/tempo.yml                       # local backend, 1h retention
```

## Extending

- **Add a scrape target**: append to `scrape_configs` in [observability/prometheus/prometheus.yml](observability/prometheus/prometheus.yml).
- **Ship logs**: point your app's log shipper (e.g. Alloy, Promtail, or the OTLP log endpoint) to `http://localhost:3100`.
- **Ship traces**: use OTLP exporter pointing to `localhost:4317` (gRPC) or `localhost:4318` (HTTP).
