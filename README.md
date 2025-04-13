
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout c043a533df406c8f6816787faa3929f54633a0ef
kustomize build ./kustomize/overlays/dev
```