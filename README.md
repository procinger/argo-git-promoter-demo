
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 4a50755818bff93efcd9af5d3fc554da4598d3de
kustomize build ./kustomize/overlays/test
```