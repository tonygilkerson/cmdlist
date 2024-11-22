# Aliases for the IL4 stuff that I don't want to push to github


# Hub ntahub (ssh only)
alias ws-il4-ntahub-start="gcloud workstations start cws-il4-nonprod-ntahub-gke-gilkerson-ext --project=afrl-il4-hub-ops-m2bl --cluster=il4-aw-ops-cws-cluster --config=ops-wolfi-ntahub --region=us-central1 --account tony.gilkerson.ext@afresearchlab.com"
alias ws-il4-ntahub-ssh="  gcloud workstations ssh   cws-il4-nonprod-ntahub-gke-gilkerson-ext --project=afrl-il4-hub-ops-m2bl --cluster=il4-aw-ops-cws-cluster --config=ops-wolfi-ntahub --region=us-central1 --account tony.gilkerson.ext@afresearchlab.com"

# Hub Dev (ssh only)
alias ws-il4-hub-dev-start="gcloud workstations start cws-il4-dev-gke-gcp-gilkerson-ext --project=afrl-il4-hub-ops-m2bl --cluster=il4-aw-ops-cws-cluster --config=ops-wolfi-dev --region=us-central1 --account tony.gilkerson.ext@afresearchlab.com"
alias ws-il4-hub-dev-ssh="  gcloud workstations ssh   cws-il4-dev-gke-gcp-gilkerson-ext --project=afrl-il4-hub-ops-m2bl --cluster=il4-aw-ops-cws-cluster --config=ops-wolfi-dev --region=us-central1 --account tony.gilkerson.ext@afresearchlab.com"

# Teams
alias ws-il4-team-start="gcloud workstations start cws-il4-team-gke-gcp-gilkerson-ext --project=afrl-il4-hub-ops-m2bl --cluster=il4-aw-ops-cws-cluster --config=ops-wolfi-team --region=us-central1 --account tony.gilkerson.ext@afresearchlab.com"
alias ws-il4-team-ssh="  gcloud workstations ssh   cws-il4-team-gke-gcp-gilkerson-ext --project=afrl-il4-hub-ops-m2bl --cluster=il4-aw-ops-cws-cluster --config=ops-wolfi-team --region=us-central1 --account tony.gilkerson.ext@afresearchlab.com"

# Hub Prod (vscode)
alias ws-il4-hub-prod-start="gcloud workstations start --project=afrl-il4-prod-cws-udwe --cluster=il4-cws-cluster --config=tf-wolfi --region=us-central1 cws-tony-gilkerson-ext"
alias ws-il4-hub-prod-ssh="  gcloud workstations ssh   --project=afrl-il4-prod-cws-udwe --cluster=il4-cws-cluster --config=tf-wolfi --region=us-central1 cws-tony-gilkerson-ext"
alias ws-il4-hub-prod-open="open https://80-cws-tony-gilkerson-ext.cluster-uiz7lwpy35gd6xmn3dxolrdhey.cloudworkstations.dev/?authuser=0"

# Ops (vscode)
alias ws-ops-vdi-start="gcloud workstations start cws-vdi-tony-gilkerson-ext --project=afrl-ops-org-ops-vdi-jpr8 --cluster=sre-cws-cluster --config=sre-cws-config-wolfi --region=us-central1 --account tony.gilkerson.ext@afresearchlab.com"
alias ws-ops-vdi-ssh="  gcloud workstations ssh   cws-vdi-tony-gilkerson-ext --project=afrl-ops-org-ops-vdi-jpr8 --cluster=sre-cws-cluster --config=sre-cws-config-wolfi --region=us-central1 --account tony.gilkerson.ext@afresearchlab.com"
alias ws-ops-vdi-open="open https://80-cws-vdi-tony-gilkerson-ext.cluster-wsjmz7e3encl4tmjfeswircom4.cloudworkstations.dev/?authuser=0"

# Corelight Dev - admin vm
alias ws-corelight-dev-admin-vm-ssh="gcloud compute ssh vm-gke-admin-console --zone us-central1-b --tunnel-through-iap --project afrl-il4-dev-corelight-8hjs"

# tgilkerson Research - admin vm
alias ws-tgilkerson-rsh-admin-vm-ssh="gcloud compute ssh vm-gke-admin-console --zone us-central1-b --tunnel-through-iap --project afrl-il4-rch-tgilkerson-a247"

# Test instance on test vpc for testing Corelight Sensor
alias ws-client-a="gcloud compute ssh --zone us-central1-a afrl-il4-rch-tgilkerson-a247-client-a --tunnel-through-iap --project afrl-il4-rch-tgilkerson-a247"

# List the cloud workstations
alias wsl="alias | grep 'ws-' | awk -F= '{print \$1}' | awk '{print \$2}' | grep -v cwsl"
