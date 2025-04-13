
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 3daaaa71802f5690acfab619a96f290e2dca0f58
kustomize build ./kustomize/overlays/test
```