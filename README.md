
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout e64d2d909077322b071579cdaf532d89798fa5ce
kustomize build ./kustomize/overlays/dev
```