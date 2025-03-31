
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 14c9a7db6042cc71c992ea60367e35a4045afa54
kustomize build ./kustomize/overlays/test
```