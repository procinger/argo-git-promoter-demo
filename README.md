
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout df51f3a7afb6369164c0ff8b87b35bbc60db11a4
kustomize build ./kustomize/overlays/dev
```