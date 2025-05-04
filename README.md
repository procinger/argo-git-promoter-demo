
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 0370ba59bf459e1e9469c987c13f4118e595f508
kustomize build ./kustomize/overlays/dev
```