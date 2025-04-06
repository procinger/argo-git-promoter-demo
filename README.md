
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout ad4dac44ac3d36dc9da0d790a82669d1145e59a3
kustomize build ./kustomize/overlays/test
```