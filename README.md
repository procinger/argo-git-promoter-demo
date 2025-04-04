
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 73e8c72daedc380641490ae0a4da790d46e7d5df
kustomize build ./kustomize/overlays/test
```