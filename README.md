
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 7ea6f7cc5a705d813d6f51f5c182e11be61c7d61
kustomize build ./kustomize/overlays/dev
```