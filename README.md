
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 363811f38566eaa35a208afd8ec2e447ed570b78
kustomize build ./kustomize/overlays/test
```