# Manual Next Steps — Deploy Bytestone Blog to OVH Managed Kubernetes

## Prerequisites

- OpenTofu (`tofu`) installed
- `kubectl` installed
- `helm` installed
- `htpasswd` installed (from `apache2-utils`)
- Podman installed
- Hugo installed
- An OVH Public Cloud project with billing enabled

## Step 1: OVH API Credentials

Create an API application at https://eu.api.ovh.com/createApp/

Then generate a consumer key:

```bash
curl -X POST https://eu.api.ovh.com/1.0/auth/credential \
  -H "X-Ovh-Application: YOUR_APP_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "accessRules": [
      {"method": "GET", "path": "/cloud/project/*"},
      {"method": "POST", "path": "/cloud/project/*"},
      {"method": "PUT", "path": "/cloud/project/*"},
      {"method": "DELETE", "path": "/cloud/project/*"}
    ]
  }'
```

Visit the validation URL returned, then export:

```bash
export OVH_APPLICATION_KEY="your-app-key"
export OVH_APPLICATION_SECRET="your-app-secret"
export OVH_CONSUMER_KEY="your-consumer-key"
```

## Step 2: Create Terraform Variables File

```bash
cp infra/ovh/terraform.tfvars.example infra/ovh/terraform.tfvars
```

Edit `infra/ovh/terraform.tfvars` and set your OVH project ID:

```hcl
ovh_project_id = "your-actual-project-id"
```

Find your project ID in the OVH Manager under Public Cloud > your project > Project settings.

## Step 3: Provision the Kubernetes Cluster

```bash
task ovh-init
task ovh-plan
```

Review the plan. It will create:
- 1x Managed Kubernetes cluster in GRA9 (France)
- 1x d2-2 node pool (2 vCPU, 2GB RAM, ~EUR6.50/mo)

Apply:

```bash
task ovh-apply
```

This takes 10-15 minutes. Wait for it to complete.

## Step 4: Extract Kubeconfig

```bash
task ovh-kubeconfig
```

This writes `infra/ovh/kubeconfig.yml` (gitignored). Verify connectivity:

```bash
export KUBECONFIG=$(pwd)/infra/ovh/kubeconfig.yml
kubectl cluster-info
kubectl get nodes
```

You should see one `Ready` node.

## Step 5: Check Available Storage Classes

```bash
kubectl get storageclass
```

If the default storage class is not `csi-cinder-high-speed`, edit `infra/ovh/k8s/registry.yaml` and update the `storageClassName` field to match.

## Step 6: Bootstrap the Cluster

Set a registry password (or let it auto-generate one):

```bash
export REGISTRY_PASSWORD="your-chosen-password"
task ovh-bootstrap
```

This installs:
1. nginx-ingress controller (creates a LoadBalancer)
2. cert-manager (for Let's Encrypt TLS)
3. `blog` namespace
4. Registry htpasswd secret
5. Let's Encrypt ClusterIssuer
6. Docker registry:2 (Deployment + PVC + Service + Ingress)
7. Blog Deployment + Service + Ingress

At the end it prints the LoadBalancer IP. Note it down.

If you let the password auto-generate, **save it now**. You'll need it as `OVH_REGISTRY_PASSWORD` for deploys.

## Step 7: Configure DNS

At your domain registrar, create A records pointing to the LoadBalancer IP:

| Record | Type | Value |
|--------|------|-------|
| `www.bytestone.uk` | A | `<LoadBalancer IP>` |
| `registry.bytestone.uk` | A | `<LoadBalancer IP>` |

Wait for DNS propagation. Verify:

```bash
dig +short www.bytestone.uk
dig +short registry.bytestone.uk
```

Both should return the LoadBalancer IP.

## Step 8: Wait for TLS Certificates

cert-manager will automatically request Let's Encrypt certificates once DNS resolves. Check progress:

```bash
kubectl get certificate -n blog
kubectl describe certificate blog-tls -n blog
kubectl describe certificate registry-tls -n blog
```

Both should show `Ready: True` within a few minutes of DNS propagation.

## Step 9: First Deploy

```bash
export OVH_REGISTRY_PASSWORD="your-registry-password"
task deploy-ovh
```

This will:
1. Build Hugo site
2. Build container image with Podman
3. Push to `registry.bytestone.uk`
4. Update the Kubernetes deployment
5. Wait for rollout

### TLS Bootstrap Chicken-and-Egg

If the registry TLS certificate hasn't issued yet (DNS not propagated), the push will fail. Workaround using port-forward:

```bash
# Terminal 1: forward registry port
kubectl port-forward svc/registry 5000:5000 -n blog

# Terminal 2: push via localhost
go run ./cmd/deploy -target ovh -registry-host localhost:5000 -dry-run=false
```

## Step 10: Verify

```bash
curl -I https://www.bytestone.uk/
curl -I https://www.bytestone.uk/nonexistent
```

Expect:
- First request: `200`, valid TLS, correct cache headers
- Second request: `404`

## Step 11: Monitor Resources

The d2-2 node has 2GB RAM. Monitor usage:

```bash
kubectl top nodes
kubectl top pods -n blog
kubectl top pods -n ingress-nginx
kubectl top pods -n cert-manager
```

Expected total usage is ~200MB. If you see memory pressure, consider upgrading to d2-4.

## Ongoing Deploys

For subsequent deploys:

```bash
export OVH_REGISTRY_PASSWORD="your-registry-password"
task deploy-ovh
```

Or as part of a release:

```bash
export OVH_REGISTRY_PASSWORD="your-registry-password"
task release-ovh:v0.2.0:"new blog post"
```

## Re-running Bootstrap

The bootstrap program is idempotent. Re-run it any time to converge state:

```bash
export REGISTRY_PASSWORD="your-registry-password"
task ovh-bootstrap
```

This is safe to run repeatedly — it uses `helm upgrade --install` and `kubectl apply`.

## Troubleshooting

### Pod stuck in ImagePullBackOff
The blog pod can't pull from the in-cluster registry. Check the registry is running:
```bash
kubectl get pods -n blog -l app=registry
kubectl logs -n blog -l app=registry
```

### LoadBalancer IP not assigned
OVH may take a few minutes. Check:
```bash
kubectl get svc ingress-nginx-controller -n ingress-nginx -w
```

### Certificate not issuing
Check cert-manager logs and challenge status:
```bash
kubectl logs -n cert-manager -l app.kubernetes.io/name=cert-manager
kubectl get challenges -n blog
```
