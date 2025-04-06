
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 71e88951492399d3fc8e992e7ef08d1674ac407c
kustomize build ./kustomize/overlays/test
```