
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 4b6951a44f0ea3b4235a8bfb46e48dfb5ba67777
kustomize build ./kustomize/overlays/dev
```