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

To check if the deployment is successfull, run the below,
```
 k get po,ksvc -n watermark 
 ```
 
 You should see something similar to this,
 ```
NAME                                              READY   STATUS    RESTARTS   AGE
pod/watermark-00001-deployment-667f4f8448-6gllv   2/2     Running   0          24s

NAME                                    URL                                                LATESTCREATED     LATESTREADY       READY   REASON
service.serving.knative.dev/watermark   http://watermark.watermark.10.95.16.209.sslip.io   watermark-00001   watermark-00001   True    
```

So far so good!. Let's test the app now by hitting the health endpoint.
```
$ curl http://watermark.watermark.10.95.16.209.sslip.io/health
{"status":"OK"}
```

You can provide image url and text to watermark as a parameter.
```
curl --request POST \
  --url http://watermark.watermark.10.95.16.209.sslip.io/watermark \
  --header 'content-type: multipart/form-data' \
  --form url=https://www.pitara.com/media/mango.jpg \
  --form text=MyWaterMark@2024 >watermarked.jpg
```

## Debugging
To check the issues with routes and configs, run the below
```
kubectl get routes,configuration 
```

## Knative serving Pods are missing after a minute !!
This is the expected behaviour, you are basically running Function as a Service(FaaS). If there is no request to service, pod will be terminated. It will spin up a new one as soon as you hit the serving endpoint. In this case, http://watermark.watermark.10.95.16.209.sslip.io 