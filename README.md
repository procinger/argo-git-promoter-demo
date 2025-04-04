
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 34f6a239ebbbdae247f8c76f149a88e14e957af4
kustomize build ./kustomize/overlays/test
```