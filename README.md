
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 2d93f859c8795a2b8dfb4dff489e3722121a8fb7
kustomize build ./kustomize/overlays/dev
```