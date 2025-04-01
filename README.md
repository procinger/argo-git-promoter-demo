
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 90bb6d3e8ca3a543b6c59a9017dd135e912f48f8
kustomize build ./kustomize/overlays/test
```