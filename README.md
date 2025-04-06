
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout bf2bbd2791dcb2c5d83fc9be0a655e1d61c88b1a
kustomize build ./kustomize/overlays/dev
```