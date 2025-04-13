
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 5bf6c998305108ba18bca4e61fcb9dbb011079bd
kustomize build ./kustomize/overlays/dev
```