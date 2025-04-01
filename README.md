
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout cdd0546501dd908868967ca4ec63078dc8617a34
kustomize build ./kustomize/overlays/dev
```