
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout cbe0647ca85374395c93a65be6d6df6c0bafde2f
kustomize build ./kustomize/overlays/test
```