
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 96479bd89299e0cb240bcd66b14b2be4b7c87641
kustomize build ./kustomize/overlays/dev
```