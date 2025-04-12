
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout e12219d326f1439062468e87eff1f9c794514638
kustomize build ./kustomize/overlays/test
```