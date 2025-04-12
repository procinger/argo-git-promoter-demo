
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 8870f5cc6184c2d48c0983edc7a44c5d1ac9d3d9
kustomize build ./kustomize/overlays/dev
```