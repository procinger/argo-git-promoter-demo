
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout a98de5bba2e9ddcd5dd20118805b818156176b79
kustomize build ./kustomize/overlays/dev
```