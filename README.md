
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout e689263a9d1ef20db58fdda1e1afe3dd24578bb5
kustomize build ./kustomize/overlays/test
```