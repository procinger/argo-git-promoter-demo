
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 95a799cc6a1f8cfd6bc5a4e277f8efc79931a3a7
kustomize build ./kustomize/overlays/dev
```