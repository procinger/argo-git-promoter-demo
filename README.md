
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 0c3c7f85f5bb45b7286005b4575799956770dd07
kustomize build ./kustomize/overlays/dev
```