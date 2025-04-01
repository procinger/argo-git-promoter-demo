
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout eb9fffa753cb876f49a792561205654946c2f722
kustomize build ./kustomize/overlays/dev
```