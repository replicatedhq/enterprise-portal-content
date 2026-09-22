---
id: 2026-09-21-kots-online-cli-downloads
title: Workstation CLIs on the default KOTS online page
published_at: 2026-09-22T19:59:14Z
impact: recommended
summary: The default Existing Cluster (KOTS) online and proxy path includes a CLI-only download list (KOTS, Preflight, Support Bundle) after the generated install commands. Air gap downloads stay in the air gap branch. Do not unwrap the air gap tag. Adopt this page only after the portal build that supports cliOnly (vandoor #10549).
affects:
  - installation
  - kots
---

The default page mounts `<KotsDownloadAssets cliOnly={true} />` on the online/proxy branch. Copy leads with Preflight and Support Bundle CLIs (those are not curl-dependent). Customers who cannot run `curl https://kots.io/install` download the KOTS CLI here, extract it, rename `kots` to `kubectl-kots`, put it on PATH, then run the `kubectl kots install` line from step 2. The generated curl command remains step 2 for environments that can use it. This list does not include the Admin Console bundle or the application air gap bundle.

The air gap branch is unchanged: full `<KotsDownloadAssets />` inside `<WhenNetwork mode="airgap">`. Do not move that tag out of the wrap. `cliOnly` keys off the mount, so a dual-entitled customer who selected **online** does not get air-gap installer files in that flow.

Do not adopt this page until the portal lists `cliOnly` on `KotsDownloadAssets` (vandoor #10549 **deployed**, not only merged). If you ship the page first, an older portal strips the prop and shows Admin Console and air-gap bundles on the online branch for customers whose license also has air gap.

## Apply this update to your repo

Your content repo was created from the Enterprise Portal template repository, not forked from it. Add the template as `upstream` if it is not already there, fetch, then take `pages/installation/kots.md` from this update's template commit. `toc.yaml` is unchanged.

### 1. Set up the upstream remote (one-time)

```shell
git remote add upstream https://github.com/replicatedhq/enterprise-portal-content.git
```

### 2. Fetch the latest template changes

```shell
git fetch upstream
```

The commands below use this update's template commit, `d361634b194cbf131fb01dc49585817646435739`.

### 3. Compare your KOTS page

```shell
git diff HEAD d361634b194cbf131fb01dc49585817646435739 -- pages/installation/kots.md
```

### 4. Take the new page

If you have not customized that file:

```shell
git checkout d361634b194cbf131fb01dc49585817646435739 -- pages/installation/kots.md
```

If you keep a customized page, add `<KotsDownloadAssets cliOnly={true} />` on the online/proxy `<WhenNetwork>` branch next to `<KotsInstallAssets />`. Leave the air-gap `<KotsDownloadAssets />` inside `<WhenNetwork mode="airgap">`. Do not unwrap it.

### 5. Review, preview, commit, and push

```shell
git diff HEAD
replicated enterprise-portal preview . --app <your-app-slug>
git add pages/installation/kots.md
git commit -m "Adopt Enterprise Portal template update: KOTS online workstation CLIs"
git push
```

> **Note:** Replace `<your-app-slug>` with your app's slug. If you maintain version branches, apply and push this update on each branch where customers should see it.
