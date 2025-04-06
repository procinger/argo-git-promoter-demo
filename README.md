
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 924c2918557266c4b38cdfe0351f2e4a63a2d923
kustomize build ./kustomize/overlays/dev
```