
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 718a3d1f7946af0cc407881da3e111645f91f9c6
kustomize build ./kustomize/overlays/dev
```