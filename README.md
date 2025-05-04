
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 14b68781c9779ebd8755889e79ada603a4b2782e
kustomize build ./kustomize/overlays/dev
```