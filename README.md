# Concurra

**Concurra** is a Kubernetes-native platform autoscaling system that allows platform engineers to schedule events with a target concurrency (in millions of users) and automatically scale all applications in the cluster ahead of time — before traffic arrives.

Built entirely in Go, Concurra is composed of two independent services and a set of shared modules, all living in a multi-module monorepo.

---

## The problem it solves

Large platforms hosting live events (sports broadcasts, product launches, live streams) experience sudden, predictable traffic spikes. Traditional reactive autoscaling kicks in too late — by the time the HPA detects load, users are already experiencing degraded service.

Concurra solves this by letting engineers schedule events in advance. Each application declares how it scales relative to platform-wide concurrency via a Kubernetes CRD (`ScalingSpec`), and Concurra ensures every application is at full capacity **before** the event starts — not after.

---

## How it works

Each application in the cluster owns a `ScalingSpec` custom resource that defines:

- The **PromQL query** that returns its current observed concurrency
- A **concurrency ratio** — how many users one replica can serve
- An **expected scaling time** — how long the app takes to scale up

When an event is scheduled, the **scaler** service computes the trigger time per app:

```
scaleUpAt = event.StartTime - scalingSpec.ExpectedScalingTime
```

At that moment, it computes the target replica count:

```
targetValue = event.TargetConcurrency / scalingSpec.ConcurrencyRatio
```

And patches each app's [KEDA](https://keda.sh) `ScaledObject` with that value. KEDA then drives the Kubernetes HPA, which scales the Deployment. By the time the event starts, every application is already at capacity.

---

## Architecture

```
┌─────────────────────────────────────────┐
│             React Web App               │
└─────────────────┬───────────────────────┘
                  │ HTTP
┌─────────────────▼───────────────────────┐
│              scheduler                  │
│  REST API · Event store (Postgres)      │
│  POST /events · GET /events             │
│  DELETE /events/:id · GET /events/:id   │
└─────────────────┬───────────────────────┘
                  │ polls every 30s
┌─────────────────▼───────────────────────┐
│               scaler                    │
│  Reconciler loop · Target calculator    │
│  Reads ScalingSpec CRDs from cluster    │
│  Patches KEDA ScaledObjects             │
└──────┬──────────────────────┬───────────┘
       │                      │
┌──────▼──────┐     ┌─────────▼──────────┐
│ ScalingSpec │     │  KEDA ScaledObject │
│    CRDs     │     │   → HPA → Pods     │
└─────────────┘     └────────────────────┘
                              ▲
                    ┌─────────┴──────────┐
                    │     Prometheus     │
                    │  scrapes app pods  │
                    └────────────────────┘
```

**scheduler** — owns the event lifecycle. Exposes a REST API backed by Postgres. Has no knowledge of Kubernetes.

**scaler** — owns the cluster reconciliation loop. Polls the scheduler for active events, reads `ScalingSpec` CRDs, computes per-app target replica counts, and patches KEDA `ScaledObjects`. Has no database and no HTTP server.

This separation means the two services fail independently. If the cluster has issues, the scheduler still accepts and stores events. If the scheduler is temporarily unreachable, the scaler can finish handling an in-progress event.

---

## The ScalingSpec CRD

Each application team owns a `ScalingSpec` custom resource in their namespace:

```yaml
apiVersion: concurra.io/v1
kind: ScalingSpec
metadata:
  name: my-app
  namespace: my-team
spec:
  targetDeployment: my-app
  concurrencyRatio: 80000        # 1 replica per 80k concurrent users
  expectedScalingTime: 10m       # start scaling 10 minutes before event
  promQL: "sum(rate(active_sessions_total{app='my-app'}[2m]))"
  minReplicas: 2
  maxReplicas: 50
```

The `ScalingSpec` CRD schema is defined in Go using [kubebuilder](https://book.kubebuilder.io) struct tags and generated into `config/crd/bases/` via `controller-gen`.

---

## Tech stack

| Layer | Technology |
|---|---|
| Services | Go 1.22 |
| HTTP router | [chi](https://github.com/go-chi/chi) |
| Database | PostgreSQL via [pgx](https://github.com/jackc/pgx) |
| Kubernetes client | [controller-runtime](https://github.com/kubernetes-sigs/controller-runtime) |
| Autoscaling | [KEDA](https://keda.sh) with Prometheus scaler |
| CRD generation | [controller-gen](https://github.com/kubernetes-sigs/controller-tools) |
| Observability | Prometheus + Grafana |
| Frontend | React + TypeScript (Vite) |

---

## Event lifecycle

```
created → pending → active → ended
```

| Status | Meaning |
|---|---|
| `pending` | Event stored, scaling not yet triggered |
| `active` | Scaler has patched all ScaledObjects, apps are scaling up |
| `ended` | Event is over, scaler has reset all apps to baseline |

---

## Inspiration

This project is a personal reconstruction and open-source reimagining of a platform autoscaling system built at a large media company to handle live event traffic at scale. The original system served millions of concurrent users across dozens of microservices. Concurra is built from scratch with the same core ideas, using idiomatic Go and modern Kubernetes tooling.