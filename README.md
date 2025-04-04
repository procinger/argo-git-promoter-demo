
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 3356bdac08b553cef510e934c316466ae0b98ef2
kustomize build ./kustomize/overlays/dev
```