
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 126b2b605f4957f4e6f72d0bd484e5987146f435
kustomize build ./kustomize/overlays/dev
```