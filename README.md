
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 495304c8f72fe7c12c4394c089aeedfa24c4df76
kustomize build ./kustomize/overlays/test
```