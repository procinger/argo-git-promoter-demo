
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 69572b756411254d1f64821a0d7f835700f616fb
kustomize build ./kustomize/overlays/dev
```