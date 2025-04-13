
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout cef3d2866d6afa60c5574be6ef12cea17d0016a6
kustomize build ./kustomize/overlays/dev
```