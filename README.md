
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 09a24f7ab1fb27c77eae59f326d6fa0eecdc0300
kustomize build ./kustomize/overlays/test
```