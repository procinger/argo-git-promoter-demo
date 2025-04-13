
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 3bd53a65899fd4914c3c1f09595dbd268e5f9eec
kustomize build ./kustomize/overlays/dev
```