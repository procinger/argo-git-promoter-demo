
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 5a59daa19477c301ba00c35834207bbd65f16800
kustomize build ./kustomize/overlays/test
```