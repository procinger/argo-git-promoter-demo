
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 35d1b51edf8b875c3c9869b93a36e958df3d508c
kustomize build ./kustomize/overlays/dev
```