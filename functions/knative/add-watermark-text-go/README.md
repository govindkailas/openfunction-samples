# Knative serving

## Prerequisites
A Kubernetes cluster with Knative installed and DNS configured. See [Install Knative Serving](https://knative.dev/docs/install/serving/install-serving-with-yaml)

## Apply the serving
This is a simple app that will add a watermark text to an image. More details about the [project here](https://github.com/govindkailas/go-watermark/tree/add-gin-router) 

Lets look at how to deploy it as a knative serving,
```
kubectl create ns watermark
kubectl apply --filename addwatermark-knative-serving.yaml
```

You can provide image url and text to watermark as an argument.
