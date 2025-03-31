
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout b95e2ef4fe370d2387fd247b190a21ccc32a863e
kustomize build ./kustomize/overlays/dev
```