
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout ea7106be5174a6ed6f66fca07edd4d1936802875
kustomize build ./kustomize/overlays/dev
```