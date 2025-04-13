
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout e3ddba9d3c223d0fd193cc07e4113aa3e9f21fac
kustomize build ./kustomize/overlays/dev
```