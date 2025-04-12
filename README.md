
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout a8e116c3ace320542007f11620c4b1ffa159a7b8
kustomize build ./kustomize/overlays/dev
```