
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 9e2944ddd6c529dd40ce765d555a50c9ab185649
kustomize build ./kustomize/overlays/test
```