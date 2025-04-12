
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 0c4ef82bfe04b9c90766c8441f4eccc6e5c18669
kustomize build ./kustomize/overlays/test
```