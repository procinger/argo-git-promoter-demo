
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout c9de145e77bf3ac6fd752bfb4722453c6ebd3d7b
kustomize build ./kustomize/overlays/dev
```