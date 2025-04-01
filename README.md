
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 3fc2b3bc9f3dcc9c109737e9319402abca72c6cd
kustomize build ./kustomize/overlays/dev
```