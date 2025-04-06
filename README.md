
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 27d5b63af18853d9e1950be069f1fd616aa4619e
kustomize build ./kustomize/overlays/test
```