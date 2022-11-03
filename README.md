eksctl create cluster -f eks.yaml

helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm upgrade --install ingress-nginx ingress-nginx/ingress-nginx \
 --namespace ingress-nginx \
 --create-namespace \
 --timeout 600s \
 --debug \
 --set controller.publishService.enabled=true

<!-- kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.8.2/cert-manager.crds.yaml -->

helm install \
 cert-manager jetstack/cert-manager \
 --namespace cert-manager \
 --create-namespace \
 --version v1.8.2 \
 --set installCRDs=true

helm install \
 --set authorizer.database_type=dynamodb \
 --set authorizer.aws_access_key_id=${AWS_ACCESS_KEY_ID} \
 --set authorizer.aws_secret_access_key=${AWS_SECRET_ACCESS_KEY} \
 --set authorizer.aws_region=us-east-1 \
 --set authorizer.authorizer_url=https://auth.aws-demo.authorizer.dev \
 --set redis.install=true \
 --set redis.storage=5Gi \
 --set redis.storageClassName=gp2 \
 --set securityContext.readOnlyRootFilesystem=false \
authorizer authorizer/authorizer

// Update hostedzoneID & email in cluster_issuer.yaml based on Route53
kubectl apply -f cluster_issuer.yaml

kubectl apply -f authorizer_ingress.yaml

---

API

- change `authorizer_client_id` to based64 encoded value as per your deployment client_id of authorizer instance
