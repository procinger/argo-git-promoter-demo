
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 2ed947c723d02ef2ee2c507f69f6b0083a105aa2
kustomize build ./kustomize/overlays/test
```