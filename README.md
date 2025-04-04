
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 8b6cad41b33443eb0cf2496ef65fcd9b2c89d071
kustomize build ./kustomize/overlays/test
```