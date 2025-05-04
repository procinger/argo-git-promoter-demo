
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 83fbf9fb0b947bb81c2f9ca623e9859443abb668
kustomize build ./kustomize/overlays/test
```