
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout aa4f21c582528f99ab2562a4ea78587c6e8b840e
kustomize build ./kustomize/overlays/dev
```