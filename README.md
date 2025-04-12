
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 28cb7f4e06c03037184cb149c75fce1ace256f6f
kustomize build ./kustomize/overlays/dev
```