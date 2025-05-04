
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout cee7201e4d0c289943bc8c27401908894a8ba9db
kustomize build ./kustomize/overlays/test
```