
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout e09d8dca67a2ad206a883b5414a325b4670e3e3a
kustomize build ./kustomize/overlays/dev
```