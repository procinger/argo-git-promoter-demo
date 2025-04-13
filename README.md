
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 2dc0627c0d9f4b55e4f4e90bdad6254e079d6c6e
kustomize build ./kustomize/overlays/dev
```