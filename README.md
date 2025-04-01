
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout e597e24c70b86ee9a04c2431060e10c3137f0969
kustomize build ./kustomize/overlays/test
```