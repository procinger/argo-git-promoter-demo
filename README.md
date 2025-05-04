
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout d6ff3b45c32426493fe3d6efdde17ebfc9ef4a08
kustomize build ./kustomize/overlays/dev
```