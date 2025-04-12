
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 79c6ea8e47e5d3167829370d370b96d0f9dd71e8
kustomize build ./kustomize/overlays/dev
```