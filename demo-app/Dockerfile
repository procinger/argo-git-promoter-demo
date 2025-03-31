####### backend-builder
FROM docker.io/golang:1.23.8-alpine3.20 AS backend-builder

COPY backend /backend
RUN cd /backend && go build -o backend

####### backend
FROM docker.io/golang:1.23.8-alpine3.20 AS backend

WORKDIR /
LABEL org.opencontainers.image.source="https://github.com/procinger/argo-git-promoter-demo"

COPY --from=backend-builder /backend/backend /backend
COPY version /version

ENTRYPOINT ["/backend"]

####### frontend-builder
FROM docker.io/golang:1.23.8-alpine3.20 AS frontend-builder

COPY frontend /frontend
RUN cd /frontend && go build -o frontend

####### frontend
FROM docker.io/golang:1.23.8-alpine3.20 AS frontend

WORKDIR /
LABEL org.opencontainers.image.source="https://github.com/procinger/argo-git-promoter-demo"

COPY --from=frontend-builder /frontend/frontend /frontend
COPY version /version

ENTRYPOINT ["/frontend"]
