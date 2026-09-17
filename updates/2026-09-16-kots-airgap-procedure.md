---
id: 2026-09-16-kots-airgap-procedure
title: KOTS existing-cluster install page (online, proxy, and air gap)
published_at: 2026-09-16T00:00:00Z
impact: recommended
summary: The default Existing Cluster (KOTS) page uses the same Network Availability radios as Helm and Linux. Online and proxy use generated install commands. Air gap uses a docs procedure with registry placeholders and a wrappable download component. The page is no longer locked to air gap.
affects:
  - installation
  - kots
---

The default Existing Cluster (KOTS) page used to stop at the license and asset downloads. Air gap customers had no way to see from the portal that they need a private registry and two accounts, and no in-portal procedure for `push-images`, `kots install`, or uploading the license and `.airgap` bundle in the Admin Console.

`pages/installation/kots.md` is now that procedure. Every vendor who follows the template gets it without writing a generator. Registry host and credentials stay documentation-style placeholders (`REGISTRY_HOST`, `RW_USERNAME`, `RO_USERNAME`) that the customer substitutes. The application slug is filled in with `{{app.slug}}` because the portal knows it. `--namespace` is the app slug, not Kubernetes `default`.

The KOTS CLI is installed from the file downloaded on the page, not `curl https://kots.io/install`. License and `.airgap` are uploaded in the Admin Console after the Admin Console is running.

Do not replace this page with `<KotsAirgapInstallAssets />`. That component is still downloads only. Do not generate `push-images` or `kots install --kotsadm-registry` with customer registry passwords.

This page uses `<WhenNetwork>`. Your portal must be on a release that includes that component (vandoor PR #10535 / the Enterprise Portal build that shipped it). If you adopt the page before that release, the tag is unknown and the page fails to render.

`pages/installation/kurl.md` is unchanged by this update. It remains the downloads-only air gap page and stays gated on `isKurlInstallEnabled` and `isAirgapSupported` in both frontmatter and `toc.yaml`.

## Wrapping the download component

`<KotsDownloadAssets />` (and `<KurlDownloadAssets />`) take an optional `stepNumber`. The default page wraps the picker in an `<InstallStep>` instead, so the surrounding copy owns the numbers:

```mdx
<InstallStep stepNumber={1} title="Your network and registry rules">
Harbor project, jump host, support contact, …
</InstallStep>

<InstallStep stepNumber={2} title="Download the installation assets">
<KotsDownloadAssets />
</InstallStep>

<InstallStep stepNumber={3} title="Push Admin Console images">
Your REGISTRY_HOST convention, then the docs-style command.
</InstallStep>
```

If you want the picker itself to draw the numbered chrome, omit the wrapping `<InstallStep>` and pass the position through:

```mdx
<KotsDownloadAssets stepNumber={2} />
```

Do not do both, or the customer sees two step headers.

## Apply this update to your repo

Your content repo was created from the Enterprise Portal template repository, not forked from it. A repo created from a template starts with its own commit history rather than a copy of the template's, so the two share no common ancestor and a plain `git pull` or `git merge` from the template fails with `refusing to merge unrelated histories`. The correct approach is to add the template as an upstream remote, fetch its changes, and then review what differs.

### 1. Set up the upstream remote (one-time)

Run this once per content repo. If `upstream` already exists, skip to step 2.

```shell
git remote add upstream https://github.com/replicatedhq/enterprise-portal-content.git
```

### 2. Fetch the latest template changes

```shell
git fetch upstream
```

The commands below use this update's template commit, `TEMPLATE_COMMIT_SHA`, so later template changes are not pulled in accidentally.

Before checking out template files, make sure you do not have uncommitted work in the target paths:

```shell
git status --short
```

Commit or stash any local changes before continuing. The restore command below returns files to `HEAD`; it cannot recover uncommitted edits overwritten by checkout.

### 3. Compare your KOTS page and nav gate

This update touches two files:

- `pages/installation/kots.md` — the online/proxy/air gap install page
- `toc.yaml` — the Existing Cluster (KOTS) nav item is gated on `isKotsInstallEnabled` **only**. This update removes `isAirgapSupported` from that entry so online-only KOTS customers see the page

```shell
git diff HEAD TEMPLATE_COMMIT_SHA -- pages/installation/kots.md toc.yaml
```

### 4. Take the new page and nav gate

If you have not customized those files, take the template versions:

```shell
git checkout TEMPLATE_COMMIT_SHA -- pages/installation/kots.md toc.yaml
```

If you started from a clean worktree and decide not to keep the copied files, restore them before committing:

```shell
git checkout HEAD -- pages/installation/kots.md toc.yaml
```

If the diff shows copy you want to keep (registry naming, network requirements, support contact), keep your surrounding steps and adopt the procedure: two credential pairs, downloaded CLI rather than `curl kots.io`, `--namespace` as the app slug, Admin Console for license and `.airgap`. Leave `<KotsDownloadAssets />` in place so customers can still pick a version. Put your copy in `<InstallStep>` blocks above and below it.

If you keep a customized `toc.yaml`, still remove `isAirgapSupported` from the Existing Cluster (KOTS) nav entry. Leaving that gate in place hides the page from online-only KOTS customers — the population this update exists to serve — and they get no nav item and a 404.

### 5. Review, preview, commit, and push

Inspect the changes, run a local preview, then commit and push. Use `git diff HEAD` rather than a bare `git diff`, because the checkout in step 4 stages the change and a bare `git diff` would show nothing:

```shell
git diff HEAD
replicated enterprise-portal preview . --app <your-app-slug>
git add pages/installation/kots.md toc.yaml
git commit -m "Adopt Enterprise Portal template update: KOTS air gap install procedure"
git push
```

> **Note:** Replace `<your-app-slug>` with your app's slug. You can find your app's preview command in Enterprise Portal > Content > Preview.
>
> If you maintain version branches, apply and push this update on each branch where customers should see it.
