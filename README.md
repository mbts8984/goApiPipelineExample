# pipeline-example

example for jamf
docker desktop
run in kind
kind create cluster --name local-test

kubectl get pods

# 🧩 Basic API Pipeline Example

_A lightweight Go web service with automated Kubernetes deployment using Kind, Helm, and GitHub Actions._

---

## 📘 Overview

This repository demonstrates a full CI/CD workflow for a simple Go-based API deployed to a local **Kind (Kubernetes in Docker)** cluster.  
The GitHub Actions pipeline handles building, pushing, deploying, and verifying your application end-to-end — making it easy to reproduce locally or debug via **tmate**.

---

## What's it do?

A minimal Go API running on port `8080`:

```go
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello World")
    })
    http.ListenAndServe(":8080", nil)
}
```

When running, http://localhost:8080

⚙️ Prerequisites

Install these tools locally to mirror the CI environment:
Tool - Version - Install
Go ≥ 1.22 - golang.org/dl
Docker ≥ 24 - docs.docker.com (I use Docker Desktop)
Kubectl ≥ 1.30 - kubernetes.io
Helm ≥ 3.14 - helm.sh
Kind ≥ 0.23 - kind.sigs.k8s.io

# Running the App Locally

## Run locally

Pull the repo down and cd into it then run it locally to verify it works

```
go run main.go
```

## The included workflow automates the full build → deploy → verify cycle.

## Make sure Kubernetes is enabled in Docker

📄 .github/workflows/ci.yml
Pipeline steps:
Build & Test Go App
Build & Push Docker Image — Builds and pushes to Docker Hub.
Create Kind Cluster — Spins up a temporary Kubernetes cluster inside GitHub Actions.
Export kubeconfig — Ensures Helm and kubectl point to the correct cluster.
Deploy Helm Chart
Verify API Response — Uses kubectl port-forward + curl to confirm Hello World.
(Optional) action-tmate — Opens a live debugging session for interactive testing. [link](https://github.com/mxschmitt/action-tmate?tab=readme-ov-file)

🔐 Required GitHub Infp
Name Description
DOCKERHUB_USERNAME - Your Docker Hub username (new variable)
DOCKERHUB_TOKEN - Docker Hub access token for docker push (new secret)
Public SSH Key - Optional, for tMate validation

# Trigger Workflow

1. Make a change to a file and push it up to Github. Create a new PR and the workflow will kick off.

   - When the workflow gets to the 'deploy-and-validate' stage, click into the workflow details.
   - Inside, wait for 'Start tMate Session' to start

     - Click in, and wait for a series of ssh code to show up, should look similar to:
       🔹 How to Connect
       When the workflow reaches the tmate step, it prints something like:
       SSH: ssh abcdefgh@nyc1.tmate.io

       SSH URL: Connect from your local terminal:
       `ssh abcdefgh@nyc1.tmate.io`

   - 🔹 Once Inside tmate
     Check your pods:
     `kubectl get pods`

   Run: `kubectl port-forward deployment/pipeline-example 9091:8080 &`
   Wait a couple of seconds, then run: `curl http://localhost:9091`

   You should see a 'Hello World' returned to you.
   To close the instance, run `exit`. This will also finish your pipeline.
   You the pipeline will continue to run until you cancel it you dont use the tMate option.

🧹 Cleanup
To reset your environment:
kind delete cluster --name local-test
docker image rm USERNAME/pipeline_example:latest

🔍 Troubleshooting
Issue - Cause - Fix
`port already in use` - Existing forward active `pkill -f "kubectl port-forward"`

## Structure

```.
├── main.go
├── Dockerfile
├── pipeline-example/
│   └─
│       ├── Chart.yaml
│       ├── templates/
│       │   ├── deployment.yaml
│       │   └── service.yaml
│       └── values.yaml
└── .github/
    └── workflows/
        └── ci.yml
```

🏁 Summary

This project is a simple, reproducible template for:
Building and deploying a Go API to Kubernetes.
Managing deployments with Helm.
Automating CI/CD with GitHub Actions.
Debugging live clusters with action-tmate.
