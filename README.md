
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 838f8d6e0aaf2573052428472fd7ae7556679271
kustomize build ./kustomize/overlays/test
```