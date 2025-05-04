
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 838ee1639b64d7ec09ebb41fbd8d67ade1bbb131
kustomize build ./kustomize/overlays/dev
```