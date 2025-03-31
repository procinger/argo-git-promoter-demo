
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout c1be6148a2c501ab03f88d993f95d4359cd6abe4
kustomize build ./kustomize/overlays/test
```