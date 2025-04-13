
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 4ad628fe716dc969075b0d3b515df70e646caacd
kustomize build ./kustomize/overlays/dev
```